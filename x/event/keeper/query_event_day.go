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

func (k Keeper) EventDayAll(ctx context.Context, req *types.QueryAllEventDayRequest) (*types.QueryAllEventDayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var eventDays []types.EventDay

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	eventDayStore := prefix.NewStore(store, types.KeyPrefix(types.EventDayKeyPrefix))

	pageRes, err := query.Paginate(eventDayStore, req.Pagination, func(key []byte, value []byte) error {
		var eventDay types.EventDay
		if err := k.cdc.Unmarshal(value, &eventDay); err != nil {
			return err
		}

		eventDays = append(eventDays, eventDay)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEventDayResponse{EventDay: eventDays, Pagination: pageRes}, nil
}

func (k Keeper) EventDay(ctx context.Context, req *types.QueryGetEventDayRequest) (*types.QueryGetEventDayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetEventDay(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEventDayResponse{EventDay: val}, nil
}
