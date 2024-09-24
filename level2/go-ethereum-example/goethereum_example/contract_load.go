package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "goethereum_example/goethereum_example/contracts/Store" // for demo
)

// 加载智能合约
func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	//一旦使用 abigen 工具将智能合约的 ABI 编译为 Go 包，下一步就是调用“New”方法，其格式为“New”，
	//所以在我们的例子中如果你 回想一下它将是_NewStore_。 此初始化方法接收智能合约的地址，并返回可以开始与之交互的合约实例。
	address := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
