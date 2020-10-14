package itempipeline

import "webcrawler/base"

// 条目处理管道接口类型
type ItemPipeline interface {
	// 发送条目
	Send(item base.Item) []error
	// FailFast方法返回一个布尔值。该值表式当前的条目处理管道是否是快速失败的。
	// 这里的快速失败是指：只要对某个条目的处理流程在某一个步骤上出错，
	// 那么条目处理管道就会忽略掉后续的所有处理步骤并报告错误。
	FailFast() bool
	// 设置是否快速失败
	SetFailFase(failFast bool)
	// 获得已发送、已接受和已处理的条目的计数值
	// 更确切的说，作为结果值的切片总会有三个元素值。这三个值会分别代表前述的三个计数
	Count() []uint64
	// 获取正在被处理的条目数量
	ProcessingNumber() uint64
	// 获取摘要信息
	Summary() string
}
