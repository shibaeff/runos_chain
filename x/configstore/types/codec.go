package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetPort{}, "configstore/SetPort", nil)
	cdc.RegisterConcrete(&MsgCreateHostsDatabase{}, "configstore/CreateHostsDatabase", nil)
	cdc.RegisterConcrete(&MsgUpdateHostsDatabase{}, "configstore/UpdateHostsDatabase", nil)
	cdc.RegisterConcrete(&MsgDeleteHostsDatabase{}, "configstore/DeleteHostsDatabase", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetPort{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHostsDatabase{},
		&MsgUpdateHostsDatabase{},
		&MsgDeleteHostsDatabase{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
