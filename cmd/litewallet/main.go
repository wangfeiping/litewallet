package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	litewallet "github.com/QOSGroup/litewallet/litewallet/mobile"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "litewallet",
		Short: "Command line interface for cosmos & ethereum wallet",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			viper.BindPFlags(cmd.Flags())
			return nil
		},
	}
	// rootCmd.PersistentFlags().String(types.FlagHome, "$HOME/.coscli/", "home dir")

	// cmdCreate := &cobra.Command{
	// 	Use:   "create",
	// 	Short: "Create an account",
	// 	RunE:  doCreate,
	// }
	// cmdCreate.Flags().Int16P(types.FlagBitSize, "b", 0, "bit size of the mnemonic")

	// cmdRecover := &cobra.Command{
	// 	Use:   "recover",
	// 	Short: "Recover an account from mnemonic",
	// 	RunE:  doRecover,
	// }
	// cmdRecover.Flags().StringP(types.FlagName, "n", "", "name")

	cmdValidators := &cobra.Command{
		Use:   "validators",
		Short: "Get all validators",
		RunE:  getAllValidators,
	}
	cmdValidators.Flags().StringP(flags.FlagNode, "n", "tcp://8.211.162.156:26657", "node address")
	cmdValidators.Flags().StringP(flags.FlagChainID, "c", "stargate-1", "chain id")

	cmdVersion := &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		RunE:  doVersion,
	}
	rootCmd.AddCommand(
		// cmdCreate, cmdRecover,
		cmdValidators, cmdVersion)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func getAllValidators(cmd *cobra.Command, _ []string) error {
	node := viper.GetString(flags.FlagNode)
	chainID := viper.GetString(flags.FlagChainID)
	fmt.Println("node: ", node)
	fmt.Println("chain id: ", chainID)
	resp := litewallet.CosmosGetAllValidators("", node, chainID)
	fmt.Println("response: ", resp)
	return nil
}

func doCreate(_ *cobra.Command, _ []string) error {

	return nil
}

func doRecover(cmd *cobra.Command, _ []string) error {

	return nil
}

func doVersion(_ *cobra.Command, _ []string) error {
	// version.ShowVersion()
	return nil
}

func showJSONString(js string) (err error) {
	var out bytes.Buffer
	if err = json.Indent(&out, []byte(js), "", "  "); err == nil {
		fmt.Println(out.String())
	}
	return
}
