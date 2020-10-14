package base

import (
	"bytes"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

// 错误类型
type ErrorType string

// 错误类型常量
const (
	DOWNLOADER_ERROR     ErrorType = "Downloader Error"
	ANALYZER_ERROR       ErrorType = "Analyzer Error"
	ITEM_PROCESSOR_ERROR ErrorType = "Item Processor Error"
)

// 爬虫错误接口
type CrawlerError interface {
	Type() ErrorType // 获得错误类型
	Error() string   // 获得错误信息
}

// 爬虫错误的实现
type myCrawlerError struct {
	errType    ErrorType // 错误类型
	errMsg     string    // 错误信息
	fullErrMsg string    // 完整错误提示信息
}

// 创建一个新的爬虫错误
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}

// 获取错误类型
func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

// 获取错误信息
func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

// 生成错误提示信息
func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler Error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprint("%s\n", buffer.String())
	return
}
