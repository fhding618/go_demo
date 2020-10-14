package downloader

import (
	base "webcrawler/base"
)

// 网页下载器接口类型
type PageDownloader interface {
	ID() uint32                                        // 获取ID
	Download(req base.Request) (*base.Response, error) // 下载网页
}
