// Copyright 2024 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
