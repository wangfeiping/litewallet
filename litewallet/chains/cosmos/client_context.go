package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdktx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/viper"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

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
	pw := b.passwd
	b.passwd = ""
	return pw
}

var box = blackBox{}

func SetPasswd(passwd string) {
	box.setPasswd(passwd)
}

func NewKeyring() (sdkkeyring.Keyring, error) {
	home := viper.GetString(types.FlagHome)

	return sdkkeyring.NewFileKeyring(name, "", home,
		func(_ string) (string, error) {
			return box.getPasswd(), nil
		})
}

// initEncodingConfig creates an EncodingConfig for an amino based test configuration.
func initEncodingConfig() params.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	return params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// reference:
// https://github.com/cosmos/cosmos-sdk/blob/v0.42.4/simapp/encoding.go
//
// makeEncodingConfig creates an EncodingConfig for mobile-sdk
func makeEncodingConfig() params.EncodingConfig {
	encodingConfig := initEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	stakingtypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	stakingtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	// ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

func NewClientContext(rootDir, node, chainID string) (client.Context, error) {
	ctx := client.Context{}
	encodingConfig := makeEncodingConfig()
	// encodingConfig := simapp.MakeTestEncodingConfig()
	keybase, err := NewKeyring()
	if err != nil {
		return ctx, err
	}

	ctx = ctx.WithSkipConfirmation(true).
		WithOutputFormat("json").
		WithHomeDir(rootDir).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(keybase).
		WithNodeURI(node).
		WithChainID(chainID)

	client, err := rpchttp.New(node, "/websocket")
	if err != nil {
		return ctx, err
	}

	return ctx.WithClient(client), nil
}

// GenerateOrBroadcastTx will either generate
// and sign it and broadcast it returning an error upon failure.
//
// reference:
// github.com/cosmos/cosmos-sdk/client/tx/tx.go
// func GenerateOrBroadcastTxCLI(...) error {
func GenerateOrBroadcastTx(
	ctx client.Context, msgs ...types.Msg,
) (*sdk.TxResponse, error) {
	txf := newFactoryCLI(ctx)
	return GenerateOrBroadcastTxWithFactory(ctx, txf, msgs...)
}

// reference:
// github.com/cosmos/cosmos-sdk/client/tx/factory.go
// func NewFactoryCLI(...) Factory {
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

// GenerateOrBroadcastTxWithFactory will either generate and print and unsigned transaction
// or sign it and broadcast it returning an error upon failure.
// reference:
// github.com/cosmos/cosmos-sdk/client/tx/tx.go
// func GenerateOrBroadcastTxWithFactory(...) error {
func GenerateOrBroadcastTxWithFactory(
	ctx client.Context, txf tx.Factory, msgs ...types.Msg,
) (*sdk.TxResponse, error) {
	// if clientCtx.GenerateOnly {
	// 	return GenerateTx(clientCtx, txf, msgs...)
	// }

	return BroadcastTx(ctx, txf, msgs...)
}

// BroadcastTx attempts to generate, sign and broadcast a transaction with the
// given set of messages. It will also simulate gas requirements if necessary.
// It will return an error upon failure.
//
//
// func BroadcastTx(...) error {
func BroadcastTx(
	ctx client.Context, txf tx.Factory, msgs ...types.Msg,
) (res *sdk.TxResponse, err error) {
	txf, err = sdktx.PrepareFactory(ctx, txf)
	if err != nil {
		return
	}

	// if txf.SimulateAndExecute() || ctx.Simulate {
	// 	_, adjusted, err := sdktx.CalculateGas(ctx.QueryWithData, txf, msgs...)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	txf = txf.WithGas(adjusted)
	// 	_, _ = fmt.Fprintf(os.Stderr, "%s\n", sdktx.GasEstimateResponse{GasEstimate: txf.Gas()})
	// }

	// if ctx.Simulate {
	// 	return nil
	// }

	tx, err := sdktx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return
	}

	// if !ctx.SkipConfirm {
	// 	out, err := ctx.TxConfig.TxJSONEncoder()(tx.GetTx())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

	// 	buf := bufio.NewReader(os.Stdin)
	// 	ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

	// 	if err != nil || !ok {
	// 		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
	// 		return err
	// 	}
	// }

	err = sdktx.Sign(txf, ctx.GetFromName(), tx, true)
	if err != nil {
		return
	}

	txBytes, err := ctx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return
	}

	// broadcast to a Tendermint node
	return ctx.BroadcastTx(txBytes)
}
