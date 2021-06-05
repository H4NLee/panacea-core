package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	topicKey := types.TopicCompositeKey{OwnerAddress: msg.OwnerAddress, TopicName: msg.TopicName}
	if !k.HasTopic(ctx, topicKey) {
		return nil, sdkerrors.Wrapf(types.ErrTopicNotFound, "topic <%s, %s>", msg.OwnerAddress, msg.TopicName)
	}
	writerKey := types.WriterCompositeKey{OwnerAddress: msg.OwnerAddress, TopicName: msg.TopicName, WriterAddress: msg.WriterAddress}
	if !k.HasWriter(ctx, writerKey) {
		return nil, sdkerrors.Wrapf(types.ErrWriterNotAuthorized, "writer <%s, %s, %s>", msg.OwnerAddress, msg.TopicName, msg.WriterAddress)
	}

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
