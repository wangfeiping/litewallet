# Changes in upgrades

## Packages has bean changed

- github.com/tendermint/tendermint/libs/common  
  some funs move to:  
  github.com/tendermint/tendermint/libs/os  

- github.com/tendermint/tendermint/libs/db  
  github.com/tendermint/tm-db  

- github.com/cosmos/cosmos-sdk/crypto/keys/keyerror  
  redefined:  
  github.com/cosmos/cosmos-sdk/types/errors/errors.go

- github.com/gogo/protobuf  
  forked:  
  github.com/regen-network/protobuf v1.3.2-alpha.regen.4

## Compilation errors

- constant 9223372036854775807 overflows int  
  !!!This error only occurs when compiling with gomobile!!!
  \# github.com/cosmos/iavl  
  ../pkg/mod/github.com/cosmos/iavl@v0.15.0-rc3/nodedb.go:241:65: constant 9223372036854775807 overflows int  
  ../pkg/mod/github.com/cosmos/iavl@v0.15.0-rc3/repair.go:44:69: constant 9223372036854775807 overflows int  
