# DevStar AI Code Review 深度集成设计文档

## 1. 项目背景与目标

### 1.1 背景
当前 AI 辅助研发已成为提升代码质量和开发效率的重要手段。DevStar 作为面向 AI 时代的 R&D 平台，需要将 AI Code Review 能力原生内嵌到 PR（Pull Request）工作流中，而非仅依赖外部 CI/CD 流水线。

### 1.2 目标
- **原生体验**：用户在 PR diff 页面即可一键请求 AI 评审，结果以标准 Review 形式呈现。
- **异步可靠**：AI API 调用通过队列异步执行，不阻塞用户 HTTP 请求。
- **可配置**：支持开关、模型选择、自动触发策略等灵活配置。
- **兼容现有体系**：AI Review 复用现有的 `Review`/`Comment` 数据模型，无需重建前端渲染逻辑。

### 1.3 非目标
- 不替换现有人工 Code Review 流程，AI 仅作为辅助。
- 第一版不支持 AI 自动修复代码并直接 push commit。
- 第一版不强制阻塞 PR 合并（可作为 Phase 2 扩展）。

---

## 2. 架构总览

```
┌─────────────────────────────────────────────────────────────────────┐
│                         Frontend (Web UI)                          │
│  PR diff page ──► "🤖 AI 评审" 按钮 ──► fetch API call             │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         API / Web Router                            │
│  POST /api/v1/repos/{owner}/{repo}/pulls/{index}/ai-review          │
│  (或 Web 路由：POST /{owner}/{repo}/pulls/{index}/ai-review)        │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         AI Service Layer                            │
│  ai_service.CreateReviewTask() ──► 生成队列任务                      │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         Queue (modules/queue)                       │
│  "ai_review" UniqueQueue ──► 异步消费                               │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         AI Worker Handler                           │
│  1. 获取 PR diff                                                    │
│  2. 调用 Claude API (或 OpenAI/DeepSeek 等)                         │
│  3. 解析返回结果                                                     │
│  4. 调用 pull_service.SubmitReview() 写入 Review                    │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         Data Layer (models)                         │
│  issues_model.Review + issues_model.Comment                         │
│  (复用现有模型，ReviewType = ReviewTypeComment)                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 3. 详细设计

### 3.1 数据模型（零新增表，复用现有模型）

AI Review 完全复用现有 `issues_model.Review` 和 `issues_model.Comment`，不需要新增数据库表。

#### 3.1.1 Review 模型
**文件**：`models/issues/review.go`

AI 提交的 Review 使用：
- `ReviewType`: `ReviewTypeComment` (值为 2)
- `ReviewerID`: 指向一个系统用户（如 `ai-reviewer` 的 UserID）
- `IssueID`: 关联的 PR IssueID
- `Content`: AI 生成的 Markdown 评审内容

#### 3.1.2 Comment 模型
**文件**：`models/issues/comment.go`

`ReviewTypeComment` 的 Review 提交后，会在 `comment` 表中自动生成一条 `CommentTypeReview` (22) 的记录。这是现有 `SubmitReview` 流程的默认行为。

> **Phase 2 扩展**：如果需要支持"行级评论"（精准定位到某行代码），可进一步创建 `CommentTypeCode` (21) 的记录，通过 `TreePath`、`LineNum`、`CommitSHA` 等字段关联到具体代码位置。

---

### 3.2 配置系统设计

#### 3.2.1 新增配置模块
**文件**：`modules/setting/ai.go`（新建）

```go
package setting

import "time"

