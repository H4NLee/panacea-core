package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/x/aol/types"
)

func (k msgServer) CreateTopic(goCtx context.Context, msg *types.MsgCreateTopic) (*types.MsgCreateTopicResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	topicKey := types.TopicCompositeKey{OwnerAddress: msg.OwnerAddress, TopicName: msg.TopicName}
	if k.HasTopic(ctx, topicKey) {
		return nil, sdkerrors.Wrapf(types.ErrTopicExists, "topic <%s, %s>", msg.OwnerAddress, msg.TopicName)
	}

	ownerKey := types.OwnerCompositeKey{OwnerAddress: msg.OwnerAddress}
	owner := k.GetOwner(ctx, ownerKey).IncreaseTotalTopics()
	k.SetOwner(ctx, ownerKey, owner)

	topic := types.Topic{Description: msg.Description}
	k.SetTopic(ctx, topicKey, topic)

	return &types.MsgCreateTopicResponse{}, nil
}
