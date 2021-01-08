package scheduler

import (
	"bytes"
	"fmt"
	"webcrawler/base"
)

// 调度器摘要信息的接口类型
type SchedSummary interface {
	String() string               // 获得摘要信息的一般表示
	Detail() string               // 获得摘要信息的详细表示
	Same(other SchedSummary) bool // 判断是否与另一份摘要信息相同
}

// 调度器摘要信息类
type mySchedSummary struct {
	prefix              string
	running             uint32
	channelArgs         base.ChannelArgs
	poolBaseArgs        base.PoolBaseArgs
	crawlDepth          uint32
	chanmanSummary      string
	reqCacheSummary     string
	dlPoolLen           uint32
	dlPoolCap           uint32
	analyzerPoolLen     uint32
	analyzerPoolCap     uint32
	itemPipelineSummary string
	urlCount            int
	urlDetail           string
	stopSignSummary     string
}

// 创建调度器摘要信息
func NewSchedSummary(sched *myScheduler, prefix string) SchedSummary {
	if sched == nil {
		return nil
	}
	urlCount := len(sched.urlMap)
	var urlDetail string
	if urlCount > 0 {
		var buffer bytes.Buffer
		buffer.WriteByte('\n')
		sched.rw.RLock()
		for k, _ := range sched.urlMap {
			buffer.WriteString(prefix)
			buffer.WriteString(prefix)
			buffer.WriteString(k)
			buffer.WriteByte('\n')
		}
		sched.rw.RUnlock()
		urlDetail = buffer.String()
	} else {
		urlDetail = "\n"
	}
	return &mySchedSummary{
		prefix:              prefix,
		running:             sched.running,
		channelArgs:         sched.channelArgs,
		poolBaseArgs:        sched.poolBaseArgs,
		crawlDepth:          sched.crawlDepth,
		chanmanSummary:      sched.chanman.Summary(),
		reqCacheSummary:     sched.reqCache.summary(),
		dlPoolCap:           sched.dlpool.Total(),
		dlPoolLen:           sched.dlpool.Used(),
		analyzerPoolCap:     sched.analyzerPool.Total(),
		analyzerPoolLen:     sched.analyzerPool.Used(),
		itemPipelineSummary: sched.itemPipeline.Summary(),
		urlCount:            urlCount,
		urlDetail:           urlDetail,
		stopSignSummary:     sched.stopSign.Summary(),
	}
}

// 获得摘要信息的一般表示
func (ss *mySchedSummary) String() string {
	return ss.getSummary(false)
}

// 获得摘要信息的详细表示
func (ss *mySchedSummary) Detail() string {
	return ss.getSummary(true)
}

// 判断是否与另一份摘要信息相同
func (ss *mySchedSummary) Same(other SchedSummary) bool {
	if other == nil {
		return false
	}
	otherSs, ok := interface{}(other).(*mySchedSummary)
	if !ok {
		return false
	}
	if ss.running != otherSs.running ||
		ss.crawlDepth != otherSs.crawlDepth ||
		ss.dlPoolLen != otherSs.dlPoolLen ||
		ss.dlPoolCap != otherSs.dlPoolCap ||
		ss.analyzerPoolLen != otherSs.analyzerPoolLen ||
		ss.analyzerPoolCap != otherSs.analyzerPoolCap ||
		ss.urlCount != otherSs.urlCount ||
		ss.stopSignSummary != otherSs.stopSignSummary ||
		ss.reqCacheSummary != otherSs.reqCacheSummary ||
		ss.poolBaseArgs.String() != otherSs.poolBaseArgs.String() ||
		ss.channelArgs.String() != otherSs.channelArgs.String() ||
		ss.itemPipelineSummary != otherSs.itemPipelineSummary ||
		ss.chanmanSummary != otherSs.chanmanSummary {
		return false
	} else {
		return true
	}
}

// 获取摘要信息
func (ss *mySchedSummary) getSummary(detail bool) string {
	prefix := ss.prefix
	template := prefix + "Running: %v \n" +
		prefix + "Channel args: %s \n" +
		prefix + "Pool base args: %s \n" +
		prefix + "Crawl depth: %d \n" +
		prefix + "Channels manager: %s \n" +
		prefix + "Request cache: %s \n" +
		prefix + "Downloader pool: %d/%d \n" +
		prefix + "Analyzer pool: %d/%d\n" +
		prefix + "Item pipeline: %s \n" +
		prefix + "Url(%d): %s" +
		prefix + "Stop sign: %s\n"
	return fmt.Sprintf(template,
		func() bool {
			return ss.running == 1
		}(),
		ss.channelArgs.String(),
		ss.poolBaseArgs.String(),
		ss.crawlDepth,
		ss.chanmanSummary,
		ss.reqCacheSummary,
		ss.dlPoolLen, ss.dlPoolCap,
		ss.analyzerPoolLen, ss.analyzerPoolCap,
		ss.itemPipelineSummary,
		ss.urlCount,
		func() string {
			if detail {
				return ss.urlDetail
			} else {
				return "<concealed>\n"
			}
		}(),
		ss.stopSignSummary)
}
