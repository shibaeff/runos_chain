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

func createNConfig(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Config {
	items := make([]types.Config, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetConfig(ctx, items[i])
	}
	return items
}

func TestConfigGet(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNConfig(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetConfig(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNConfig(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveConfig(ctx,
			item.Index,
		)
		_, found := keeper.GetConfig(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestConfigGetAll(t *testing.T) {
	keeper, ctx := keepertest.ConfigstoreKeeper(t)
	items := createNConfig(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllConfig(ctx)),
	)
}
