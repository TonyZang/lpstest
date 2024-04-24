package base

import "lpstest/log/field"

// 代表日志记录器的选项
type Option interface {
	Name() string
}

// 代表一个日志记录器选项的类型
type OptWithLocation struct {
	Value bool
}

func (opt OptWithLocation) Name() string {
	return "with location"
}

// MyLogger 代表日志记录器接口
type MyLogger interface {
	Name() string
	Options() []Option
	Format() LogFormat
	Level() LogLevel

	Debug(v ...any)
	Debugf(format string, v ...any)
	Debugln(v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
	Errorln(v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Infoln(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Warnln(v ...any)

	WithFields(fields ...field.Field) MyLogger
}
