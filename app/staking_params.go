package app

import (
	"cosmossdk.io/math"
)

// Tickfy Staking Parameters
var (
	// MinValidatorStake is the minimum stake required to be an active validator
	// 200,000 TKFY = 200,000,000,000 ustake (with 6 decimals)
	MinValidatorStake = math.NewInt(200_000_000_000)

	// MaxValidators is the maximum number of validators in the active set
	MaxValidators = uint32(250)

	// MinCommissionRate is the minimum commission rate for validators (5%)
	MinCommissionRate = math.LegacyNewDecWithPrec(5, 2) // 0.05 = 5%
)
