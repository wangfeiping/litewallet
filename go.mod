module github.com/QOSGroup/litewallet

go 1.14

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200829041113-200e88ba075b
	github.com/ethereum/go-ethereum v1.9.20
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.0-rc3
	github.com/tendermint/tm-db v0.6.2
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.34.4-0.20200829041113-200e88ba075b

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
