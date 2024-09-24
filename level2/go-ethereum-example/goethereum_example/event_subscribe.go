package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 订阅事件日志
func main() {
	//为了订阅事件日志，我们需要做的第一件事就是拨打启用 websocket 的以太坊客户端。 幸运的是，Infura 支持 websockets。
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	//下一步是创建筛选查询。 在这个例子中，我们将阅读来自我们在之前课程中创建的示例合约中的所有事件。
	contractAddress := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	//我们接收事件的方式是通过 Go channel。 让我们从 go-ethereum core/types 包创建一个类型为 Log 的 channel。
	logs := make(chan types.Log)

	//现在我们所要做的就是通过从客户端调用 SubscribeFilterLogs 来订阅，它接收查询选项和输出通道。
	//这将返回包含 unsubscribe 和 error 方法的订阅结构。
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	//最后，我们要做的就是使用 select 语句设置一个连续循环来读入新的日志事件或订阅错误。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
