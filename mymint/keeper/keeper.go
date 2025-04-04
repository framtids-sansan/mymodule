package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"encoding/binary"
	//"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/framtids-sansan/mymodule/mymint/types"
)

// Keeper 负责数据存储
type MintKeeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

// 存储数据
func (mk MintKeeper) SetNumber(ctx sdk.Context, addr string, number int64) {
	// 获取 KVStore，并用 types.KeyPrefixAddr 作为前缀，创建前缀存储空间。
	store := prefix.NewStore(ctx.KVStore(mk.storeKey), types.KeyPrefixAddr)
	// 生成存储键 []byte格式
	key := types.Key(addr)
	// 创建一个 8 字节的切片，用于存储 int64 类型的值。
	value := make([]byte, 8)
	binary.BigEndian.PutUint64(value, uint64(number))
	store.Set(key, value)
}

// 获取数据
func (mk MintKeeper) GetNumber(ctx sdk.Context, addr string) (int64, bool) {
	// 获取带前缀的 KVStore
	store := prefix.NewStore(ctx.KVStore(mk.storeKey), types.KeyPrefixAddr)
	// 生成存储键 []byte格式
	key := types.Key(addr)
	// 获取value []byte格式
	value := store.Get(key)
	if value == nil {
		return 0, false
	}
	return int64(binary.BigEndian.Uint64(value)), true
}
