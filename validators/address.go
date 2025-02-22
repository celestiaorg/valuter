package validators

import (
	"strings"

	"github.com/celestiaorg/valuter/configs"
)

func IsConsAddr(address string) bool {

	return strings.HasPrefix(address, configs.Configs.Bech32Prefix.Consensus.Address)

}

func IsOprAddr(address string) bool {

	return strings.HasPrefix(address, configs.Configs.Bech32Prefix.Validator.Address)

}
