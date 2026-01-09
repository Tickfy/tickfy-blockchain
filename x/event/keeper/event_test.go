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

func createNEvent(keeper keeper.Keeper, ctx context.Context, n int) []types.Event {
	items := make([]types.Event, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEvent(ctx, items[i])
	}
	return items
}

func TestEventGet(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEvent(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEvent(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEventRemove(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEvent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEvent(ctx,
			item.Index,
		)
		_, found := keeper.GetEvent(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEventGetAll(t *testing.T) {
	keeper, ctx := keepertest.EventKeeper(t)
	items := createNEvent(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEvent(ctx)),
	)
}
