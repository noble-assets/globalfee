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

package globalfee

import (
	"math"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/noble-assets/globalfee/keeper"
)

// TxFeeChecker returns a custom ante.TxFeeChecker that ensures the fees for a
// given transaction respect the gas prices set in the GlobalFee module.
func TxFeeChecker(keeper *keeper.Keeper) ante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return nil, 0, errors.Wrap(errorstypes.ErrTxDecode, "Tx must be a FeeTx")
		}
		fees := feeTx.GetFee()

		if ctx.IsCheckTx() {
			allBypassMessages := true
			for _, msg := range feeTx.GetMsgs() {
				if has, _ := keeper.BypassMessages.Has(ctx, sdk.MsgTypeURL(msg)); !has {
					allBypassMessages = false
					break
				}
			}
			if allBypassMessages && fees.IsZero() {
				return sdk.Coins{}, 0, nil
			}

			requiredFees, err := keeper.GetRequiredFees(ctx, feeTx)
			if err != nil {
				return nil, 0, err
			}
			if len(requiredFees) == 0 {
				return sdk.Coins{}, 0, nil
			}

			sufficientFees := false
			for _, fee := range fees {
				found, requiredFee := requiredFees.Find(fee.Denom)
				if found && fee.Amount.GTE(requiredFee.Amount) {
					sufficientFees = true
					break
				}
			}

			if !sufficientFees {
				return nil, 0, errors.Wrapf(errorstypes.ErrInsufficientFee, "expected at least one of %s", requiredFees)
			}
		}

		return fees, getTxPriority(fees, int64(feeTx.GetGas())), nil
	}
}

// getTxPriority is copied from the Cosmos SDK as it's not exported for reuse.
// https://github.com/cosmos/cosmos-sdk/blob/v0.50.10/x/auth/ante/validator_tx_fee.go#L50-L68
func getTxPriority(fee sdk.Coins, gas int64) int64 {
	var priority int64
	for _, c := range fee {
		p := int64(math.MaxInt64)
		gasPrice := c.Amount.QuoRaw(gas)
		if gasPrice.IsInt64() {
			p = gasPrice.Int64()
		}
		if priority == 0 || p < priority {
			priority = p
		}
	}

	return priority
}
