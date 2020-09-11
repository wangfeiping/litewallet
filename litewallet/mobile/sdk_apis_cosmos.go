// Gomobile-based cosmos litewallet interface implementation.
package litewallet

import (
	"encoding/json"

	"github.com/QOSGroup/litewallet/litewallet/chains/cosmos"
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

//recover key
func CosmosRecoverKey(rootDir, name, password, seed string) string {
	output := ""
	return output
}

//update password
func CosmosUpdateKey(rootDir, name, oldpass, newpass string) string {
	output := ""
	return output
}

//get account info
func CosmosGetAccount(rootDir, node, chainID, addr string) string {
	output := ""
	return output
}

//transfer
func CosmosTransfer(rootDir, node, chainId,
	fromName, password, toStr,
	coinStr, feeStr, broadcastMode string) string {
	output := ""
	return output
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

//get all the validators
func CosmosGetAllValidators(rootDir, node, chainID string) string {
	return cosmos.CosmosGetAllValidators(rootDir, node, chainID)
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
