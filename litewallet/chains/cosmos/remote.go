package cosmos

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetAccount query account info from remote
func GetAccount(rootDir, node, chainID,
	addr string) (authtypes.BaseAccount, sdk.Coins, error) {
	var acc authtypes.BaseAccount
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return acc, nil, err
	}

	clientCtx := NewClientContext(rootDir, node, chainID)
	queryClient := authtypes.NewQueryClient(clientCtx)
	res, err := queryClient.Account(context.Background(),
		&authtypes.QueryAccountRequest{Address: accAddr})
	if err != nil {
		return acc, nil, err
	}

	out, err := clientCtx.JSONMarshaler.MarshalJSON(res.Account)
	// out, err := res.Account.MarshalAmino()
	if err != nil {
		return acc, nil, err
	}
	err = clientCtx.JSONMarshaler.UnmarshalJSON(out, &acc)
	if err != nil && strings.Index(err.Error(), "unknown field") < 0 {
		return acc, nil, err
	}
	bankQueryClient := banktypes.NewQueryClient(clientCtx)
	pageReq := &query.PageRequest{
		Key:        []byte(""),
		Offset:     0,
		Limit:      100,
		CountTotal: false}
	params := types.NewQueryAllBalancesRequest(accAddr, pageReq)
	resBalances, err := bankQueryClient.AllBalances(context.Background(), params)
	if err != nil {
		return acc, nil, err
	}
	coins := append(resBalances.Balances, sdk.Coin{})
	return acc, coins, nil
}

// GetAllValidators returns all the validators
func GetAllValidators(rootDir, node, chainID string) (
	stakingtypes.Validators, error) {
	clientCtx := NewClientContext(rootDir, node, chainID)
	resKVs, _, err := clientCtx.QuerySubspace(stakingtypes.ValidatorsKey, stakingtypes.StoreKey)
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
