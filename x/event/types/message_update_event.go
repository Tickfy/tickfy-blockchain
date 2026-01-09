package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateEvent{}

func NewMsgUpdateEvent(creator string, index string, name string, description string, location string, imageUrl string, creatorFee uint64) *MsgUpdateEvent {
	return &MsgUpdateEvent{
		Creator:     creator,
		Index:       index,
		Name:        name,
		Description: description,
		Location:    location,
		ImageUrl:    imageUrl,
		CreatorFee:  creatorFee,
	}
}

func (msg *MsgUpdateEvent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
