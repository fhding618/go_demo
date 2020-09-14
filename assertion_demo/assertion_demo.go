package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	data unsafe.Pointer
	typz unsafe.Pointer
}

type readOnly struct {
	m       map[interface{}]int
	amended bool // true if the dirty map contains some key not in m.
}

func main() {
	var x interface{}

	xr, ok := x.(readOnly)
	println(ok)
	fmt.Printf("%v\n", xr)
}
