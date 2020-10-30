package downloader

import (
	"log"
	"net/http"
	base "webcrawler/base"
	mdw "webcrawler/middleware"
)

// ID 生成器
var downloaderIdGenerator mdw.IdGenerator = mdw.NewIdGenerator()

func genDownloaderId() uint32 {
	return downloaderIdGenerator.GetUint32()
}

// 网页下载器接口类型
type PageDownloader interface {
	Id() uint32                                        // 获取ID
	Download(req base.Request) (*base.Response, error) // 下载网页
}

// 实现类
type myPageDownloader struct {
	id         uint32      // ID
	httpClient http.Client // HTTP 客户端
}

// 创建下载器
func NewPageDownloader(client *http.Client) PageDownloader {
	id := genDownloaderId()
	if client == nil {
		client = &http.Client{}
	}
	return &myPageDownloader{
		id:         id,
		httpClient: *client,
	}
}

func (dl *myPageDownloader) Id() uint32 {
	return dl.id
}

func (dl *myPageDownloader) Download(req base.Request) (*base.Response, error) {
	httpReq := req.HttpReq()
	log.Printf("Do the request (url=%s)... \n", httpReq.URL)
	httpResp, err := dl.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return base.NewResponse(httpResp, req.Depth()), nil
}