var AI = struct {
    Enabled              bool
    APIKey               string        `ini:"API_KEY"`
    APIBaseURL           string        `ini:"API_BASE_URL"`
    Model                string        `ini:"MODEL"`
    MaxDiffSize          int           `ini:"MAX_DIFF_SIZE"`
    AutoReviewOnPROpen   bool          `ini:"AUTO_REVIEW_ON_PR_OPEN"`
    AutoReviewOnPRSync   bool          `ini:"AUTO_REVIEW_ON_PR_SYNC"`
    ReviewerUserID       int64         `ini:"REVIEWER_USER_ID"`
    RequestTimeout       time.Duration `ini:"REQUEST_TIMEOUT"`
    PromptSystem         string        `ini:"-"`
}{
    Enabled:            false,
    APIBaseURL:         "https://api.anthropic.com/v1",
    Model:              "claude-3-5-sonnet-20241022",
    MaxDiffSize:        150000,
    AutoReviewOnPROpen: false,
    AutoReviewOnPRSync: false,
    RequestTimeout:     5 * time.Minute,
}

func loadAIFrom(rootCfg ConfigProvider) error {
    sec := rootCfg.Section("ai")
    if err := sec.MapTo(&AI); err != nil {
        return fmt.Errorf("failed to map AI settings: %v", err)
    }
    // 读取多行 prompt
    AI.PromptSystem = sec.Key("PROMPT_SYSTEM").String()
    if AI.PromptSystem == "" {
        AI.PromptSystem = defaultAIPrompt()
    }
    return nil
}

func defaultAIPrompt() string {
    return `You are an expert code reviewer. Analyze the git diff and provide a structured review in Markdown.
Use Chinese. Sections: 1) Summary 2) Potential Bugs 3) Security Issues 4) Performance Concerns 5) Readability & Best Practices 6) Suggested Fixes.
Be concise and actionable.`
}
```

#### 3.2.2 注册配置加载
**文件**：`modules/setting/setting.go`

在 `loadSettings()` 函数中增加：
```go
if err := loadAIFrom(cfg); err != nil {
    return err
}
```

#### 3.2.3 管理员配置示例
在 `data/app.ini` 中可配置：
```ini
[ai]
ENABLED = true
API_KEY = sk-ant-api03-xxxxxxxx
API_BASE_URL = https://api.anthropic.com/v1
MODEL = claude-3-5-sonnet-20241022
MAX_DIFF_SIZE = 150000
AUTO_REVIEW_ON_PR_OPEN = false
AUTO_REVIEW_ON_PR_SYNC = false
REVIEWER_USER_ID = 1
REQUEST_TIMEOUT = 5m
```

---

### 3.3 服务层设计（核心）

#### 3.3.1 新增 AI Service 包
**目录**：`services/ai/`（新建）

##### `services/ai/task.go` —— 任务结构体
```go
package ai

type ReviewTask struct {
    PullRequestID int64
    IssueID       int64
    RepoID        int64
    CommitID      string
    DoerID        int64  // 请求发起者，用于日志和审计
}
```

##### `services/ai/queue.go` —— 队列注册与处理
```go
package ai

import (
    "code.gitea.io/gitea/modules/graceful"
    "code.gitea.io/gitea/modules/queue"
)

var reviewQueue queue.UniqueQueue

func Init() error {
    reviewQueue = queue.CreateUniqueQueue(
        graceful.GetManager().ShutdownContext(),
        "ai_review",
        handleReviewTask,
    )
    return nil
}

func CreateReviewTask(task *ReviewTask) error {
    return reviewQueue.Push(task)
}

func handleReviewTask(tasks ...*ReviewTask) []*ReviewTask {
    var failed []*ReviewTask
    for _, task := range tasks {
        if err := processReviewTask(task); err != nil {
            log.Error("AI review task failed: %v", err)
            failed = append(failed, task)
        }
    }
    return failed
}
```

##### `services/ai/review.go` —— 核心处理逻辑
```go
package ai

import (
    "context"
    "fmt"
    "strings"

    issues_model "code.gitea.io/gitea/models/issues"
    repo_model "code.gitea.io/gitea/models/repo"
    user_model "code.gitea.io/gitea/models/user"
    "code.gitea.io/gitea/modules/git"
    "code.gitea.io/gitea/modules/gitrepo"
    "code.gitea.io/gitea/modules/json"
    "code.gitea.io/gitea/modules/log"
    "code.gitea.io/gitea/modules/setting"
    pull_service "code.gitea.io/gitea/services/pull"
)

