package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

// 代币的转账
func main() {
	//连接客户端
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	//加载私钥
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//代币传输不需要传输 ETH，因此将交易“值”设置为“0”。
	value := big.NewInt(0) // in wei (0 eth)
	//配置gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//先将您要发送代币的地址存储在变量中。
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//让我们将代币合约地址分配给变量。
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	//函数名将是传递函数的名称，即 ERC-20 规范中的 transfer 和参数类型。
	//第一个参数类型是 address（令牌的接收者），第二个类型是 uint256（要发送的代币数量）。
	//不需要没有空格和参数名称。 我们还需要用字节切片格式。
	transferFnSignature := []byte("transfer(address,uint256)")

	// 导入 crypto/sha3 包以生成函数签名的 Keccak256 哈希。 然后我们只使用前 4 个字节来获取方法 ID。
	//hash := sha3.NewKeccak256()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	//接下来，我们需要将给我们发送代币的地址左填充到 32 字节。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	//接下来我们确定要发送多少个代币，在这个例子里是 1,000 个，并且我们需要在 big.Int 中格式化为 wei。

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens

	//代币量也需要左填充到 32 个字节。
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	//接下来我们只需将方法 ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片。
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gas上限制将取决于交易数据的大小和智能合约必须执行的计算步骤。
	//幸运的是，客户端提供了 EstimateGas 方法，它可以为我们估算所需的燃气量。
	//这个函数从 ethereum 包中获取 CallMsg 结构，我们在其中指定数据和地址。
	//它将返回我们估算的完成交易所需的估计燃气上限。
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	//接下来我们需要做的是构建交易事务类型，这类似于您在 ETH 转账部分中看到的，
	//除了_to_字段将是代币智能合约地址。 这个常让人困惑。
	//我们还必须在调用中包含 0 ETH 的值字段和刚刚生成的数据字节。
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	//下一步是使用发件人的私钥对事务进行签名。
	//SignTx 方法需要 EIP155 签名器(EIP155 signer)，这需要我们从客户端拿到链 ID。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//最后广播交易。
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
