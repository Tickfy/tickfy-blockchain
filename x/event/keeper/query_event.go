package keeper

import (
	"context"

	"tickfy-blockchain/x/event/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EventAll(ctx context.Context, req *types.QueryAllEventRequest) (*types.QueryAllEventResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var events []types.Event

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	eventStore := prefix.NewStore(store, types.KeyPrefix(types.EventKeyPrefix))

	pageRes, err := query.Paginate(eventStore, req.Pagination, func(key []byte, value []byte) error {
		var event types.Event
		if err := k.cdc.Unmarshal(value, &event); err != nil {
			return err
		}

		events = append(events, event)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEventResponse{Event: events, Pagination: pageRes}, nil
}

func (k Keeper) Event(ctx context.Context, req *types.QueryGetEventRequest) (*types.QueryGetEventResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetEvent(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEventResponse{Event: val}, nil
}
