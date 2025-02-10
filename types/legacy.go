// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
