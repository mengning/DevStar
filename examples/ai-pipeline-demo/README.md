# AI Pipeline Demo（AI 开发流程示例）

这是一个展示如何在 **DevStar / Gitea** 中集成 AI 开发流程的示例 Go 项目。

## 功能演示

本项目包含三个 Gitea Actions 工作流，分别演示 AI 在开发流程中的三种应用场景：

### 1. AI Code Review（`.gitea/workflows/ai-code-review.yml`）
- **触发时机**：创建或更新 Pull Request
- **效果**：Claude 自动分析 PR diff，并在 PR 下发表评论，指出潜在 bug、安全问题、性能隐患和可读性建议

### 2. AI 生成单元测试（`.gitea/workflows/ai-test-gen.yml`）
- **触发时机**：Push 到 `main` 分支
- **效果**：Claude 自动为本次修改的 `.go` 文件生成对应的单元测试，并**提交一个新的 PR**供人工 review 后合并
- **安全设计**：AI 代码不直接 push 到主分支

### 3. AI Lint 缺陷分析（`.gitea/workflows/ai-lint-analyzer.yml`）
- **触发时机**：Push 或 Pull Request
- **效果**：先执行 `golangci-lint`，再由 Claude 将报错"翻译"成人话并给出修复代码，结果发布为 PR 评论或 Issue

---

## 前置准备

### 1. 在 DevStar 中启用 Actions

确保你的 DevStar 实例已开启 Gitea Actions：
- 管理员面板 → 配置 → Actions → 启用 Actions
- 配置 Actions Runner（`act_runner`）并注册到你的实例

### 2. 配置 API Key

在仓库设置中添加 Secret：
- 进入仓库 → 设置 → Secrets
- 添加 `ANTHROPIC_API_KEY`，值为你的 Claude API Key

> 如果你使用其他模型（如 OpenAI、DeepSeek），可以修改工作流中的 `curl` 请求部分。

### 3. 代码中的已知缺陷（供 AI 发现）

`main.go` 中故意预留了几个典型问题：
- **除零漏洞**：`Divide` 方法未检查 `b == 0`
- **参数校验缺失**：HTTP handler 未处理 `strconv.Atoi` 错误
- **无优雅关闭**：`http.ListenAndServe` 没有超时和 graceful shutdown

---

## 快速体验

### 步骤 1：将此项目推送到你的 DevStar 实例

```bash
git init
git add .
git commit -m "init ai pipeline demo"
git remote add origin http://localhost:3000/你的用户名/ai-pipeline-demo.git
git push -u origin main
```

### 步骤 2：创建 Pull Request 触发 AI Code Review

修改任意代码（例如在 `main.go` 里加一行注释），然后：

```bash
git checkout -b feature/test-ai-review
git commit -am "test ai review"
git push origin feature/test-ai-review
```

在 DevStar Web UI 中基于该分支创建 Pull Request，稍等片刻即可看到 AI 的 Review 评论。

### 步骤 3：合并后观察 AI 生成测试

PR 合并到 `main` 后，Actions 会自动运行 `ai-test-gen.yml`，并创建一个新的 PR（如 `ai/tests-1234567890`），里面包含了 Claude 自动生成的 `main_test.go`。

---

## 自定义扩展建议

| 场景 | 改造方向 |
|------|---------|
| 前端项目 | 将 `golangci-lint` 替换为 `eslint` + `stylelint`，生成 `.test.ts` / `.spec.js` |
| Python 项目 | 使用 `pytest` + `ruff` / `pylint` 替代 Go 工具链 |
| 更严格的 Review | 在工作流中加入 "必须 AI 通过才能合并" 的分支保护规则 |
| 自动修复 | 让 AI 直接输出 patch，Action 自动 apply 并 push commit |

---

## 注意事项

1. **Token 权限**：Gitea Actions 中默认使用 `secrets.GITHUB_TOKEN`，它在 Gitea 中会被映射为当前仓库的访问 token。如果你的 DevStar 版本不同，可能需要改为 `secrets.GITEA_TOKEN`。
2. **API 成本**：每次 PR 都会调用 Claude API，频繁提交可能产生一定费用。建议在正式使用时加入文件大小限制和缓存机制。
3. **Diff 截断**：工作流中已将 diff 限制在约 150KB，避免超大 PR 导致 API token 超限。
