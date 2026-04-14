# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在操作本代码仓库时提供指导。

## 项目概述

本项目是 **DevStar**，[Gitea](https://github.com/go-gitea/gitea) 的商业发行版。它是一个基于 **Go** 语言编写、前端采用 **JavaScript/TypeScript** 的自托管 Git 服务。

## 构建系统与依赖

- **Go**：1.25+（定义在 `go.mod` 中）
- **Node.js**：22.6.0+ 和 **pnpm**：10.0.0+（定义在 `package.json` 中）
- **构建工具**：`make`

安装依赖：
```bash
make deps          # 安装所有依赖
make deps-backend  # go mod download
make deps-frontend # pnpm install --frozen-lockfile
```

构建命令：
```bash
TAGS="bindata" make build                                      # 完整构建
TAGS="bindata sqlite sqlite_unlock_notify" make build          # 带 SQLite 支持的构建
make backend                                                   # 仅构建后端
make frontend                                                  # 仅构建前端 JS/CSS 资源
make clean                                                     # 清理构建产物
```

运行二进制文件：
```bash
./gitea web
```

## 开发与监听模式

```bash
make watch              # 同时监听前端与后端变化
make watch-frontend     # webpack --watch
make watch-backend      # air（Go 实时重载）
```

## 测试

### 后端单元测试
```bash
make test                    # 运行所有后端单元测试
make test#TestSpecificName   # 运行指定的单元测试
```

### 前端测试
```bash
make test-frontend           # vitest（配置位于 vitest.config.ts）
```

### 集成测试
运行前需要先执行干净构建（`make clean build`）。默认使用 SQLite，也可通过环境变量连接 MySQL/PostgreSQL/MSSQL。
```bash
make test-sqlite             # SQLite 集成测试
make test-sqlite#GPG         # 运行指定的集成测试
make test-mysql              # 需要 MySQL 服务正在运行
make test-pgsql              # 需要 PostgreSQL 服务正在运行
make test-mssql              # 需要 MSSQL 服务正在运行
```

### E2E 测试
使用 Playwright（配置位于 `playwright.config.ts`）。
```bash
make test-e2e-sqlite         # SQLite e2e 测试
make test-e2e-sqlite#example # 运行指定的 e2e 测试文件
make test-e2e-mysql
```

## 代码检查与格式化

提交前务必运行 `make fmt`。常用检查命令：
```bash
make fmt                     # 格式化 Go 代码（gofumpt）与模板
make lint                    # 检查所有代码
make lint-backend            # golangci-lint + gitea-vet + gopls + editorconfig
make lint-frontend           # eslint + vue-tsc + stylelint
make lint-go-fix             # 自动修复 Go 的 lint 问题
make lint-js-fix             # 自动修复 JS 的 lint 问题
make generate-swagger        # 根据代码注释重新生成 Swagger 规范
```

## 代码架构

### 后端（Go）

后端采用分层架构：

- **`cmd/`** —— CLI 命令与入口（使用 `urfave/cli/v3`）。`main.go` 负责启动应用。
- **`routers/`** —— HTTP 处理器与路由。
  - `routers/web/` —— Web UI 路由
  - `routers/api/v1/` —— REST API 路由
  - `routers/api/actions/` —— 兼容 GitHub Actions 的 API
  - `routers/api/packages/` —— 包注册表 API
  - `routers/private/` —— 内部 API 路由
  - `routers/common/` —— 共享中间件与协议处理
  - `routers/init.go` —— 全局初始化顺序（`InitWebInstalled`）
- **`services/`** —— 业务逻辑层（例如 `services/pull`、`services/issue`、`services/repository`、`services/migrations`）。
- **`models/`** —— 数据库模型与 ORM（XORM）。
  - `models/migrations/` —— 数据库迁移
  - `models/fixtures/` —— 测试夹具
- **`modules/`** —— 可复用库与基础设施。
  - `modules/setting/` —— 配置系统（`app.ini`）
  - `modules/git/` —— Git 操作封装
  - `modules/templates/` —— HTML 模板引擎
  - `modules/markup/` —— 标记渲染（markdown 等）
  - `modules/public/` —— 静态资源服务
  - `modules/storage/` —— 抽象文件存储
  - `modules/json/` —— JSON 序列化封装（请使用此包，而非 `encoding/json`）
  - `modules/structs/` —— API 请求/响应结构体（用于 Swagger 文档）
- **`tests/`** —— 测试套件。
  - `tests/integration/` —— 基于 Go 的集成测试
  - `tests/e2e/` —— Playwright 端到端测试

### 前端（JS/TS）

- **`web_src/js/`** —— JavaScript/TypeScript 源码。
  - 组件使用 Vue 3，服务器端交互使用 HTMX，打包使用 Webpack。
  - 测试文件：`web_src/**/*.test.ts`（Vitest + happy-dom）。
- **`web_src/css/`** —— CSS 源码（Tailwind CSS v3，自定义主题变量）。
- **`web_src/svg/`** —— SVG 图标。
- **`templates/`** —— Go HTML 模板（`*.tmpl`）。
- **`public/assets/`** —— 生成的静态资源（请勿直接编辑）。

### 关键配置文件

- `go.mod` / `go.sum` —— Go 依赖
- `package.json` / `pnpm-lock.yaml` —— Node.js 依赖
- `.golangci.yml` —— Go 代码检查规则
- `eslint.config.ts` —— ESLint 配置
- `stylelint.config.ts` —— Stylelint 配置
- `tailwind.config.ts` —— Tailwind CSS 配置
- `vitest.config.ts` —— 前端单元测试配置
- `playwright.config.ts` —— E2E 测试配置
- `.air.toml` —— Air 实时重载配置

## 重要约定

- **导入限制**（由 `.golangci.yml` 中的 `depguard` 强制执行）：
  - 请使用 `code.gitea.io/gitea/modules/json`，而不是 `encoding/json`
  - 请勿使用 `io/ioutil`；请使用 `os` 或 `io`
  - 请勿使用 `gopkg.in/ini.v1`；请使用 Gitea 的配置系统
  - 请勿导入 `modules/git/internal`；请使用 `AddXxx` 函数
- **模板**：使用 `*.tmpl` 扩展名。`make fmt` 会自动去除 `{{ }}` 和 `( )` 内部的空白。
- **Swagger**：API 端点必须文档化预期结果。JSON 结构体放在 `modules/structs/` 中，并必须在 `routers/api/v1/swagger/` 中引用。
- **新 Go 文件的版权头**：
  ```go
  // Copyright 2026 The Gitea Authors. All rights reserved.
  // SPDX-License-Identifier: MIT
  ```
- **PR 合并**：需要 2 位维护者批准（一周后的重构 PR 或仅文档修改的 PR 除外）。