func processReviewTask(task *ReviewTask) error {
    ctx := context.Background()

    // 1. 加载关联数据
    pr, err := issues_model.GetPullRequestByID(ctx, task.PullRequestID)
    if err != nil {
        return err
    }
    issue := pr.Issue
    if issue == nil {
        if err := pr.LoadIssue(ctx); err != nil {
            return err
        }
        issue = pr.Issue
    }

    repo, err := repo_model.GetRepositoryByID(ctx, task.RepoID)
    if err != nil {
        return err
    }

    // 2. 获取 diff
    gitRepo, closer, err := gitrepo.RepositoryFromContextOrOpen(ctx, repo)
    if err != nil {
        return err
    }
    defer closer.Close()

    diff, err := getPRDiff(ctx, gitRepo, pr)
    if err != nil {
        return err
    }

    // 3. 调用 AI API
    reviewContent, err := generateReviewContent(diff)
    if err != nil {
        return err
    }

    // 4. 获取 AI Reviewer 用户
    reviewer, err := user_model.GetUserByID(ctx, setting.AI.ReviewerUserID)
    if err != nil {
        return err
    }

    // 5. 提交 Review
    _, _, err = pull_service.SubmitReview(
        ctx, reviewer, gitRepo, issue,
        issues_model.ReviewTypeComment,
        reviewContent,
        task.CommitID,
        nil,
    )
    return err
}

func getPRDiff(ctx context.Context, gitRepo *git.Repository, pr *issues_model.PullRequest) (string, error) {
    // 使用 git diff merge_base..head_commit 获取完整 diff
    // 参考 services/pull/comment.go 中 getCommitIDsFromRepo 的实现
    // 以及 gitdiff 包的相关逻辑
    
    if pr.MergeBase == "" {
        if err := pr.LoadBaseRepo(ctx); err != nil {
            return "", err
        }
        // 尝试计算 merge base
        // ... 省略具体实现，可参考现有代码
    }
    
    cmd := git.NewCommand(ctx, "diff", "--max-count=1", "--")
    cmd.AddDynamicArguments(pr.MergeBase + ".." + pr.GetGitHeadRefName())
    // 实际应使用更安全的封装方法
    return cmd.RunInDir(gitRepo.Path)
}

func generateReviewContent(diff string) (string, error) {
    if len(diff) > setting.AI.MaxDiffSize {
        diff = diff[:setting.AI.MaxDiffSize] + "\n\n[Diff truncated due to size limit]"
    }

    payload := map[string]interface{}{
        "model":      setting.AI.Model,
        "max_tokens": 4096,
        "system":     setting.AI.PromptSystem,
        "messages": []map[string]string{
            {
                "role":    "user",
                "content": "Please review the following code diff:\n\n```diff\n" + diff + "\n```",
            },
        },
    }

    body, _ := json.Marshal(payload)
    
    // 实际 HTTP 调用逻辑
    resp, err := http.Post(
        setting.AI.APIBaseURL+"/messages",
        "application/json",
        bytes.NewReader(body),
    )
    // ... 解析响应
}
```

---

### 3.4 路由层设计

#### 3.4.1 API 路由（推荐）
**文件**：`routers/api/v1/repo/pull_review.go`（新建或在现有文件中追加）

新增 API 端点：
```go
// TriggerAIReview 触发 AI 评审
func TriggerAIReview(ctx *context.APIContext) {
    // 1. 权限检查：当前用户是否有权限评论此 PR
    // 2. 检查 AI 功能是否开启
    if !setting.AI.Enabled {
        ctx.Error(http.StatusForbidden, "", "AI review is disabled")
        return
    }
    
    // 3. 加载 PR
    pr, err := issues_model.GetPullRequestByIndex(ctx, ctx.Repo.Repository.ID, ctx.ParamsInt64(":index"))
    if err != nil {
        ctx.Error(http.StatusInternalServerError, "GetPullRequestByIndex", err)
        return
    }
    
    // 4. 入队
    task := &ai.ReviewTask{
        PullRequestID: pr.ID,
        IssueID:       pr.IssueID,
        RepoID:        ctx.Repo.Repository.ID,
        CommitID:      pr.HeadCommitID,
        DoerID:        ctx.Doer.ID,
    }
    if err := ai.CreateReviewTask(task); err != nil {
        ctx.Error(http.StatusInternalServerError, "CreateReviewTask", err)
        return
    }
    
    ctx.JSON(http.StatusAccepted, map[string]string{
        "message": "AI review task has been queued",
    })
}
```

**注册路由**：
**文件**：`routers/api/v1/api.go`

在 PR 相关路由组中增加：
```go
m.Combo("/ai-review").Post(repo.TriggerAIReview)
```

#### 3.4.2 Web 路由（备选）
如果不希望走 API，也可以在 `routers/web/repo/pull_review.go` 中新增 Web handler，表单提交后重定向回 PR 页面。

---

### 3.5 前端/UI 设计

#### 3.5.1 页面入口
**文件**：`templates/repo/diff/new_review.tmpl`

在现有 Review 按钮附近新增 AI 评审按钮：
```html
{{if and .PageIsPullFiles $.SignedUserID setting.AI.Enabled}}
<button id="btn-ai-review" class="ui tiny button" data-link="{{.Issue.Link}}/ai-review">
    🤖 AI 评审
