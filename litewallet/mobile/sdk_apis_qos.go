package litewallet

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

//for QOSCreateSignedTransfer
func QOSCreateSignedTransfer(addrto, coinstr, privkey, chainid string) string {
	output := ""
	return output
}
