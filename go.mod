module github.com/QOSGroup/litewallet

go 1.14

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200921130040-27db2cf89772
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.9.20
	github.com/gogo/protobuf v1.3.3
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/mobile v0.0.0-20210220033013-bdb1ca9a1e08 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/wangfeiping/cosmos-sdk v0.42.5-0.20210418085920-5649229dfd9b
