package treasury_test

import (
	"testing"

	keepertest "tickfy-blockchain/testutil/keeper"
	"tickfy-blockchain/testutil/nullify"
	treasury "tickfy-blockchain/x/treasury/module"
	"tickfy-blockchain/x/treasury/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TreasuryKeeper(t)
	treasury.InitGenesis(ctx, k, genesisState)
	got := treasury.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
