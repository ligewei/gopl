package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "regli")
	sum := h.Sum(nil)
	fmt.Printf("sum hex: %x\n", sum)

	// hex.Encode，把数字1变成'1'(31)，2变成'2'(32) ...
	// 2f8053139278aa952172a32a9da94471，逐个encode
	hexText := make([]byte, 32)
	hex.Encode(hexText, sum)

	fmt.Println("sum dec:", sum)
	fmt.Println("sum str:", string(sum))
	fmt.Println("b16 dec:", hexText)
	fmt.Printf("b16 hex: %x\n", hexText)
	fmt.Println("sum str:", string(hexText))

}
