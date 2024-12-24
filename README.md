# DevStar Studio

DevStar Studio is a Gitea distribution that provides a development environment execution engine (DevEnv) based on Git code repository hosting. It deeply integrates with VS Code plugins or custom IDEs, forming an ecosystem platform that flexibly adapts basic software tools, thereby offering intelligent (AI-powered code models), secure (fully cloud-native), and one-stop ready-to-use CI/CD lifecycle development platforms for developer users.

DevStar Studio is a general-purpose, one-stop software development platform, but its initial goal is to serve developers in embedded software development scenarios such as automotive software, consumer electronics, and intelligent manufacturing.

Vision of DevStar Studio: To serve software developers worldwide!

If you want to try an online demo or use the free DevStar service (with usage limits), please visit [devstar.cn](https://devstar.cn/) or search for DevStar in the VS Code plugin marketplace.

If you encounter any issues while using DevStar, feel free to submit a [Bug Report](https://github.com/mengning/DevStar/issues/new).

If you are a cloud service provider looking to offer DevStar instances to your customers, please contact contact@mengning.com.cn.

DevStar Studio 是一个Gitea 发行版，在Git代码仓库托管的基础上提供了开发环境DevEnv执行引擎，与VS Code插件或自定义IDE深度融合，形成灵活适配基础软件工具的生态平台，从而为开发者用户提供智能（代码大模型AI+）、安全（完全云原生）、一站式开箱即用的CI/CD全生命周期研发平台。

DevStar Studio是一个通用的一站式软件研发平台，但它最初的目标是服务于汽车软件、消费电子、智能制造等嵌入式软件研发场景中的开发者。

DevStar Studio的愿景：服务全球软件开发者！

如果你想试用在线演示或者使用免费的DevStar服务（有数量限制），请访问 [devstar.cn](https://devstar.cn/) 或者在VS Code插件市场搜索DevStar。

如果您在使用DevStar中有任何问题欢迎提交[Bug Report](https://github.com/mengning/DevStar/issues/new)。

如果你是云服务厂商想为您的客户提供DevStar实例请联系contact@mengning.com.cn

## Quick Start 快速开始

如果您是在Windows环境下，请在cmd命令行下先运行如下命令：

```
powershell wsl --install -d Ubuntu-20.04
```
如果以上命令无法自动完成安装WSL，可以手动执行如下两条命令完成安装：
```
dism.exe /Online /Enable-Feature /FeatureName:Microsoft-Windows-Subsystem-Linux /All /LimitAccess /All
dism.exe /Online /Enable-Feature /FeatureName:VirtualMachinePlatform /All /LimitAccess
```

在Ubuntu-20.04下完成安装：

```bash
wget -c https://devstar.cn/assets/install.sh && chmod +x install.sh && sudo ./install.sh
devstar help
------------------------------------------------------------------------
DevStar usage help:
  help, -h, --help,     Help information
  start                 Start DevStar Studio
      --port=<arg>      Specify the port number (default port is 8080)
      --version=<arg>   Specify the DevStar Studio Image Version (default verson is latest)
  stop                  Stop the running DevStar Studio
  logs                  View the logs of the devstar-studio container
  clean                 Clean up the running DevStar Studio, including deleting user data. Please use with caution.
------------------------------------------------------------------------
sudo devstar start
...
-------------------------------------------------------
DevStar started in http://localhost:8080 successfully!
-------------------------------------------------------
```

然后通过浏览器打开http://localhost:8080 完成后配置安装，默认第一个注册用户为管理员账户。

## License Agreement

* FREE USE - Only one instance is allowed, and the number of users is limited to less than 50, excluding technical support services.
* The standalone version of DevStar Studio can be deployed for free, while deployment based on Kubernetes in a cloud-native environment requires commercial authorization . It is especially noted that both the standalone version and the Kubernetes cloud-native deployment have the same features, but the standalone version is limited in capacity and scalability.
