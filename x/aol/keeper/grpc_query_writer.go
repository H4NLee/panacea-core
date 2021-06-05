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

func (k Keeper) Writer(c context.Context, req *types.QueryGetWriterRequest) (*types.QueryGetWriterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	writerKey := types.WriterCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName, WriterAddress: req.WriterAddress}
	if !k.HasWriter(ctx, writerKey) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	writer := k.GetWriter(ctx, writerKey)
	return &types.QueryGetWriterResponse{Writer: &writer}, nil
}

func (k Keeper) Writers(c context.Context, req *types.QueryListWritersRequest) (*types.QueryListWritersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var writerAddresses []string
	ctx := sdk.UnwrapSDKContext(c)

	keyPrefix := bytes.Join([][]byte{
		types.KeyPrefix(types.WriterKey),
		[]byte(types.WriterCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName, WriterAddress: ""}.Marshal()),
	}, []byte{})

	store := ctx.KVStore(k.storeKey)
	writerStore := prefix.NewStore(store, keyPrefix)

	pageRes, err := query.Paginate(writerStore, req.Pagination, func(key []byte, value []byte) error {
		writerAddress := string(key)
		writerAddresses = append(writerAddresses, writerAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListWritersResponse{WriterAddresses: writerAddresses, Pagination: pageRes}, nil
}
