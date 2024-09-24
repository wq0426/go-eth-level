package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// 账户

func main() {
	//使用 go-ethereum 的账户地址，必须先将它们转化为 go-ethereum 中的 common.Address 类型
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	//fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
