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
	"fmt"

	"cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		GasPrices:      sdk.DecCoins{},
		BypassMessages: []string{},
	}
}

func (genesis *GenesisState) Validate(registry codectypes.InterfaceRegistry) error {
	gasPrices := GasPrices{Value: genesis.GasPrices}
	if err := gasPrices.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate gas prices")
	}

	for _, bypassMessage := range genesis.BypassMessages {
		resolved, err := registry.Resolve(bypassMessage)
		if err != nil || resolved == nil {
			return fmt.Errorf("invalid bypass message %s", bypassMessage)
		}
	}

	return nil
}
