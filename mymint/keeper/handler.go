package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = MintKeeper{}.SetNumber

// 处理交易消息
func HandleMsgSetMint(ctx sdk.Context, mk MintKeeper, addr string, number int64) error {
	mk.SetNumber(ctx, addr, number)
	return nil
}
