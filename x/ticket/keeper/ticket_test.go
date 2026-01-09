package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	"tickfy-blockchain/x/ticket/keeper"
	"tickfy-blockchain/x/ticket/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTicket(keeper keeper.Keeper, ctx context.Context, n int) []types.Ticket {
	items := make([]types.Ticket, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTicket(ctx, items[i])
	}
	return items
}

func TestTicketGet(t *testing.T) {
	keeper, ctx := keepertest.TicketKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTicket(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicketRemove(t *testing.T) {
	keeper, ctx := keepertest.TicketKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicket(ctx,
			item.Index,
		)
		_, found := keeper.GetTicket(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTicketGetAll(t *testing.T) {
	keeper, ctx := keepertest.TicketKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicket(ctx)),
	)
}
