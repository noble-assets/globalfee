package e2e

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/globalfee/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
)

func TestGlobalFee(t *testing.T) {

	ctx := context.Background()
	cw := ChainSpinUp(t, ctx)

	chain := cw.Chain
	val := chain.Validators[0]

	sender := interchaintest.GetAndFundTestUsers(t, ctx, "wallet", math.NewInt(100), chain)[0]

	// ARRANGE: gas-prices are 0stake and bypass-messages are currently empty
	receiver1, err := chain.BuildRelayerWallet(ctx, "receiver1")
	require.NoError(t, err)

	// ACT: Send coins with no gas while gas prices are set to 0
	sendRes := bankSendWithFees(t, ctx, chain, sender, receiver1.FormattedAddress(), "1ustake", "0ustake")

	// ASSERT: successful transctaion
	require.Contains(t, sendRes, "code: 0")

	// ARRANGE: set gas-prices to 10ustake
	broadcaster := cosmos.NewBroadcaster(t, chain)
	tx, err := cosmos.BroadcastTx(
		ctx,
		broadcaster,
		cw.AuthWall,
		&types.MsgUpdateGasPrices{
			Signer: cw.AuthWall.FormattedAddress(),
			GasPrices: sdk.DecCoins{
				sdk.DecCoin{Denom: "ustake", Amount: math.LegacyNewDec(10)},
			},
		},
	)
	require.NoError(t, err)
	require.Zero(t, tx.Code)

	// ASSERT: gas prices are set to 10ustake
	res, _, err := val.ExecQuery(ctx, "globalfee", "gas-prices")
	require.NoError(t, err)
	// TODO: Unmarshall
	// expectedGasPrices := types.QueryGasPricesResponse{
	// 	GasPrices: sdk.DecCoins{
	// 		sdk.DecCoin{Denom: "ustake", Amount: math.LegacyNewDec(10)},
	// 	},
	// }
	// var gasPricesResponse types.QueryGasPricesResponse
	// err = json.Unmarshal(res, &gasPricesResponse)
	require.NoError(t, err)
	// require.Equal(t, expectedGasPrices, gasPricesResponse)
	// TODO: remove once successful unmarshall check
	require.Contains(t, string(res), "10.000000000000000000ustake")

	// ACT: Send coins with no gas while gas prices are set to 10
	sendRes = bankSendWithFees(t, ctx, chain, sender, receiver1.FormattedAddress(), "1ustake", "0ustake")
	require.Contains(t, sendRes, "insufficient fees")

}

// We ignore some of the safeguards interchaintest puts in place (such as gas prices and adjustment, since we are testing fees)
func bankSendWithFees(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, from ibc.Wallet, toAddr, coins, feeCoin string) string {
	cmd := []string{chain.Config().Bin, "tx", "bank", "send", from.KeyName(), toAddr, coins,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--gas", "200000",
		"--fees", feeCoin,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}

	return string(stdout)
}
