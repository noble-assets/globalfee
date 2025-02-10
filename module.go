// Copyright 2025 NASD Inc. All Rights Reserved.
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

package globalfee

import (
	"context"
	"encoding/json"
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	modulev1 "github.com/noble-assets/globalfee/api/module/v1"
	globalfeev1 "github.com/noble-assets/globalfee/api/v1"
	"github.com/noble-assets/globalfee/client/cli"
	"github.com/noble-assets/globalfee/keeper"
	"github.com/noble-assets/globalfee/types"
	"github.com/spf13/cobra"
)

// ConsensusVersion defines the current GlobalFee module consensus version.
const ConsensusVersion = 2

var (
	_ module.AppModuleBasic      = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasGenesisBasics    = AppModuleBasic{}
	_ module.HasServices         = AppModule{}
)

//

type AppModuleBasic struct {
	registry codectypes.InterfaceRegistry
}

func NewAppModuleBasic(registry codectypes.InterfaceRegistry) AppModuleBasic {
	return AppModuleBasic{registry: registry}
}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (b AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genesis types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesis); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return genesis.Validate(b.registry)
}

//

type AppModule struct {
	AppModuleBasic

	subspace paramstypes.Subspace
	keeper   *keeper.Keeper
}

func NewAppModule(registry codectypes.InterfaceRegistry, subspace paramstypes.Subspace, keeper *keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(registry),
		subspace:       subspace,
		keeper:         keeper,
	}
}

func (AppModule) IsOnePerModuleType() {}

func (AppModule) IsAppModule() {}

func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

func (m AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
	var genesis types.GenesisState
	cdc.MustUnmarshalJSON(bz, &genesis)

	InitGenesis(ctx, m.keeper, genesis)
}

func (m AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesis := ExportGenesis(ctx, m.keeper)
	return cdc.MustMarshalJSON(genesis)
}

func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(m.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(m.keeper))

	migrator := NewMigrator(m.subspace, m.keeper)
	if err := cfg.RegisterMigration(types.ModuleName, 1, migrator.Migrate1to2); err != nil {
		panic(fmt.Sprintf("failed to migrate GlobalFee from version 1 to 2: %v", err))
	}
}

//

func (AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: globalfeev1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					// TODO: For now we overwrite this with a custom implementation because AutoCLI throws errors when parsing DecCoins. Remove once hearing back from the Cosmos SDK team.
					Skip:      true,
					RpcMethod: "UpdateGasPrices",
				},
				{
					RpcMethod: "UpdateBypassMessages",
					Use:       "update-bypass-messages [bypass-messages ...]",
					Short:     "Update the messages that are allowed to bypass required gas prices",
					Example:   "update-bypass-messages /ibc.core.client.v1.MsgUpdateClient /noble.globalfee.v1.MsgUpdateGasPrices",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{
							ProtoField: "bypass_messages",
							Varargs:    true,
						},
					},
				},
			},
			EnhanceCustomCommand: true,
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: globalfeev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GasPrices",
					Use:       "gas-prices",
				},
				{
					RpcMethod: "BypassMessages",
					Use:       "bypass-messages",
				},
			},
		},
	}
}

func (AppModule) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

//

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config   *modulev1.Module
	Service  store.KVStoreService
	Registry codectypes.InterfaceRegistry
	Cdc      codec.Codec
	Subspace paramstypes.Subspace
}

type ModuleOutputs struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	if in.Config.Authority == "" {
		panic("authority for GlobalFee module must be set")
	}

	authority := authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	k := keeper.NewKeeper(
		authority.String(),
		in.Registry,
		in.Service,
		in.Cdc,
	)
	m := NewAppModule(in.Registry, in.Subspace, k)

	return ModuleOutputs{Keeper: k, Module: m}
}
