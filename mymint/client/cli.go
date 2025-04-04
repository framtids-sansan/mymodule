package client

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/framtids-sansan/mymodule/mymint/keeper"
	"strconv"
)

// CLI 命令：存储
func SetMintCmd(mk keeper.MintKeeper) *cobra.Command {
	return &cobra.Command{
		Use:   "set-mint [address] [number]",
		Short: "Set addr -> number",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			number := args[1]
			amount, _ := strconv.ParseInt(number, 10, 64)

			ctx := sdk.Context{} // 这里需要真实的 ctx
			mk.SetNumber(ctx, string(addr), amount)
			fmt.Println("数据已存储！")
			return nil
		},
	}
}

// CLI 命令：查询
func GetMintCmd(mk keeper.MintKeeper) *cobra.Command {
	return &cobra.Command{
		Use:   "get-mint [address]",
		Short: "查询 addr -> number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			ctx := sdk.Context{} // 这里需要真实的 ctx
			mint, found := mk.GetNumber(ctx, string(addr))
			if !found {
				fmt.Println("未找到数据")
				return nil
			}
			fmt.Printf("地址: %s, Number: %d\n", string(addr), mint)
			return nil
		},
	}
}
