/**
 * 懒汉式
 *
 **/

package singleton

import "sync"

type singleton_02 struct {
	count int
}

var (
	instance *singleton_02
	mutex    sync.Mutex
)

// 没有「双重检查」
// 每次Lock性能损耗
func New_01() *singleton_02 {
	mutex.Lock()
	if instance == nil {
		instance = new(singleton_02)
	}
	mutex.Unlock()

	return instance
}

// 增加「双重检查」，相较于每次Lock，性能会有很大提升
// 避免多个goroutine 同时走到Lock，后分别初始化了自己的instance实例
func New_02() *singleton_02 {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(singleton_02)
		}
		mutex.Unlock()
	}
	return instance
}

func (s *singleton_02) Add() int {
	s.count++
	return s.count
}

var once sync.Once

func New_03() *singleton_02 {
	once.Do(func() {
		instance = new(singleton_02)
	})

	return instance
}
