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

package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/globalfee/types"
)

type Keeper struct {
	authority string
	registry  codectypes.InterfaceRegistry

	GasPrices      collections.Item[types.GasPrices]
	BypassMessages collections.KeySet[string]
}

func NewKeeper(
	authority string,
	registry codectypes.InterfaceRegistry,
	service store.KVStoreService,
	cdc codec.Codec,
) *Keeper {
	builder := collections.NewSchemaBuilder(service)

	keeper := &Keeper{
		authority: authority,
		registry:  registry,

		GasPrices:      collections.NewItem(builder, types.GasPricesKey, "gas_prices", codec.CollValue[types.GasPrices](cdc)),
		BypassMessages: collections.NewKeySet(builder, types.BypassMessagePrefix, "bypass_messages", collections.StringKey),
	}

	_, err := builder.Build()
	if err != nil {
		panic(err)
	}

	return keeper
}

// GetRequiredFees is a utility that returns the required fees for a given
// transaction using the gas prices configured in this module.
func (k *Keeper) GetRequiredFees(ctx context.Context, tx sdk.FeeTx) (sdk.Coins, error) {
	rawGasPrices, err := k.GasPrices.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get minimum gas prices from state")
	}
	gasPrices := rawGasPrices.Value

	gasLimit := tx.GetGas()
	requiredFees := sdk.NewCoins()

	for _, gasPrice := range gasPrices {
		fee := gasPrice.Amount.MulInt64(int64(gasLimit))
		// TODO: Is the Ceil needed?
		requiredFees = requiredFees.Add(sdk.NewCoin(gasPrice.Denom, fee.Ceil().RoundInt()))
	}

	return requiredFees.Sort(), nil
}
