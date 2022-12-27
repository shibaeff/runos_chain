package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "runos_chain/testutil/keeper"
	"runos_chain/testutil/nullify"
	"runos_chain/x/configstore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHostsDatabaseQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHostsDatabase(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetHostsDatabaseRequest
		response *types.QueryGetHostsDatabaseResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetHostsDatabaseRequest{
				Dpid: msgs[0].Dpid,
				Mac:  msgs[0].Mac,
			},
			response: &types.QueryGetHostsDatabaseResponse{HostsDatabase: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetHostsDatabaseRequest{
				Dpid: msgs[1].Dpid,
				Mac:  msgs[1].Mac,
			},
			response: &types.QueryGetHostsDatabaseResponse{HostsDatabase: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetHostsDatabaseRequest{
				Dpid: strconv.Itoa(100000),
				Mac:  strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.HostsDatabase(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestHostsDatabaseQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHostsDatabase(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHostsDatabaseRequest {
		return &types.QueryAllHostsDatabaseRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HostsDatabaseAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HostsDatabase), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HostsDatabase),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HostsDatabaseAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HostsDatabase), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HostsDatabase),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HostsDatabaseAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HostsDatabase),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HostsDatabaseAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
