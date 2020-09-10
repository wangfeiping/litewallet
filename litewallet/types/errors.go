package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	//nolint
	ErrWrongPassword = sdkerrors.Register(sdkerrors.UndefinedCodespace, 1000, "wrong password")

	ErrKeyNotFound = sdkerrors.Register(sdkerrors.UndefinedCodespace, 1001, "key not found")
)
