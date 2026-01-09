package keeper

import (
	"context"

	"tickfy-blockchain/x/ticket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTicket(goCtx context.Context, msg *types.MsgCreateTicket) (*types.MsgCreateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTicketResponse{}, nil
}
