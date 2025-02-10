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
