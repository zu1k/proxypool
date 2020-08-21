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

![Star](https://img.shields.io/github/stars/zu1k/proxypool.svg?style=social&label=Star) 来都来了，不点个Star再走？ 

## 支持

- 支持ss、ssr、vmess、trojan节点链接与订阅
- 任意 Telegram 频道抓取
- 机场订阅地址抓取解析
- 公开互联网页面模糊抓取
- 翻墙党论坛新节点信息
- 其他节点分享网站
- 定时抓取更新
- 使用配置文件提供抓取源
- 自动检测节点可用性

## 安装

### 使用Heroku

点击按钮进入部署页面，填写基本信息然后运行

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

### 使用docker

```sh
docker pull docker.pkg.github.com/zu1k/proxypool/proxypool:latest
```

## 使用

### 修改配置文件

首先修改 source.yaml 中的必要配置信息，其中域名修改为你自己的域名，cf开头的配置信息可以留空

### 启动程序

```shell
proxypool -c source.yaml
```

### 用户使用

目前公开版本： https://proxy.tgbot.co

访问页面，按照相关指导进行使用

## 截图

![Speedtest](docs/speedtest.png)

![Fast](docs/fast.png)

## 声明

本项目遵循 GNU General Public License v3.0 开源，在此基础上，所有使用本项目提供服务者都必须在网站首页保留指向本项目的链接

禁止使用本项目进行营利和做其他违法事情，产生的一切后果本项目概不负责
