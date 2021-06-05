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

	var topic types.Topic
	ctx := sdk.UnwrapSDKContext(c)

	topicKey := types.TopicCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName}
	if !k.HasTopic(ctx, topicKey) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopicKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(topicKey.Marshal())), &topic)

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
		[]byte(types.TopicCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: ""}.Marshal()),
	}, []byte{})

	store := ctx.KVStore(k.storeKey)
	topicStore := prefix.NewStore(store, keyPrefix)

	pageRes, err := query.Paginate(topicStore, req.Pagination, func(key []byte, value []byte) error {
		var topicKey types.TopicCompositeKey
		if err := topicKey.Unmarshal(string(key)); err != nil {
			return err
		}

		topicNames = append(topicNames, topicKey.TopicName)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListTopicsResponse{TopicNames: topicNames, Pagination: pageRes}, nil
}
