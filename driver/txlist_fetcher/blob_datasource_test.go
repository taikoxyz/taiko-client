package txlistdecoder

import (
	"context"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/testutils"
)

type BlobDataSourceTestSuite struct {
	testutils.ClientTestSuite
	ds *BlobDataSource
}

func (s *BlobDataSourceTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	// Init BlobDataSource
	blobServerEndpoint, err := url.Parse("https://blob.hekla.taiko.xyz")
	s.Nil(err)
	s.ds = NewBlobDataSource(context.Background(), s.RPCClient, blobServerEndpoint)
}

func (s *BlobDataSourceTestSuite) TestGetBlobs() {
	meta := &bindings.TaikoDataBlockMetadata{
		BlobUsed:  true,
		BlobHash:  common.HexToHash("0x019101fb28118ceccaabca22a47e35b9c3f12eb2dcb25e5c543d5b75e6cd841f"),
		Timestamp: 1,
	}
	sidecars, err := s.ds.GetBlobs(context.Background(), meta)
	s.Nil(err)
	s.Greater(len(sidecars), 0)
}

func TestBlobDataSourceTestSuite(t *testing.T) {
	suite.Run(t, new(BlobDataSourceTestSuite))
}
