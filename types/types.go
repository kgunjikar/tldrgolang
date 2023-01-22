package types

import (
	"fmt"
	"unsafe"
)

func printSz() {
	var a int
	var b int32
	var c int64
	var d float32
	var e float64
	var f rune
	var g byte
	var h uintptr
	var i uint
	var j uint16
	var k uint32
	var l uint64

	fmt.Println(unsafe.Sizeof(1), unsafe.Sizeof("c"), unsafe.Sizeof('c'), unsafe.Sizeof(1.0))
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d), unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f), unsafe.Sizeof(g), unsafe.Sizeof(h), unsafe.Sizeof(i), unsafe.Sizeof(j), unsafe.Sizeof(k), unsafe.Sizeof(l))
}
