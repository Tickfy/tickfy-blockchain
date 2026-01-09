package keeper

import (
	"tickfy-blockchain/x/event/types"
)

var _ types.QueryServer = Keeper{}
