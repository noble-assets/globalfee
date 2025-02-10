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
