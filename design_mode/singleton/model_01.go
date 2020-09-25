/**
 * 单例模式
 * 饿汉式
 * singleton struct 小写，不允许导出
 * 使用：
 * c := singleton.Instance.Add()
 **/

package singleton

type singleton_01 struct {
	count int
}

var Instance = new(singleton_01)

func (s *singleton_01) Add() int {
	s.count++
	return s.count
}
