package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1200, "black address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1201, "red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1202, "game cannot be parsed")
)
