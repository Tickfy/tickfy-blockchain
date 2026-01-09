package keeper

import (
	"context"

	"tickfy-blockchain/x/event/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEventDay(goCtx context.Context, msg *types.MsgCreateEventDay) (*types.MsgCreateEventDayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateEventDayResponse{}, nil
}
