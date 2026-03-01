# DevStar

AI 赋能研发的“最后一公里”。让一切皆可编程。

DevStar 提供真正 AI 原生的全栈研发生态，专为 AI 与人类开发者组成的混合团队设计，是面向 AI 时代打造的新一代智能平台。

-   **开箱即用**。DevStar 提供一站式研发系统，安装流程简洁，助你快速搭建开发环境，立即投入编程！
-   **AI 原生**。从 Copilot、Claude Code、OpenCode 到 OpenClawd、Trae IDE、Cursor IDE，DevStar 无缝集成最新的 AI 工具。它深度嵌入 AI 代码审查、AI 助手和 MCP 服务器，提供真正 AI 原生的一体化研发生态。专为 AI 与人类开发者组成的混合团队设计，是面向 AI 时代打造的新一代智能平台。
-   **云原生**。提供云原生开发环境，支持 DevContainer，可在 Docker 和 Kubernetes 环境中一键部署云原生研发工具，如 CI/CD 流水线 Runner、Cloudbuild 分布式编译系统、私有化代码大语言模型等！
-   **一切即代码**。支持通过 Dockerfile 将 Ubuntu、openEuler、Alpine Linux 等主流操作系统作为 DevContainers 和 Actions Runners 运行，实现“一切即代码”。让一切皆可编程。

