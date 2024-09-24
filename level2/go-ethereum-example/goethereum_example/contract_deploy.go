package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	/*
		将 interface 智能合约编译为 JSON ABI，并使用 abigen 从 ABI 创建 Go 包
			安装 Solidity 编译器 (solc).
			npm install -g solc

			安装abigen
			go get -u github.com/ethereum/go-ethereum
			cd $GOPATH/src/github.com/ethereum/go-ethereum/
			make
			make devtools

			执行命令,从一个 solidity 文件生成 ABI
			solcjs --abi Store.sol

			用 abigen 将 ABI 转换为我们可以导入的 Go 文件
			abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go

			为了从 Go 部署智能合约，我们还需要将 solidity 智能合约编译为 EVM 字节码。
			EVM 字节码将在事务的数据字段中发送。 在 Go 文件上生成部署方法需要 bin 文件。
			solcjs --bin Store.sol

			编译 Go 合约文件，其中包括 deploy 方法，因为我们包含了 bin 文件
			abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go

		将新的_store_包导入我们的应用程序
		1、将文件放到C:\Program Files\Go\src\store目录下(GOPATCH)
		2、将文件上传到github，类似上面导入的写法"github.com/youname/project/store"
		3、根据go.mod里面的module名加上目录，moduleName/directory
	*/
	store "goethereum_example/goethereum_example/contracts/Store" // for demo
)

// 部署智能合约
func main() {
	//设置 ethclient
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	//加载私钥
	//fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
	//841163b4778beddbf90f4e2ca9545138ed9377b5a870e8f93a299130d83b24b7
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

	//设置通常的属性，如 nonce，燃气价格，燃气上线限制和 ETH 值。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//创建一个有配置密匙的交易发送器(tansactor)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//deploy 函数接受有密匙的事务处理器，ethclient，以及智能合约构造函数可能接受的任何输入参数。
	//我们测试的智能合约接受一个版本号的字符串参数。 此函数将返回新部署的合约地址，事务对象，我们可以交互的合约实例，还有错误（如果有）
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	// 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	//0x65E57942b2e20da8C463f7F236268fB43b9FC8f7
	fmt.Println(address.Hex())
	// 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
	//0x4fb732257ffd3a78de59b393979699b06fd2d6b60c8c7d43e746a952b325bd74
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
