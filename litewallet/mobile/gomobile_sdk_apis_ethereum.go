package litewallet

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
