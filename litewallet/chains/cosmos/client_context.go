package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdkkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/viper"

	"github.com/QOSGroup/litewallet/litewallet/types"
)

const (
	name = "litewallet"

// 	signModeDirect    = "direct"
// 	signModeAminoJSON = "amino-json"
)

type blackBox struct {
	passwd string
}

func (b *blackBox) setPasswd(passwd string) {
	b.passwd = passwd
}
func (b *blackBox) getPasswd() string {
	return b.passwd
}

var box = blackBox{}

func SetPasswd(passwd string) {
	box.setPasswd(passwd)
}

func NewKeyring() (sdkkeyring.Keyring, error) {
	home := viper.GetString(types.FlagHome)

	return sdkkeyring.NewFileKeyring(name, "", home,
		func(_ string) (string, error) {
			fmt.Println("passwd: ", box.getPasswd())
			return box.getPasswd(), nil
		})
}

func NewClientContext(rootDir, node, chainID string) (client.Context, error) {
	home := viper.GetString(types.FlagHome)
	ctx := client.Context{}
	encodingConfig := simapp.MakeEncodingConfig()
	keybase, err := NewKeyring()
	if err != nil {
		return ctx, err
	}
	ctx = ctx.WithOutputFormat("json").
		WithHomeDir(home).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(keybase).
		WithNodeURI(node).
		WithChainID(chainID)
	return ctx, nil
}

// GenerateOrBroadcastTx will either generate
// and sign it and broadcast it returning an error upon failure.
//
// github.com/cosmos/cosmos-sdk/client/tx/tx.go
// func GenerateOrBroadcastTxCLI(...) error {
func GenerateOrBroadcastTx(ctx client.Context, msgs ...types.Msg) error {
	// github.com/cosmos/cosmos-sdk/client/tx/factory.go
	// func NewFactoryCLI(...) Factory {
	txf := newFactoryCLI(ctx)
	return tx.GenerateOrBroadcastTxWithFactory(ctx, txf, msgs...)
}

func newFactoryCLI(ctx client.Context) tx.Factory {
	// signModeStr := viper.GetString(flags.FlagSignMode)

	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	// switch signModeStr {
	// case signModeDirect:
	// 	signMode = signing.SignMode_SIGN_MODE_DIRECT
	// case signModeAminoJSON:
	// 	signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	// }

	// accNum := viper.GetUint64(flags.FlagAccountNumber) // offline mode only
	// accSeq := viper.GetUint64(flags.FlagSequence)      // offline mode only
	// memo := viper.GetString(flags.FlagMemo)
	memo := ""
	// timeoutHeight := viper.GetUint64(flags.FlagTimeoutHeight)
	// feesStr := viper.GetString(flags.FlagFees)
	feesStr := "" // 10uatom
	// gasAdj := viper.GetFloat64(flags.FlagGasAdjustment)
	gasAdj := float64(1)
	// gasPricesStr := viper.GetString(flags.FlagGasPrices)
	gasPricesStr := ""
	// gasStr := viper.GetString(flags.FlagGas)
	gasStr := ""
	gasSetting, _ := types.ParseGasSetting(gasStr)

	f := tx.Factory{}.WithTxConfig(ctx.TxConfig).
		WithAccountRetriever(ctx.AccountRetriever).
		WithKeybase(ctx.Keyring).
		WithChainID(ctx.ChainID).
		WithGas(gasSetting.Gas).
		WithSimulateAndExecute(gasSetting.Simulate).
		// WithAccountNumber(accNum).
		// WithSequence(accSeq).
		// WithTimeoutHeight(timeoutHeight).
		WithGasAdjustment(gasAdj).
		WithMemo(memo).
		WithSignMode(signMode).
		WithFees(feesStr).
		WithGasPrices(gasPricesStr)

	return f
}
