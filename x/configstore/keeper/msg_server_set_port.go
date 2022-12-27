package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"runos_chain/x/configstore/types"
)

func (k msgServer) SetPort(goCtx context.Context, msg *types.MsgSetPort) (*types.MsgSetPortResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetPortResponse{}, nil
}
