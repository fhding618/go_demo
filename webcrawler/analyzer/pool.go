package analyzer

import (
	"errors"
	"fmt"
	"reflect"
	mdw "webcrawler/middleware"
)

// 生成分析器的函数类型
type GenAnalyzer func() Analyzer

// 分析器池接口类型
type AnalyzerPool interface {
	Take() (Analyzer, error)        // 从池中取出一个分析器
	Return(analyzer Analyzer) error // 将分析器放回池中
	Total() uint32                  // 获得池子总量
	Used() uint32                   // 获得正在使用的分析器总量
}

// 实体类
type myAnalyzerPool struct {
	pool  mdw.Pool     // 实体池
	etype reflect.Type // 池内实体类型
}

// 实例化
func NewAnalyzerPool(
	total uint32,
	gen GenAnalyzer) (AnalyzerPool, error) {
	etype := reflect.TypeOf(gen())
	genEntity := func() mdw.Entity { return gen() }
	pool, err := mdw.NewPool(total, etype, genEntity)
	if err != nil {
		return nil, err
	}
	dlpool := &myAnalyzerPool{pool: pool, etype: etype}
	return dlpool, nil
}

// 从池中取出一个分析器
func (spdpool *myAnalyzerPool) Take() (Analyzer, error) {
	entity, err := spdpool.pool.Take()
	if err != nil {
		return nil, err
	}
	analyzer, ok := entity.(Analyzer)
	if !ok {
		errMsg := fmt.Sprintf("The type of entity is NOT %s!\n", spdpool.etype)
		panic(errors.New(errMsg))
	}
	return analyzer, nil
}

// 将分析器放回池中
func (spdpool *myAnalyzerPool) Return(analyzer Analyzer) error {
	return spdpool.pool.Return(analyzer)
}

// 获得池子总量
func (spdpool *myAnalyzerPool) Total() uint32 {
	return spdpool.pool.Total()
}

// 获得正在使用的分析器总量
func (spdpool *myAnalyzerPool) Used() uint32 {
	return spdpool.pool.Used()
}
