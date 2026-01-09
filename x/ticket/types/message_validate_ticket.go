package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgValidateTicket{}

func NewMsgValidateTicket(creator string, index string) *MsgValidateTicket {
  return &MsgValidateTicket{
		Creator: creator,
    Index: index,
	}
}

func (msg *MsgValidateTicket) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

