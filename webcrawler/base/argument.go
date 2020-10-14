package base

import (
	"errors"
	"fmt"
)

// 参数容器的接口
type Args interface {
	// 自检参数的有效性，并在必要时返回可以说明问题的错误值
	// 若结果值为nil，则说明未发现问题，否则就意味着自检未通过
	Check() error
	// 获得参数容器的字符串表现形式
	String() string
}

// 通道参数的容器的描述模板
var chanelArgsTemplate string = "{ reqChanLen: %d, respChanLen: %d," +
	" itemChanLen: %d, errorChanLen: %d }"

// 通道参数的容器。
type ChannelArgs struct {
	reqChanLen   uint   // 请求通道的长度
	respChanLen  uint   // 响应通道的长度
	itemChanLen  uint   // 条目通道的长度
	errorChanLen uint   // 错误通道的长度
	description  string // 描述
}

// 创建通道参数的容器
func NewChannelArgs(
	reqChanLen uint,
	respChanLen uint,
	itemChanLen uint,
	errorChanLen uint) ChannelArgs {
	return ChannelArgs{
		reqChanLen:   reqChanLen,
		respChanLen:  respChanLen,
		itemChanLen:  itemChanLen,
		errorChanLen: errorChanLen,
	}
}

func (args *ChannelArgs) Check() error {
	if args.reqChanLen == 0 {
		return errors.New("The request channel max length (capacity) can not be 0!\n")
	}
	if args.respChanLen == 0 {
		return errors.New("The response channel max length (capacity) can not be 0!\n")
	}
	if args.itemChanLen == 0 {
		return errors.New("The item channel max length (capacity) can not be 0!\n")
	}
	if args.errorChanLen == 0 {
		return errors.New("The error channel max length (capacity) can not be 0!\n")
	}
	return nil
}

func (args *ChannelArgs) String() string {
	if args.description == "" {
		args.description = fmt.Sprintf(chanelArgsTemplate,
			args.reqChanLen,
			args.respChanLen,
			args.itemChanLen,
			args.errorChanLen)
	}
	return args.description
}

func (args *ChannelArgs) ReqChanLen() uint {
	return args.respChanLen
}

func (args *ChannelArgs) RespChanLen() uint {
	return args.respChanLen
}

func (args *ChannelArgs) ItemChanLen() uint {
	return args.itemChanLen
}

func (args *ChannelArgs) ErrorChanLen() uint {
	return args.errorChanLen
}
