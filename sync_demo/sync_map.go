package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

var sm sync.Map

func main()  {
	var p unsafe.Pointer
	data := "abc"
	p = unsafe.Pointer(&data)

	atomic.StorePointer(&p, unsafe.Pointer(&data))
	fmt.Printf("value is %s", *(*string)(atomic.LoadPointer(&p)))
}
