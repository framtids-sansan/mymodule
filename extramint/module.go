package extramint

import (
	"context"
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/extramint/keeper"
	"github.com/cosmos/cosmos-sdk/x/extramint/types"
)

// AppModuleBasic defines the basic application module used by the yourmodule module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the yourmodule module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the mint module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis 返回默认的 genesis 状态
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return json.RawMessage("{}") // 空的 genesis 状态
}

// ValidateGenesis 校验 genesis 状态
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ module.ClientTxEncodingConfig, bz json.RawMessage) error {
	return nil // 目前没有额外的校验
}

// AppModule implements an application module for the yourmodule module.
type AppModule struct {
	AppModuleBasic
	//keeper keeper.Keeper
}

// IsOnePerModuleType 确保该模块在应用中只注册一次
func (AppModule) IsOnePerModuleType() {}

// IsAppModule 确保该模块符合 `appmodule.AppModule` 接口
func (AppModule) IsAppModule() {}

// RegisterServices 注册 gRPC 服务（目前为空）
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis 初始化模块的 genesis 状态
func (am AppModule) InitGenesis(ctx context.Context, cdc codec.JSONCodec, data json.RawMessage) {
	//fmt.Println("extramint module initialized")
}

// ExportGenesis 导出模块的 genesis 状态
func (am AppModule) ExportGenesis(ctx context.Context, cdc codec.JSONCodec) json.RawMessage {
	return json.RawMessage("{}")
}

// ConsensusVersion 返回模块的版本号
func (AppModule) ConsensusVersion() uint64 { return 1 }

// 依赖注入注册
func init() {
	appmodule.Register(&module.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In
	Cdc codec.Codec
}

type ModuleOutputs struct {
	depinject.Out
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	return ModuleOutputs{Module: AppModule{AppModuleBasic: AppModuleBasic{cdc: in.Cdc}}}
}
