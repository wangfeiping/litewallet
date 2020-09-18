package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func NewClientContext(rootDir, node, chainID string) client.Context {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	ctx := client.Context{}
	ctx = ctx.WithOutputFormat("json").
		WithNodeURI(node).
		WithChainID(chainID).
		WithLegacyAmino(codec.NewLegacyAmino()).
		WithJSONMarshaler(marshaler)
	return ctx
}
