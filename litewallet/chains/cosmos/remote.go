package cosmos

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type validatorWithShare struct {
	Validator stakingtypes.Validator `json:"validator,omitempty"`
}

//get all the validators
func CosmosGetAllValidators(rootDir, node, chainID string) string {
	clientCtx := NewClientContext(rootDir, node, chainID)
	resKVs, _, err := clientCtx.QuerySubspace(stakingtypes.ValidatorsKey, stakingtypes.StoreKey)
	if err != nil {
		return err.Error()
	}

	// var validators stakingtypes.Validators
	var all []validatorWithShare
	for _, kv := range resKVs {
		validator, err := stakingtypes.UnmarshalValidator(
			stakingtypes.ModuleCdc, kv.Value)
		if err != nil {
			return err.Error()
		}
		validatorWithShare := validatorWithShare{Validator: validator}
		all = append(all, validatorWithShare)
	}

	// return clientCtx.PrintOutputLegacy(validators)
	out, err := clientCtx.LegacyAmino.MarshalJSON(all)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
