package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"runos_chain/x/configstore/types"
)

func (k msgServer) GetPort(goCtx context.Context, msg *types.MsgGetPort) (*types.MsgGetPortResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGetPortResponse{}, nil
}
