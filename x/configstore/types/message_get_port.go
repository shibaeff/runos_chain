package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGetPort = "get_port"

var _ sdk.Msg = &MsgGetPort{}

func NewMsgGetPort(creator string, dpid uint64, mac string) *MsgGetPort {
	return &MsgGetPort{
		Creator: creator,
		Dpid:    dpid,
		Mac:     mac,
	}
}

func (msg *MsgGetPort) Route() string {
	return RouterKey
}

func (msg *MsgGetPort) Type() string {
	return TypeMsgGetPort
}

func (msg *MsgGetPort) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGetPort) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGetPort) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
