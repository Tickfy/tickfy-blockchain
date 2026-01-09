package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	"tickfy-blockchain/x/event/keeper"
	"tickfy-blockchain/x/event/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEventDay(keeper keeper.Keeper, ctx context.Context, n int) []types.EventDay {
	items := make([]types.EventDay, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEventDay(ctx, items[i])
	}
	return items
}

func TestEventDayGet(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEventDay(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEventDay(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEventDayRemove(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEventDay(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEventDay(ctx,
			item.Index,
		)
		_, found := keeper.GetEventDay(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEventDayGetAll(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEventDay(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEventDay(ctx)),
	)
}
