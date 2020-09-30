/**
 * Go 的结构体中可以使用一种叫做内嵌的类型
 * 和常规的具有类型和名字的字段不同，我们可以仅仅使用类型(匿名字段)。
 * 内嵌类型的方法或者字段如果不和其他的冲突的话，则可以使用一个简短的选择器来引用
 *
 * 包 bufio 使用内嵌的方式来定义 ReadWriter。 它由 Reader 和 Writer 构成:
 * type ReadWriter struct {
 *		*Reader
 *		*Writer
 * }
 **/

package bufio_func

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type T1 struct {
	t1 string
}

func (t1 *T1) f1() {
	fmt.Println("T1.f1")
}

type T2 struct {
	t2 string
}

func (t2 *T2) f2() {
	fmt.Printf("T2.f2")
}

type U struct {
	*T1
	*T2
}

//可以简单的使用 u.t1 来代替 u.T1.t1
func t1_t2_demo() {
	u := U{T1: &T1{"foo"}, T2: &T2{"bar"}}
	u.f1()
	u.f2()
	fmt.Println(u.t1)
	fmt.Println(u.t2)
}

func readwriter_func() {
	s := strings.NewReader("abcd")
	br := bufio.NewReader(s)
	w := new(bytes.Buffer)
	bw := bufio.NewWriter(w)
	rw := bufio.NewReadWriter(br, bw)
	buf := make([]byte, 2)
	_, err := rw.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf)
	buf = []byte("efgh")
	_, err = rw.Write(buf)
	if err != nil {
		panic(err)
	}
	err = rw.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(w.String())
}
