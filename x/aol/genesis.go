package aol

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/keeper"
	"github.com/medibloc/panacea-core/x/aol/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for keyStr, owner := range genState.Owners {
		var key types.OwnerCompositeKey
		types.MustUnmarshalCompositeKey(&key, keyStr)
		k.SetOwner(ctx, key, *owner)
	}

	for keyStr, topic := range genState.Topics {
		var key types.TopicCompositeKey
		types.MustUnmarshalCompositeKey(&key, keyStr)
		k.SetTopic(ctx, key, *topic)
	}

	for keyStr, writer := range genState.Writers {
		var key types.WriterCompositeKey
		types.MustUnmarshalCompositeKey(&key, keyStr)
		k.SetWriter(ctx, key, *writer)
	}

	for keyStr, record := range genState.Records {
		var key types.RecordCompositeKey
		types.MustUnmarshalCompositeKey(&key, keyStr)
		k.SetRecord(ctx, key, *record)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	for _, ownerAddr := range k.GetAllOwnerAddrs(ctx) {
		ownerKey := types.OwnerCompositeKey{
			OwnerAddress: ownerAddr,
		}
		owner := k.GetOwner(ctx, ownerKey)
		genesis.Owners[ownerKey.Marshal()] = &owner

		for _, topicName := range k.GetAllTopicNames(ctx, ownerAddr) {
			topicKey := types.TopicCompositeKey{
				OwnerAddress: ownerAddr,
				TopicName:    topicName,
			}
			topic := k.GetTopic(ctx, topicKey)
			genesis.Topics[topicKey.Marshal()] = &topic

			for _, writerAddr := range k.GetAllWriterAddrs(ctx, ownerAddr, topicName) {
				writerKey := types.WriterCompositeKey{
					OwnerAddress:  ownerAddr,
					TopicName:     topicName,
					WriterAddress: writerAddr,
				}
				writer := k.GetWriter(ctx, writerKey)
				genesis.Writers[writerKey.Marshal()] = &writer
			}

			var offset uint64
			for offset = 0; offset < topic.TotalRecords; offset++ {
				recordKey := types.RecordCompositeKey{
					OwnerAddress: ownerAddr,
					TopicName:    topicName,
					Offset:       offset,
				}
				record := k.GetRecord(ctx, recordKey)
				genesis.Records[recordKey.Marshal()] = &record
			}
		}

	}

	return genesis
}
