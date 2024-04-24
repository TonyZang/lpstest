package base

type LogLevel uint8

const (
	// 调试级别，是最低的日志等级。
	LEVEL_DEBUG LogLevel = iota + 1
	// 信息级别，是最常用的日志等级。
	LEVEL_INFO
	// 警告级别，是适合输出到错误输出的日志等级。
	LEVEL_WARN
	// 普通错误级别，是建议输出到错误输出的日志等级。
	LEVEL_ERROR
	// 致命错误级别，是建议输出到错误输出的日志等级。
	// 此级别的日志一旦输出就意味着 `os.Exit(1)`立即会被调用。
	LEVEL_FATAL
	// 恐慌级别，是最高的日志等级。
	// 此级别的日志一旦输出就意味着运行时恐慌立即会被引发。
	LEVEL_PANIC
)
