package main

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mymint/keeper"
	"mymint/types"
)

// 模块结构体
type AppModule struct {
	Keeper keeper.MintKeeper
}

// 初始化模块
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) {
	var mints []types.MintData
	json.Unmarshal(data, &mints)

	for _, mint := range mints {
		am.Keeper.SetNumber(ctx, string(mint.Address), mint.Number)
	}
}

// 导出模块数据
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return json.RawMessage("[]")
}
