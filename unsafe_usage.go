package main

import (
	"fmt"
	"unsafe"
)

type st struct {
	i32 int32
	i64 int64
}

type stst struct {
	s st
	i8 int8
}

func main() {
	i := 0
	fmt.Println(unsafe.Sizeof(i))

	var s st
	fmt.Println(unsafe.Sizeof(s))

	var ss stst
	fmt.Println(unsafe.Sizeof(ss))
}
