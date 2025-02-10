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

package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/noble-assets/globalfee/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/stretchr/testify/require"
)

func TestGlobalFee(t *testing.T) {
	ctx, wrapper := GlobalFeeSuite(t)
	validator := wrapper.Chain.Validators[0]

	var (
		gasPrice = sdk.DecCoin{Denom: "ustake", Amount: math.LegacyNewDec(10)}
		gas      = math.NewInt(200_000) // default gas
		fee      = fmt.Sprintf("%sustake", gasPrice.Amount.MulInt(gas).String())
	)

	// ARRANGE: Generate a sender and recipient wallet.
	sender := interchaintest.GetAndFundTestUsers(t, ctx, "sender", math.NewInt(100_000_000_000), wrapper.Chain)[0]
	recipient, err := wrapper.Chain.BuildRelayerWallet(ctx, "recipient")
	require.NoError(t, err)

	// ACT: Attempt a transaction with no fees.
	err = bankSendWithFees(ctx, validator, sender, recipient, "1ustake", "0ustake")
	// ASSERT: The transaction was successful due to no required fees.
	require.NoError(t, err)

	// ACT: Set required gas prices to 10ustake.
	_, err = validator.ExecTx(ctx, wrapper.Authority.KeyName(), "globalfee", "update-gas-prices", gasPrice.String())
	// ASSERT: The transaction was successful and updated the required gas prices.
	require.NoError(t, err)
	gasPrices, err := GasPrices(ctx, validator)
	require.NoError(t, err)
	require.Equal(t, sdk.NewDecCoins(gasPrice), gasPrices)

	// ACT: Attempt a transaction with no fees.
	err = bankSendWithFees(ctx, validator, sender, recipient, "1ustake", "0ustake")
	// ASSERT: The transaction failed due to insufficient fees.
	require.ErrorContains(t, err, "insufficient fee")

	// ACT: Set x/bank MsgSend as a bypassed message.
	bankSendType := sdk.MsgTypeURL(&banktypes.MsgSend{})
	_, err = validator.ExecTx(ctx, wrapper.Authority.KeyName(), "globalfee", "update-bypass-messages", bankSendType, "--fees", fee)
	// ASSERT: The transaction was successful and updated the bypass messages.
	require.NoError(t, err)
	raw, _, err := validator.ExecQuery(ctx, "globalfee", "bypass-messages")
	require.NoError(t, err)
	var res types.QueryBypassMessagesResponse
	require.NoError(t, json.Unmarshal(raw, &res))
	require.Len(t, res.BypassMessages, 1)
	require.Contains(t, res.BypassMessages, bankSendType)

	// ACT: Attempt a transaction with no fees.
	err = bankSendWithFees(ctx, validator, sender, recipient, "1ustake", "0ustake")
	// ASSERT: The transaction was successful due to bypassed message.
	require.NoError(t, err)
}
