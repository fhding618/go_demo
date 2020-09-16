/**
 * sync.map
 * 初始Store数据到dirty，加锁
 * 初始Load则从dirty中读取数据，加锁，且miss+1
 * 当多次读累加的Miss数不小于len时，将dirty数据加载到read
 * 当读足够多时，数据读写全部在原子操作read中完成，极少需要Mutex锁，极大的提高性能
 * sync.map主要应用于读远大于写的情况
 * 若是读与写的数量差不多，则与带锁的map性能差异并不会太明显
 *
 * https://studygolang.com/articles/30385?fr=sidebar
 **/

package main

import (
	"fmt"
	"sync"
)

var sm sync.Map

func main() {

	sm.Store("key01", 10)         // 加锁，存储到dirty
	sm.Store("key01", 100)        // 加锁，存储到dirty
	sm.Store("key02", "ab")       // 加锁，存储到dirty
	fmt.Println(sm.Load("key01")) // 加锁，读取dirty， miss+1
	fmt.Println(sm.Load("key02")) // 加锁，读取dirty， miss+1, dirty数据copy到read，dirty清空
	sm.Store("key03", 1000)       // 加锁，新key，存至dirty
	sm.Load("key02")              // 不加锁，直接从read读取
	sm.Delete("key03")
	fmt.Println(sm.Load("key01"))

}
