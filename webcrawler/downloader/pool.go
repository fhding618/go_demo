package downloader

// 网页下载器池接口类型
type PageDownloaderPool interface {
	Take() (PageDownloader, error)  // 从池中取出一个下载器
	Return(dl PageDownloader) error // 将一个下载器放回池中
	Total() uint32                  // 获取池子总量
	Used() uint32                   // 获取使用中的下载器数量
}
