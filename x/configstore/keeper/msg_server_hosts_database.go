package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"runos_chain/x/configstore/types"
)

func (k msgServer) CreateHostsDatabase(goCtx context.Context, msg *types.MsgCreateHostsDatabase) (*types.MsgCreateHostsDatabaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetHostsDatabase(
		ctx,
		msg.Dpid,
		msg.Mac,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var hostsDatabase = types.HostsDatabase{
		Creator: msg.Creator,
		Dpid:    msg.Dpid,
		Mac:     msg.Mac,
		Inport:  msg.Inport,
	}

	k.SetHostsDatabase(
		ctx,
		hostsDatabase,
	)
	return &types.MsgCreateHostsDatabaseResponse{}, nil
}

func (k msgServer) UpdateHostsDatabase(goCtx context.Context, msg *types.MsgUpdateHostsDatabase) (*types.MsgUpdateHostsDatabaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHostsDatabase(
		ctx,
		msg.Dpid,
		msg.Mac,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var hostsDatabase = types.HostsDatabase{
		Creator: msg.Creator,
		Dpid:    msg.Dpid,
		Mac:     msg.Mac,
		Inport:  msg.Inport,
	}

	k.SetHostsDatabase(ctx, hostsDatabase)

	return &types.MsgUpdateHostsDatabaseResponse{}, nil
}

func (k msgServer) DeleteHostsDatabase(goCtx context.Context, msg *types.MsgDeleteHostsDatabase) (*types.MsgDeleteHostsDatabaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHostsDatabase(
		ctx,
		msg.Dpid,
		msg.Mac,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHostsDatabase(
		ctx,
		msg.Dpid,
		msg.Mac,
	)

	return &types.MsgDeleteHostsDatabaseResponse{}, nil
}
