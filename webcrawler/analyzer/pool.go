package analyzer

// 分析器池接口类型
type AnalyzerPool interface {
	Take() (Analyzer, error)        // 从池中取出一个分析器
	Return(analyzer Analyzer) error // 将分析器放回池中
	Total() uint32                  // 获得池子总量
	Used() uint32                   // 获得正在使用的分析器总量
}
