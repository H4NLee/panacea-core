package types

const (
	// module name
	ModuleName = "wasm"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used in router
	RouterKey = ModuleName

	// QuerierRoute is the query router key
	QuerierRoute = ModuleName
)

var (
	KeyLastCodeID = []byte("lastCodeId")
)
