package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "goethereum_example/goethereum_example/contracts/Store" // for demo
)

// 写入智能合约
func main() {
	//写入智能合约需要我们用私钥来对交易事务进行签名。
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("841163b4778beddbf90f4e2ca9545138ed9377b5a870e8f93a299130d83b24b7")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//还需要先查到 nonce 和燃气价格。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//接下来，我们创建一个新的 keyed transactor，它接收私钥。
	auth := bind.NewKeyedTransactor(privateKey)
	//然后我们需要设置 keyed transactor 的标准交易选项。
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//现在我们加载一个智能合约的实例
	address := common.HexToAddress("0x65E57942b2e20da8C463f7F236268fB43b9FC8f7")
	//要初始化它，我们只需调用合约包的_New_方法，并提供智能合约地址和 ethclient，它返回我们可以使用的合约实例。
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//创建的智能合约有一个名为_SetItem_的外部方法，它接受 solidity“bytes32”格式的两个参数（key，value）。
	//这意味着 Go 合约包要求我们传递一个长度为 32 个字节的字节数组。
	//调用_SetItem_方法需要我们传递我们之前创建的 auth 对象（keyed transactor）。
	//在幕后，此方法将使用它的参数对此函数调用进行编码，将其设置为事务的 data 属性，并使用私钥对其进行签名。
	//结果将是一个已签名的事务对象。
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	//要验证键/值是否已设置，我们可以读取智能合约中的值。
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
