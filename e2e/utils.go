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
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
)

// We ignore some of the safeguards interchaintest puts in place (such as gas prices and adjustment, since we are testing fees)
func bankSendWithFees(
	ctx context.Context,
	validator *cosmos.ChainNode,
	sender, recipient ibc.Wallet,
	amount, fees string,
) error {
	_, err := validator.ExecTx(
		ctx, sender.KeyName(),
		"bank", "send",
		sender.FormattedAddress(),
		recipient.FormattedAddress(), amount,
		"--gas", "200000",
		"--fees", fees,
	)

	return err
}

func GasPrices(ctx context.Context, validator *cosmos.ChainNode) (sdk.DecCoins, error) {
	type QueryGasPricesResponse struct {
		GasPrices []string `json:"gas_prices"`
	}

	raw, _, err := validator.ExecQuery(ctx, "globalfee", "gas-prices")
	if err != nil {
		return nil, err
	}

	var res QueryGasPricesResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, err
	}

	var gasPrices sdk.DecCoins
	for _, rawGasPrice := range res.GasPrices {
		gasPrice, _ := sdk.ParseDecCoin(rawGasPrice)
		gasPrices = append(gasPrices, gasPrice)
	}

	return gasPrices, nil
}
