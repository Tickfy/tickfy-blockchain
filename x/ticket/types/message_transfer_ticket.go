package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferTicket{}

func NewMsgTransferTicket(creator string, index string, newOwner string, price uint64) *MsgTransferTicket {
	return &MsgTransferTicket{
		Creator:  creator,
		Index:    index,
		NewOwner: newOwner,
		Price:    price,
	}
}

func (msg *MsgTransferTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
