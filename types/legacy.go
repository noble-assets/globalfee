// Copyright 2024 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
