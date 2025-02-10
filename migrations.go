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
