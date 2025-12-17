# Logger - 日志框架

WechatGo项目使用的结构化日志框架，基于zap实现。

## 特性

- 🚀 **高性能**: 基于zap的高性能日志实现
- 📝 **结构化日志**: 支持JSON格式的结构化日志输出
- 🎯 **灵活配置**: 支持多种日志级别、输出格式和配置选项
- 🔍 **开发友好**: 开发模式提供彩色输出和详细堆栈信息
- 🌐 **Context支持**: 支持从context中获取logger

## 快速开始

### 基本用法

```go
import "github.com/wechatpy/wechatgo/logger"

// 创建默认logger
log := logger.New()

// 记录不同级别的日志
log.Debug("调试信息", logger.String("key", "value"))
log.Info("系统启动", logger.String("component", "wechatgo"))
log.Warn("警告信息", logger.String("reason", "deprecated"))
log.Error("发生错误", fmt.Errorf("连接失败"), logger.String("host", "localhost"))
```

### 开发模式

```go
// 创建开发模式logger（彩色输出，详细信息）
log := logger.NewDevelopment()

// 创建生产模式logger（JSON格式，性能优化）
log := logger.NewProduction()
```

### 自定义配置

```go
import (
    "os"
    "github.com/wechatpy/wechatgo/logger"
)

// 设置日志级别
log := logger.New(logger.WithLevel(logger.DebugLevel))

// 设置输出
log := logger.New(logger.WithOutput(os.Stdout))

// 组合配置
log := logger.New(
    logger.WithLevel(logger.DebugLevel),
    logger.WithOutput(os.Stdout),
    logger.WithDevelopment(true),
    logger.WithName("myapp"),
)
```

## 日志级别

```go
const (
    DebugLevel  // 调试信息
    InfoLevel   // 一般信息
    WarnLevel   // 警告信息
    ErrorLevel  // 错误信息
    DPanicLevel // 严重错误（开发模式会panic）
    PanicLevel  // 致命错误（会panic）
    FatalLevel  // 致命错误（会调用os.Exit）
)
```

## 字段类型

### 字符串字段
```go
logger.String("username", "alice")
logger.String("url", "https://example.com")
```

### 数字字段
```go
logger.Int("count", 42)
logger.Int64("timestamp", 1234567890)
```

### 布尔字段
```go
logger.Bool("success", true)
logger.Bool("enabled", false)
```

### 错误字段
```go
err := fmt.Errorf("connection failed")
logger.Error("API调用失败", err, logger.String("url", "https://api.example.com"))
```

### 组合字段
```go
fields := logger.FormatFields(
    logger.String("user", "alice"),
    logger.Int("age", 30),
)
log.Info("用户信息", fields)
```

## 高级用法

### WithFields - 添加默认字段

```go
// 创建一个带有默认字段的logger
appLog := logger.New().WithFields(logger.String("app", "wechatgo"))

// 所有日志都会包含app字段
appLog.Info("系统启动")
appLog.Error("发生错误", fmt.Errorf("test"))
```

### WithContext - Context支持

```go
import "context"

ctx := context.Background()
log := logger.New()

// 将logger放入context
ctx = log.WithContext(ctx)

// 从context中获取logger
ctxLog := logger.FromContext(ctx)
ctxLog.Info("在context中的日志")
```

### StartTimer - 性能计时

```go
// 开始计时
timer := logger.StartTimer()

// ... 执行一些操作 ...

// 记录执行时间
fields := make(logger.Fields)
timer(fields)
log.Info("操作完成", fields)
```

### ParseFormat - 格式化消息

```go
msg, fields := logger.ParseFormat("用户 %s 发送了 %d 条消息", "Alice", 42)
log.Info(msg, fields)
```

## 在WechatGo中使用

### HTTP客户端日志

```go
import (
    "github.com/wechatpy/wechatgo/client"
    "github.com/wechatpy/wechatgo/logger"
)

// 创建带有日志的客户端
client := client.NewClient(appID, secret, storage)
client.WithLogger(logger.NewDevelopment())

// API调用会自动记录日志
result, err := client.User.Get("openid")
if err != nil {
    client.GetLogger().Error("获取用户信息失败", err,
        logger.String("openid", "openid"))
}
```

### 错误处理日志

```go
func handleError(err error) {
    log := logger.FromContext(ctx)
    log.Error("处理请求时发生错误", err,
        logger.String("component", "wechatgo"),
        logger.String("operation", "api_call"))
}
```

## 配置示例

### 开发环境配置

```go
log := logger.New(
    logger.WithLevel(logger.DebugLevel),
    logger.WithDevelopment(true),
    logger.WithName("wechatgo-dev"),
)
```

### 生产环境配置

```go
log := logger.New(
    logger.WithLevel(logger.InfoLevel),
    logger.WithDevelopment(false),
    logger.WithName("wechatgo-prod"),
)
```

### 性能监控配置

```go
log := logger.New(
    logger.WithLevel(logger.InfoLevel),
    logger.WithName("performance"),
)

// 记录API响应时间
start := time.Now()
result := callAPI()
duration := time.Since(start).Milliseconds()

log.Info("API调用完成",
    logger.String("endpoint", "/api/users"),
    logger.Int("duration_ms", int(duration)),
    logger.Int("status_code", 200))
```

## 输出格式

### 开发模式输出
```
2024-01-15T10:30:45.123Z	INFO	wechatgo	开始发送HTTP请求	method=GET url=https://api.weixin.qq.com/cgi-bin/token caller=base.go:113
```

### 生产模式输出
```json
{"level":"INFO","time":"2024-01-15T10:30:45.123Z","logger":"wechatgo","msg":"开始发送HTTP请求","method":"GET","url":"https://api.weixin.qq.com/cgi-bin/token","caller":"base.go:113"}
```

## 最佳实践

1. **选择合适的日志级别**
   - Debug: 调试信息，生产环境通常关闭
   - Info: 重要的业务流程信息
   - Warn: 警告信息，但不影响系统运行
   - Error: 错误信息，需要关注

2. **结构化日志**
   - 使用结构化字段而不是拼接字符串
   - 包含足够的上下文信息
   - 避免在日志中包含敏感信息

3. **性能考虑**
   - 在性能敏感的代码中谨慎使用Debug级别
   - 使用StartTimer记录关键操作耗时
   - 避免在高频循环中记录过多日志

4. **错误处理**
   - 记录错误时包含错误信息
   - 使用context传递logger
   - 为不同的错误场景添加适当的字段

## 贡献

如果您发现任何问题或有改进建议，欢迎提交Issue或Pull Request。
