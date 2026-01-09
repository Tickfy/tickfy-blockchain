package keeper

import (
	"tickfy-blockchain/x/ticket/types"
)

var _ types.QueryServer = Keeper{}
