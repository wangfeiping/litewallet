package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	litewallet "github.com/QOSGroup/litewallet/litewallet/mobile"
	"github.com/QOSGroup/litewallet/litewallet/types"
)

func buildTxCMD() *cobra.Command {
	cmdTx := &cobra.Command{
		Use:   "tx",
		Short: "Create or submit a Tx for ...",
	}

	cmdTxSend := &cobra.Command{
		Use:   "send [from_address] [to_address] [amount]",
		Short: "Send funds from one account to another.(Tx build -> sign -> send)",
		Args:  cobra.ExactArgs(3),
		RunE:  doTxSend,
	}

	cmdTx.AddCommand(cmdTxSend)
	cmdTx.PersistentFlags().StringP(types.FlagNode, "n", "tcp://8.211.162.156:26657", "node address")
	cmdTx.PersistentFlags().StringP(types.FlagChainID, "c", "stargate-1", "chain id")

	return cmdTx
}

func doTxSend(cmd *cobra.Command, args []string) error {
	node := viper.GetString(types.FlagNode)
	chainID := viper.GetString(types.FlagChainID)
	from := args[0]
	to := args[1]
	amount := args[2]
	passwd := ""
	fee := ""
	broadcastMode := ""
	fmt.Println("from: ", from, "; to: ", to, "; amount: ", amount)
	resp := litewallet.CosmosTransfer("", node, chainID, from, passwd, to,
		amount, fee, broadcastMode)
	fmt.Println("response: ", resp)
	return nil
}
