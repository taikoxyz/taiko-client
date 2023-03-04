package encoding

import (
	"errors"
	"strings"
)

func TryParseCustomError(err error) error {
	for _, customError := range TaikoL1ABI.Errors {
		// Hardhat node custom errors:
		// VM Exception while processing transaction: reverted with an unrecognized custom error (return data: 0xb6d363fd)
		if strings.Contains(err.Error(), "reverted with an unrecognized custom error") {
			if strings.HasPrefix(customError.ID.Hex(), err.Error()[len(err.Error())-11:len(err.Error())-1]) {
				return errors.New(customError.Name)
			}
		}
	}

	return err
}
