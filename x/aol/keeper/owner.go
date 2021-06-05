package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

// SetOwner set a specific owner in the store
func (k Keeper) SetOwner(ctx sdk.Context, key types.OwnerCompositeKey, owner types.Owner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKey))
	b := k.cdc.MustMarshalBinaryBare(&owner)
	store.Set([]byte(key.Marshal()), b)
}

// GetOwner returns a owner from its id
func (k Keeper) GetOwner(ctx sdk.Context, key types.OwnerCompositeKey) types.Owner {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKey))
	var owner types.Owner
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(key.Marshal())), &owner)
	return owner
}

// HasOwner checks if the owner exists in the store
func (k Keeper) HasOwner(ctx sdk.Context, key types.OwnerCompositeKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKey))
	return store.Has([]byte(key.Marshal()))
}

func (k Keeper) GetAllOwnerAddrs(ctx sdk.Context) (addrs []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var key types.OwnerCompositeKey
		types.MustUnmarshalCompositeKey(&key, string(iterator.Key()))
		addrs = append(addrs, key.OwnerAddress)
	}

	return
}
