package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/medibloc/panacea-core/x/aol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Record(c context.Context, req *types.QueryGetRecordRequest) (*types.QueryGetRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	recordKey := types.RecordCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName, Offset: req.Offset}
	if !k.HasRecord(ctx, recordKey) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	record := k.GetRecord(ctx, recordKey)
	return &types.QueryGetRecordResponse{Record: &record}, nil
}
