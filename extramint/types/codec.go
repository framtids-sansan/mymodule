package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterLegacyAminoCodec 注册 LegacyAmino 编解码器（主要用于 JSON 序列化）
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// 未来如果 extramint 需要支持 Amino 序列化的消息类型，可以在这里注册
}

// RegisterInterfaces 注册 extramint 相关的接口和消息类型
func RegisterInterfaces(registry types.InterfaceRegistry) {
	// 未来如果 extramint 有自定义接口或消息类型，可以在这里注册
}

var (
	ModuleCdc = codec.NewLegacyAmino()
)
