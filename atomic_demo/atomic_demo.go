/**
 * 原子操作没有锁，基本是在硬件层面实现的，经常被用来实现其他同步技术
 **/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 并发原子操作：加 1
func add_demo() {
	var n int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("%v\n", atomic.LoadInt32(&n))
}

type Page struct {
	views uint32
}

//setter getter
func (page *Page) SetViews(n uint32) {
	atomic.StoreUint32(&page.views, n)
}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}

// 原子加减操作
// 对无符号变量v， -v是合法的，可以把 -v 传给 AddT 的第二个参数
// 对应常量整数c, -c作为AddT的第二个参数是非法的，但可以使用^T(c-1)作为AddT的第二个参数
func count_demo() {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	atomic.AddUint64(&n, -m)
	fmt.Println(n) // 97-1 = 96

	atomic.AddUint64(&n, -uint64(k))
	fmt.Println(n) // 96-2=94

	atomic.AddUint64(&n, ^uint64(a-1))
	fmt.Println(n) // 94-3=91

	atomic.AddUint64(&n, ^(b - 1))
	fmt.Println(n) // 91-4=87

	atomic.AddUint64(&n, ^uint64(c-1))
	fmt.Println(n) // 87-5=82

	atomic.AddUint64(&n, ^uint64(d-1))
	fmt.Println(n) // 82-6=76

	x := b
	atomic.AddUint64(&n, -x)
	fmt.Println(n) // 76-4=72

	atomic.AddUint64(&n, ^(m - 1))
	fmt.Println(n) // 72-1=71

	atomic.AddUint64(&n, ^uint64(k-1))
	fmt.Println(n) // 71-2=69
}

// swap CompareAndSwapT Demo
func swap_demo() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old)                                   // 789 123
	fmt.Println(atomic.CompareAndSwapInt64(&n, 123, 456)) // false
	fmt.Println(n)                                        // 789
	fmt.Println(atomic.CompareAndSwapInt64(&n, 789, 456)) // true
	fmt.Println(n)                                        // 456
}

// atomic.Value Demo
// 对于可寻址的Value 值v，一旦v.Store被调用过，接下来的v.Store必须采用同样类型的参数作为参数，否则将会panic
// 用nil作为参数也会导致panic
func val_demo() {
	type T struct {
		a, b, c int
	}
	var ta = T{1, 2, 3}
	var v atomic.Value
	v.Store(ta)
	var tb = v.Load().(T)
	fmt.Println(tb)
	fmt.Println(ta == tb)
	v.Store("hello")
}

func main() {
	val_demo()
}
