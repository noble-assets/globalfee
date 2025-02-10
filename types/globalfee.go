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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate is a utility that verifies the provided gas prices. It ensures that
// prices are not negative, and that denoms are both valid and not repeated.
func (raw *GasPrices) Validate() error {
	gasPrices := raw.Value

	switch gasPrices.Len() {
	case 0:
		return nil
	default:
		if err := ValidateGasPrice(gasPrices[0]); err != nil {
			return err
		}

		seen := make(map[string]bool)
		seen[gasPrices[0].Denom] = true

		for _, gasPrice := range gasPrices[1:] {
			if seen[gasPrice.Denom] {
				return fmt.Errorf("denom %s is repeated", gasPrice.Denom)
			}

			if err := ValidateGasPrice(gasPrice); err != nil {
				return err
			}

			seen[gasPrice.Denom] = true
		}

		return nil
	}
}

func ValidateGasPrice(gasPrice sdk.DecCoin) error {
	if err := sdk.ValidateDenom(gasPrice.Denom); err != nil {
		return err
	}

	if gasPrice.IsNegative() {
		return fmt.Errorf("gas price for %s is negative", gasPrice.Denom)
	}

	return nil
}
