/**
 * 通过它，可以从底层的 io.Reader 中更大批量的读取数据，这会使底层读取操作变少
 * 如果数据读取时的块数量是固定合适的，底层媒体设备将会有更好的表现，也因此会提高程序的性能
 * io.Reader --> buffer --> consumer
 *
 * Func:
 *	Peek: 方法可以查看缓存的前 n 个字节而不会真的『吃掉』它
 *  Reset:
 *  Discard: 这个方法将会丢弃 n 个字节的，返回时也不会返回被丢弃的 n 个字节
 *	Read:
 *	{Read, Unread}Byte
 *	{Read, Unread}Rune
 *	ReadSlice
 *	ReadBytes
 *	ReadString
 *	ReadLine
 *	WriteTo
 **/

package bufio_func

import (
	"bufio"
	"fmt"
	"strings"
)

/**
 * Peek:
 * - 如果缓存不满，而且缓存中缓存的数据少于 n 个字节，其将会尝试从 io.Reader 中读取
 * - 如果请求的数据量大于缓存的容量，将会返回 bufio.ErrBufferFull
 * - 如果 n 大于流的大小，将会返回 EOF
 **/
func Peek_func() {
	s1 := strings.NewReader(strings.Repeat("a", 20))
	r := bufio.NewReaderSize(s1, 16)
	b, err := r.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b) // output: aaa
	b, err = r.Peek(17)   // 请求的数据量大于缓存的容量
	if err != nil {
		fmt.Println(err) // output: bufio: buffer full
	}
	fmt.Printf("%q\n", b)
	s2 := strings.NewReader("aaa")
	r.Reset(s2)
	b, err = r.Peek(10) // 如果 n 大于流的大小
	if err != nil {
		fmt.Println(err) // output: EOF
	}
	fmt.Printf("%q\n", b)

	// 返回的切片和被 bufio.Reader 使用的内部缓存底层使用相同的数组
	// 因此引擎底层在执行任何读取操作之后内部返回的切片将会变成无效的
	// 这是由于其将有可能被其他的缓存数据覆盖
	fmt.Printf("\n\n bufio 与 返回切片Case:\n")
	s3 := strings.NewReader(strings.Repeat("a", 16) + strings.Repeat("b", 16))
	r3 := bufio.NewReaderSize(s3, 16)
	b3, _ := r3.Peek(3)
	fmt.Printf("%q\n", b3)
	r3.Read(make([]byte, 16))
	r3.Read(make([]byte, 15))
	fmt.Printf("%q\n", b3)
}
