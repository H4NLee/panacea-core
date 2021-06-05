package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	topicKey := types.TopicCompositeKey{OwnerAddress: msg.OwnerAddress, TopicName: msg.TopicName}
	topic := k.GetTopic(ctx, topicKey)
	offset := topic.NextRecordOffset()
	k.SetTopic(ctx, topicKey, topic.IncreaseTotalRecords())

	recordKey := types.RecordCompositeKey{OwnerAddress: msg.OwnerAddress, TopicName: msg.TopicName, Offset: offset}
	record := types.Record{
		Key:           msg.Key,
		Value:         msg.Value,
		NanoTimestamp: ctx.BlockTime().UnixNano(),
		WriterAddress: msg.WriterAddress,
	}
	k.SetRecord(ctx, recordKey, record)

	return &types.MsgAddRecordResponse{
		OwnerAddress: msg.OwnerAddress,
		TopicName:    msg.TopicName,
		Offset:       offset,
	}, nil
}