</button>
{{end}}
```

> 说明：模板中通过 `setting.AI.Enabled` 控制按钮显示，仅对登录用户且 AI 功能开启时可见。

#### 3.5.2 前端交互脚本
**文件**：`web_src/js/features/repo-diff.ts`（追加）

```typescript
function initAIReviewButton() {
  const btn = document.querySelector('#btn-ai-review');
  if (!btn) return;
  
  btn.addEventListener('click', async () => {
    const link = btn.getAttribute('data-link');
    btn.classList.add('loading', 'disabled');
    
    try {
      const resp = await fetch(link, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          // 从页面 meta/csrf 中获取 token
        },
      });
      
      if (resp.status === 202) {
        showToast('AI 评审任务已提交，请稍后刷新页面查看结果');
      } else {
        const data = await resp.json();
        showToast(data.message || '提交失败', 'error');
      }
    } catch (e) {
      showToast('网络错误', 'error');
    } finally {
      btn.classList.remove('loading', 'disabled');
    }
  });
}

// 在页面初始化时调用
initAIReviewButton();
```

> 注：`showToast` 函数如果项目中已有 toast 组件可直接复用（DevStar/Gitea 前端通常有 `toast` 相关封装）。

#### 3.5.3 AI Review 展示效果
由于 AI Review 是以 `ReviewTypeComment` 写入的，它会**自动复用现有 PR Review UI**：
- 出现在 PR 的 Conversation Timeline 中
- 带有 Reviewer 头像（AI 系统用户的头像）
- 显示 "commented" 标签
- 支持 Markdown 渲染、代码高亮

**Phase 2 视觉增强**：
- 可在 reviewer 用户名旁边加一个 "AI" 徽章
- 这需要修改 `templates/repo/issue/view_content/comments.tmpl` 中对 reviewer 的渲染逻辑

---

### 3.6 PR 自动触发机制（Phase 1.5）

**文件**：`services/pull/pull.go` 或 `services/pull/update.go`

在 PR 创建（`pull_request opened`）或更新（`synchronize`）时，自动检查配置并触发：

```go
func triggerAutoAIReview(ctx context.Context, pr *issues_model.PullRequest) {
    if !setting.AI.Enabled {
        return
    }
    
    // 避免重复触发：检查最近是否已有 AI Review
    hasRecent, err := issues_model.HasRecentAIReview(ctx, pr.IssueID, time.Hour)
    if err != nil || hasRecent {
        return
    }
    
    task := &ai.ReviewTask{
        PullRequestID: pr.ID,
        IssueID:       pr.IssueID,
        RepoID:        pr.BaseRepoID,
        CommitID:      pr.HeadCommitID,
        DoerID:        setting.AI.ReviewerUserID,
    }
    _ = ai.CreateReviewTask(task)
}
```

> 注入点：在 `pull_service.NewPullRequest()` 或 PR patch checker 完成后的回调中调用。

---

### 3.7 Diff 获取策略

Gitea 内部获取 PR diff 有多种方式，推荐以下方案：

#### 方案 A：使用 `gitdiff.GetDiff()`（最稳定）
**参考**：`services/gitdiff/gitdiff.go`

```go
import "code.gitea.io/gitea/services/gitdiff"

