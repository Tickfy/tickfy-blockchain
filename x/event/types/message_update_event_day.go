package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateEventDay{}

func NewMsgUpdateEventDay(creator string, index string, eventIndex string, name string, startTime string, endTime string) *MsgUpdateEventDay {
	return &MsgUpdateEventDay{
		Creator:    creator,
		Index:      index,
		EventIndex: eventIndex,
		Name:       name,
		StartTime:  startTime,
		EndTime:    endTime,
	}
}

func (msg *MsgUpdateEventDay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
