package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"runos_chain/testutil/network"
	"runos_chain/x/configstore/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateHostsDatabase(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	for _, tc := range []struct {
		desc   string
		idDpid string
		idMac  string

		args []string
		err  error
		code uint32
	}{
		{
			idDpid: strconv.Itoa(0),
			idMac:  strconv.Itoa(0),

			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idDpid,
				tc.idMac,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateHostsDatabase(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateHostsDatabase(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{
		"0",
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateHostsDatabase(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc   string
		idDpid string
		idMac  string

		args []string
		code uint32
		err  error
	}{
		{
			desc:   "valid",
			idDpid: strconv.Itoa(0),
			idMac:  strconv.Itoa(0),

			args: common,
		},
		{
			desc:   "key not found",
			idDpid: strconv.Itoa(100000),
			idMac:  strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idDpid,
				tc.idMac,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateHostsDatabase(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestDeleteHostsDatabase(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{
		"0",
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateHostsDatabase(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc   string
		idDpid string
		idMac  string

		args []string
		code uint32
		err  error
	}{
		{
			desc:   "valid",
			idDpid: strconv.Itoa(0),
			idMac:  strconv.Itoa(0),

			args: common,
		},
		{
			desc:   "key not found",
			idDpid: strconv.Itoa(100000),
			idMac:  strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idDpid,
				tc.idMac,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteHostsDatabase(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
