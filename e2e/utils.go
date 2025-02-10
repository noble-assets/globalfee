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
