package simulation

import (
	"math/rand"

	"tickfy-blockchain/x/ticket/keeper"
	"tickfy-blockchain/x/ticket/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgValidateTicket(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgValidateTicket{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ValidateTicket simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ValidateTicket simulation not implemented"), nil, nil
	}
}
