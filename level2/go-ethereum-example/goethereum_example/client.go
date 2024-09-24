package main

//首先，导入 go-etherem 的 ethclient 包
import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

//初始化客户端

func main() {
	//通过调用接收区块链服务提供者 URL 的 Dial 来初始化它。
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
