package types

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	FlagHome    = flags.FlagHome
	FlagNode    = flags.FlagNode
	FlagChainID = flags.FlagChainID
	FlagName    = flags.FlagName
	FlagBitSize = "bit-size"

	AccAddressFromBech32 = sdk.AccAddressFromBech32
	ParseCoins           = sdk.ParseCoins
	ParseGasSetting      = flags.ParseGasSetting
)

type (
	Msg = sdk.Msg
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

type ValidatorWithShare struct {
	Validator      stakingtypes.Validator `json:"validator,omitempty"`
	SelfBondShares string                 `json:"selfbond_shares,omitempty"`
}

type BankBalances struct {
	Address       string    `json:"address,omitempty"`
	PubKey        []byte    `json:"public_key,omitempty"`
	AccountNumber uint64    `json:"account_number,omitempty"`
	Sequence      uint64    `json:"sequence,omitempty"`
	Coins         sdk.Coins `json:"coins,omitempty"`
}

type AccountWithBalances struct {
	Type  string       `json:"type,omitempty"`
	Value BankBalances `json:"value,omitempty"`
}
