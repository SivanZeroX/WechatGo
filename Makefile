# WechatGo Makefile

.PHONY: help build test clean fmt vet lint coverage install deps

# 默认目标
.DEFAULT_GOAL := help

# 变量
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOFMT := gofmt
GOVET := $(GOCMD) vet
GOLINT := golangci-lint
BINARY_NAME := wechatgo
BINARY_UNIX := $(BINARY_NAME)_unix

# 帮助信息
help: ## 显示帮助信息
	@echo "WechatGo - 微信 SDK for Go"
	@echo ""
	@echo "可用命令:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# 构建项目
build: ## 构建项目
	@echo "构建项目..."
	$(GOBUILD) -o $(BINARY_NAME) -v ./...

# 交叉编译 Linux 版本
build-linux: ## 构建 Linux 版本
	@echo "构建 Linux 版本..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./...

# 清理构建文件
clean: ## 清理构建文件
	@echo "清理构建文件..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# 运行测试
test: ## 运行所有测试
	@echo "运行测试..."
	$(GOTEST) -v ./...

# 运行测试并生成覆盖率报告
test-coverage: ## 运行测试并生成覆盖率报告
	@echo "运行测试并生成覆盖率报告..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 代码格式化
fmt: ## 格式化代码
	@echo "格式化代码..."
	$(GOFMT) -s -w .

# 静态分析
vet: ## 运行静态分析
	@echo "运行静态分析..."
	$(GOVET) ./...

# 代码检查
lint: ## 运行代码检查
	@echo "运行代码检查..."
	@if command -v $(GOLINT) > /dev/null; then \
		$(GOLINT) run ./...; \
	else \
		echo "golangci-lint 未安装，跳过代码检查"; \
		echo "安装命令: go install github.com/golangci/golangci-lint/cmd/[email protected]"; \
	fi

# 安装依赖
deps: ## 安装依赖
	@echo "安装依赖..."
	$(GOCMD) mod download
	$(GOCMD) mod tidy

# 安装项目
install: ## 安装项目
	@echo "安装项目..."
	$(GOCMD) install ./...

# 运行基准测试
bench: ## 运行基准测试
	@echo "运行基准测试..."
	$(GOTEST) -bench=. -benchmem ./...

# 生成文档
doc: ## 生成文档
	@echo "生成文档..."
	$(GOCMD) doc ./...

# 检查更新
update: ## 检查并更新依赖
	@echo "检查更新..."
	$(GOCMD) get -u ./...
	$(GOCMD) mod tidy

# 安全检查
security: ## 运行安全检查
	@echo "运行安全检查..."
	@if command -v gosec > /dev/null; then \
		gosec ./...; \
	else \
		echo "gosec 未安装，跳过安全检查"; \
		echo "安装命令: go install github.com/securecodewarrior/gosec/v2/cmd/[email protected]"; \
	fi

# 生成版本信息
version: ## 显示版本信息
	@echo "Go 版本: $(shell go version)"
	@echo "项目路径: $(PWD)"

# 完整检查（格式化、静态分析、测试、覆盖率）
check: fmt vet test test-coverage ## 完整检查
	@echo "完整检查完成!"

# CI/CD 流水线
ci: deps fmt vet test lint ## CI/CD 流水线
	@echo "CI/CD 流水线完成!"

# Docker 构建
docker-build: ## 构建 Docker 镜像
	@echo "构建 Docker 镜像..."
	docker build -t wechatgo:latest .

# 发布准备
release: clean build-linux test ## 发布准备
	@echo "发布准备完成!"
	@echo "二进制文件: $(BINARY_UNIX)"

# 开发环境搭建
dev-setup: ## 开发环境搭建
	@echo "搭建开发环境..."
	go install github.com/golangci/golangci-lint/cmd/[email protected]
	go install github.com/securecodewarrior/gosec/v2/cmd/[email protected]
	go install github.com/swaggo/swag/cmd/[email protected]
	@echo "开发环境搭建完成!"
