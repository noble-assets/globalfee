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

package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/globalfee/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	*Keeper
}

func NewQueryServer(keeper *Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

func (k queryServer) GasPrices(ctx context.Context, req *types.QueryGasPrices) (*types.QueryGasPricesResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	rawGasPrices, err := k.Keeper.GasPrices.Get(ctx)
	if err != nil {
		return nil, err
	}
	gasPrices := rawGasPrices.Value

	return &types.QueryGasPricesResponse{GasPrices: gasPrices}, nil
}

func (k queryServer) BypassMessages(ctx context.Context, req *types.QueryBypassMessages) (*types.QueryBypassMessagesResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	bypassMessages, err := k.GetBypassMessages(ctx)

	return &types.QueryBypassMessagesResponse{BypassMessages: bypassMessages}, err
}
