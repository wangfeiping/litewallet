package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	signModeDirect    = "direct"
	signModeAminoJSON = "amino-json"
)

func NewClientContext(rootDir, node, chainID string) client.Context {
	encodingConfig := simapp.MakeEncodingConfig()
	ctx := client.Context{}
	ctx = ctx.WithOutputFormat("json").
		// WithHomeDir("/tmp/").
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(keyring.NewInMemory()).
		WithNodeURI(node).
		WithChainID(chainID)
	return ctx
}

// GenerateOrBroadcastTx will either generate
// and sign it and broadcast it returning an error upon failure.
//
// github.com/cosmos/cosmos-sdk/client/tx/tx.go
// func GenerateOrBroadcastTxCLI(...) error {
func GenerateOrBroadcastTx(ctx client.Context, msgs ...sdk.Msg) error {
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
	gasSetting, _ := flags.ParseGasSetting(gasStr)

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
