package configstore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "runos_chain/testutil/keeper"
	"runos_chain/testutil/nullify"
	"runos_chain/x/configstore"
	"runos_chain/x/configstore/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		HostsDatabaseList: []types.HostsDatabase{
			{
				Dpid: "0",
				Mac:  "0",
			},
			{
				Dpid: "1",
				Mac:  "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ConfigstoreKeeper(t)
	configstore.InitGenesis(ctx, *k, genesisState)
	got := configstore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HostsDatabaseList, got.HostsDatabaseList)
	// this line is used by starport scaffolding # genesis/test/assert
}
