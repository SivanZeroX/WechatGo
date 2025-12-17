package logger

import (
	"context"
	"fmt"
	"io"
	"maps"
	"strings"
)

// Level 日志级别
type Level int

const (
	// DebugLevel 调试级别
	DebugLevel Level = iota
	// InfoLevel 信息级别
	InfoLevel
	// WarnLevel 警告级别
	WarnLevel
	// ErrorLevel 错误级别
	ErrorLevel
	// DPaniclevel 严重错误级别（会导致panic）
	DPanicLevel
	// PanicLevel 致命错误级别（会导致panic）
	PanicLevel
	// FatalLevel 致命错误级别（会调用os.Exit）
	FatalLevel
)

// String 实现String接口
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case DPanicLevel:
		return "DPANIC"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Fields 字段映射
type Fields map[string]any

// Logger 日志接口
type Logger interface {
	// Debug 调试级别日志
	Debug(msg string, fields ...Fields)
	// Info 信息级别日志
	Info(msg string, fields ...Fields)
	// Warn 警告级别日志
	Warn(msg string, fields ...Fields)
	// Error 错误级别日志
	Error(msg string, err error, fields ...Fields)
	// DPanic 严重错误级别日志（开发模式会panic）
	DPanic(msg string, err error, fields ...Fields)
	// Panic 致命错误级别日志（会panic）
	Panic(msg string, err error, fields ...Fields)
	// Fatal 致命错误级别日志（会调用os.Exit）
	Fatal(msg string, err error, fields ...Fields)

	// WithFields 添加默认字段
	WithFields(fields Fields) Logger

	// WithContext 添加context
	WithContext(ctx context.Context) context.Context

	// Sync 刷新日志缓冲区
	Sync() error

	// Clone 克隆logger
	Clone() Logger
}

// New 创建默认logger
func New(opts ...Option) Logger {
	// 默认使用zap实现
	return NewZapLogger(opts...)
}

// Option 配置选项
type Option func(*options)

// WithLevel 设置日志级别
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// WithOutput 设置输出
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithDevelopment 设置开发模式
func WithDevelopment(development bool) Option {
	return func(o *options) {
		o.development = development
	}
}

// WithName 设置logger名称
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

type options struct {
	level       Level
	output      io.Writer
	development bool
	name        string
}

// FromContext 从context获取logger
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return logger
	}
	return New()
}

// ToContext 将logger放入context
func ToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// ContextLogger context中的logger
type ContextLogger struct {
	logger Logger
}

// NewContextLogger 创建context logger
func NewContextLogger(logger Logger) *ContextLogger {
	return &ContextLogger{logger: logger}
}

// Get 获取logger
func (cl *ContextLogger) Get() Logger {
	return cl.logger
}

type loggerKey struct{}

// String 格式化日志
func String(key, value string) Fields {
	return Fields{
		key: value,
	}
}

// Int 格式化整数
func Int(key string, value int) Fields {
	return Fields{
		key: value,
	}
}

// Int64 格式化64位整数
func Int64(key string, value int64) Fields {
	return Fields{
		key: value,
	}
}

// Bool 格式化布尔值
func Bool(key string, value bool) Fields {
	return Fields{
		key: value,
	}
}

// Error 格式化错误
func Error(err error) Fields {
	return Fields{
		"error": err.Error(),
	}
}

// Duration 格式化持续时间
func Duration(key string, value any) Fields {
	return Fields{
		key: value,
	}
}

// FormatFields 合并多个Fields
func FormatFields(fieldsList ...Fields) Fields {
	result := make(Fields)
	for _, fields := range fieldsList {
		maps.Copy(result, fields)
	}
	return result
}

// ParseFormat 解析格式化字符串
// 支持: %s 字符串, %d 整数, %v 任意值
func ParseFormat(format string, args ...any) (string, Fields) {
	if len(args) == 0 {
		return format, nil
	}

	fields := make(Fields)
	argIndex := 0
	var result strings.Builder

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i+1 < len(format) {
			if argIndex < len(args) {
				switch format[i+1] {
				case 's':
					fields[fmt.Sprintf("arg%d", argIndex+1)] = args[argIndex]
					fmt.Fprintf(&result, "%s", args[argIndex])
					i++ // 跳过格式字符
					argIndex++
				case 'd':
					fields[fmt.Sprintf("arg%d", argIndex+1)] = args[argIndex]
					fmt.Fprintf(&result, "%d", args[argIndex])
					i++ // 跳过格式字符
					argIndex++
				case 'v':
					fields[fmt.Sprintf("arg%d", argIndex+1)] = args[argIndex]
					fmt.Fprintf(&result, "%v", args[argIndex])
					i++ // 跳过格式字符
					argIndex++
				default:
					result.WriteByte(format[i])
				}
			} else {
				result.WriteByte(format[i])
			}
		} else {
			result.WriteByte(format[i])
		}
	}

	return result.String(), fields
}
