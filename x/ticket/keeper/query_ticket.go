package keeper

import (
	"context"

	"tickfy-blockchain/x/ticket/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TicketAll(ctx context.Context, req *types.QueryAllTicketRequest) (*types.QueryAllTicketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickets []types.Ticket

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ticketStore := prefix.NewStore(store, types.KeyPrefix(types.TicketKeyPrefix))

	pageRes, err := query.Paginate(ticketStore, req.Pagination, func(key []byte, value []byte) error {
		var ticket types.Ticket
		if err := k.cdc.Unmarshal(value, &ticket); err != nil {
			return err
		}

		tickets = append(tickets, ticket)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTicketResponse{Ticket: tickets, Pagination: pageRes}, nil
}

func (k Keeper) Ticket(ctx context.Context, req *types.QueryGetTicketRequest) (*types.QueryGetTicketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTicket(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTicketResponse{Ticket: val}, nil
}
