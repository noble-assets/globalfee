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

func TestThis(t *testing.T) {
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

	// ACT: Add bank transfer messages to bypass fee
	bankSendType := sdk.MsgTypeURL(&banktypes.MsgSend{}) // /cosmos.bank.v1beta1.MsgSend
	o := sdk.MsgTypeURL(&types.MsgUpdateGasPrices{})
	_, err = validator.ExecTx(ctx, wrapper.Authority.KeyName(), "globalfee", "update-bypass-messages", "/cosmos.bank.v1beta1.MsgSend", o, "--fees", fee)
	require.NoError(t, err)
	// ASSERT: The transaction successfully updated the bypass messages
	raw, _, err := validator.ExecQuery(ctx, "globalfee", "bypass-messages")
	require.NoError(t, err)
	fmt.Println("HERERERE", string(raw))
	var bypassMessagesResponse types.QueryBypassMessagesResponse
	err = json.Unmarshal(raw, &bypassMessagesResponse)
	require.NoError(t, err)
	expectedBypassMessages := types.QueryBypassMessagesResponse{
		BypassMessages: []string{bankSendType},
	}
	require.Equal(t, expectedBypassMessages, bypassMessagesResponse)

	// ACT: Attempt a bank send with no fees now that that message types is bypassed
	err = bankSendWithFees(ctx, validator, sender, recipient, "1ustake", "0ustake")
	// ASSERT: The transaction was successful with no fees.
	require.NoError(t, err)
}
