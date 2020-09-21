package main

import (
	"fmt"

	litewallet "github.com/QOSGroup/litewallet/litewallet/mobile"
	"github.com/QOSGroup/litewallet/litewallet/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func buildQueryCMD() *cobra.Command {
	cmdQuery := &cobra.Command{
		Use:   "query",
		Short: "Query for ...",
	}

	cmdValidators := &cobra.Command{
		Use:   "validators",
		Short: "Get all validators",
		RunE:  getAllValidators,
	}

	cmdAccount := &cobra.Command{
		Use:   "account [address]",
		Short: "Get account",
		Args:  cobra.ExactArgs(1),
		RunE:  getAccount,
	}

	cmdQuery.AddCommand(cmdAccount, cmdValidators)

	return cmdQuery
}

func getAccount(cmd *cobra.Command, args []string) error {
	node := viper.GetString(types.FlagNode)
	chainID := viper.GetString(types.FlagChainID)
	fmt.Println("node: ", node)
	fmt.Println("chain id: ", chainID)
	addr := args[0]
	resp := litewallet.CosmosGetAccount("", node, chainID, addr)
	fmt.Println("response: ", resp)
	return nil
}

func getAllValidators(cmd *cobra.Command, _ []string) error {
	node := viper.GetString(types.FlagNode)
	chainID := viper.GetString(types.FlagChainID)
	fmt.Println("node: ", node)
	fmt.Println("chain id: ", chainID)
	resp := litewallet.CosmosGetAllValidators("", node, chainID)
	fmt.Println("response: ", resp)
	return nil
}
