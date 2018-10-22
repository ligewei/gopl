package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("hello, world")
	var i int
	fmt.Println(unsafe.Sizeof(i))
}
