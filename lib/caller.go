package lib

import "time"

// 调用器接口
type Caller interface {
	// 构建请求
	BuildRed() RawReq
	// 调用
	Call(req []byte, timeoutNS time.Duration) ([]byte, error)
	// 检查响应
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}
