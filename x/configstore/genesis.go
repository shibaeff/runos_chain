package configstore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"runos_chain/x/configstore/keeper"
	"runos_chain/x/configstore/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the hostsDatabase
	for _, elem := range genState.HostsDatabaseList {
		k.SetHostsDatabase(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.HostsDatabaseList = k.GetAllHostsDatabase(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
