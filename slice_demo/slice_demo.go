/**
 * 1. 切片的长度和容量不同，长度表示左指针至右指针之间的距离，容量表示左指针至底层数组末尾的距离
 * 2. 切片的扩容机制，append的时候，如果长度增加后超过容量，则将容量增加2倍，同时变换了底层数组
 * 3. 切片append机制，是把slice后边数据逐个覆盖写入
 * 4. 切片copy机制，按其中较小的那个数组切片的元素个数进行复制
 * 5. Go数组是值类型，赋值和传参操作都会复制整个数组数据。（slice 传递的是地址）
 **/

package main

import (
	"fmt"
	"unsafe"
)

// 将数据追加到slice尾部，必要的话增加容量
// 实现了容量的完全控制
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		fmt.Printf("slice ptr=%p newSlice ptr=%p\n", slice, newSlice)
		slice = newSlice
		fmt.Printf("slice ptr=%p newSlice ptr=%p\n", slice, newSlice)
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// 检查与 slice 关联的内存
func InspectSlice(slice []string) {
	// Capture the address to the slice structure
	address := unsafe.Pointer(&slice)
	addrSize := unsafe.Sizeof(address)

	// Capture the address where the length and cap size is stored
	lenAddr := uintptr(address) + addrSize
	capAddr := uintptr(address) + (addrSize * 2)

	// Create pointers to the length and cap size
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}

	fmt.Printf("\n\n")
}

func slice_demo() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := []int{10, 20, 5: 30}
	s4 := []int{10: 20}
	slice := x[:2:5]

	fmt.Printf("len:%d cap:%d val=%v\n", len(s3), cap(s3), s3)
	fmt.Printf("len:%d cap:%d val=%v\n", len(s4), cap(s4), s4)
	fmt.Printf("len:%d cap:%d val=%v\n", len(x[:]), cap(x[:]), x[:])
	fmt.Printf("len:%d cap:%d val=%v\n", len(slice), cap(slice), slice)
	fmt.Printf("x_ptr=%p, slice_ptr=%p\n", x[:], slice)
}

func main() {
	// InspectSlice Demo
	orgSlice := make([]string, 5)
	orgSlice[0] = "Apple"
	orgSlice[1] = "Orange"
	orgSlice[2] = "Banana"
	orgSlice[3] = "Grape"
	orgSlice[4] = "Plum"

	InspectSlice(orgSlice)

	// AppendByte Demo
	p := []byte{1, 2, 3}
	p = AppendByte(p, 4, 5, 6)
	fmt.Printf("len=%d cap=%d ptr=%p val=%v", len(p), cap(p), p, p)

	// Slice Demo
	slice_demo()
}
