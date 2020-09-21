// Gomobile-based cosmos litewallet interface implementation.
package litewallet

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/viper"

	"github.com/QOSGroup/litewallet/litewallet/chains/cosmos"
	"github.com/QOSGroup/litewallet/litewallet/types"
)

// CosmosCreateAccount returns the account info that created with name, password and mnemonic input
func CosmosCreateAccount(rootDir, name, passwd, seed string) string {
	acc, err := cosmos.CreateAccount(rootDir, name, passwd, seed)
	if err != nil {
		acc.Error = err.Error()
	}
	bytes, _ := json.Marshal(acc)
	return string(bytes)
}

// CosmosRecoverKey recover the account with name, password and mnemonic input
func CosmosRecoverKey(rootDir, name, passwd, seed string) string {
	acc, err := cosmos.CreateAccount(rootDir, name, passwd, seed)
	if err != nil {
		acc.Error = err.Error()
	}
	// acc.Seed = ""
	bytes, _ := json.Marshal(acc)
	return string(bytes)
}

//update password
func CosmosUpdateKey(rootDir, name, oldpass, newpass string) string {
	output := ""
	return output
}

// CosmosGetAccount query account info & balances from remote
func CosmosGetAccount(rootDir, node, chainID, addr string) string {
	accAddr, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return err.Error()
	}

	ctx, err := cosmos.NewClientContext(rootDir, node, chainID)
	if err != nil {
		return err.Error()
	}
	acc, err := cosmos.GetAccount(ctx, accAddr)
	if err != nil {
		return err.Error()
	}
	coins, err := cosmos.GetBalances(ctx, accAddr)
	if err != nil {
		return err.Error()
	}
	balances := types.BankBalances{
		Address:       acc.Address.String(),
		PubKey:        acc.PubKey,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
		Coins:         coins}
	ab := types.AccountWithBalances{Value: balances}
	output, err := json.Marshal(ab)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

//transfer
func CosmosTransfer(rootDir, node, chainID,
	from, passwd, toAddrStr,
	amountStr, feeStr, broadcastMode string) string {
	viper.Set(types.FlagHome, rootDir)
	// accAddr, err := types.AccAddressFromBech32(accountName)
	// if err != nil {
	// 	return err.Error()
	// }
	toAccAddr, err := types.AccAddressFromBech32(toAddrStr)
	if err != nil {
		return err.Error()
	}
	funds, err := types.ParseCoins(amountStr)
	if err != nil {
		return err.Error()
	}
	ctx, err := cosmos.NewClientContext(rootDir, node, chainID)
	if err != nil {
		return err.Error()
	}
	ctx = ctx.WithBroadcastMode(broadcastMode)
	cosmos.SetPasswd(passwd)
	// info, err := ctx.Keyring.Key("test")
	// if err != nil {
	// 	return err.Error()
	// }
	// fmt.Println("info: ", info.GetAddress().String())

	// ctx = ctx.WithFrom("test").WithFromAddress(accAddr).WithFromName("test")

	fromAddr, fromName, err := client.GetFromFields(ctx.Keyring, from, ctx.GenerateOnly)
	if err != nil {
		return err.Error()
	}
	ctx = ctx.WithFrom(from).WithFromAddress(fromAddr).WithFromName(fromName)
	fmt.Println("from: ", from, "; name: ", fromName, "; ", fromAddr.String())

	msg := banktypes.NewMsgSend(ctx.GetFromAddress(), toAccAddr, funds)
	if err := msg.ValidateBasic(); err != nil {
		return err.Error()
	}
	resp, err := cosmos.GenerateOrBroadcastTx(ctx, msg)
	if err != nil {
		return err.Error()
	}
	out, err := ctx.JSONMarshaler.MarshalJSON(resp)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

//delegate
func CosmosDelegate(rootDir, node, chainID,
	delegatorName, password, delegatorAddr, validatorAddr,
	delegationCoinStr, feeStr, broadcastMode string) string {
	output := ""
	return output
}

//get a specific delegation shares
func CosmosGetDelegationShares(rootDir, node, chainID,
	delegatorAddr, validatorAddr string) string {
	output := ""
	return output
}

//for unbond delegation shares from specific validator
func CosmosUnbondingDelegation(rootDir, node, chainID,
	delegatorName, password, delegatorAddr, validatorAddr,
	Ubdshares, feeStr, broadcastMode string) string {
	output := ""
	return output
}

//get all unbonding delegations from a specific delegator
func CosmosGetAllUnbondingDelegations(rootDir, node, chainID,
	delegatorAddr string) string {
	output := ""
	return output
}

//Get bonded validators
func CosmosGetBondValidators(rootDir, node, chainID,
	delegatorAddr string) string {
	output := ""
	return output
}

// CosmosGetAllValidators returns all the validators
func CosmosGetAllValidators(rootDir, node, chainID string) string {
	ctx, err := cosmos.NewClientContext(rootDir, node, chainID)
	if err != nil {
		return err.Error()
	}
	validators, err := cosmos.GetAllValidators(ctx)
	if err != nil {
		return err.Error()
	}
	var all []types.ValidatorWithShare
	for _, validator := range validators {
		vShare := types.ValidatorWithShare{Validator: validator}
		all = append(all, vShare)
	}

	out, err := json.Marshal(all)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

//get all delegations from the delegator
func CosmosGetAllDelegations(rootDir, node, chainID,
	delegatorAddr string) string {
	output := ""
	return output
}

//Withdraw rewards from a specific validator
func CosmosWithdrawDelegationReward(rootDir, node, chainID,
	delegatorName, password, delegatorAddr, validatorAddr,
	feeStr, broadcastMode string) string {
	output := ""
	return output
}

//get a delegation reward between delegator and validator
func CosmosGetDelegationRewards(rootDir, node, chainID,
	delegatorAddr, validatorAddr string) string {
	output := ""
	return output
}

//query the tx result by txHash generated via async broadcast
func CosmosQueryTx(rootDir, node, chainId, txHash string) string {
	output := ""
	return output
}

func CosmosGetValSelfBondShares(rootDir, node, chainID,
	validatorAddr string) string {
	output := ""
	return output
}

func CosmosGetDelegtorRewardsShares(rootDir, node, chainId,
	delegatorAddr string) string {
	output := ""
	return output
}

func CosmosWithdrawDelegatorAllRewards(rootDir, node, chainID,
	delegatorName, password, delegatorAddr,
	feeStr, broadcastMode string) string {
	output := ""
	return output
}

func CosmosQueryQueryTxsWithTags(rootDir, node, chainID,
	addr string, page, limit int) string {
	output := ""
	return output
}
