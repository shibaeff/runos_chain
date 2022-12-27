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

func (k Keeper) HostsDatabaseAll(c context.Context, req *types.QueryAllHostsDatabaseRequest) (*types.QueryAllHostsDatabaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostsDatabases []types.HostsDatabase
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hostsDatabaseStore := prefix.NewStore(store, types.KeyPrefix(types.HostsDatabaseKeyPrefix))

	pageRes, err := query.Paginate(hostsDatabaseStore, req.Pagination, func(key []byte, value []byte) error {
		var hostsDatabase types.HostsDatabase
		if err := k.cdc.Unmarshal(value, &hostsDatabase); err != nil {
			return err
		}

		hostsDatabases = append(hostsDatabases, hostsDatabase)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHostsDatabaseResponse{HostsDatabase: hostsDatabases, Pagination: pageRes}, nil
}

func (k Keeper) HostsDatabase(c context.Context, req *types.QueryGetHostsDatabaseRequest) (*types.QueryGetHostsDatabaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetHostsDatabase(
		ctx,
		req.Dpid,
		req.Mac,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetHostsDatabaseResponse{HostsDatabase: val}, nil
}