opts := gitdiff.DiffOptions{
    BeforeCommitID: pr.MergeBase,
    AfterCommitID:  pr.HeadCommitID,
    // ... 其他字段
}

diff, err := gitdiff.GetDiff(ctx, repo, opts, nil)
```

获取后需要遍历 `diff.Files` 将其转换为纯文本 diff 字符串供 AI 消费。也可以直接用 `git.NewCommand()` 跑 `git diff` 拿到最原始的 diff 文本。

#### 方案 B：直接执行 git diff（最简单）
```go
cmd := git.NewCommand(ctx, "diff", pr.MergeBase, pr.HeadCommitID)
stdout, _, err := cmd.RunStdString(&git.RunOpts{Dir: gitRepo.Path})
```

**推荐方案 B**，因为 Claude API 最擅长解析标准 `git diff` 格式。

---

## 4. 时序图

```
用户          浏览器          API Router      AI Service      Queue      AI Worker      Claude API      DB/Review
 |              |                 |               |             |            |               |             |
 |--点击按钮--> |                 |               |             |            |               |             |
 |              |----POST /ai-review------------> |             |            |               |             |
 |              |                 |---入队()----->|             |            |               |             |
 |              |<---202 Accepted--|               |             |            |               |             |
 |<---提示提交成功-|               |               |             |            |               |             |
 |              |                 |               |---Push()--->|            |               |             |
 |              |                 |               |             |---出队---->|               |             |
 |              |                 |               |             |            |--获取 diff-->  |             |
 |              |                 |               |             |            |<--diff文本-----|             |
 |              |                 |               |             |            |---调用API-------------------->|
 |              |                 |               |             |            |<--Review内容------------------|
 |              |                 |               |             |            |---SubmitReview()----------->|
 |              |                 |               |             |            |                             |--写入--
 |              |                 |               |             |            |                             |        |
 |--刷新页面--> |                 |               |             |            |                             |        |
 |              |----GET PR page----------------> |             |            |                             |        |
 |              |<---显示 AI Review 评论---------|               |            |                             |        |
 |<---看到结果--|                 |               |             |            |                             |        |
