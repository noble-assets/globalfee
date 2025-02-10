package types

import (
	"fmt"

	"cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		GasPrices:      sdk.DecCoins{},
		BypassMessages: []string{},
	}
}

func (genesis *GenesisState) Validate(registry codectypes.InterfaceRegistry) error {
	gasPrices := GasPrices{Value: genesis.GasPrices}
	if err := gasPrices.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate gas prices")
	}

	for _, bypassMessage := range genesis.BypassMessages {
		resolved, err := registry.Resolve(bypassMessage)
		if err != nil || resolved == nil {
			return fmt.Errorf("invalid bypass message %s", bypassMessage)
		}
	}

	return nil
}
