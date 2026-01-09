package keeper

import (
	"context"

	"tickfy-blockchain/x/event/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetEventDay set a specific eventDay in the store from its index
func (k Keeper) SetEventDay(ctx context.Context, eventDay types.EventDay) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventDayKeyPrefix))
	b := k.cdc.MustMarshal(&eventDay)
	store.Set(types.EventDayKey(
		eventDay.Index,
	), b)
}

// GetEventDay returns a eventDay from its index
func (k Keeper) GetEventDay(
	ctx context.Context,
	index string,

) (val types.EventDay, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventDayKeyPrefix))

	b := store.Get(types.EventDayKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEventDay removes a eventDay from the store
func (k Keeper) RemoveEventDay(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventDayKeyPrefix))
	store.Delete(types.EventDayKey(
		index,
	))
}

// GetAllEventDay returns all eventDay
func (k Keeper) GetAllEventDay(ctx context.Context) (list []types.EventDay) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EventDayKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EventDay
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
