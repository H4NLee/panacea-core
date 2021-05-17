package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const DefaultCodespace sdk.CodespaceType = ModuleName

const (
	CodeCreateFailed      sdk.CodeType = 101
	CodeAccountExists     sdk.CodeType = 102
	CodeInstantiateFailed sdk.CodeType = 103
	CodeExecuteFailed     sdk.CodeType = 104
	CodeGasLimit          sdk.CodeType = 105
	CodeInvalidGenesis    sdk.CodeType = 106
	CodeNotFound          sdk.CodeType = 107
	CodeQueryFailed       sdk.CodeType = 108
	CodeInvalidMsg        sdk.CodeType = 109
	CodeMigrationFailed   sdk.CodeType = 110
	CodeEmpty             sdk.CodeType = 111
	CodeLimit             sdk.CodeType = 112
	CodeInvalid           sdk.CodeType = 113
	CodeDuplicate         sdk.CodeType = 114
)

func ErrCreateFailed(err error) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCreateFailed, "create wasm contract failed: %v", err)
}

func ErrAccountExists() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountExists, "contract account already exists")
}

func ErrInstantiateFailed() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInstantiateFailed, "instantiate wasm contract failed")
}

func ErrExecuteFailed() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeExecuteFailed, "execute wasm contract failed")
}

func ErrGasLimit() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeGasLimit, "insufficient gas")
}

func ErrInvalidGenesis() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidGenesis, "invalid genesis")
}

func ErrNotFound(name string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeNotFound, "%v is not found", name)
}

func ErrQueryFailed() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeNotFound, "query wasm contract failed")
}

func ErrInvalidMsg() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidMsg, "invalid CosmosMsg from the contract")
}

func ErrMigrationFailed() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeMigrationFailed, "migrate wasm contract failed")
}

func ErrEmpty(name string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeEmpty, "%v is empty", name)
}

func ErrLimit(err error) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeLimit, "exceeds limit: %v", err)
}

func ErrInvalid(err error) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalid, "invalid: %v", err)
}

func ErrDuplicate(err error) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeDuplicate, "duplicate: %v", err)
}
