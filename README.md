<h1 align="center">
  <br>proxypool<br>
</h1>

<h5 align="center">自动抓取tg频道、订阅地址、公开互联网上的ss、ssr、vmess、trojan节点信息，聚合去重测试可用性后提供节点列表</h5>

<p align="center">
  <a href="https://github.com/zu1k/proxypool/actions">
    <img src="https://img.shields.io/github/workflow/status/zu1k/proxypool/Go?style=flat-square" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/report/github.com/zu1k/proxypool">
    <img src="https://goreportcard.com/badge/github.com/zu1k/proxypool?style=flat-square">
  </a>
  <a href="https://github.com/zu1k/proxypool/releases">
    <img src="https://img.shields.io/github/release/zu1k/proxypool/all.svg?style=flat-square">
  </a>
</p>

## 支持

- 支持ss、ssr、vmess、trojan多种类型
- Telegram频道抓取
- 订阅地址抓取解析
- 公开互联网页面模糊抓取
- 定时抓取自动更新
- 通过配置文件设置抓取源
- 自动检测节点可用性

## 安装

### 使用Heroku

点击按钮进入部署页面，填写基本信息然后运行

其中 `DOMAIN` 需要填写为你需要绑定的域名，`CONFIG_FILE` 需要填写你的配置文件路径，配置文件模板见 config/config.yaml 文件

`CF` 开头的选项暂不需要填写，不影响程序运行

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

> 因为爬虫程序需要持续运行，所以至少选择 $7/月 的配置
> 免费配置长时间无人访问会被heroku强制停止

### 从源码编译

需要安装Golang

```sh
$ go get -u -v github.com/zu1k/proxypool
```

### 下载预编译程序

从这里下载预编译好的程序 [release](https://github.com/zu1k/proxypool/releases)

除了编译好的二进制程序，你还需要下载仓库 `assets` 文件夹，放置在可执行文件同目录

### 使用docker

```sh
docker pull docker.pkg.github.com/zu1k/proxypool/proxypool:latest
```

## 使用

### 修改配置文件

首先修改 config.yaml 中的必要配置信息，cf开头的选项不需要填写

source.yaml 文件中定义了抓取源，需要定期手动维护更新

### 启动程序

使用 `-c` 参数指定配置文件路径，支持http链接

```shell
proxypool -c config.yaml
```

## 截图

![Speedtest](docs/speedtest.png)

![Fast](docs/fast.png)

## 声明

本项目遵循 GNU General Public License v3.0 开源，在此基础上，所有使用本项目提供服务者都必须在网站首页保留指向本项目的链接

禁止使用本项目进行营利和做其他违法事情，产生的一切后果本项目概不负责