```

---

## 5. 安全与权限设计

### 5.1 权限控制
- 触发 AI 评审的用户必须对当前 PR 有 **Read** 权限（至少能看 diff）。
- 建议额外要求用户有 **Write** 或 **Collaborator** 权限，防止外部人员滥用 API 配额。

### 5.2 API Key 安全
- `API_KEY` 仅存储在服务端 `app.ini` 配置中，不暴露给前端。
- 建议支持环境变量注入：`API_KEY = $ANTHROPIC_API_KEY`

### 5.3 Diff 大小限制
- 通过 `MAX_DIFF_SIZE` 控制，默认 150KB。
- 超大 PR 直接截断并提示 AI，避免 token 超限和费用爆炸。

### 5.4 速率控制
- UniqueQueue 天然防重复入队。
- 可扩展在 `CreateReviewTask` 中增加仓库级/PR 级的冷却时间检查（如 1 小时内同一 PR 最多触发 3 次）。

---

## 6. 测试策略

### 6.1 单元测试
- `services/ai/review_test.go`：
  - Mock `gitRepo` 和 Claude API HTTP 调用
  - 测试 diff 截断逻辑
  - 测试队列任务成功/失败重试

### 6.2 集成测试
- 使用 `tests/integration/` 框架：
  - 模拟登录用户触发 `/api/v1/repos/.../ai-review`
  - 验证队列是否正确入队
  - 验证 Review 是否最终出现在 PR 页面

### 6.3 前端测试
- E2E：在 PR diff 页面点击 AI 评审按钮，验证 toast 提示和后续页面刷新后 Review 存在。

---

## 7. 实施阶段（Roadmap）

### Phase 1：MVP — 按钮触发 + 基础 Review
**目标**：让用户能在 PR 页面点击按钮触发 AI 评审，看到结果。
**改动**：
1. `modules/setting/ai.go`
2. `services/ai/` 包（queue + review）
3. `routers/api/v1/repo/pull_review.go` + `api.go` 注册路由
4. `templates/repo/diff/new_review.tmpl` 加按钮
5. `web_src/js/features/repo-diff.ts` 加交互
6. `routers/init.go` 的 `InitWebInstalled` 中注册 `ai.Init()`

### Phase 2：自动触发 + 行级评论
**目标**：PR 创建/更新后自动触发；AI 评审能定位到具体代码行。
**改动**：
1. 在 `services/pull/` 的 PR 生命周期钩子中接入 `triggerAutoAIReview`
2. 解析 Claude API 返回中的行级定位信息
3. 创建 `CommentTypeCode` 行级评论

### Phase 3：PR 合并卡控 + 多模型支持
**目标**：AI 评审不通过则阻止合并；支持 OpenAI/DeepSeek 切换。
**改动**：
1. 在 `services/pull/check.go` 或 `services/pull/merge.go` 中检查最新 AI Review 结论
2. 抽象 `ai.Provider` 接口，支持多模型后端

---

## 8. 附录：需要修改/新增的文件清单

### 新增文件（8 个）
| 文件路径 | 说明 |
|---------|------|
| `modules/setting/ai.go` | AI 配置模块 |
| `services/ai/task.go` | 队列任务定义 |
| `services/ai/queue.go` | 队列注册与 handler |
| `services/ai/review.go` | AI 评审核心逻辑 |
| `services/ai/client.go` | HTTP API 客户端（可选，可合并到 review.go） |
| `services/ai/review_test.go` | 单元测试 |
| `routers/api/v1/repo/pull_review.go` | API 路由 handler（如果现有文件无法复用） |

### 修改文件（8 个）
| 文件路径 | 修改内容 |
|---------|---------|
| `modules/setting/setting.go` | 注册 `loadAIFrom()` |
| `routers/api/v1/api.go` | 注册 `POST .../ai-review` 路由 |
| `routers/init.go` | `InitWebInstalled` 中调用 `ai.Init()` |
| `templates/repo/diff/new_review.tmpl` | 新增 AI 评审按钮 |
| `web_src/js/features/repo-diff.ts` | 新增按钮交互逻辑 |
| `services/pull/pull.go` 或 `update.go` | Phase 2 自动触发注入点 |

---

## 9. 风险与建议

| 风险 | 影响 | 缓解措施 |
|------|------|---------|
| Claude API 调用慢/失败 | 队列任务堆积 | 设置超时、失败重试、错误日志 |
| 超大 PR diff 超出 token 限制 | API 报错或费用高 | MAX_DIFF_SIZE 截断 |
| 同一 PR 被多次触发 | API 费用浪费 | UniqueQueue + 冷却时间检查 |
| AI 评审结果质量不稳定 | 用户体验差 | 持续优化 prompt，第一阶段定位为"辅助" |
| 系统用户不存在 | Review 无法写入 | 启动时检查 `REVIEWER_USER_ID`，不存在则 warn |

---

*文档版本：v1.0*
*编写日期：2026-04-14*
