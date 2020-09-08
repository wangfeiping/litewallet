package cosmos

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

//get all the validators
func CosmosGetAllValidators(rootDir, node, chainID string) string {
	clientCtx := NewClientContext(rootDir, node, chainID)
	resKVs, _, err := clientCtx.QuerySubspace(stakingtypes.ValidatorsKey, stakingtypes.StoreKey)
	if err != nil {
		return err
	}

	var validators stakingtypes.Validators
	for _, kv := range resKVs {
		validator, err := stakingtypes.UnmarshalValidator(
			stakingtypes.ModuleCdc, kv.Value)
		if err != nil {
			return err
		}

		validators = append(validators, validator)
	}

	return clientCtx.PrintOutputLegacy(validators)
}
