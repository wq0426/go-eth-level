package localsha256

import (
	"crypto/sha256"
	"fmt"
)

func Sha256() {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
