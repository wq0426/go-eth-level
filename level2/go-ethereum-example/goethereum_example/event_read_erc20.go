package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "goethereum_example/goethereum_example/contracts/erc20" // for demo
)

// 现在在我们的 Go 应用程序中，让我们创建与 ERC-20 事件日志签名类型相匹配的结构类型：
// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

// 读取 ERC-20 代币的事件日志
func main() {
	//初始化以太坊客户端
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	//按照 ERC-20 智能合约地址和所需的块范围创建一个“FilterQuery”。这个例子我们会用ZRX 代币:
	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//用 FilterLogs 来过滤日志：
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//接下来我们将解析 JSON abi，稍后我们将使用解压缩原始日志数据：
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	//为了按某种日志类型进行过滤，我们需要弄清楚每个事件日志函数签名的 keccak256 哈希值。
	//事件日志函数签名哈希始终是 topic [0]，我们很快就会看到。
	//以下是使用 go-ethereum crypto 包计算 keccak256 哈希的方法：
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	//现在我们将遍历所有日志并设置 switch 语句以按事件日志类型进行过滤：
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():

			//现在要解析 Transfer 事件日志，我们将使用 abi.Unpack 将原始日志数据解析为我们的日志类型结构。
			//解包不会解析 indexed 事件类型，因为它们存储在 topics 下，所以对于那些我们必须单独解析，如下例所示：
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			//err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			//Approval 日志也是类似的方法：
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			//err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
