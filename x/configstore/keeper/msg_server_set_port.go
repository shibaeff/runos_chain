package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"runos_chain/x/configstore/types"
)

func (k msgServer) SetPort(goCtx context.Context, msg *types.MsgSetPort) (*types.MsgSetPortResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	switchConfig, isFound := k.GetHostsDatabase(ctx, msg.Dpid, msg.Mac)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "Key already exists")
	}
	k.SetHostsDatabase(ctx, switchConfig)
	return &types.MsgSetPortResponse{}, nil
}
