package e2e

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

type ChainWrapper struct {
	Chain     *cosmos.CosmosChain
	Authority ibc.Wallet
}

func GlobalFeeSuite(t *testing.T) (ctx context.Context, wrapper ChainWrapper) {
	ctx = context.Background()
	reporter := testreporter.NewNopReporter()
	execReporter := reporter.RelayerExecReporter(t)
	client, network := interchaintest.DockerSetup(t)

	numValidators, numFullNodes := 1, 0

	factory := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			NumValidators: &numValidators,
			NumFullNodes:  &numFullNodes,
			ChainConfig: ibc.ChainConfig{
				Type:           "cosmos",
				Name:           "globalfee",
				ChainID:        "global-1",
				Bin:            "simd",
				Denom:          "ustake",
				Bech32Prefix:   "noble",
				CoinType:       "118",
				GasPrices:      "0.0ustake",
				GasAdjustment:  1.2,
				TrustingPeriod: "504hr",
				Images: []ibc.DockerImage{
					{
						Repository: "noble-globalfee-simd",
						Version:    "local",
						UidGid:     "1025:1025",
					},
				},
				PreGenesis: preGenAuthority(ctx, &wrapper),
			},
		},
	})

	chains, err := factory.Chains(t.Name())
	require.NoError(t, err)

	wrapper.Chain = chains[0].(*cosmos.CosmosChain)
	interchain := interchaintest.NewInterchain().AddChain(wrapper.Chain)

	require.NoError(t, interchain.Build(ctx, execReporter, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))

	t.Cleanup(func() {
		_ = interchain.Close()
	})

	return
}

func preGenAuthority(ctx context.Context, wrapper *ChainWrapper) func(ibc.Chain) error {
	return func(chain ibc.Chain) (err error) {
		validator := wrapper.Chain.Validators[0]

		wrapper.Authority, err = validator.Chain.BuildWallet(ctx, "authority", "usual parade country forward clerk group ripple dust upset sun spike dish business foster lawn jealous panther junior kite sail erosion bean armed soup")
		if err != nil {
			return err
		}

		return validator.AddGenesisAccount(ctx, wrapper.Authority.FormattedAddress(), []sdk.Coin{sdk.NewCoin(chain.Config().Denom, math.NewInt(10000000))})
	}
}
