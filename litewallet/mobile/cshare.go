package litewallet

//create the seed(mnemonic) for the account generation
func CreateSeed() string {
	output := ""
	return output
}

//WalletAddressCheck for different chains
func WalletAddressCheck(addr string) string {
	output := ""
	return output
}

//create account
func CosmosCreateAccount(rootDir, name, password, seed string) string {
	output := ""
	return output
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

//for QOSCreateSignedTransfer
func QOSCreateSignedTransfer(addrto, coinstr, privkey, chainid string) string {
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
	output := ""
	return output
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

//QOS wallet part begin from here
func QOSAccountCreate(password string) string {
	output := ""
	return output
}

func QOSAccountCreateFromSeed(mncode string) string {
	output := ""
	return output
}

//for QSCKVStoreset
func QSCKVStoreSet(k, v, privkey, chain string) string {
	output := ""
	return output
}

//for QSCKVStoreGet
func QSCKVStoreGet(k string) string {
	output := ""
	return output
}

//for QSCQueryAccount
func QSCQueryAccount(addr string) string {
	output := ""
	return output
}

//for QOSQueryAccount
func QOSQueryAccount(addr string) string {
	output := ""
	return output
}

//for AccountRecovery
func QOSAccountRecover(mncode, password string) string {
	output := ""
	return output
}

//for IP input
func QOSSetBlockchainEntrance(sh, mh string) {

}

//for PubAddrRetrieval
func QOSPubAddrRetrieval(priv string) string {
	output := ""
	return output
}

//for QSCtransferSend
func QSCtransferSend(addrto, coinstr, privkey, chainid string) string {
	output := ""
	return output
}

//for QOSCommitResultCheck
func QOSCommitResultCheck(txhash, height string) string {
	output := ""
	return output
}

func QOSJQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) string {
	output := ""
	return output
}

func QOSAesEncrypt(key, plainText string) string {
	output := ""
	return output
}

func QOSAesDecrypt(key, cipherText string) string {
	output := ""
	return output
}

func QOSTransferRecordsQuery(chainid, addr, cointype, offset, limit string) string {
	output := ""
	return output
}

func CosmosTransferB4send(rootDir, node, chainID,
	fromName, password, toStr,
	coinStr, feeStr string) string {
	output := ""
	return output
}

func CosmosBroadcastTransferTx(rootDir, node, chainID,
	txString, broadcastMode string) string {
	output := ""
	return output
}

//for AdvertisersTrue
func QOSAdvertisersTrue(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := ""
	return output
}

//for AdvertisersFalse
func QOSAdvertisersFalse(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := ""
	return output
}

//for GetTx
func QOSGetTx(tx string) string {
	output := ""
	return output
}

func QOSGetBlance(addrs string) string {
	output := ""
	return output
}

func QOSGetBlanceByCointype(addrs, cointype string) string {
	return "0"
}

// acutionAd 竞拍广告
//articleHash            //广告位标识
//privatekey             //用户私钥
//coinsType              //竞拍币种
//coinAmount             //竞拍数量
//qscchainid             //chainid
func QOSAcutionAd(articleHash, privatekey,
	coinsType, coinAmount, qscchainid string) string {
	output := ""
	return output
}

//for Extract
func QOSExtract(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := ""
	return output
}

// 提交到联盟链上
func QOSBroadcastTransferTxToQSC(txstring, broadcastModes string) string {
	output := ""
	return output
}

func QOSCommHandler(funcName, privatekey, args, qscchainid string) string {
	output := ""
	return output
}

//From here, Eth wallet part start
func EthCreateAccount(rootDir, name, password, seed string) string {
	output := ""
	return output
}

func EthRecoverAccount(rootDir, name, password, seed string) string {
	output := ""
	return output
}

func EthGetAccount(node, addr string) string {
	output := ""
	return output
}

func EthGetErc20Account(node, addr, tokenAddr string) string {
	output := ""
	return output
}

func EthTransferETH(rootDir, node,
	name, password, toAddr,
	gasPrice, amount string, gasLimit int64) string {
	output := ""
	return output
}

func EthTransferErc20(rootDir, node,
	name, password, toAddr,
	tokenAddr, tokenValue, gasPrice string, gasLimit int64) string {
	output := ""
	return output
}

//Deprecated!
//func EthGetPendingNonceAt(rootDir, node, fromName, password string) int64 {
//	output := eth.GetPendingNonceAt(rootDir, node, fromName, password)
//	return output
//}

func EthSpeedTransferETH(rootDir, node,
	fromName, password, toAddr,
	gasPrice, amount string, GasLimit, pendingNonce int64) string {
	output := ""
	return output
}

func EthSpeedTransferERC20(rootDir, node,
	fromName, password, toAddr,
	tokenAddr, tokenValue, gasPrice string, GasLimit, pendingNonce int64) string {
	output := ""
	return output
}

//EthGetNonceAt provide the nonce at the latest block
func EthGetNonceAt(rootDir, node, fromName, password string) int64 {
	output := int64(0)
	return output
}
