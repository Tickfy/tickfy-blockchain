package keeper

import (
	"context"

	"tickfy-blockchain/x/ticket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateTicket(goCtx context.Context, msg *types.MsgUpdateTicket) (*types.MsgUpdateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateTicketResponse{}, nil
}
