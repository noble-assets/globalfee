package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
	"github.com/noble-assets/globalfee/utils"
	"github.com/noble-assets/globalfee/utils/mocks"
	"github.com/stretchr/testify/require"
)

var (
	usdAmount, _ = math.LegacyNewDecFromStr("0.100000000000000000")
	eurAmount, _ = math.LegacyNewDecFromStr("0.090000000000000000")

	USDC = sdk.NewDecCoinFromDec("uusdc", usdAmount)
	EURe = sdk.NewDecCoinFromDec("ueure", eurAmount)
)

func TestUpdateGasPrices(t *testing.T) {
	k, ctx := mocks.GlobalFeeKeeper()
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to update gas prices with invalid signer.
	_, err := server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorIs(t, err, types.ErrInvalidAuthority)

	// ACT: Attempt to update gas prices. Len(GasPrices) = 0.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{},
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	rawGasPrices, err := k.GasPrices.Get(ctx)
	require.NoError(t, err)
	gasPrices := rawGasPrices.Value
	require.Empty(t, gasPrices)

	// ACT: Attempt to update gas prices. Len(GasPrices) = 1. Invalid denom.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{sdk.DecCoin{Denom: "-", Amount: usdAmount}},
	})
	// ASSERT: The action should've failed due to invalid denom.
	require.ErrorContains(t, err, "invalid denom: -")

	// ACT: Attempt to update gas prices. Len(GasPrices) = 1. Invalid amount.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{sdk.DecCoin{Denom: "uusdc", Amount: usdAmount.MulInt64(-1)}},
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "gas price for uusdc is negative")

	// ACT: Attempt to update gas prices. Len(GasPrices) = 1. Valid.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{USDC},
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.NoError(t, err)
	rawGasPrices, err = k.GasPrices.Get(ctx)
	require.NoError(t, err)
	gasPrices = rawGasPrices.Value
	require.Len(t, gasPrices, 1)
	require.Equal(t, USDC, gasPrices[0])

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Invalid first denom.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer: "authority",
		GasPrices: sdk.DecCoins{
			sdk.DecCoin{Denom: "-", Amount: usdAmount},
			USDC,
		},
	})
	// ASSERT: The action should've failed due to invalid denom.
	require.ErrorContains(t, err, "invalid denom: -")

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Invalid first amount.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer: "authority",
		GasPrices: sdk.DecCoins{
			sdk.DecCoin{Denom: "uusdc", Amount: usdAmount.MulInt64(-1)},
			EURe,
		},
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "gas price for uusdc is negative")

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Repeated denom.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{USDC, USDC},
	})
	// ASSERT: The action should've failed due to repeated denom.
	require.ErrorContains(t, err, "denom uusdc is repeated")

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Invalid second denom.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer: "authority",
		GasPrices: sdk.DecCoins{
			EURe,
			sdk.DecCoin{Denom: "-", Amount: usdAmount.MulInt64(-1)},
		},
	})
	// ASSERT: The action should've failed due to invalid denom.
	require.ErrorContains(t, err, "invalid denom: -")

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Invalid second amount.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer: "authority",
		GasPrices: sdk.DecCoins{
			EURe,
			sdk.DecCoin{Denom: "uusdc", Amount: usdAmount.MulInt64(-1)},
		},
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.ErrorContains(t, err, "gas price for uusdc is negative")

	// ACT: Attempt to update gas prices. Len(GasPrices) > 1. Valid.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{EURe, USDC},
	})
	// ASSERT: The action should've failed due to invalid amount.
	require.NoError(t, err)
	rawGasPrices, err = k.GasPrices.Get(ctx)
	require.NoError(t, err)
	gasPrices = rawGasPrices.Value
	require.Len(t, gasPrices, 2)
	require.Equal(t, EURe, gasPrices[0])
	require.Equal(t, USDC, gasPrices[1])

	// ARRANGE: Set up a failing store.
	cfg := testutil.MakeTestEncodingConfig()
	k.GasPrices = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		types.GasPricesKey, "gas_prices", codec.CollValue[types.GasPrices](cfg.Codec),
	)

	// ACT: Attempt to update gas prices with failing store.
	_, err = server.UpdateGasPrices(ctx, &types.MsgUpdateGasPrices{
		Signer:    "authority",
		GasPrices: sdk.DecCoins{EURe, USDC},
	})
	// ASSERT: The action should've failed due to failing store.
	require.ErrorIs(t, err, mocks.ErrorStoreAccess)
}

func TestUpdateBypassMessages(t *testing.T) {
	k, ctx := mocks.GlobalFeeKeeper()
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to update bypass messages with invalid signer.
	_, err := server.UpdateBypassMessages(ctx, &types.MsgUpdateBypassMessages{})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorIs(t, err, types.ErrInvalidAuthority)

	// ARRANGE: Set a bypass message in state.
	require.NoError(t, k.BypassMessages.Set(ctx, sdk.MsgTypeURL(&types.MsgUpdateGasPrices{})))

	// ARRANGE: Set up a failing store.
	tmpBypassMessages := k.BypassMessages
	k.BypassMessages = collections.NewKeySet(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Delete, utils.GetKVStore(ctx, types.ModuleName))),
		types.BypassMessagePrefix, "bypass_messages", collections.StringKey,
	)

	// ACT: Attempt to update bypass messages with failing store.
	_, err = server.UpdateBypassMessages(ctx, &types.MsgUpdateBypassMessages{
		Signer: "authority",
	})
	// ASSERT: The action should've failed due to failing store.
	require.ErrorIs(t, err, mocks.ErrorStoreAccess)
	k.BypassMessages = tmpBypassMessages

	// ACT: Attempt to update bypass messages with unregistered message.
	_, err = server.UpdateBypassMessages(ctx, &types.MsgUpdateBypassMessages{
		Signer:         "authority",
		BypassMessages: []string{sdk.MsgTypeURL(&banktypes.MsgSend{})},
	})
	// ASSERT: The action should've failed due to unregistered message.
	require.ErrorContains(t, err, "invalid bypass message /cosmos.bank.v1beta1.MsgSend")

	// ARRANGE: Set up a failing store.
	tmpBypassMessages = k.BypassMessages
	k.BypassMessages = collections.NewKeySet(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		types.BypassMessagePrefix, "bypass_messages", collections.StringKey,
	)

	// ACT: Attempt to update bypass messages with failing store.
	_, err = server.UpdateBypassMessages(ctx, &types.MsgUpdateBypassMessages{
		Signer:         "authority",
		BypassMessages: []string{sdk.MsgTypeURL(&types.MsgUpdateBypassMessages{})},
	})
	// ASSERT: The action should've failed due to failing store.
	require.ErrorIs(t, err, mocks.ErrorStoreAccess)
	k.BypassMessages = tmpBypassMessages

	// ACT: Attempt to update bypass messages.
	_, err = server.UpdateBypassMessages(ctx, &types.MsgUpdateBypassMessages{
		Signer:         "authority",
		BypassMessages: []string{sdk.MsgTypeURL(&types.MsgUpdateBypassMessages{})},
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	bypassMessages, err := k.GetBypassMessages(ctx)
	require.NoError(t, err)
	require.Len(t, bypassMessages, 1)
	require.Contains(t, bypassMessages, "/noble.globalfee.v1.MsgUpdateBypassMessages")
}
