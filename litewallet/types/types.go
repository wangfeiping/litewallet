package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	FlagHome    = flags.FlagHome
	FlagNode    = flags.FlagNode
	FlagChainID = flags.FlagChainID
	FlagName    = flags.FlagName
	FlagBitSize = "bit-size"

	AccAddressFromBech32 = sdk.AccAddressFromBech32
	ParseCoins           = sdk.ParseCoinsNormalized
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

// https://github.com/cosmos/cosmos-sdk/blob/v0.34.4/x/staking/types/validator.go
// stakingtype.Validator不适用于移动端，因此自定义ValidatorWithShare 用于移动端
type ValidatorWithShare struct {
	OperatorAddress   string                                 `json:"operator_address,omitempty"`
	ConsensusPubkey   string                                 `json:"consensus_pubkey,omitempty"`
	Jailed            bool                                   `json:"jailed,omitempty"`
	Status            stakingtype.BondStatus                 `json:"status,omitempty"`
	Tokens            github_com_cosmos_cosmos_sdk_types.Int `json:"tokens"`
	DelegatorShares   github_com_cosmos_cosmos_sdk_types.Dec `json:"delegator_shares"`
	Description       stakingtype.Description                `json:"description"`
	UnbondingHeight   int64                                  `json:"unbonding_height,omitempty"`
	UnbondingTime     time.Time                              `json:"unbonding_time"`
	Commission        stakingtype.Commission                 `json:"commission"`
	MinSelfDelegation github_com_cosmos_cosmos_sdk_types.Int `json:"min_self_delegation"`
	SelfBondShares    string                                 `json:"selfbond_shares,omitempty"`
}

func (m *ValidatorWithShare) Reset()         { *m = ValidatorWithShare{} }
func (*ValidatorWithShare) ProtoMessage()    {}
func (m *ValidatorWithShare) String() string { return "" }

func (m *ValidatorWithShare) Set(v *stakingtype.Validator) {
	m.OperatorAddress = v.OperatorAddress
}

type ValidatorsWithShare []ValidatorWithShare

func (ValidatorsWithShare) Reset()         {}
func (ValidatorsWithShare) ProtoMessage()  {}
func (ValidatorsWithShare) String() string { return "" }

type Pubkey struct {
	Type  string `json:"type,omitempty"`
	Value []byte `json:"value,omitempty"`
}

type BankBalances struct {
	Address string `json:"address,omitempty"`
	// PubKey        Pubkey    `json:"public_key,omitempty"`
	AccountNumber uint64    `json:"account_number,omitempty"`
	Sequence      uint64    `json:"sequence,omitempty"`
	Coins         sdk.Coins `json:"coins,omitempty"`
}

type AccountWithBalances struct {
	Type  string       `json:"type,omitempty"`
	Value BankBalances `json:"value,omitempty"`
}
