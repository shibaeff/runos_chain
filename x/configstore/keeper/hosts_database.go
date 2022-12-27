package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"runos_chain/x/configstore/types"
)

// SetHostsDatabase set a specific hostsDatabase in the store from its index
func (k Keeper) SetHostsDatabase(ctx sdk.Context, hostsDatabase types.HostsDatabase) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostsDatabaseKeyPrefix))
	b := k.cdc.MustMarshal(&hostsDatabase)
	store.Set(types.HostsDatabaseKey(
		hostsDatabase.Dpid,
		hostsDatabase.Mac,
	), b)
}

// GetHostsDatabase returns a hostsDatabase from its index
func (k Keeper) GetHostsDatabase(
	ctx sdk.Context,
	dpid string,
	mac string,

) (val types.HostsDatabase, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostsDatabaseKeyPrefix))

	b := store.Get(types.HostsDatabaseKey(
		dpid,
		mac,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHostsDatabase removes a hostsDatabase from the store
func (k Keeper) RemoveHostsDatabase(
	ctx sdk.Context,
	dpid string,
	mac string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostsDatabaseKeyPrefix))
	store.Delete(types.HostsDatabaseKey(
		dpid,
		mac,
	))
}

// GetAllHostsDatabase returns all hostsDatabase
func (k Keeper) GetAllHostsDatabase(ctx sdk.Context) (list []types.HostsDatabase) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostsDatabaseKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostsDatabase
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
