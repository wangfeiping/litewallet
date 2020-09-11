package types

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var (
	FlagHome    = flags.FlagHome
	FlagNode    = flags.FlagNode
	FlagChainID = flags.FlagChainID
	FlagName    = flags.FlagName
	FlagBitSize = "bit-size"
)

type SeedOutput struct {
	Seed  string `json:"seed"`
	Error string `json:"error,omitempty"`
}

type KeyOutput struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Address      string `json:"address"`
	PubKey       string `json:"pub_key"`
	PrivKeyArmor string `json:"priv_key_armor,omitempty"`
	Seed         string `json:"seed,omitempty"`
	Denom        string `json:"denom"`
	Error        string `json:"error,omitempty"`
}
