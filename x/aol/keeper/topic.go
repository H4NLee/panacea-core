package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

// SetTopic set a specific topic in the store
func (k Keeper) SetTopic(ctx sdk.Context, key types.TopicCompositeKey, topic types.Topic) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopicKey))
	b := k.cdc.MustMarshalBinaryBare(&topic)
	store.Set([]byte(key.Marshal()), b)
}

// GetTopic returns a topic from its id
func (k Keeper) GetTopic(ctx sdk.Context, key types.TopicCompositeKey) types.Topic {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopicKey))
	var topic types.Topic
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(key.Marshal())), &topic)
	return topic
}

// HasTopic checks if the topic exists in the store
func (k Keeper) HasTopic(ctx sdk.Context, key types.TopicCompositeKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopicKey))
	return store.Has([]byte(key.Marshal()))
}

// GetAllTopicNames returns all topic names created by ownerAddress
func (k Keeper) GetAllTopicNames(ctx sdk.Context, ownerAddress string) (names []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopicKey))

	prefix := types.TopicCompositeKey{OwnerAddress: ownerAddress, TopicName: ""}
	iterator := sdk.KVStorePrefixIterator(store, []byte(prefix.Marshal()))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var key types.TopicCompositeKey
		types.MustUnmarshalCompositeKey(&key, string(iterator.Key()))
		names = append(names, key.TopicName)
	}

	return
}
