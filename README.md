# DevStar Studio

DevStar Studio 是一个Gitea 发行版，在Git代码仓库托管的基础上提供了开发环境DevEnv执行引擎，与VS Code插件或自定义IDE深度融合，形成灵活适配基础软件工具的生态平台，从而为开发者用户提供智能（代码大模型AI+）、安全（完全云原生）、一站式开箱即用的CI/CD全生命周期研发平台。

DevStar Studio是一个通用的一站式软件研发平台，但它最初的目标是服务于汽车软件、消费电子、智能制造等嵌入式软件研发场景中的开发者。

DevStar Studio的愿景：服务全球软件开发者！

如果你想试用在线演示或者使用免费的DevStar服务（有数量限制），请访问 [devstar.cn](https://devstar.cn/)。

如果你想快速本地部署自己的DevStar实例免费试用或者报告问题，请访问 [https://github.com/mengning/DevStar](https://github.com/mengning/DevStar)。

如果你是云服务厂商想为您的客户提供DevStar实例请联系contact@mengning.com.cn

## Quick Start 快速开始

如果您是在Windows环境下，请先运行如下命令：

```
powershell wsl --install -d Ubuntu-20.04
```

在Ubuntu-20.04下完成安装：

```bash
sudo apt update
sudo apt install docker.io
sudo docker pull devstar.cn/devstar/devstar-studio:latest
sudo docker run --restart=always --name devstar-studio -d  -p 8080:3000 -v /var/run/docker.sock:/var/run/docker.sock devstar.cn/devstar/devstar-studio:latest
```

打开 `http://localhost:8080` 完成安装。
