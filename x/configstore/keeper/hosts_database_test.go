package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "runos_chain/testutil/keeper"
	"runos_chain/testutil/nullify"
	"runos_chain/x/configstore/keeper"
	"runos_chain/x/configstore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNHostsDatabase(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HostsDatabase {
	items := make([]types.HostsDatabase, n)
	for i := range items {
		items[i].Dpid = strconv.Itoa(i)
		items[i].Mac = strconv.Itoa(i)

		keeper.SetHostsDatabase(ctx, items[i])
	}
	return items
}

func TestHostsDatabaseGet(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNHostsDatabase(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetHostsDatabase(ctx,
			item.Dpid,
			item.Mac,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestHostsDatabaseRemove(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNHostsDatabase(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHostsDatabase(ctx,
			item.Dpid,
			item.Mac,
		)
		_, found := keeper.GetHostsDatabase(ctx,
			item.Dpid,
			item.Mac,
		)
		require.False(t, found)
	}
}

func TestHostsDatabaseGetAll(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNHostsDatabase(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHostsDatabase(ctx)),
	)
}
