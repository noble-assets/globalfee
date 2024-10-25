package e2e

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/noble-assets/globalfee/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

type ChainWrapper struct {
	Chain    *cosmos.CosmosChain
	AuthWall ibc.Wallet
}

func ChainSpinUp(t *testing.T, ctx context.Context) (cw ChainWrapper) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("noble", "noblepub")

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	client, network := interchaintest.DockerSetup(t)

	numValidators := 1
	numFullNodes := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
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
						Repository: "globalfee-simd",
						Version:    "local",
						UidGid:     "1025:1025",
					},
				},
				EncodingConfig: appEncoding(),
				PreGenesis:     preGenAuthority(ctx, &cw),
				ModifyGenesis: cosmos.ModifyGenesis(
					[]cosmos.GenesisKV{
						cosmos.NewGenesisKV("app_state.globalfee.gas_prices", sdk.DecCoins{sdk.DecCoin{Denom: "ustake", Amount: math.LegacyZeroDec()}}),
					},
				),
			},
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	cw.Chain = chains[0].(*cosmos.CosmosChain)

	ic := interchaintest.NewInterchain().
		AddChain(cw.Chain)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	return
}

func preGenAuthority(ctx context.Context, cw *ChainWrapper) func(ibc.Chain) (err error) {
	return func(c ibc.Chain) error {
		val := cw.Chain.Validators[0]

		var err error
		cw.AuthWall, err = val.Chain.BuildWallet(ctx, "auth-wallet", "market ready pilot lunch host cancel drive script remove brief lunch entry worth giant unknown grain romance gym tide perfect short because envelope sentence")
		if err != nil {
			return err
		}

		return val.AddGenesisAccount(ctx, cw.AuthWall.FormattedAddress(), []sdk.Coin{sdk.NewCoin(c.Config().Denom, math.ZeroInt())})
	}
}

func appEncoding() *sdktestutil.TestEncodingConfig {
	enc := cosmos.DefaultEncoding()
	types.RegisterInterfaces(enc.InterfaceRegistry)

	return &enc
}
