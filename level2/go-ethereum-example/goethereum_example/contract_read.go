package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "goethereum_example/goethereum_example/contracts/Store" // for demo
)

// 查询智能合约
func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//在部署过程中设置的合约中有一个名为 version 的全局变量。 因为它是公开的，这意味着它们将成为我们自动创建的 getter 函数。
	//常量和 view 函数也接受 bind.CallOpts 作为第一个参数。了解可用的具体选项要看相应类的文档 一般情况下我们可以用 nil。
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
