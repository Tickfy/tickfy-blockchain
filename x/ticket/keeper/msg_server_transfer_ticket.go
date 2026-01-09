package keeper

import (
	"context"

	"tickfy-blockchain/x/ticket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferTicket(goCtx context.Context, msg *types.MsgTransferTicket) (*types.MsgTransferTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferTicketResponse{}, nil
}
