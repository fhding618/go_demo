/**
 * bufio 用于优化频繁的写入操作
 * 频繁的写入容易造成CPU已经硬盘的资源浪费
 * bufio 用于将写入进行缓存成整个的数据块，再批量进行写入
 * producer --> buffer --> io.Writer
 *
 * 方法：
 * 	Reset
 * 	使用 Reset 方法，Writer 可以用于不同的目的对象
 * 	重复使用 Writer 缓存减少了内存的分配
 * 	而且减少了额外的垃圾回收工作
 *
 *	Buffered: 已用空间
 *  ReadFrom
 * https://zhuanlan.zhihu.com/p/129781512
 **/

package bufio_func

import (
	"bufio"
	"fmt"
	"strings"
)

type writer int

func (*writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}

func bufio_demo() {
	fmt.Println("Unbuffered I/O")
	w := new(writer)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})

	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
}

/**
 * Reset 方法 Demo
 **/
type writer1 int

func (*writer1) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#1: %q\n", p)
	return len(p), nil
}

type writer2 int

func (*writer2) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#2:%q\n", p)
	return len(p), nil
}

func reset_demo() {
	w1 := new(writer1)
	bw := bufio.NewWriterSize(w1, 2)
	fmt.Println(bw.Available()) //检测缓存剩余空间
	bw.Write([]byte("ab"))
	bw.Write([]byte("cd"))
	w2 := new(writer2)
	bw.Reset(w2) // 在Reset前未Flush，"cd"会丢失
	bw.Write([]byte("ef"))
	bw.Flush()
}

/**
 * ReadFrom 方法
 **/
type readWriter int

func (*readWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%q\n", p)
	return len(p), nil
}
func readfrom_demo() {
	s := strings.NewReader("onetwothree")
	w := new(readWriter)
	bw := bufio.NewWriterSize(w, 3)
	bw.ReadFrom(s)
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
}
