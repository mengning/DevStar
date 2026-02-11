# DevStar

The Last Mile of Al for R&D. Make it possible to code everything.

Delivering a truly AI‑Native, all‑in‑one R&D ecosystem, and designed for hybrid teams of AI and human developers, DevStar is the next‑generation intelligent platform built for the AI era.

- Out-of-the-box. DevStar provides a one-stop R&D system with a simple installation process, allowing you to quickly set up your development environment and start coding immediately!
- AI-Native. From Copilot, Claude Code, and OpenCode to OpenClawd, Trae IDE, and Cursor IDE, DevStar seamlessly integrates with the latest AI tools. It deeply embeds AI Code Review, AI assistants, and MCP Server to deliver a truly AI-Native, all-in-one R&D ecosystem. Designed for hybrid teams of AI and human developers, DevStar is the next-generation intelligent platform built for the AI era.
- Cloud-Native. Provides a Cloud-Native development environment with DevContainer, supports one-click deployment of Cloud-Native R&D tools in Docker and Kubernetes environments, such as CI/CD pipeline Runners, Cloudbuild distributed compilation systems, private code LLM, and more!
- Everything as Code. Supporting mainstream operating systems such as Ubuntu, openEuler, and Alpine Linux as DevContainers and Actions Runners via Dockerfile enables "Everything as Code." Make it possible to code everything.

DevStar is a commercial distribution of Gitea. As [Gitea](https://github.com/go-gitea/gitea) is written in Go, it works across **all** the platforms and
architectures that are supported by Go, including Linux, macOS, and
Windows on x86, amd64, ARM and PowerPC architectures.
This project has been
[forked](https://blog.gitea.com/welcome-to-gitea/) from
[Gogs](https://gogs.io) since November of 2016, but a lot has changed.

| Feature Type | Capabilities | License | Pricing Model |
|-------------|-------------|---------|--------------|
| **Core Features** | ✅ Git Repository Hosting<br>✅ Issues & Pull Requests<br>✅ Wiki & Project Boards<br>✅ Built-in CI/CD (GitHub Actions Compatible)<br>✅ Granular Permissions<br>✅ User & Organization Management<br>✅ Multi-Package Registry<br>✅ REST API<br>✅ Multi-DB Support / Lightweight Deployment | MIT License | Free & Open Source |
| **Enhanced Features** | ✅ Built-in DevContainer Environment<br>✅ Built-in MCP Server<br>✅ Built-in AI Assistant<br>✅ Integration with AI CLI Tools (Claude Code/OpenCode/etc.)<br>✅ Integration with AI IDEs (Copilot/Cursor/Trae/etc.)<br>✅ One-click Actions Runner Deployment<br>✅ Online CI/CD Script Debugging<br>✅ Project Template Ecosystem (Best Practice Templates) | Commercial License | **Personal Use:** Free Forever (Non-commercial)<br>**Enterprise Use:** Annual License<br>✅ Standard: CNY 1,888/user/year (Human or AI)<br>✅ Self-Report Discount: CNY 188/user/year (90% off)<br>✅ First-Year Trial: CNY 1.88/user |

## Quick Start

```bash
curl -fsSL https://devstar.cn/install  | bash
devstar start # Run command to deploy DevStar Studio
```

For online demonstrations, you can visit [DevStar.cn](https://DevStar.cn).

You can find comprehensive documentation on our official [documentation website](https://mengning.com.cn).

## Building

From the root of the source tree, run:

    TAGS="bindata" make build

or if SQLite support is required:

    TAGS="bindata sqlite sqlite_unlock_notify" make build

The `build` target is split into two sub-targets:

- `make backend` which requires [Go Stable](https://go.dev/dl/), the required version is defined in [go.mod](/go.mod).
- `make frontend` which requires [Node.js LTS](https://nodejs.org/en/download/) or greater and [pnpm](https://pnpm.io/installation).

Internet connectivity is required to download the go and npm modules. When building from the official source tarballs which include pre-built frontend files, the `frontend` target will not be triggered, making it possible to build without Node.js.

More info: https://docs.gitea.com/installation/install-from-source

## Using

After building, a binary file named `gitea` will be generated in the root of the source tree by default. To run it, use:

    ./gitea web

> [!NOTE]
> If you're interested in using our APIs, we have experimental support with [documentation](https://docs.gitea.com/api).

## Contributing

Expected workflow is: Fork -> Patch -> Push -> Pull Request

> [!NOTE]
>
> 1. **YOU MUST READ THE [CONTRIBUTORS GUIDE](CONTRIBUTING.md) BEFORE STARTING TO WORK ON A PULL REQUEST.**
> 2. If you have found a vulnerability in the project, please write privately to **security@gitea.io**. Thanks!

## Translating

[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://translate.gitea.com)

Translations are done through [Crowdin](https://translate.gitea.com). If you want to translate to a new language, ask one of the managers in the Crowdin project to add a new language there.

You can also just create an issue for adding a language or ask on Discord on the #translation channel. If you need context or find some translation issues, you can leave a comment on the string or ask on Discord. For general translation questions there is a section in the docs. Currently a bit empty, but we hope to fill it as questions pop up.

Get more information from [documentation](https://docs.gitea.com/contributing/localization).

## Official and Third-Party Projects

We provide an official [go-sdk](https://gitea.com/gitea/go-sdk), a CLI tool called [tea](https://gitea.com/gitea/tea) and an [action runner](https://gitea.com/gitea/act_runner) for Gitea Action.

We maintain a list of Gitea-related projects at [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea), where you can discover more third-party projects, including SDKs, plugins, themes, and more.

## Communication

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

If you have questions that are not covered by the [documentation](https://docs.gitea.com/), you can get in contact with us on our [Discord server](https://discord.gg/Gitea) or create a post in the [discourse forum](https://forum.gitea.com/).

## Authors

- [Maintainers](https://github.com/orgs/go-gitea/people)
- [Contributors](https://github.com/go-gitea/gitea/graphs/contributors)
- [Translators](options/locale/TRANSLATORS)

## Backers

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/gitea#backer)]

<a href="https://opencollective.com/gitea#backers" target="_blank"><img src="https://opencollective.com/gitea/backers.svg?width=890"></a>

## Sponsors

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/gitea#sponsor)]

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

## FAQ

**How do you pronounce Gitea?**

Gitea is pronounced [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY) as in "gi-tea" with a hard g.

**Why is this not hosted on a Gitea instance?**

We're [working on it](https://github.com/go-gitea/gitea/issues/1029).

**Where can I find the security patches?**

In the [release log](https://github.com/go-gitea/gitea/releases) or the [change log](https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md), search for the keyword `SECURITY` to find the security patches.

## License

this project is licensed under the MIT License. The project name DevStar, together with the registered trademark and official domain devstar.cn, constitutes its brand identity. Derivative works are not permitted to use the same or similar names without permission, in order to avoid public confusion.

## Further information

<details>
<summary>Looking for an overview of the interface? Check it out!</summary>

### Login/Register Page

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### User Dashboard

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### User Profile

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### Explore

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### Repository

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### Repository Issue

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### Repository Pull Requests

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### Repository Actions

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### Repository Activity

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### Organization

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
