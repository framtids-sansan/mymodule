package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// 存储结构
type MintData struct {
	Address sdk.AccAddress `json:"address"`
	Number  int64          `json:"number"`
}

// 解析数据
func (m MintData) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func UnmarshalMintData(data []byte) (MintData, error) {
	var mint MintData
	err := json.Unmarshal(data, &mint)
	return mint, err
}

// 存储 Key
const (
	ModuleName = "mymint"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

// 定义前缀
var (
	KeyPrefixAddr = []byte{0x01} // 前缀
)

func Key(addr string) []byte {
	return append(KeyPrefixAddr, []byte(addr)...)
}
