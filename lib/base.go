package lib

import (
	"time"
)

// 调用结果的结构
type CallResult struct {
	ID     int64
	Req    RawReq
	Resp   RawResp
	Code   RetCode
	Msg    string
	Elapse time.Duration
}

// 请求结构
type RawReq struct {
	ID  int64
	Req []byte
}

// 响应结构
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

// 结果代码的类型
type RetCode int

// 声明代表载荷发生器状态的常量
const (
	STATUS_ORIGINAL uint32 = iota
	STATUS_STARTING
	STATUS_STARTED
	STATUS_STOPPING
	STATUS_STOPPED
)

// 载荷发生器的接口
type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}

const (
	RET_CODE_SUCCESS              RetCode = 0
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // 调用超时警告
	RET_CODE_ERROR_CALL                   = 2001 // 调用错误
	RET_CODE_ERROR_RESPONSE               = 2002 // 响应内容错误
	RET_CODE_ERROR_CALEE                  = 2003 // 被动用方的内部错误
	RET_CODE_FATAL_CALL                   = 3001 // 调用过程中发生了致命错误
)

func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "Success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}
