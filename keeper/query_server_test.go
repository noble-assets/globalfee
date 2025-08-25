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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
	"github.com/noble-assets/globalfee/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestGasPricesQuery(t *testing.T) {
	k, ctx := mocks.GlobalFeeKeeper()
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query gas prices with invalid request.
	_, err := server.GasPrices(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query gas prices with no state.
	_, err = server.GasPrices(ctx, &types.QueryGasPrices{})
	// ASSERT: The query should've failed due to no state.
	require.ErrorIs(t, err, collections.ErrNotFound)

	// ARRANGE: Set gas prices in state.
	require.NoError(t, k.GasPrices.Set(ctx, types.GasPrices{
		Value: sdk.DecCoins{EURe, USDC},
	}))

	// ACT: Attempt to query gas prices.
	res, err := server.GasPrices(ctx, &types.QueryGasPrices{})
	// ASSERT: The query should've succeeded, and returned gas prices.
	require.NoError(t, err)
	require.Len(t, res.GasPrices, 2)
	require.Contains(t, res.GasPrices, EURe)
	require.Contains(t, res.GasPrices, USDC)
}

func TestBypassMessagesQuery(t *testing.T) {
	k, ctx := mocks.GlobalFeeKeeper()
	server := keeper.NewQueryServer(k)

	// ACT: Attempt to query bypass messages with invalid request.
	_, err := server.BypassMessages(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query bypass messages with no state.
	res, err := server.BypassMessages(ctx, &types.QueryBypassMessages{})
	// ASSERT: The query should've succeeded, and returned bypass messages.
	require.NoError(t, err)
	require.Empty(t, res.BypassMessages)

	// ARRANGE: Set bypass messages in state.
	updateGasPrice := sdk.MsgTypeURL(&types.MsgUpdateGasPrices{})
	require.NoError(t, k.BypassMessages.Set(ctx, updateGasPrice))
	updateBypassMessages := sdk.MsgTypeURL(&types.MsgUpdateBypassMessages{})
	require.NoError(t, k.BypassMessages.Set(ctx, updateBypassMessages))

	// ACT: Attempt to query bypass messages.
	res, err = server.BypassMessages(ctx, &types.QueryBypassMessages{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Len(t, res.BypassMessages, 2)
	require.Contains(t, res.BypassMessages, updateGasPrice)
	require.Contains(t, res.BypassMessages, updateBypassMessages)
}
