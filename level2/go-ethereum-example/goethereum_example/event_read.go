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

	store "goethereum_example/goethereum_example/contracts/Store" // for demo
)

// 读取事件日志
func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")

	//智能合约可以可选地释放“事件”，其作为交易收据的一部分存储日志。读取这些事件相当简单。首先我们需要构造一个过滤查询。
	//我们从 go-ethereum 包中导入 FilterQuery 结构体并用过滤选项初始化它。
	//我们告诉它我们想过滤的区块范围并指定从中读取此日志的合约地址。
	//在示例中，我们将从在智能合约章节创建的智能合约中读取特定区块所有日志。
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//下一步是调用 ethclient 的 FilterLogs，它接收我们的查询并将返回所有的匹配事件日志。
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//返回的所有日志将是 ABI 编码，因此它们本身不会非常易读。为了解码日志，我们需要导入我们智能合约的 ABI。
	//为此，我们导入编译好的智能合约 Go 包，它将包含名称格式为 <Contract>ABI 的外部属性。
	//之后，我们使用 go-ethereum 中的 accounts/abi 包的 abi.JSON 函数返回一个我们可以在 Go 应用程序中使用的解析过的 ABI 接口。
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	//现在我们可以通过日志进行迭代并将它们解码为我么可以使用的类型。
	//若您回忆起我们的样例合约释放的日志在 Solidity 中是类型为 bytes32，那么 Go 中的等价物将是 [32]byte。
	//我们可以使用这些类型创建一个匿名结构体，并将指针作为第一个参数传递给解析后的 ABI 接口的 Unpack 函数，以解码原始的日志数据。
	//第二个参数是我们尝试解码的事件名称，最后一个参数是编码的日志数据。
	for _, vLog := range logs {
		//此外，日志结构体包含附加信息，例如，区块摘要，区块号和交易摘要。
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		//err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		//若您的 solidity 事件包含 indexed 事件类型，那么它们将成为_主题_而不是日志的数据属性的一部分。
		//在 solidity 中您最多只能有 4 个主题，但只有 3 个可索引的事件类型。第一个主题总是事件的签名。
		//我们的示例合约不包含可索引的事件，但如果它确实包含，这是如何读取事件主题。
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
