package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

// SetRecord set a specific record in the store
func (k Keeper) SetRecord(ctx sdk.Context, key types.RecordCompositeKey, record types.Record) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	b := k.cdc.MustMarshalBinaryBare(&record)
	store.Set([]byte(key.Marshal()), b)
}

// GetRecord returns a record from its id
func (k Keeper) GetRecord(ctx sdk.Context, key types.RecordCompositeKey) types.Record {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	var record types.Record
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(key.Marshal())), &record)
	return record
}

// HasRecord checks if the record exists in the store
func (k Keeper) HasRecord(ctx sdk.Context, key types.RecordCompositeKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	return store.Has([]byte(key.Marshal()))
}
