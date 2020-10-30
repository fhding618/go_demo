package downloader

import (
	"errors"
	"fmt"
	"reflect"
	mdw "webcrawler/middleware"
)

type GenPageDownloader func() PageDownloader

// 网页下载器池接口类型
type PageDownloaderPool interface {
	Take() (PageDownloader, error)  // 从池中取出一个下载器
	Return(dl PageDownloader) error // 将一个下载器放回池中
	Total() uint32                  // 获取池子总量
	Used() uint32                   // 获取使用中的下载器数量
}

type myDownloaderPool struct {
	pool  mdw.Pool
	etype reflect.Type
}

func NewPageDownloaderPool(
	total uint32,
	gen GenPageDownloader) (PageDownloaderPool, error) {
	etype := reflect.TypeOf(gen())
	genEntity := func() mdw.Entity {
		return gen()
	}
	pool, err := mdw.NewPool(total, etype, genEntity)
	if err != nil {
		return nil, err
	}
	dlpool := &myDownloaderPool{
		pool:  pool,
		etype: etype,
	}
	return dlpool, nil
}

// 从池中取出一个下载器
func (dlpool *myDownloaderPool) Take() (PageDownloader, error) {
	entity, err := dlpool.pool.Take()
	if err != nil {
		return nil, err
	}
	dl, ok := entity.(PageDownloader)
	if !ok {
		errMsg := fmt.Sprintf("The type of entity is NOT %s!\n", dlpool.etype)
		return nil, errors.New(errMsg)
	}
	return dl, nil
}

// 将一个下载器放回池中
func (dlpool *myDownloaderPool) Return(dl PageDownloader) error {
	return dlpool.pool.Return(dl)
}

// 获取池子总量
func (dlpool *myDownloaderPool) Total() uint32 {
	return dlpool.pool.Total()
}

// 获取使用中的下载器数量
func (dlpool *myDownloaderPool) Used() uint32 {
	return dlpool.pool.Used()
}
