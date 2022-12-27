package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"runos_chain/x/configstore/keeper"
	"runos_chain/x/configstore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateHostsDatabase(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateHostsDatabase{
			Creator: simAccount.Address.String(),
			Dpid:    strconv.Itoa(i),
			Mac:     strconv.Itoa(i),
		}

		_, found := k.GetHostsDatabase(ctx, msg.Dpid, msg.Mac)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "HostsDatabase already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateHostsDatabase(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			hostsDatabase    = types.HostsDatabase{}
			msg              = &types.MsgUpdateHostsDatabase{}
			allHostsDatabase = k.GetAllHostsDatabase(ctx)
			found            = false
		)
		for _, obj := range allHostsDatabase {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				hostsDatabase = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "hostsDatabase creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Dpid = hostsDatabase.Dpid
		msg.Mac = hostsDatabase.Mac

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteHostsDatabase(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			hostsDatabase    = types.HostsDatabase{}
			msg              = &types.MsgUpdateHostsDatabase{}
			allHostsDatabase = k.GetAllHostsDatabase(ctx)
			found            = false
		)
		for _, obj := range allHostsDatabase {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				hostsDatabase = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "hostsDatabase creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Dpid = hostsDatabase.Dpid
		msg.Mac = hostsDatabase.Mac

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
