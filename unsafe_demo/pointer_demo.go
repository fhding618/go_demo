/**
 * 通过将任意变量转换为interface{}
 * 将interface断言为ifaceWords
 * 可以实现将任意变量存储到标准Value中的功能
 * 参考 sync.map
 **/

package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func main() {
	var num uint32 = 32
	var numI interface{}

	numI = num
	fmt.Printf("%v\n", numI)

	numIface := (*ifaceWords)(unsafe.Pointer(&numI))

	fmt.Printf("%v\n", *(*uint32)(atomic.LoadPointer(&numIface.data)))
}
