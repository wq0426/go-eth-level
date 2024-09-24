package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	/*
		将 interface 智能合约编译为 JSON ABI，并使用 abigen 从 ABI 创建 Go 包
			安装 Solidity 编译器 (solc).
			npm install -g solc
			执行命令
			solc --abi erc20.sol

			安装abigen，从solidity 智能合约生成 ABI
			go install github.com/ethereum/go-ethereum/cmd/abigen@latest
			执行命令
			abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go

		将新的_token_包导入我们的应用程序
		1、将文件放到C:\Program Files\Go\src\erc20目录下(GOPATCH)
		2、将文件上传到github，类似上面导入的写法"github.com/youname/project/contracts_erc20"
		3、根据go.mod里面的module名加上目录，moduleName/directory
	*/
	token "goethereum_example/goethereum_example/contracts/erc20" // for demo (用3的方法)
)

// 账户代币余额
func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	//将新的_token_实例化
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//我们现在可以调用任何 ERC20 的方法。 例如，我们可以查询用户的代币余额。
	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	//我们还可以读 ERC20 智能合约的公共变量。
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	//我们可以做一些简单的数学运算将余额转换为可读的十进制格式。
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
