package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

//发送原始交易事务

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b33bf7f45ab84ffcb357517d3b433ca4")
	if err != nil {
		log.Fatal(err)
	}

	//首先将原始事务十六进制解码为字节格式。
	rawTx := "f8683e8503a3331a2d825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d01808401546d71a0eb333683a779ce4f462ca556195b29df04f4b437e2f91ad0056583835554a91aa03c999fe1e540d942391e6a1349ccb70ccea4698afe40120b2cb3af935c684b2f"

	rawTxBytes, err := hex.DecodeString(rawTx)

	//接下来初始化一个新的 types.Transaction 指针并从 go-ethereum rlp 包中调用 DecodeBytes，
	//将原始事务字节和指针传递给以太坊事务类型。 RLP 是以太坊用于序列化和反序列化数据的编码方法。
	tx := new(types.Transaction)

	rlp.DecodeBytes(rawTxBytes, &tx)

	//现在，我们可以使用我们的以太坊客户端轻松地广播交易。
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
