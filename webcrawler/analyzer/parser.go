package analyzer

import (
	"net/http"
	"webcrawler/base"
)

// 用于解析HTTP响应的函数类型
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)
