package error

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// copy form https://github.com/ethereum/go-ethereum/blob/6428663faf50f8368cedf0297063154483cce72b/rpc/json.go#L51
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

func (err *JsonRPCError) ErrorCode() int {
	return err.Code
}

func (err *JsonRPCError) ErrorData() interface{} {
	return err.Data
}

func CheckProtocolRevertReason(errString string, err *JsonRPCError) bool {
	reason, ok := err.ErrorData().(string)
	if !ok {
		return false
	}
	expect := fmt.Sprintf("%#x", crypto.Keccak256([]byte(errString)))[:10]
	return expect == reason
}
