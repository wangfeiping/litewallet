package cosmos

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetAccount query account info from remote
func GetAccount(ctx client.Context,
	accAddr sdk.AccAddress) (acc authtypes.BaseAccount, err error) {
	queryClient := authtypes.NewQueryClient(ctx)
	res, err := queryClient.Account(context.Background(),
		&authtypes.QueryAccountRequest{Address: accAddr})
	if err != nil {
		return
	}

	out, err := ctx.JSONMarshaler.MarshalJSON(res.Account)
	if err != nil {
		return
	}
	err = ctx.JSONMarshaler.UnmarshalJSON(out, &acc)
	if err != nil && strings.Index(err.Error(), "unknown field") < 0 {
		return acc, err
	}
	return acc, nil
}

// GetBalances query account balances from remote
func GetBalances(ctx client.Context,
	accAddr sdk.AccAddress) (sdk.Coins, error) {
	bankQueryClient := banktypes.NewQueryClient(ctx)
	pageReq := &query.PageRequest{
		Key:        []byte(""),
		Offset:     0,
		Limit:      100,
		CountTotal: false}
	params := types.NewQueryAllBalancesRequest(accAddr, pageReq)
	res, err := bankQueryClient.AllBalances(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return res.Balances, nil
}

// GetAllValidators returns all the validators
func GetAllValidators(ctx client.Context) (
	stakingtypes.Validators, error) {
	resKVs, _, err := ctx.QuerySubspace(
		stakingtypes.ValidatorsKey, stakingtypes.StoreKey)
	if err != nil {
		return nil, err
	}

	var all stakingtypes.Validators
	for _, kv := range resKVs {
		validator, err := stakingtypes.UnmarshalValidator(
			stakingtypes.ModuleCdc, kv.Value)
		if err != nil {
			return nil, err
		}
		all = append(all, validator)
	}

	return all, nil
}
