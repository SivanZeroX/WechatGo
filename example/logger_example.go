package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wechatpy/wechatgo/logger"
)

func main() {
	fmt.Println("=== WechatGo Logger 示例 ===")

	// 1. 基本日志记录
	fmt.Println("1. 基本日志记录:")
	basicLogging()

	// 2. 不同日志级别
	fmt.Println("\n2. 不同日志级别:")
	differentLevels()

	// 3. 结构化日志
	fmt.Println("\n3. 结构化日志:")
	structuredLogging()

	// 4. WithFields 使用
	fmt.Println("\n4. WithFields 使用:")
	withFields()

	// 5. Context 使用
	fmt.Println("\n5. Context 使用:")
	contextLogging()

	// 6. 性能计时
	fmt.Println("\n6. 性能计时:")
	performanceTiming()

	fmt.Println("\n=== 示例完成 ===")
}

func basicLogging() {
	log := logger.New()

	log.Debug("这是一条调试信息")
	log.Info("这是一条信息", logger.String("component", "wechatgo"))
	log.Warn("这是一条警告", logger.String("reason", "deprecated"))
	log.Error("这是一条错误", fmt.Errorf("模拟错误"))
}

func differentLevels() {
	// 创建开发模式logger
	devLog := logger.NewDevelopment()
	devLog.Info("开发模式日志 - 彩色输出，详细堆栈")

	// 创建生产模式logger
	prodLog := logger.NewProduction()
	prodLog.Info("生产模式日志 - JSON格式")
}

func structuredLogging() {
	log := logger.New()

	// 记录用户操作
	log.Info("用户登录",
		logger.String("username", "alice"),
		logger.String("ip", "192.168.1.1"),
		logger.String("user_agent", "Mozilla/5.0"),
	)

	// 记录API调用
	log.Info("API调用",
		logger.String("method", "POST"),
		logger.String("url", "/api/users"),
		logger.Int("status_code", 200),
		logger.Int("response_time_ms", 150),
	)

	// 记录错误
	log.Error("API调用失败",
		fmt.Errorf("connection timeout"),
		logger.String("url", "https://api.example.com"),
		logger.Int("timeout_ms", 5000),
		logger.Int("retry_count", 3),
	)
}

func withFields() {
	// 创建带有默认字段的logger
	appLog := logger.NewDevelopment().WithFields(logger.Fields{
		"app":     "wechatgo",
		"version": "1.0.0",
	})

	// 所有日志都会包含默认字段
	appLog.Info("系统启动")
	appLog.Info("配置加载完成", logger.String("config_path", "/etc/wechatgo/config.yaml"))
	appLog.Error("连接数据库失败", fmt.Errorf("connection refused"),
		logger.String("host", "localhost"),
		logger.Int("port", 5432),
	)
}

func contextLogging() {
	log := logger.New()

	// 创建context并放入logger
	ctx := context.Background()
	ctx = log.WithContext(ctx)

	// 模拟处理请求
	processRequest(ctx, "alice", "192.168.1.100")
}

func processRequest(ctx context.Context, username, ip string) {
	log := logger.FromContext(ctx)

	log.Info("开始处理请求",
		logger.String("username", username),
		logger.String("ip", ip),
		logger.String("request_id", "req-12345"),
	)

	// 模拟处理时间
	time.Sleep(100 * time.Millisecond)

	log.Info("请求处理完成",
		logger.String("username", username),
		logger.String("request_id", "req-12345"),
		logger.Int("duration_ms", 100),
	)
}

func performanceTiming() {
	log := logger.New()

	// 模拟API调用
	timer := logger.StartTimer()
	time.Sleep(150 * time.Millisecond)

	fields := make(logger.Fields)
	timer(fields)

	log.Info("API调用性能统计",
		logger.String("endpoint", "/api/users"),
		logger.String("method", "GET"),
		logger.Int("duration_ms", parseDuration(fields["duration"].(string))),
		logger.String("duration", fields["duration"].(string)),
	)
}

// parseDuration 解析持续时间字符串
func parseDuration(s string) int {
	// 简化实现，实际应该使用time.ParseDuration
	var ms int
	fmt.Sscanf(s, "%dms", &ms)
	return ms
}
