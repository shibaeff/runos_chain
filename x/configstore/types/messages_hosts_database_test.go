package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"runos_chain/testutil/sample"
)

func TestMsgCreateHostsDatabase_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHostsDatabase
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHostsDatabase{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHostsDatabase{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateHostsDatabase_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHostsDatabase
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHostsDatabase{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHostsDatabase{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteHostsDatabase_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHostsDatabase
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHostsDatabase{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHostsDatabase{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
