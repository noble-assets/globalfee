package mocks

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
)

func GlobalFeeKeeper() (*keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_globalfee")

	cfg := moduletestutil.MakeTestEncodingConfig()
	types.RegisterInterfaces(cfg.InterfaceRegistry)

	k := keeper.NewKeeper(
		"authority",
		cfg.InterfaceRegistry,
		runtime.NewKVStoreService(key),
		cfg.Codec,
	)

	return k, testutil.DefaultContext(key, tkey)
}
