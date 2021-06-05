package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/aol module sentinel errors
var (
	ErrMessageTooLarge     = sdkerrors.Register(ModuleName, 1, "too large")
	ErrInvalidTopic        = sdkerrors.Register(ModuleName, 2, "not valid")
	ErrInvalidMoniker      = sdkerrors.Register(ModuleName, 3, "not valid")
	ErrTopicExists         = sdkerrors.Register(ModuleName, 4, "already exists")
	ErrWriterExists        = sdkerrors.Register(ModuleName, 5, "already exists")
	ErrTopicNotFound       = sdkerrors.Register(ModuleName, 6, "not found")
	ErrWriterNotFound      = sdkerrors.Register(ModuleName, 7, "not found")
	ErrWriterNotAuthorized = sdkerrors.Register(ModuleName, 8, "not authorized")
)
