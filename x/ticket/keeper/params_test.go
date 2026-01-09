package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/x/ticket/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TicketKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
