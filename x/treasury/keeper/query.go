package keeper

import (
	"tickfy-blockchain/x/treasury/types"
)

var _ types.QueryServer = Keeper{}