DevStar 是 Gitea 的商业发行版。由于 [Gitea](https://github.com/go-gitea/gitea) 采用 Go 语言编写，它可在 Go 支持的所有平台和架构上运行，包括 Linux、macOS、Windows 的 x86、amd64、ARM 和 PowerPC 架构。该项目自 2016 年 11 月从 [Gogs](https://gogs.io) [分叉](https://blog.gitea.com/welcome-to-gitea/) 而来，但已发生了巨大变化。

| 功能类型 | 能力 | 许可证 | 定价模式 |
| :--- | :--- | :--- | :--- |
| **核心功能** | ✅ Git 仓库托管<br>✅ 议题与拉取请求<br>✅ Wiki 与项目看板<br>✅ 内置 CI/CD（兼容 GitHub Actions）<br>✅ 细粒度权限控制<br>✅ 用户与组织管理<br>✅ 多包注册表<br>✅ REST API<br>✅ 多数据库支持 / 轻量级部署 | MIT 许可证 | 免费开源 |
| **增强功能** | ✅ 内置 DevContainer 环境<br>✅ 内置 MCP 服务器<br>✅ 内置 AI 助手<br>✅ 集成 AI CLI 工具（Claude Code/OpenCode 等）<br>✅ 集成 AI IDE（Copilot/Cursor/Trae 等）<br>✅ 一键部署 Actions Runner<br>✅ 在线调试 CI/CD 脚本<br>✅ 项目模板生态（最佳实践模板） | 商业许可证 | **个人使用：** 永久免费（非商业用途）<br>**企业使用：** 年度许可<br>✅ 标准版：CNY 1,888/用户/年（人类或 AI 用户）<br>✅ 自主申报折扣：CNY 188/用户/年（9折优惠）<br>✅ 首年试用：CNY 1.88/用户 |

## 快速开始

```bash
curl -fsSL https://devstar.cn/install  | bash
devstar start # 运行命令部署 DevStar Studio
```

如需在线演示，请访问 [DevStar.cn](https://DevStar.cn)。

你可以在我们的官方 [文档网站](https://github.com/mengning/DevStar/wiki) 找到完整文档。

## 構建

從源代碼樹的根目錄運行：

    TAGS="bindata" make build

如果需要 SQLite 支援：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目標分為兩個子目標：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本在 [go.mod](/go.mod) 中定義。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更高版本。

需要互聯網連接來下載 go 和 npm 模塊。從包含預構建前端文件的官方源代碼壓縮包構建時，不會觸發 `frontend` 目標，因此可以在沒有 Node.js 的情況下構建。

更多信息：https://docs.gitea.com/installation/install-from-source

## 使用

構建後，默認情況下會在源代碼樹的根目錄生成一個名為 `gitea` 的二進制文件。要運行它，請使用：

    ./gitea web

> [!注意]
> 如果您對使用我們的 API 感興趣，我們提供了實驗性支援，並附有 [文件](https://docs.gitea.com/api)。

## 貢獻

預期的工作流程是：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **在開始進行 Pull Request 之前，您必須閱讀 [貢獻者指南](CONTRIBUTING.md)。**
> 2. 如果您在項目中發現了漏洞，請私下寫信給 **security@gitea.io**。謝謝！

## 翻譯

[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com)

翻譯通過 [Crowdin](https://translate.gitea.com) 進行。如果您想翻譯成新的語言，請在 Crowdin 項目中請求管理員添加新語言。

您也可以創建一個 issue 來添加語言，或者在 discord 的 #translation 頻道上詢問。如果您需要上下文或發現一些翻譯問題，可以在字符串上留言或在 Discord 上詢問。對於一般的翻譯問題，文檔中有一個部分。目前有點空，但我們希望隨著問題的出現而填充它。

更多信息請參閱 [文件](https://docs.gitea.com/contributing/localization)。

## 官方和第三方項目

我們提供了一個官方的 [go-sdk](https://gitea.com/gitea/go-sdk)，一個名為 [tea](https://gitea.com/gitea/tea) 的 CLI 工具和一個 Gitea Action 的 [action runner](https://gitea.com/gitea/act_runner)。

我們在 [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea) 維護了一個 Gitea 相關項目的列表，您可以在那裡發現更多的第三方項目，包括 SDK、插件、主題等。

## 通訊

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

如果您有任何文件未涵蓋的問題，可以在我們的 [Discord 服務器](https://discord.gg/Gitea) 上與我們聯繫，或者在 [discourse 論壇](https://forum.gitea.com/) 上創建帖子。

## 作者

- [維護者](https://github.com/orgs/go-gitea/people)
- [貢獻者](https://github.com/go-gitea/gitea/graphs/contributors)
- [翻譯者](options/locale/TRANSLATORS)

## 支持者

感謝所有支持者！ 🙏 [[成為支持者](https://opencollective.com/gitea#backer)]

<a href="https://opencollective.com/gitea#backers" target="_blank"><img src="https://opencollective.com/gitea/backers.svg?width=890"></a>

## 贊助商

通過成為贊助商來支持這個項目。您的標誌將顯示在這裡，並帶有鏈接到您的網站。 [[成為贊助商](https://opencollective.com/gitea#sponsor)]

<a href="https://opencollective.com/gitea/sponsor/0/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/1/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/2/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/3/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/4/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/5/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/6/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/7/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/8/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/gitea/sponsor/9/website" target="_blank"><img src="https://opencollective.com/gitea/sponsor/9/avatar.svg"></a>

## 常見問題

**Gitea 怎麼發音？**

Gitea 的發音是 [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY)，就像 "gi-tea" 一樣，g 是硬音。

**為什麼這個項目沒有託管在 Gitea 實例上？**

我們正在 [努力](https://github.com/go-gitea/gitea/issues/1029)。

**在哪裡可以找到安全補丁？**

在 [發佈日誌](https://github.com/go-gitea/gitea/releases) 或 [變更日誌](https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md) 中，搜索關鍵詞 `SECURITY` 以找到安全補丁。

## 許可證

這個項目是根據 MIT 許可證授權的。
项目名称DevStar与已注册相关商标及官方域名devstar.cn共同构成品牌标识，未经许可，衍生产品不得使用相同或近似名称，以免造成公众混淆。翻译成英语

## 進一步信息

<details>
<summary>尋找界面概述？查看這裡！</summary>

### 登錄/註冊頁面

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### 用戶儀表板

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### 用戶資料

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### 探索

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### 倉庫

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### 倉庫問題

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### 倉庫拉取請求

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### 倉庫操作

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### 倉庫活動

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### 組織

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
