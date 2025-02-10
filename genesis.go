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

package globalfee

import (
	"context"

	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
)

func InitGenesis(ctx context.Context, k *keeper.Keeper, genesis types.GenesisState) {
	gasPrices := types.GasPrices{Value: genesis.GasPrices}
	if err := k.GasPrices.Set(ctx, gasPrices); err != nil {
		panic(err)
	}

	for _, bypassMessage := range genesis.BypassMessages {
		if err := k.BypassMessages.Set(ctx, bypassMessage); err != nil {
			panic(err)
		}
	}
}

func ExportGenesis(ctx context.Context, k *keeper.Keeper) *types.GenesisState {
	gasPrices, err := k.GasPrices.Get(ctx)
	if err != nil {
		panic(err)
	}

	bypassMessages, err := k.GetBypassMessages(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		GasPrices:      gasPrices.Value,
		BypassMessages: bypassMessages,
	}
}
