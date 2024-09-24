package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 读取智能合约的字节码
func main() {
	//首先设置客户端和要读取的字节码的智能合约地址。
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")

	//现在你需要调用客户端的 codeAt 方法。 codeAt 方法接受智能合约地址和可选的块编号，并以字节格式返回字节码。
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
