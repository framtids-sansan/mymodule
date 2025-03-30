package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState 定义 extramint 模块的初始状态
type GenesisState struct {
	// 未来可以在这里添加初始化参数
}

// NewGenesisState 创建一个新的 GenesisState
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState 返回默认 GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// Validate 进行基本的 GenesisState 校验
func (gs GenesisState) Validate() error {
	// 目前没有具体数据需要校验
	return nil
}

// MarshalJSON 将 GenesisState 序列化
func (gs GenesisState) MarshalJSON() ([]byte, error) {
	return json.Marshal(gs)
}

// UnmarshalJSON 解析 JSON 数据
func (gs *GenesisState) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, gs)
}

// GetGenesisStateFromAppState 从 appState 获取 extramint 的 GenesisState
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}
	return genesisState
}
