package configstore

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"runos_chain/testutil/sample"
	configstoresimulation "runos_chain/x/configstore/simulation"
	"runos_chain/x/configstore/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = configstoresimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSetPort = "op_weight_msg_set_port"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetPort int = 100

	opWeightMsgCreateHostsDatabase = "op_weight_msg_hosts_database"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHostsDatabase int = 100

	opWeightMsgUpdateHostsDatabase = "op_weight_msg_hosts_database"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHostsDatabase int = 100

	opWeightMsgDeleteHostsDatabase = "op_weight_msg_hosts_database"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHostsDatabase int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	configstoreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		HostsDatabaseList: []types.HostsDatabase{
			{
				Creator: sample.AccAddress(),
				Dpid:    "0",
				Mac:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Dpid:    "1",
				Mac:     "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&configstoreGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetPort int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetPort, &weightMsgSetPort, nil,
		func(_ *rand.Rand) {
			weightMsgSetPort = defaultWeightMsgSetPort
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetPort,
		configstoresimulation.SimulateMsgSetPort(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHostsDatabase int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHostsDatabase, &weightMsgCreateHostsDatabase, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHostsDatabase = defaultWeightMsgCreateHostsDatabase
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHostsDatabase,
		configstoresimulation.SimulateMsgCreateHostsDatabase(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHostsDatabase int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHostsDatabase, &weightMsgUpdateHostsDatabase, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHostsDatabase = defaultWeightMsgUpdateHostsDatabase
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHostsDatabase,
		configstoresimulation.SimulateMsgUpdateHostsDatabase(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHostsDatabase int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHostsDatabase, &weightMsgDeleteHostsDatabase, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHostsDatabase = defaultWeightMsgDeleteHostsDatabase
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHostsDatabase,
		configstoresimulation.SimulateMsgDeleteHostsDatabase(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
