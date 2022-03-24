package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "datapool"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_datapool"

	// this line is used by starport scaffolding # ibc/keys/name
)

var (
	// KeyPrefixDataValidators defines key to store data validator
	KeyPrefixDataValidators = []byte{0x01}
)

func GetKeyPrefixDataValidator(dataValidatorAddr sdk.AccAddress) []byte {
	return append(KeyPrefixDataValidators, dataValidatorAddr.Bytes()...)
}
