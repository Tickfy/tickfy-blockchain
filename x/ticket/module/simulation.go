package ticket

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"tickfy-blockchain/testutil/sample"
	ticketsimulation "tickfy-blockchain/x/ticket/simulation"
	"tickfy-blockchain/x/ticket/types"
)

// avoid unused import issue
var (
	_ = ticketsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTicket = "op_weight_msg_create_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTicket int = 100

	opWeightMsgUpdateTicket = "op_weight_msg_update_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTicket int = 100

	opWeightMsgTransferTicket = "op_weight_msg_transfer_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferTicket int = 100

	opWeightMsgValidateTicket = "op_weight_msg_validate_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgValidateTicket int = 100

	opWeightMsgValidateTicket = "op_weight_msg_validate_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgValidateTicket int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ticketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ticketGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTicket, &weightMsgCreateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTicket = defaultWeightMsgCreateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTicket,
		ticketsimulation.SimulateMsgCreateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTicket, &weightMsgUpdateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTicket = defaultWeightMsgUpdateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTicket,
		ticketsimulation.SimulateMsgUpdateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferTicket, &weightMsgTransferTicket, nil,
		func(_ *rand.Rand) {
			weightMsgTransferTicket = defaultWeightMsgTransferTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferTicket,
		ticketsimulation.SimulateMsgTransferTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgValidateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgValidateTicket, &weightMsgValidateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgValidateTicket = defaultWeightMsgValidateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgValidateTicket,
		ticketsimulation.SimulateMsgValidateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgValidateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgValidateTicket, &weightMsgValidateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgValidateTicket = defaultWeightMsgValidateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgValidateTicket,
		ticketsimulation.SimulateMsgValidateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTicket,
			defaultWeightMsgCreateTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketsimulation.SimulateMsgCreateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTicket,
			defaultWeightMsgUpdateTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketsimulation.SimulateMsgUpdateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTransferTicket,
			defaultWeightMsgTransferTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketsimulation.SimulateMsgTransferTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
	opWeightMsgValidateTicket,
	defaultWeightMsgValidateTicket,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		ticketsimulation.SimulateMsgValidateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
simulation.NewWeightedProposalMsg(
	opWeightMsgValidateTicket,
	defaultWeightMsgValidateTicket,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		ticketsimulation.SimulateMsgValidateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
