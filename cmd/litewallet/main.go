package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QOSGroup/litewallet/cmd/version"
	litewallet "github.com/QOSGroup/litewallet/litewallet/mobile"
	"github.com/QOSGroup/litewallet/litewallet/types"
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
	rootCmd.PersistentFlags().String(types.FlagHome, "$HOME/.coscli/", "home dir")
	rootCmd.PersistentFlags().StringP(types.FlagNode, "n", "tcp://8.211.162.156:26657", "node address")
	rootCmd.PersistentFlags().StringP(types.FlagChainID, "c", "stargate-3a", "chain id")

	cmdCreate := &cobra.Command{
		Use:   "create [name]",
		Short: "Create an account",
		Args:  cobra.ExactArgs(1),
		RunE:  doCreate,
	}
	cmdCreate.Flags().Int16P(types.FlagBitSize, "b", 0, "bit size of the mnemonic")

	cmdRecover := &cobra.Command{
		Use:   "recover [name]",
		Short: "Recover an account from mnemonic",
		Args:  cobra.ExactArgs(1),
		RunE:  doRecover,
	}

	cmdQuery := buildQueryCMD()
	cmdTx := buildTxCMD()

	cmdVersion := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v", "V"},
		Short:   "Show version info",
		RunE:    doVersion,
	}
	rootCmd.AddCommand(
		cmdCreate, cmdRecover, cmdQuery, cmdTx, cmdVersion)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func doCreate(cmd *cobra.Command, args []string) error {
	home := os.ExpandEnv(viper.GetString(types.FlagHome))
	viper.Set(types.FlagHome, home)

	name := args[0]
	fmt.Println("do create...name: ", name)
	buf := bufio.NewReader(cmd.InOrStdin())
	passwd, err := input.GetPassword("> password:\n", buf)
	if err != nil {
		return err
	}
	var seed types.SeedOutput
	ret := litewallet.CreateSeed()
	if err := json.Unmarshal([]byte(ret), &seed); err != nil {
		fmt.Println("error: ", err)
		return err
	}
	showJSONString(ret)
	ret = litewallet.CosmosCreateAccount(home, name, passwd, seed.Seed)
	fmt.Println("create account: ", ret)
	showJSONString(ret)
	return nil
}

func doRecover(cmd *cobra.Command, args []string) error {
	home := os.ExpandEnv(viper.GetString(types.FlagHome))
	viper.Set(types.FlagHome, home)
	name := args[0]
	fmt.Println("do recover...name: ", name)
	buf := bufio.NewReader(cmd.InOrStdin())
	passwd, err := input.GetPassword("> password:\n", buf)
	if err != nil {
		return err
	}
	mnem, err := input.GetString("mnemonic:", buf)
	if err != nil {
		return err
	}

	// var ko types.KeyOutput
	fmt.Println("--- cosmos")
	ret := litewallet.CosmosRecoverKey(home, name, passwd, mnem)
	showJSONString(ret)
	return nil
}

func doVersion(_ *cobra.Command, _ []string) error {
	version.ShowVersion()
	return nil
}

func showJSONString(js string) (err error) {
	var out bytes.Buffer
	if err = json.Indent(&out, []byte(js), "", "  "); err == nil {
		fmt.Println(out.String())
	}
	return
}
