package main

import (
	"github.com/QOSGroup/litewallet/litewallet/types"
	"github.com/spf13/cobra"
)

func buildTxCMD() *cobra.Command {
	cmdTx := &cobra.Command{
		Use:   "tx",
		Short: "Create or submit a Tx for ...",
	}

	cmdTxSend := &cobra.Command{
		Use:   "send [from_address] [to_address] [amount]",
		Short: "Send funds from one account to another.",
		Args:  cobra.ExactArgs(3),
		RunE:  doTxSend,
	}

	cmdTx.AddCommand(cmdTxSend)
	cmdTx.PersistentFlags().StringP(types.FlagNode, "n", "tcp://8.211.162.156:26657", "node address")
	cmdTx.PersistentFlags().StringP(types.FlagChainID, "c", "stargate-1", "chain id")

	return cmdTx
}

func doTxSend(cmd *cobra.Command, args []string) error {
	return nil
}
