package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/proofofexistence module sentinel errors
var (
	ErrInvalidHash = sdkerrors.Register(ModuleName, 2, "invalid hash")
	ErrHashExists  = sdkerrors.Register(ModuleName, 3, "hash already registered")
)
