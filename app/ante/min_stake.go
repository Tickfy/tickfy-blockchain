package ante

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// MinValidatorStake is the minimum stake required to be an active validator
// 200,000 TKFY = 200,000,000,000 ustake (with 6 decimals)
var MinValidatorStake = math.NewInt(200_000_000_000)

// MaxCommissionRate is the maximum commission rate allowed for validators (20%)
var MaxCommissionRate = math.LegacyNewDecWithPrec(20, 2) // 0.20 = 20%

// ValidatorStakeDecorator checks that validators meet stake and commission requirements
type ValidatorStakeDecorator struct{}

// NewMinStakeDecorator creates a new ValidatorStakeDecorator
func NewMinStakeDecorator() ValidatorStakeDecorator {
	return ValidatorStakeDecorator{}
}

// AnteHandle checks minimum stake and maximum commission for validator transactions
func (vsd ValidatorStakeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		switch m := msg.(type) {
		case *stakingtypes.MsgCreateValidator:
			// Check if the initial stake meets the minimum requirement
			if m.Value.Amount.LT(MinValidatorStake) {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"validator stake %s is below minimum required %s (200,000 TKFY)",
					m.Value.Amount.String(),
					MinValidatorStake.String(),
				)
			}
			// Check if commission rate is within allowed range
			if m.Commission.Rate.GT(MaxCommissionRate) {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"validator commission rate %s exceeds maximum allowed 20%%",
					m.Commission.Rate.String(),
				)
			}
			// Check if max commission rate is within allowed range
			if m.Commission.MaxRate.GT(MaxCommissionRate) {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"validator max commission rate %s exceeds maximum allowed 20%%",
					m.Commission.MaxRate.String(),
				)
			}

		case *stakingtypes.MsgEditValidator:
			// Check if new commission rate is within allowed range
			if m.CommissionRate != nil && m.CommissionRate.GT(MaxCommissionRate) {
				return ctx, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"validator commission rate %s exceeds maximum allowed 20%%",
					m.CommissionRate.String(),
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}
