package globalfee

import (
	"context"

	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
)

func InitGenesis(ctx context.Context, k *keeper.Keeper, genesis types.GenesisState) {
	gasPrices := types.GasPrices{Value: genesis.GasPrices}
	if err := k.GasPrices.Set(ctx, gasPrices); err != nil {
		panic(err)
	}

	for _, bypassMessage := range genesis.BypassMessages {
		if err := k.BypassMessages.Set(ctx, bypassMessage); err != nil {
			panic(err)
		}
	}
}

func ExportGenesis(ctx context.Context, k *keeper.Keeper) *types.GenesisState {
	gasPrices, err := k.GasPrices.Get(ctx)
	if err != nil {
		panic(err)
	}

	bypassMessages, err := k.GetBypassMessages(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		GasPrices:      gasPrices.Value,
		BypassMessages: bypassMessages,
	}
}
