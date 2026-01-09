package event

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"tickfy-blockchain/x/event/keeper"
	"tickfy-blockchain/x/event/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the event
	for _, elem := range genState.EventList {
		k.SetEvent(ctx, elem)
	}
	// Set all the eventDay
	for _, elem := range genState.EventDayList {
		k.SetEventDay(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EventList = k.GetAllEvent(ctx)
	genesis.EventDayList = k.GetAllEventDay(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
