package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/input"
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

	return cmdTx
}

func doTxSend(cmd *cobra.Command, args []string) error {
	home := os.ExpandEnv(viper.GetString(types.FlagHome))

	node := viper.GetString(types.FlagNode)
	chainID := viper.GetString(types.FlagChainID)
	from := args[0]
	to := args[1]
	amount := args[2]
	buf := bufio.NewReader(cmd.InOrStdin())
	passwd, err := input.GetPassword("> password:\n", buf)
	if err != nil {
		return err
	}
	fee := ""
	broadcastMode := "block" // "sync" // "async"
	fmt.Println("from: ", from, "; to: ", to, "; amount: ", amount)
	resp := litewallet.CosmosTransfer(home, node, chainID,
		from, passwd, to,
		amount, fee, broadcastMode)
	fmt.Println("response: ", resp)
	return nil
}
