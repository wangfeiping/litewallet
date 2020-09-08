// Gomobile-based ethereum litewallet interface implementation.
package litewallet

import "github.com/QOSGroup/litewallet/litewallet/eth"

func EthCreateAccount(rootDir, name, password, seed string) string {
	output := eth.CreateAccount(rootDir, name, password, seed)
	return output
}

func EthRecoverAccount(rootDir, name, password, seed string) string {
	output := eth.RecoverAccount(rootDir, name, password, seed)
	return output
}

func EthGetAccount(node, addr string) string {
	output := eth.GetAccount(node, addr)
	return output
}

func EthGetErc20Account(node, addr, tokenAddr string) string {
	output := eth.GetAccountERC20(node, addr, tokenAddr)
	return output
}

func EthTransferETH(rootDir, node, name, password, toAddr, gasPrice, amount string, gasLimit int64) string {
	output := eth.TransferETH(rootDir, node, name, password, toAddr, gasPrice, amount, gasLimit)
	return output
}

func EthTransferErc20(rootDir, node, name, password, toAddr, tokenAddr, tokenValue, gasPrice string, gasLimit int64) string {
	output := eth.TransferERC20(rootDir, node, name, password, toAddr, tokenAddr, tokenValue, gasPrice, gasLimit)
	return output
}

//Deprecated!
//func EthGetPendingNonceAt(rootDir, node, fromName, password string) int64 {
//	output := eth.GetPendingNonceAt(rootDir, node, fromName, password)
//	return output
//}

func EthSpeedTransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount string, GasLimit, pendingNonce int64) string {
	output := eth.SpeedTransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount, GasLimit, pendingNonce)
	return output
}

func EthSpeedTransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice string, GasLimit, pendingNonce int64) string {
	output := eth.SpeedTransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice, GasLimit, pendingNonce)
	return output
}

//EthGetNonceAt provide the nonce at the latest block
func EthGetNonceAt(rootDir, node, fromName, password string) int64 {
	output := eth.GetNonceAt(rootDir, node, fromName, password)
	return output
}
