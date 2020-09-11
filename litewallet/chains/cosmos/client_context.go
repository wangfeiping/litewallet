package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewClientContext(rootDir, node, chainID string) client.Context {
	ctx := client.Context{}
	ctx = ctx.WithOutputFormat("json").
		WithNodeURI(node).
		WithChainID(chainID).
		WithLegacyAmino(codec.NewLegacyAmino())
	return ctx
}
