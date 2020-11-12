package main

import (
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			println("Defer")
			wg.Done()
		}()
		runtime.Goexit()
		println("never run ")
	}()
	wg.Wait()
}
