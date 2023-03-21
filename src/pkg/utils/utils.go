package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func HexToString(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

func CompareValidatorSet(validatorSet1, validatorSet2 []common.Address) bool {
	if len(validatorSet1) != len(validatorSet2) {
		return false
	}

	// make a deep copy of validatorSet2 because we will change it when check
	validatorSet2Copy := make([]common.Address, len(validatorSet2))
	copy(validatorSet2Copy, validatorSet2)
CHECK_LOOP:
	for _, validator1 := range validatorSet1 {
		i := 0
		for i < len(validatorSet2Copy) {
			if validatorSet2Copy[i] == validator1 {
				validatorSet2Copy = append(validatorSet2Copy[:i], validatorSet2Copy[i+1:]...)
				continue CHECK_LOOP
			}
			i++
		}
		return false
	}
	return true
}

func ParseValidatorSet(header *types.Header) ([]common.Address, error) {
	extraData := header.Extra
	extraSuffix := len(extraData) - 65
	validatorsBytes := extraData[32:extraSuffix]

	if len(validatorsBytes)%20 != 0 {
		return []common.Address{}, fmt.Errorf("GetValidatorSetFromParliaBlock: mismatching validator list")
	}

	var validatorSet []common.Address
	totalValidators := len(validatorsBytes) / 20
	for i := 0; i < totalValidators; i++ {
		validator := validatorsBytes[i*20 : (i+1)*20]
		validatorSet = append(validatorSet, common.BytesToAddress(validator))
	}
	return validatorSet, nil
}
