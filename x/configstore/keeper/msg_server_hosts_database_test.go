package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "runos_chain/testutil/keeper"
	"runos_chain/x/configstore/keeper"
	"runos_chain/x/configstore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHostsDatabaseMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ConfigstoreKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateHostsDatabase{Creator: creator,
			Dpid: strconv.Itoa(i),
			Mac:  strconv.Itoa(i),
		}
		_, err := srv.CreateHostsDatabase(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetHostsDatabase(ctx,
			expected.Dpid,
			expected.Mac,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestHostsDatabaseMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateHostsDatabase
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateHostsDatabase{Creator: "B",
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(100000),
				Mac:  strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ConfigstoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			}
			_, err := srv.CreateHostsDatabase(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateHostsDatabase(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetHostsDatabase(ctx,
					expected.Dpid,
					expected.Mac,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestHostsDatabaseMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteHostsDatabase
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteHostsDatabase{Creator: "B",
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(100000),
				Mac:  strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ConfigstoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateHostsDatabase(wctx, &types.MsgCreateHostsDatabase{Creator: creator,
				Dpid: strconv.Itoa(0),
				Mac:  strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteHostsDatabase(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetHostsDatabase(ctx,
					tc.request.Dpid,
					tc.request.Mac,
				)
				require.False(t, found)
			}
		})
	}
}
