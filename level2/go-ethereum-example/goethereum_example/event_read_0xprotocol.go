package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	exchange "goethereum_example/goethereum_example/contracts/Exchange" // for demo
)

//要读取0x Protocol事件日志，我们必须首先将 solidity 智能合约编译为一个 Go 包。
//安装 solc 版本 0.4.26
//	npm i -g solc@0.4.26

//接着给定 abi，使用 abigen 来创建 Go exchange 包：
//	solcjs --abi Exchange.sol
//	abigen --abi="Exchange.sol:Exchange.abi" --pkg=exchange --out=Exchange.go

// 为例如 Exchange.sol 的事件日志创建 0x Protocol 交易所智能合约接口:
// LogFill ...
type LogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
}

// LogCancel ...
type LogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
}

// LogError ...
type LogError struct {
	ErrorID   uint8
	OrderHash [32]byte
}

// 读取0x Protocol事件日志
func main() {
	//初始化以太坊客户端：
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	//创建一个 FilterQuery，并为其传递 0x Protocol 智能合约地址和所需的区块范围：
	// 0x Protocol Exchange smart contract address
	contractAddress := common.HexToAddress("0x12459C951127e0c374FF9105DdA097662A027093")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383482),
		ToBlock:   big.NewInt(6383488),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//用 FilterLogs 查询日志：
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//接下来我们将解析 JSON abi，我们后续将使用解压缩原始日志数据：
	contractAbi, err := abi.JSON(strings.NewReader(string(exchange.ExchangeABI)))
	if err != nil {
		log.Fatal(err)
	}

	//为了按某种日志类型过滤，我们需要知晓每个事件日志函数签名的 keccak256 摘要。
	//正如我们很快所见到的那样，事件日志函数签名摘要总是 topic[0]：
	// NOTE: keccak256("LogFill(address,address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes32)")
	logFillEvent := common.HexToHash("0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3")

	// NOTE: keccak256("LogCancel(address,address,address,address,uint256,uint256,bytes32,bytes32)")
	logCancelEvent := common.HexToHash("67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131")

	// NOTE: keccak256("LogError(uint8,bytes32)")
	logErrorEvent := common.HexToHash("36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90")

	//现在我们迭代所有的日志并设置一个 switch 语句来按事件日志类型过滤：
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		//现在要解析 LogFill，我们将使用 abi.Unpack 将原始数据类型解析为我们自定义的日志类型结构体。
		//Unpack 不会解析 indexed 事件类型，因为这些它们存储在 topics 下，所以对于那些我们必须单独解析，如下例所示：
		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			fmt.Printf("Log Name: LogFill\n")

			var fillEvent LogFill

			//err := contractAbi.Unpack(&fillEvent, "LogFill", vLog.Data)
			err := contractAbi.UnpackIntoInterface(&fillEvent, "LogFill", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", fillEvent.Maker.Hex())
			fmt.Printf("Taker: %s\n", fillEvent.Taker.Hex())
			fmt.Printf("Fee Recipient: %s\n", fillEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", fillEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", fillEvent.TakerToken.Hex())
			fmt.Printf("Filled Maker Token Amount: %s\n", fillEvent.FilledMakerTokenAmount.String())
			fmt.Printf("Filled Taker Token Amount: %s\n", fillEvent.FilledTakerTokenAmount.String())
			fmt.Printf("Paid Maker Fee: %s\n", fillEvent.PaidMakerFee.String())
			fmt.Printf("Paid Taker Fee: %s\n", fillEvent.PaidTakerFee.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(fillEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(fillEvent.OrderHash[:]))

		case logCancelEvent.Hex():
			fmt.Printf("Log Name: LogCancel\n")

			var cancelEvent LogCancel

			//err := contractAbi.Unpack(&cancelEvent, "LogCancel", vLog.Data)
			err := contractAbi.UnpackIntoInterface(&cancelEvent, "LogCancel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			cancelEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			cancelEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			cancelEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", cancelEvent.Maker.Hex())
			fmt.Printf("Fee Recipient: %s\n", cancelEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", cancelEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", cancelEvent.TakerToken.Hex())
			fmt.Printf("Cancelled Maker Token Amount: %s\n", cancelEvent.CancelledMakerTokenAmount.String())
			fmt.Printf("Cancelled Taker Token Amount: %s\n", cancelEvent.CancelledTakerTokenAmount.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(cancelEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))

		case logErrorEvent.Hex():
			fmt.Printf("Log Name: LogError\n")

			errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			errorEvent := &LogError{
				ErrorID:   uint8(errorID),
				OrderHash: vLog.Topics[2],
			}

			fmt.Printf("Error ID: %d\n", errorEvent.ErrorID)
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))
		}

		fmt.Printf("\n\n")
	}
}
