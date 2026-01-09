package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTicket{}

func NewMsgCreateTicket(creator string, index string, eventIndex string, eventDayIndex string, owner string, price uint64, metadata string) *MsgCreateTicket {
	return &MsgCreateTicket{
		Creator:       creator,
		Index:         index,
		EventIndex:    eventIndex,
		EventDayIndex: eventDayIndex,
		Owner:         owner,
		Price:         price,
		Metadata:      metadata,
	}
}

func (msg *MsgCreateTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
