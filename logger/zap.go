package logger

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger zap实现
type zapLogger struct {
	logger      *zap.Logger
	sugar       *zap.SugaredLogger
	level       Level
	development bool
	fields      Fields
}

// NewZapLogger 创建zap logger
func NewZapLogger(opts ...Option) Logger {
	// 默认选项
	o := &options{
		level:       InfoLevel,
		output:      os.Stdout,
		development: false,
		name:        "wechatgo",
	}

	// 应用选项
	for _, opt := range opts {
		opt(o)
	}

	// 创建zap配置
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(toZapLevel(o.level)),
		Development:      o.development,
		Encoding:         getEncoding(o.development),
		EncoderConfig:    getEncoderConfig(o.development),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 如果指定了自定义output，使用console模式并创建core
	var core zapcore.Core
	if o.output != nil && o.output != os.Stdout {
		config.Encoding = "console"
		config.EncoderConfig = getEncoderConfig(true)

		// 创建自定义core来包装io.Writer
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(o.output),
			zap.NewAtomicLevelAt(toZapLevel(o.level)),
		)
	}

	// 创建logger
	var zapLog *zap.Logger
	var err error

	if core != nil {
		// 使用自定义core创建logger
		zapLog = zap.New(core, zap.AddCallerSkip(2))
	} else {
		zapLog, err = config.Build(zap.AddCallerSkip(2))
		if err != nil {
			// 如果创建失败，使用默认配置
			zapLog, err = zap.NewDevelopment(zap.AddCallerSkip(2))
			if err != nil {
				// 如果仍然失败，使用空操作logger避免panic
				zapLog = zap.NewNop()
			}
		}
	}

	sugar := zapLog.Sugar()

	return &zapLogger{
		logger:      zapLog,
		sugar:       sugar,
		level:       o.level,
		development: o.development,
		fields:      make(Fields),
	}
}

// Debug 实现Logger接口
func (l *zapLogger) Debug(msg string, fields ...Fields) {
	if l.level > DebugLevel {
		return
	}
	l.sugar.Debugw(msg, toKeyValuePairs(l.fields, fields...)...)
}

// Info 实现Logger接口
func (l *zapLogger) Info(msg string, fields ...Fields) {
	if l.level > InfoLevel {
		return
	}
	l.sugar.Infow(msg, toKeyValuePairs(l.fields, fields...)...)
}

// Warn 实现Logger接口
func (l *zapLogger) Warn(msg string, fields ...Fields) {
	if l.level > WarnLevel {
		return
	}
	l.sugar.Warnw(msg, toKeyValuePairs(l.fields, fields...)...)
}

// Error 实现Logger接口
func (l *zapLogger) Error(msg string, err error, fields ...Fields) {
	if l.level > ErrorLevel {
		return
	}
	allFields := toKeyValuePairs(l.fields, fields...)
	if err != nil {
		allFields = append(allFields, "error", err.Error())
	}
	l.sugar.Errorw(msg, allFields...)
}

// DPanic 实现Logger接口
func (l *zapLogger) DPanic(msg string, err error, fields ...Fields) {
	if l.level > DPanicLevel {
		return
	}
	allFields := toKeyValuePairs(l.fields, fields...)
	if err != nil {
		allFields = append(allFields, "error", err.Error())
	}
	l.sugar.DPanicw(msg, allFields...)
}

// Panic 实现Logger接口
func (l *zapLogger) Panic(msg string, err error, fields ...Fields) {
	if l.level > PanicLevel {
		return
	}
	allFields := toKeyValuePairs(l.fields, fields...)
	if err != nil {
		allFields = append(allFields, "error", err.Error())
	}
	l.sugar.Panicw(msg, allFields...)
}

// Fatal 实现Logger接口
func (l *zapLogger) Fatal(msg string, err error, fields ...Fields) {
	if l.level > FatalLevel {
		return
	}
	allFields := toKeyValuePairs(l.fields, fields...)
	if err != nil {
		allFields = append(allFields, "error", err.Error())
	}
	l.sugar.Fatalw(msg, allFields...)
}

// WithFields 实现Logger接口
func (l *zapLogger) WithFields(fields Fields) Logger {
	newFields := make(Fields)
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}

	return &zapLogger{
		logger:      l.logger,
		sugar:       l.sugar,
		level:       l.level,
		development: l.development,
		fields:      newFields,
	}
}

// WithContext 实现Logger接口
func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return ToContext(ctx, l)
}

// Sync 实现Logger接口
func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}

// Clone 实现Logger接口
func (l *zapLogger) Clone() Logger {
	fields := make(Fields)
	for k, v := range l.fields {
		fields[k] = v
	}

	return &zapLogger{
		logger:      l.logger,
		sugar:       l.sugar,
		level:       l.level,
		development: l.development,
		fields:      fields,
	}
}

// toZapLevel 转换日志级别到zap级别
func toZapLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case DPanicLevel:
		return zapcore.DPanicLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// toKeyValuePairs 转换Fields到key-value对
func toKeyValuePairs(baseFields Fields, fieldsList ...Fields) []interface{} {
	keysValues := make([]interface{}, 0)

	// 添加基础字段
	for k, v := range baseFields {
		keysValues = append(keysValues, k, v)
	}

	// 合并其他字段
	for _, fields := range fieldsList {
		for k, v := range fields {
			keysValues = append(keysValues, k, v)
		}
	}

	return keysValues
}

// getEncoding 获取编码器
func getEncoding(development bool) string {
	if development {
		return "console"
	}
	return "json"
}

// getEncoderConfig 获取编码器配置
func getEncoderConfig(development bool) zapcore.EncoderConfig {
	if development {
		return zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		}
	}

	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewDevelopment 创建开发模式logger
func NewDevelopment(opts ...Option) Logger {
	opts = append(opts, WithDevelopment(true))
	return NewZapLogger(opts...)
}

// NewProduction 创建生产模式logger
func NewProduction(opts ...Option) Logger {
	opts = append(opts, WithDevelopment(false))
	return NewZapLogger(opts...)
}

// StartTimer 开始计时器
func StartTimer() func(fields Fields) {
	start := time.Now()
	return func(fields Fields) {
		fields["duration"] = time.Since(start).String()
	}
}
