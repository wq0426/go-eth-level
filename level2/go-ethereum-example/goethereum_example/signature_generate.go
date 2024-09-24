package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 生成签名
//数字签名允许不可否认性，因为这意味着签署消息的人必须拥有私钥，来证明消息是真实的。
//任何人都可以验证消息的真实性，只要它们具有原始数据的散列和签名者的公钥即可。
//签名是区块链的基本组成部分，我们将在接下来的几节课中学习如何生成和验证签名。

// 用于生成签名的组件是：签名者私钥，以及将要签名的数据的哈希。 只要输出为 32 字节，就可以使用任何哈希算法。
// 我们将使用 Keccak-256 作为哈希算法，这是以太坊常常使用的算法。
func main() {
	//首先，我们将加载私钥。
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	//接下来我们将获取我们希望签名的数据的 Keccak-256，在这个例子里，它将是_hello_。
	//go-ethereum crypto 包提供了一个方便的 Keccak256Hash 方法来实现这一目的。
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	//最后，我们使用私钥签名哈希，得到签名。
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
}
