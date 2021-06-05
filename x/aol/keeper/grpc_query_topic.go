package keeper

import (
	"bytes"
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/medibloc/panacea-core/x/aol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Topic(c context.Context, req *types.QueryGetTopicRequest) (*types.QueryGetTopicResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	topicKey := types.TopicCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName}
	if !k.HasTopic(ctx, topicKey) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	topic := k.GetTopic(ctx, topicKey)
	return &types.QueryGetTopicResponse{Topic: &topic}, nil
}

func (k Keeper) Topics(c context.Context, req *types.QueryListTopicsRequest) (*types.QueryListTopicsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var topicNames []string
	ctx := sdk.UnwrapSDKContext(c)

	keyPrefix := bytes.Join([][]byte{
		types.KeyPrefix(types.TopicKey),
		types.MarshalBytes(&types.TopicCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: ""}),
	}, []byte{})

	store := ctx.KVStore(k.storeKey)
	topicStore := prefix.NewStore(store, keyPrefix)

	pageRes, err := query.Paginate(topicStore, req.Pagination, func(key []byte, value []byte) error {
		topicName := string(key)
		topicNames = append(topicNames, topicName)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListTopicsResponse{TopicNames: topicNames, Pagination: pageRes}, nil
}
