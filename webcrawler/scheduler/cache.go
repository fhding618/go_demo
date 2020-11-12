package scheduler

import (
	"fmt"
	"sync"
	"webcrawler/base"
)

var statusMap = map[byte]string{
	0: "running",
	1: "closed",
}

// 缓存接口
type requestCache interface {
	put(req *base.Request) bool
	get() *base.Request
	capacity() int
	length() int
	close()
	summary() string
}

// 实现类
type reqCacheBySlice struct {
	cache  []*base.Request // 请求的存储介质
	mutex  sync.Mutex      // 互斥锁
	status byte            // 缓存状态。0表示正在运行，1表示已关闭
}

// 实例化
func newRequestCache() requestCache {
	return &reqCacheBySlice{
		cache: make([]*base.Request, 0),
	}
}

func (rcache *reqCacheBySlice) put(req *base.Request) bool {
	if req == nil {
		return false
	}
	if rcache.status == 1 {
		return false
	}
	rcache.mutex.Lock()
	defer rcache.mutex.Unlock()
	rcache.cache = append(rcache.cache, req)
	return true
}

func (rcache *reqCacheBySlice) get() *base.Request {
	if len(rcache.cache) == 0 {
		return nil
	}
	if rcache.status == 1 {
		return nil
	}
	rcache.mutex.Lock()
	defer rcache.mutex.Unlock()
	req := rcache.cache[0]
	rcache.cache = rcache.cache[1:]
	return req
}

func (rcache *reqCacheBySlice) capacity() int {
	return cap(rcache.cache)
}

func (rcache *reqCacheBySlice) length() int {
	return len(rcache.cache)
}

func (rcache *reqCacheBySlice) close() {
	if rcache.status == 1 {
		return
	}
	rcache.status = 1
	return
}

var summaryTemplate = "status: %s, " + "length: %d, " + "capacity: %d"

func (rcache *reqCacheBySlice) summary() string {
	summary := fmt.Sprintf(summaryTemplate,
		statusMap[rcache.status],
		rcache.length(),
		rcache.capacity())
	return summary
}
