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

package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
	"github.com/noble-assets/globalfee/utils"
	"github.com/noble-assets/globalfee/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestNewKeeper(t *testing.T) {
	// ARRANGE: Set the BypassMessagePrefix to an already existing key
	types.BypassMessagePrefix = types.GasPricesKey

	// ACT: Attempt to create a new Keeper with overlapping prefixes
	require.Panics(t, func() {
		cfg := testutil.MakeTestEncodingConfig()
		key := storetypes.NewKVStoreKey(types.ModuleName)

		keeper.NewKeeper(
			"",
			cfg.InterfaceRegistry,
			runtime.NewKVStoreService(key),
			cfg.Codec,
		)
	})
	// ASSERT: The function should've panicked.

	// ARRANGE: Restore the original BypassMessagePrefix
	types.BypassMessagePrefix = []byte("bypass_message")
}

func TestGetRequiredFees(t *testing.T) {
	k, ctx := mocks.GlobalFeeKeeper()

	// ARRANGE: Set up a failing GasPrices store.
	tmpGasPrices := k.GasPrices
	cfg := testutil.MakeTestEncodingConfig()
	k.GasPrices = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Get, utils.GetKVStore(ctx, types.ModuleName))),
		types.GasPricesKey, "gas_prices", codec.CollValue[types.GasPrices](cfg.Codec),
	)

	// ACT: Attempt to get required fees with failing GasPrices store.
	_, err := k.GetRequiredFees(ctx, nil)
	// ASSERT: The action should've failed due to store error.
	require.ErrorIs(t, err, mocks.ErrorStoreAccess)
	k.GasPrices = tmpGasPrices

	// ARRANGE: Create a mock transaction with a default gas limit of 200k.
	builder := cfg.TxConfig.NewTxBuilder()
	builder.SetGasLimit(200_000)
	tx := builder.GetTx()

	// ARRANGE: Set gas prices in state. Len(GasPrices) = 0.
	require.NoError(t, k.GasPrices.Set(ctx, types.GasPrices{Value: sdk.DecCoins{}}))

	// ACT: Attempt to get required fees.
	fees, err := k.GetRequiredFees(ctx, tx)
	// ARRANGE: The action should've succeeded, and returned required fees.
	require.NoError(t, err)
	require.True(t, fees.IsZero())

	// ARRANGE: Set gas prices in state. Len(GasPrices) = 1.
	require.NoError(t, k.GasPrices.Set(ctx, types.GasPrices{Value: sdk.DecCoins{USDC}}))

	// ACT: Attempt to get required fees.
	fees, err = k.GetRequiredFees(ctx, tx)
	// ARRANGE: The action should've succeeded, and returned required fees.
	require.NoError(t, err)
	require.Len(t, fees, 1)
	require.Equal(t, math.NewInt(20_000), fees.AmountOf(USDC.Denom))

	// ARRANGE: Set gas prices in state. Len(GasPrices) > 1.
	require.NoError(t, k.GasPrices.Set(ctx, types.GasPrices{Value: sdk.DecCoins{EURe, USDC}}))

	// ACT: Attempt to get required fees.
	fees, err = k.GetRequiredFees(ctx, tx)
	// ARRANGE: The action should've succeeded, and returned required fees.
	require.NoError(t, err)
	require.Len(t, fees, 2)
	require.Equal(t, math.NewInt(18_000), fees.AmountOf(EURe.Denom))
	require.Equal(t, math.NewInt(20_000), fees.AmountOf(USDC.Denom))
}
