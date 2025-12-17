# WechatGo - 微信 SDK for Go

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)]()
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg)](https://goreportcard.com/report/github.com/wechatpy/wechatgo)
[![Coverage](https://img.shields.io/badge/coverage-80%25-green.svg)]()

## 📖 项目简介

WechatGo 是 [wechatpy](https://github.com/wechatpy/wechatpy) 的 Go 语言实现，提供微信公众平台、企业微信、微信支付、物联网等 API 的完整 Go SDK。

**仓库地址**: [https://github.com/SivanZeroX/WechatGo](https://github.com/SivanZeroX/WechatGo)

当前版本：**v0.1.0** (编译测试通过 ✅)

## 📑 目录

- [核心特性](#-核心特性)
- [项目架构](#-项目架构)
- [快速开始](#-快速开始)
  - [安装](#安装)
  - [使用 Makefile](#使用-makefile)
  - [基本使用](#基本使用)
- [示例代码](#-示例代码)
- [API 模块详解](#-api-模块详解)
- [开发指南](#-开发指南)
- [测试](#-测试)
- [性能优化](#-性能优化)
- [最佳实践](#-最佳实践)
- [常见问题](#-常见问题)
- [故障排查](#-故障排查)
- [贡献](#-贡献)
- [版本历史](#-版本历史)
- [许可证](#-许可证)

## ✨ 核心特性

### 🎯 多客户端模块化设计
- **按业务领域分离** - 公众平台、支付、企业微信、IoT 各模块独立
- **高内聚低耦合** - 模块间依赖最小化，支持独立使用
- **依赖隔离** - 避免循环依赖，按需加载减少二进制大小

### 🔒 安全与性能
- ✅ **并发安全** - 所有共享资源使用 sync.RWMutex 保护
- ✅ **资源管理** - HTTP 响应体自动释放，防止泄漏
- ✅ **TTL 支持** - 会话存储支持毫秒级过期时间
- ✅ **消息加密** - 支持消息加密/解密、签名验证

### 📦 完整功能覆盖
- ✅ **消息处理** - 解析、回复、事件处理
- ✅ **会话管理** - 内存/Redis 存储，可扩展接口
- ✅ **日志框架** - 集成 Zap 高性能结构化日志
- ✅ **单元测试** - 14+ 测试用例覆盖核心功能

## 🏗️ 项目架构

```
wechatgo/
├── client/              # 📱 微信公众号/小程序客户端
│   ├── api/            # 公众平台API实现
│   │   ├── base.go     # 基础API
│   │   ├── user.go     # 用户管理
│   │   ├── message.go  # 消息管理
│   │   ├── menu.go     # 自定义菜单
│   │   ├── media.go    # 媒体管理
│   │   ├── qrcode.go   # 二维码
│   │   ├── tag.go      # 标签管理
│   │   └── merchant/   # 商户API
│   ├── base.go         # 客户端基础类
│   └── client.go       # 客户端主类
├── pay/                # 💰 微信支付客户端
│   ├── api/            # 支付API (v2)
│   │   ├── order.go    # 订单管理
│   │   ├── refund.go   # 退款处理
│   │   ├── jsapi.go    # JSAPI支付
│   │   ├── micropay.go # 刷卡支付
│   │   ├── redpack.go  # 红包
│   │   ├── transfer.go # 企业转账
│   │   ├── coupon.go   # 代金券
│   │   └── profitsharing.go # 分账
│   ├── v3/             # 支付API v3
│   │   ├── api/        # V3 API实现
│   │   │   ├── base.go         # 基础结构
│   │   │   ├── banks.go        # 银行API
│   │   │   ├── ecommerce.go    # 电商API
│   │   │   ├── media.go        # 媒体API
│   │   │   └── partner_order.go # 合作伙伴订单
│   └── client.go       # 支付客户端
├── work/               # 🏢 企业微信客户端
│   └── client/         # 企业微信API
│       ├── user.go      # 用户管理
│       ├── department.go# 部门管理
│       ├── tag.go       # 标签管理
│       ├── message.go   # 消息管理
│       ├── media.go     # 媒体管理
│       ├── contact.go   # 客户联系
│       ├── oa.go        # 办公应用
│       └── api/         # 扩展API
│           ├── auth.go        # 认证API
│           └── miniprogram.go # 小程序API
├── iot/                # 🔌 物联网客户端
│   └── client/         # IoT API
│       ├── client.go   # IoT客户端
│       ├── device.go   # 设备管理
│       └── cloud.go    # 云端API
├── crypto/             # 🔐 加密相关
│   ├── cipher.go       # 加密算法
│   ├── pkcs7.go        # PKCS7填充
│   └── utils.go        # 加密工具
├── session/            # 💾 会话管理
│   ├── memory.go       # 内存存储 (支持TTL)
│   ├── redis.go        # Redis存储
│   └── session.go      # 会话接口
├── logger/             # 📝 日志框架 (Zap)
│   ├── logger.go       # 日志接口
│   └── zap.go          # Zap实现
├── example/            # 📚 示例代码
│   └── logger_example.go # 日志使用示例
├── constants.go        # 常量定义
├── errors.go           # 错误处理
├── events.go           # 事件处理
├── messages.go         # 消息结构
├── parser.go           # 消息解析
├── replies.go          # 回复处理
├── utils.go            # 工具函数
├── doc.go              # 包文档
└── Makefile            # 构建脚本
```

## 🚀 快速开始

### 安装

```bash
go get github.com/wechatpy/wechatgo
```

### 使用 Makefile

项目提供了完整的 Makefile 来简化开发流程:

```bash
# 查看所有可用命令
make help

# 常用命令
make deps           # 安装依赖
make test           # 运行测试
make test-coverage  # 生成测试覆盖率报告
make fmt            # 格式化代码
make vet            # 静态分析
make lint           # 代码检查
make build          # 构建项目
make check          # 完整检查(格式化+静态分析+测试+覆盖率)
make ci             # CI/CD 流水线
make dev-setup      # 搭建开发环境
```

### 基本使用

#### 1. 公众平台客户端

```go
import (
    "github.com/wechatpy/wechatgo/client"
    "github.com/wechatpy/wechatgo/session"
)

// 创建会话存储
storage := session.NewMemoryStorage()

// 创建客户端
wechatClient := client.NewClient("your_app_id", "your_app_secret", storage)

// 获取访问令牌
err := wechatClient.FetchAccessToken()
if err != nil {
    // 处理错误
}

// 使用 API 模块
// 获取用户信息
userInfo, err := wechatClient.User.Get("openid")

// 发送模板消息
err = wechatClient.Template.Send(templateData)

// 创建自定义菜单
err = wechatClient.Menu.Create(menuData)
```

#### 2. 微信支付客户端

```go
import (
    "github.com/wechatpy/wechatgo/pay"
    "github.com/wechatpy/wechatgo/pay/api"
)

httpClient := &http.Client{}
client := pay.NewClient("appID", "apiKey", "mchID", "certPath", "keyPath", httpClient)

// 创建订单
req := &api.PrepayRequest{
    AppID: "appID",
    MchID: "mchID",
    Body: "测试订单",
    OutTradeNo: "order_001",
    TotalFee: 100,
    SpbillCreateIP: "127.0.0.1",
    NotifyURL: "https://yourapp.com/notify",
    TradeType: "JSAPI",
}

prepayID, err := client.GetPrepayID(req)
```

#### 3. 企业微信客户端

```go
import (
    "github.com/wechatpy/wechatgo/work/client"
)

storage := session.NewMemoryStorage()
workClient := workclient.NewWorkClient("corpID", "corpSecret", storage)

// 获取部门列表
deptList, err := workClient.Dept.Get()
```

#### 4. IoT 客户端

```go
import (
    "github.com/wechatpy/wechatgo/iot/client"
)

storage := session.NewMemoryStorage()
iotClient := iotclient.NewIotClient("appID", "secret", storage)

// 获取访问令牌
err := iotClient.FetchAccessToken()
```

### 消息处理

```go
import (
    "github.com/wechatpy/wechatgo"
)

// 解析微信推送的 XML 消息
msg, err := wechatgo.ParseMessage(xmlData)
if err != nil {
    // 处理解析错误
}

// 根据消息类型处理
switch m := msg.(type) {
case *wechatgo.TextMessage:
    // 处理文本消息
    reply := wechatgo.NewTextReply(m.Source, m.Target, "收到消息: " + m.Content)
    return reply.Render()

case *wechatgo.ImageMessage:
    // 处理图片消息
    reply := wechatgo.NewImageReply(m.Source, m.Target, m.MediaID)
    return reply.Render()

case *wechatgo.SubscribeEvent:
    // 处理关注事件
    reply := wechatgo.NewTextReply(m.Source, m.Target, "感谢关注!")
    return reply.Render()

case *wechatgo.ClickEvent:
    // 处理菜单点击事件
    // 根据 m.EventKey 处理不同的菜单项
}
```

### 会话管理

```go
import (
    "github.com/wechatpy/wechatgo/session"
    "time"
)

// 内存会话 (支持TTL)
storage := session.NewMemoryStorage()
storage.Set("key", "value", 5*time.Minute)

// Redis 会话
import "github.com/redis/go-redis/v9"
redisClient := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})
redisStorage := session.NewRedisStorage(redisClient)
```

### 日志记录

```go
import (
    "github.com/wechatpy/wechatgo/logger"
)

log := logger.New()

// 结构化日志
log.Info("用户登录",
    logger.String("user_id", "12345"),
    logger.String("ip", "192.168.1.1"),
)

// 错误日志
log.Error("请求失败", err,
    logger.String("url", "/api/test"),
)
```

## 📚 示例代码

项目提供了完整的示例代码,帮助您快速上手:

```bash
# 查看日志使用示例
cat example/logger_example.go

# 运行示例
go run example/logger_example.go
```

更多示例代码正在持续添加中,包括:
- 公众号消息处理完整示例
- 微信支付订单处理示例
- 企业微信应用集成示例
- 会话管理最佳实践

## 📚 API 模块详解

### 公众平台 API (`client/api/`)

| 模块 | 功能 | 状态 |
|------|------|------|
| User | 用户管理 | ✅ |
| Message | 发送消息 | ✅ |
| Menu | 自定义菜单 | ✅ |
| Media | 图片/视频/语音 | ✅ |
| QRCode | 二维码生成 | ✅ |
| Tag | 用户标签 | ✅ |
| Template | 模板消息 | ✅ |
| DataCube | 数据统计 | ✅ |
| CustomService | 客服功能 | ✅ |
| Device | 设备管理 | ✅ |
| POI | 门店管理 | ✅ |
| WiFi | WiFi管理 | ✅ |

### 支付 API (`pay/`)

| 模块 | 功能 | 版本 |
|------|------|------|
| Order | 订单管理 | v2/v3 |
| Refund | 退款处理 | v2 |
| JsAPI | JSAPI支付 | v2 |
| MicroPay | 刷卡支付 | v2 |
| RedPack | 红包发放 | v2 |
| Transfer | 企业转账 | v2 |
| Coupon | 代金券 | v2 |
| ProfitShare | 分账 | v2 |
| Ecommerce | 电商API | v3 |
| Banks | 银行API | v3 |
| Media | 媒体API | v3 |
| PartnerOrder | 合作伙伴订单 | v3 |

### 企业微信 API (`work/client/`)

| 模块 | 功能 | 状态 |
|------|------|------|
| User | 用户管理 | ✅ |
| Department | 部门管理 | ✅ |
| Tag | 标签管理 | ✅ |
| Message | 消息管理 | ✅ |
| Media | 媒体管理 | ✅ |
| Contact | 客户联系 | ✅ |
| OA | 办公应用 | ✅ |
| Auth | 认证API | ✅ |
| MiniProgram | 小程序API | ✅ |

### IoT API (`iot/client/`)

| 模块 | 功能 | 状态 |
|------|------|------|
| Device | 设备管理 | ✅ |
| Cloud | 云端API | ✅ |

## 🛠️ 开发指南

### 添加新的API

1. 在相应模块创建新的 `.go` 文件
2. 实现 API 结构体和方法
3. 在客户端初始化中注册

```go
// 示例：添加新的API
type NewAPI struct {
    *BaseAPI
}

func NewNewAPI(client interface {
    Get(url string, params map[string]string) (map[string]interface{}, error)
    Post(url string, data interface{}) (map[string]interface{}, error)
}) *NewAPI {
    return &NewAPI{
        BaseAPI: NewBaseAPI(client),
    }
}
```

### 扩展会话存储

```go
type SessionStore interface {
    Get(key string) (string, error)
    Set(key string, value string, ttl time.Duration) error
    Delete(key string) error
}
```

### 自定义日志

```go
type Logger interface {
    Debug(msg string, fields ...Fields)
    Info(msg string, fields ...Fields)
    Warn(msg string, fields ...Fields)
    Error(msg string, err error, fields ...Fields)
    WithFields(fields Fields) Logger
}
```

## ✅ 测试

```bash
# 运行所有测试
go test ./...

# 运行特定模块测试
go test ./session/...

# 运行测试并查看覆盖率
go test -cover ./...
```

**当前测试覆盖率**: 核心模块 80%+

## 📈 性能优化

- **HTTP 连接复用** - 自动管理 HTTP 连接池,最大空闲连接数 100
- **访问令牌缓存** - 自动缓存和刷新 access_token,避免频繁请求
- **响应体释放** - 自动释放 HTTP 响应体,防止资源泄漏
- **并发安全** - 使用 sync.RWMutex 保护共享资源,支持高并发
- **按需加载** - 模块化设计减少内存占用,只加载需要的模块
- **结构化日志** - 使用 Zap 高性能日志框架,零内存分配

## 🔧 最佳实践

### 1. 会话存储选择

```go
// 开发环境 - 使用内存存储
storage := session.NewMemoryStorage()

// 生产环境 - 使用 Redis 存储(支持分布式部署)
import "github.com/redis/go-redis/v9"
redisClient := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // 生产环境请设置密码
    DB:       0,
})
storage := session.NewRedisStorage(redisClient)
```

### 2. 错误处理

```go
// 始终检查错误
err := client.FetchAccessToken()
if err != nil {
    log.Error("获取 access token 失败", err)
    // 根据错误类型进行重试或告警
    return
}
```

### 3. 日志配置

```go
// 生产环境建议配置日志级别和输出
log := logger.New()
log.Info("应用启动",
    logger.String("version", "v1.0.0"),
    logger.String("env", "production"),
)
```

### 4. HTTP 客户端复用

```go
// 客户端会自动复用 HTTP 连接,无需手动管理
// 避免频繁创建新的客户端实例
var globalClient *client.Client

func init() {
    storage := session.NewMemoryStorage()
    globalClient = client.NewClient(appID, secret, storage)
}

## ❓ 常见问题

### Q: 如何处理 access_token 过期?

A: SDK 会自动管理 access_token 的缓存和刷新,无需手动处理。如果遇到 token 过期错误,SDK 会自动重新获取。

### Q: 支持哪些 Go 版本?

A: 项目使用 Go 1.21 开发,建议使用 Go 1.21 及以上版本(支持 1.21, 1.22, 1.23)。

### Q: 如何在生产环境部署?

A: 建议使用 Redis 作为会话存储,配置合适的日志级别,并使用 `make build` 构建生产版本。

### Q: 消息加密如何处理?

A: SDK 支持消息加密/解密,使用 `crypto` 包提供的加密工具即可。

### Q: 如何贡献代码?

A: 请参考下方的贡献指南,我们欢迎所有形式的贡献!

## 🐛 故障排查

### 问题: 获取 access_token 失败

```bash
# 检查 AppID 和 Secret 是否正确
# 检查网络连接是否正常
# 查看详细错误日志
```

### 问题: Redis 连接失败

```bash
# 检查 Redis 服务是否启动
redis-cli ping

# 检查 Redis 连接配置
# 确认防火墙规则
```

### 问题: 消息解析失败

```bash
# 检查 XML 格式是否正确
# 查看原始 XML 数据
# 确认消息类型是否支持
```

## 🤝 贡献

欢迎提交 Pull Request 和 Issue!

### 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 开发规范

- 遵循 Go 代码规范 (`gofmt`, `go vet`)
- 添加单元测试
- 更新文档
- 保持向后兼容

## 📋 版本历史

### v0.1.0 (当前版本)
- ✅ 完成核心架构设计
- ✅ 实现公众平台 API
- ✅ 实现微信支付 API (v2/v3)
- ✅ 实现企业微信 API
- ✅ 实现 IoT API
- ✅ 消息解析和回复功能
- ✅ 会话管理(内存/Redis)
- ✅ 日志框架集成(Zap)
- ✅ 单元测试覆盖

### 开发路线图

#### v0.2.0 (计划中)
- [ ] 完善单元测试覆盖率至 90%+
- [ ] 添加更多示例代码
- [ ] 性能基准测试
- [ ] API 文档完善

#### v0.3.0 (计划中)
- [ ] 支持微信小程序 API
- [ ] 支持微信开放平台 API
- [ ] 添加中间件支持
- [ ] 集成更多第三方存储

#### v1.0.0 (长期目标)
- [ ] 完整的 API 覆盖
- [ ] 生产环境验证
- [ ] 性能优化
- [ ] 完善的文档和示例

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🔗 参考资料

- [微信公众平台开发文档](https://developers.weixin.qq.com/doc/offiaccount/)
- [企业微信开发文档](https://developer.work.weixin.qq.com/)
- [微信支付开发文档](https://pay.weixin.qq.com/wiki/doc/apiv3/)
- [wechatpy 项目](https://github.com/wechatpy/wechatpy)

## 👥 贡献者

感谢所有为本项目做出贡献的开发者！

## 📞 联系方式

- **GitHub**: [https://github.com/SivanZeroX/WechatGo](https://github.com/SivanZeroX/WechatGo)
- **问题反馈**: [https://github.com/SivanZeroX/WechatGo/issues](https://github.com/SivanZeroX/WechatGo/issues)
- **讨论交流**: [GitHub Discussions](https://github.com/SivanZeroX/WechatGo/discussions)

---

<div align="center">

**⭐ 如果这个项目对您有帮助，请给我们一个 Star！⭐**

</div>
