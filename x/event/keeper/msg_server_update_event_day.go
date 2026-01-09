package keeper

import (
	"context"

	"tickfy-blockchain/x/event/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateEventDay(goCtx context.Context, msg *types.MsgUpdateEventDay) (*types.MsgUpdateEventDayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateEventDayResponse{}, nil
}
