package keeper

import (
	"context"

	"tickfy-blockchain/x/event/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetEvent set a specific event in the store from its index
func (k Keeper) SetEvent(ctx context.Context, event types.Event) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventKeyPrefix))
	b := k.cdc.MustMarshal(&event)
	store.Set(types.EventKey(
		event.Index,
	), b)
}

// GetEvent returns a event from its index
func (k Keeper) GetEvent(
	ctx context.Context,
	index string,

) (val types.Event, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventKeyPrefix))

	b := store.Get(types.EventKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEvent removes a event from the store
func (k Keeper) RemoveEvent(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventKeyPrefix))
	store.Delete(types.EventKey(
		index,
	))
}

// GetAllEvent returns all event
func (k Keeper) GetAllEvent(ctx context.Context) (list []types.Event) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Event
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
