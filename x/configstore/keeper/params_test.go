package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "runos_chain/testutil/keeper"
	"runos_chain/x/configstore/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ConfigstoreKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
