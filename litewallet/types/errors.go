package types

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	//nolint
	ErrWrongPassword = sdkerrors.Register(sdkerrors.UndefinedCodespace, 1000, "wrong password")

	ErrKeyNotFound = sdkerrors.Register(sdkerrors.UndefinedCodespace, 1001, "key not found")

	// errors on account creation
	ErrKeyNameConflict = fmt.Errorf("acount with the name already exists")
	ErrMissingName     = fmt.Errorf("you have to specify a name for the locally stored account")
	ErrMissingPassword = fmt.Errorf("you have to specify a password for the locally stored account")
	ErrMissingSeed     = fmt.Errorf("you have to specify seed for key recover")
)
