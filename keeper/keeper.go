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
