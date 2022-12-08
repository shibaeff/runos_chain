package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runos_chain/x/configstore/types"
)

func (k Keeper) ConfigAll(c context.Context, req *types.QueryAllConfigRequest) (*types.QueryAllConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var configs []types.Config
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	configStore := prefix.NewStore(store, types.KeyPrefix(types.ConfigKeyPrefix))

	pageRes, err := query.Paginate(configStore, req.Pagination, func(key []byte, value []byte) error {
		var config types.Config
		if err := k.cdc.Unmarshal(value, &config); err != nil {
			return err
		}

		configs = append(configs, config)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllConfigResponse{Config: configs, Pagination: pageRes}, nil
}

func (k Keeper) Config(c context.Context, req *types.QueryGetConfigRequest) (*types.QueryGetConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetConfig(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetConfigResponse{Config: val}, nil
}
