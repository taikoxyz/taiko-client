package guardianproversender

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/db"
)

type healthCheckReq struct {
	ProverAddress      string `json:"prover"`
	HeartBeatSignature []byte `json:"heartBeatSignature"`
}

type signedBlockReq struct {
	BlockID   uint64         `json:"blockID"`
	BlockHash string         `json:"blockHash"`
	Signature []byte         `json:"signature"`
	Prover    common.Address `json:"proverAddress"`
}

type GuardianProverBlockSender struct {
	privateKey                *ecdsa.PrivateKey
	healthCheckServerEndpoint *url.URL
	db                        ethdb.KeyValueStore
	rpc                       *rpc.Client
	proverAddress             common.Address
}

func NewGuardianProverBlockSender(
	privateKey *ecdsa.PrivateKey,
	healthCheckServerEndpoint *url.URL,
	db ethdb.KeyValueStore,
	rpc *rpc.Client,
	proverAddress common.Address,
) *GuardianProverBlockSender {
	return &GuardianProverBlockSender{
		privateKey:                privateKey,
		healthCheckServerEndpoint: healthCheckServerEndpoint,
		db:                        db,
		rpc:                       rpc,
		proverAddress:             proverAddress,
	}
}

func (s *GuardianProverBlockSender) post(ctx context.Context, route string, req interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%v/%v", s.healthCheckServerEndpoint.String(), route),
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"unable to contact health check server endpoint, status code: %v", resp.StatusCode)
	}

	return nil
}

func (s *GuardianProverBlockSender) SignAndSendBlock(ctx context.Context, blockID *big.Int) error {
	signed, blockHash, err := s.sign(ctx, blockID)
	if err != nil {
		return nil
	}

	if signed == nil {
		return nil
	}

	if err := s.sendSignedBlockReq(ctx, signed, blockHash, blockID); err != nil {
		return err
	}

	return nil
}

func (s *GuardianProverBlockSender) sendSignedBlockReq(
	ctx context.Context,
	signed []byte,
	hash common.Hash,
	blockID *big.Int,
) error {
	if s.healthCheckServerEndpoint == nil {
		log.Info("No health check server endpoint set, returning early")
		return nil
	}

	req := &signedBlockReq{
		BlockID:   blockID.Uint64(),
		BlockHash: hash.Hex(),
		Signature: signed,
		Prover:    s.proverAddress,
	}

	if err := s.post(ctx, "signedBlock", req); err != nil {
		return err
	}

	log.Info("Guardian prover successfully signed block", "blockID", blockID.Uint64())

	return nil
}

func (s *GuardianProverBlockSender) sign(ctx context.Context, blockID *big.Int) ([]byte, common.Hash, error) {
	log.Info("Guardian prover signing block", "blockID", blockID.Uint64())

	head, err := s.rpc.L2.BlockNumber(ctx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	for head < blockID.Uint64() {
		log.Info(
			"Guardian prover block signing waiting for chain",
			"latestBlock", head,
			"eventBlockID", blockID.Uint64(),
		)

		if _, err := s.rpc.WaitL1Origin(ctx, blockID); err != nil {
			return nil, common.Hash{}, err
		}

		head, err = s.rpc.L2.BlockNumber(ctx)
		if err != nil {
			return nil, common.Hash{}, err
		}
	}

	header, err := s.rpc.L2.HeaderByNumber(ctx, blockID)
	if err != nil {
		return nil, common.Hash{}, err
	}

	exists, err := s.db.Has(db.BuildBlockKey(header.Time))
	if err != nil {
		return nil, common.Hash{}, err
	}

	if exists {
		log.Info("Guardian prover already signed block", "blockID", blockID.Uint64())
		return nil, common.Hash{}, nil
	}

	log.Info(
		"Guardian prover block signing caught up",
		"latestBlock", head,
		"eventBlockID", blockID.Uint64(),
	)

	signed, err := crypto.Sign(header.Hash().Bytes(), s.privateKey)
	if err != nil {
		return nil, common.Hash{}, err
	}

	if err := s.db.Put(
		db.BuildBlockKey(header.Time),
		db.BuildBlockValue(header.Hash().Bytes(),
			signed,
			blockID,
		),
	); err != nil {
		return nil, common.Hash{}, err
	}

	return signed, header.Hash(), nil
}

func (s *GuardianProverBlockSender) Close() error {
	return s.db.Close()
}

func (s *GuardianProverBlockSender) SendHeartbeat(ctx context.Context) error {
	sig, err := crypto.Sign(crypto.Keccak256Hash([]byte("HEART_BEAT")).Bytes(), s.privateKey)
	if err != nil {
		return err
	}

	req := &healthCheckReq{
		HeartBeatSignature: sig,
		ProverAddress:      s.proverAddress.Hex(),
	}

	if err := s.post(ctx, "healthCheck", req); err != nil {
		return err
	}

	log.Info("Successfully sent heartbeat", "signature", common.Bytes2Hex(sig))

	return nil
}
