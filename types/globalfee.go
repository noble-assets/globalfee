// Copyright 2025 NASD Inc. All Rights Reserved.
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
