package litewallet

import (
	"encoding/json"

	"github.com/QOSGroup/litewallet/litewallet/chains/cosmos"
	"github.com/QOSGroup/litewallet/litewallet/types"
)

//create the seed(mnemonic) for the account generation
func CreateSeed() string {
	mnem, err := cosmos.CreateSeed()
	// json output the seed
	var seed types.SeedOutput
	if err != nil {
		seed.Error = err.Error()
	} else {
		seed.Seed = mnem
	}
	bytes, _ := json.Marshal(seed)
	return string(bytes)
}

//WalletAddressCheck for different chains
func WalletAddressCheck(addr string) string {
	output := ""
	return output
}
