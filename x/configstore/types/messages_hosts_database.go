package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateHostsDatabase = "create_hosts_database"
	TypeMsgUpdateHostsDatabase = "update_hosts_database"
	TypeMsgDeleteHostsDatabase = "delete_hosts_database"
)

var _ sdk.Msg = &MsgCreateHostsDatabase{}

func NewMsgCreateHostsDatabase(
	creator string,
	dpid string,
	mac string,
	inport string,

) *MsgCreateHostsDatabase {
	return &MsgCreateHostsDatabase{
		Creator: creator,
		Dpid:    dpid,
		Mac:     mac,
		Inport:  inport,
	}
}

func (msg *MsgCreateHostsDatabase) Route() string {
	return RouterKey
}

func (msg *MsgCreateHostsDatabase) Type() string {
	return TypeMsgCreateHostsDatabase
}

func (msg *MsgCreateHostsDatabase) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHostsDatabase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHostsDatabase) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHostsDatabase{}

func NewMsgUpdateHostsDatabase(
	creator string,
	dpid string,
	mac string,
	inport string,

) *MsgUpdateHostsDatabase {
	return &MsgUpdateHostsDatabase{
		Creator: creator,
		Dpid:    dpid,
		Mac:     mac,
		Inport:  inport,
	}
}

func (msg *MsgUpdateHostsDatabase) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHostsDatabase) Type() string {
	return TypeMsgUpdateHostsDatabase
}

func (msg *MsgUpdateHostsDatabase) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHostsDatabase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHostsDatabase) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHostsDatabase{}

func NewMsgDeleteHostsDatabase(
	creator string,
	dpid string,
	mac string,

) *MsgDeleteHostsDatabase {
	return &MsgDeleteHostsDatabase{
		Creator: creator,
		Dpid:    dpid,
		Mac:     mac,
	}
}
func (msg *MsgDeleteHostsDatabase) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHostsDatabase) Type() string {
	return TypeMsgDeleteHostsDatabase
}

func (msg *MsgDeleteHostsDatabase) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHostsDatabase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHostsDatabase) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
