package ticket_test

import (
	"testing"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	ticket "tickfy-blockchain/x/ticket/module"
	"tickfy-blockchain/x/ticket/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TicketKeeper(t)
	ticket.InitGenesis(ctx, k, genesisState)
	got := ticket.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
