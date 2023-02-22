package error

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// JsonRPCError is copied from
// https://github.com/ethereum/go-ethereum/blob/6428663faf50f8368cedf0297063154483cce72b/rpc/json.go#L51
type JsonRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *JsonRPCError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}

// GetRevertReasonHash parses solidity contract revert reason hash
// from ethereum rpc calls.
func GetRevertReasonHash(err error) (string, error) {
	bytes, err := json.Marshal(errors.Unwrap(err))
	if err != nil {
		return "", err
	}
	rErr := new(JsonRPCError)
	if err = json.Unmarshal(bytes, rErr); err != nil {
		return "", err
	}
	reasonHash, ok := rErr.Data.(string)
	if !ok {
		return "", fmt.Errorf("invalid revert error, %T", rErr.Data)
	}
	return reasonHash, nil
}

// CheckExpectRevertReason checks if the revert reason in solidity contracts matches the expectation.
func CheckExpectRevertReason(expect string, revertErr error) (bool, error) {
	reason, err := GetRevertReasonHash(revertErr)
	if err != nil {
		return false, err
	}
	hash := fmt.Sprintf("%#x", crypto.Keccak256([]byte(expect)))[:10]
	if hash != reason {
		fmt.Printf("expect=%s,reason=%s", hash, reason)
	}
	return hash == reason, nil
}
