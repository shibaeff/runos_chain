package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"runos_chain/x/configstore/types"
)

// SetConfig set a specific config in the store from its index
func (k Keeper) SetConfig(ctx sdk.Context, config types.Config) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfigKeyPrefix))
	b := k.cdc.MustMarshal(&config)
	store.Set(types.ConfigKey(
		config.Index,
	), b)
}

// GetConfig returns a config from its index
func (k Keeper) GetConfig(
	ctx sdk.Context,
	index string,

) (val types.Config, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfigKeyPrefix))

	b := store.Get(types.ConfigKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveConfig removes a config from the store
func (k Keeper) RemoveConfig(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfigKeyPrefix))
	store.Delete(types.ConfigKey(
		index,
	))
}

// GetAllConfig returns all config
func (k Keeper) GetAllConfig(ctx sdk.Context) (list []types.Config) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConfigKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Config
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
