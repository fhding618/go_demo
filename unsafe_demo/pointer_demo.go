/**
 * 通过将任意变量转换为interface{}
 * 将interface断言为ifaceWords
 * 可以实现将任意变量存储到标准Value中的功能
 * 参考 sync.map
 **/

package main

import (
	"fmt"
	"unsafe"
)

type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

type sizeStruct struct {
	x int8
	s string
}
type size01 struct {
	x bool
	y float64
	z int16
}
type size02 struct {
	x float64
	y int16
	z bool
}

func size_demo() {
	tmp := sizeStruct{
		0,
		"x",
	}
	size01 := size01{
		false,
		0,
		0,
	}
	size02 := size02{
		0,
		0,
		false,
	}
	fmt.Println(unsafe.Sizeof(tmp))
	fmt.Println(unsafe.Sizeof(tmp.x))
	fmt.Println(unsafe.Sizeof(tmp.s))
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Printf("size01 %v\n", unsafe.Sizeof(size01))
	fmt.Printf("size02 %v\n", unsafe.Sizeof(size02))

	fmt.Printf("x align:%v offset:%v\n", unsafe.Alignof(tmp.x), unsafe.Offsetof(tmp.x))
	fmt.Printf("s align:%v offset:%v\n", unsafe.Alignof(tmp.s), unsafe.Offsetof(tmp.s))

	fmt.Printf("%p\n", &tmp.x)
	fmt.Printf("%p\n", &tmp.s)
}

// 地址偏移计算
type num struct {
	i string
	j int64
}

func offset_demo() {
	n := num{i: "EMPTY", j: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "中文"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("%v\n", n)
}

func main() {
	size_demo()

	/*
		var num uint32 = 32
		var numI interface{}

		numI = num
		fmt.Printf("%v\n", numI)

		numIface := (*ifaceWords)(unsafe.Pointer(&numI))

		fmt.Printf("%v\n", *(*uint32)(atomic.LoadPointer(&numIface.data)))*/
}
