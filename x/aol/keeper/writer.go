package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

// SetWriter set a specific writer in the store
func (k Keeper) SetWriter(ctx sdk.Context, key types.WriterCompositeKey, writer types.Writer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WriterKey))
	b := k.cdc.MustMarshalBinaryBare(&writer)
	store.Set([]byte(key.Marshal()), b)
}

// GetWriter returns a writer from its id
func (k Keeper) GetWriter(ctx sdk.Context, key types.WriterCompositeKey) types.Writer {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WriterKey))
	var writer types.Writer
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(key.Marshal())), &writer)
	return writer
}

// HasWriter checks if the writer exists in the store
func (k Keeper) HasWriter(ctx sdk.Context, key types.WriterCompositeKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WriterKey))
	return store.Has([]byte(key.Marshal()))
}

// RemoveWriter removes a writer from the store
func (k Keeper) RemoveWriter(ctx sdk.Context, key types.WriterCompositeKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WriterKey))
	store.Delete([]byte(key.Marshal()))
}

// GetAllWriterAddrs returns all writer addresses of <ownerAddress, topicName>
func (k Keeper) GetAllWriterAddrs(ctx sdk.Context, ownerAddress, topicName string) (addrs []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WriterKey))

	prefix := types.WriterCompositeKey{OwnerAddress: ownerAddress, TopicName: topicName, WriterAddress: ""}
	iterator := sdk.KVStorePrefixIterator(store, []byte(prefix.Marshal()))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var key types.WriterCompositeKey
		types.MustUnmarshalCompositeKey(&key, string(iterator.Key()))
		addrs = append(addrs, key.WriterAddress)
	}

	return
}
