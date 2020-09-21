package litewallet

import (
	"encoding/json"
	"fmt"
	"regexp"

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

//WalletAddressCheck to differentiate the addresses from various wallets, e.g. cosmos, ETH, qos, .etc
func WalletAddressCheck(addr string) string {
	//split the address with prefix, e.g. "0x", "cosmos", "address" for ETH, cosmos, qos respectively
	//regexp for ETH address
	ethre := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	//regexp for cosmos address
	cosre := regexp.MustCompile("^cosmos1[0-9a-z]{38}$")

	//regexp for qos address
	qosre := regexp.MustCompile("address1[0-9a-z]{38}$")

	switch {
	case ethre.MatchString(addr):
		return "ETH"

	case cosre.MatchString(addr):
		return "COSMOS"

	case qosre.MatchString(addr):
		return "QOS"

	default:
		return fmt.Sprintf("None")
	}
}
