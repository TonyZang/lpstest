package base

type LogFormat string

const (
	// 普通文件日志格式
	FORMAT_TEXT LogFormat = "text"
	// JSON日志格式
	FORMAT_JSON LogFormat = "json"
)

const (
	TIMESTAMP_FORMAT = "2006-01-02T15:04:05.999"
)
