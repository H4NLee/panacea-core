package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
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

	var record types.Record
	ctx := sdk.UnwrapSDKContext(c)

	recordKey := types.RecordCompositeKey{OwnerAddress: req.OwnerAddress, TopicName: req.TopicName, Offset: req.Offset}
	if !k.HasRecord(ctx, recordKey) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(recordKey.Marshal())), &record)

	return &types.QueryGetRecordResponse{Record: &record}, nil
}
