package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// LegacyParams defines the parameters of the GlobalFee v1 module.
type LegacyParams struct {
	GasPrices      sdk.DecCoins
	BypassMessages []string
}

// Deprecated: ParamKeyTable returns the key table of the GlobalFee v1 module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&LegacyParams{})
}

func (p *LegacyParams) ParamSetPairs() paramstypes.ParamSetPairs {
	// NOTE: It is assumed that the GlobalFee v1 parameters are valid.

	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(
			[]byte("MinimumGasPricesParam"), &p.GasPrices, func(_ interface{}) error {
				return nil
			},
		),
		paramstypes.NewParamSetPair(
			[]byte("BypassMinFeeMsgTypesParam"), &p.BypassMessages, func(_ interface{}) error {
				return nil
			},
		),
	}
}
