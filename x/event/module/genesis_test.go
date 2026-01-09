package event_test

import (
	"testing"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	event "tickfy-blockchain/x/event/module"
	"tickfy-blockchain/x/event/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EventList: []types.Event{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		EventDayList: []types.EventDay{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EventKeeper(t)
	event.InitGenesis(ctx, k, genesisState)
	got := event.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EventList, got.EventList)
	require.ElementsMatch(t, genesisState.EventDayList, got.EventDayList)
	// this line is used by starport scaffolding # genesis/test/assert
}
