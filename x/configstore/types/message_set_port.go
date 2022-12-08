package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetPort = "set_port"

var _ sdk.Msg = &MsgSetPort{}

func NewMsgSetPort(creator string, dpid uint64, mac string, inPort uint64) *MsgSetPort {
	return &MsgSetPort{
		Creator: creator,
		Dpid:    dpid,
		Mac:     mac,
		InPort:  inPort,
	}
}

func (msg *MsgSetPort) Route() string {
	return RouterKey
}

func (msg *MsgSetPort) Type() string {
	return TypeMsgSetPort
}

func (msg *MsgSetPort) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPort) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPort) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
