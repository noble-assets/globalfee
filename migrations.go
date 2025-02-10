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
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
)

type Migrator struct {
	*keeper.Keeper
	subspace paramstypes.Subspace
}

func NewMigrator(subspace paramstypes.Subspace, keeper *keeper.Keeper) Migrator {
	return Migrator{
		Keeper:   keeper,
		subspace: subspace,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	var params types.LegacyParams
	m.subspace.GetParamSet(ctx, &params)

	err := m.GasPrices.Set(ctx, types.GasPrices{Value: params.GasPrices})
	if err != nil {
		return errors.Wrap(err, "failed to set gas prices in state")
	}

	for _, bypassMessage := range params.BypassMessages {
		err := m.BypassMessages.Set(ctx, bypassMessage)
		if err != nil {
			return errors.Wrapf(err, "failed to set bypass message %s in state", bypassMessage)
		}
	}

	return nil
}
