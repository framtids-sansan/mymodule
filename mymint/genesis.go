package main

import (
	"encoding/json"
	//"mymint/keeper"
	"github.com/framtids-sansan/mymodule/mymint/types"
)

// 初始状态
type GenesisState struct {
	Mints []types.MintData `json:"mints"`
}

// 解析
func DefaultGenesis() GenesisState {
	return GenesisState{Mints: []types.MintData{}}
}

func (g GenesisState) MarshalJSON() ([]byte, error) {
	return json.Marshal(g)
}
