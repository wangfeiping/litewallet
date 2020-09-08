package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewClientContext(rootDir, node, chainID string) client.Context {
	ctx := client.Context{}
	ctx.WithOutputFormat("json").
		WithNodeURI(node).
		WithChainID(chainID).
		WithLegacyAmino(codec.New())
	return ctx
}
