package itempipeline

import "webcrawler/base"

// 被用来处理条目的函数类型
type ProcessItem func(item base.Item) (result base.Data, err error)
