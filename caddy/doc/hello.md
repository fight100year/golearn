# caddy 介绍

- 2015年4月开始的一个开源项目，截止目前为止，star数 2w+，与目前如日中天的etcd的star差不多
- 主要由个人维护，v1.0.0已发布，目前主要以稳定和解决bug为主，更新放缓
- 一个通用的 http2 web server，特征是默认使用https
- 基本支持全平台
- cncf 毕业项目coredns 底层就使用了caddy，说明在go实现的webserver中，还是有一定优势的
- 数字海洋DigitalOcean 赞助主题给caddy运行，数字海洋是新兴云服务提供商，和亚马逊云、google云、微软云一起号称4大国外公有云

## 特征

- 易于配置，直接使用caddyfile配置即可
- 默认自动支持https，通过let's encrypt 自动申请证书，自动续期
- 默认使用http2
- 支持virtual host，这样多个站点也可以很好工作
- quic支持，实现性质的，quic是http3,2018年年底更名为http3，属于尖端技术
- 为了安全连接，使用tls会话的key rotation
- 插件可扩展，所以功能可以通过插件进行扩展
- 基于go，所以没有额外依赖


quic ：
- cutting-edge technology 尖端传输技术，尖端表示最新最前沿的技术
- Quick UDP Internet Connection 谷歌制定的一种基于UDP的低时延的互联网传输层协议
- 传输层除了tcp/udp, 新贵就是quic，2016.11 ietf开始了标准化之路
- 牛在哪：融合了tcp/tls/http2等特性，基于udp，在效率和部署上都有很大的优势
- 2018.10.28，标准化组织ietf，正式更名为http/3协议，再一次贴了一个标签:牛
- 腾讯的qq空间已经开始支持quic，利用的正式caddy库

## 安装

二进制安装，直接去官网下载

编译：
- v1.0.0，需要go >= 1.12
  - 从1.11升级到1.12： 
  - sudo rm -rf /usr/local/go
  - sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
  - go version
- 开启mod模式，这是新的包管理工具
  - export GO111MODULE=on
  - export GOPROXY=https://goproxy.io
  - 开启mod之后，用go get下载的模块都丢到pkg/mod/下了
- 下载caddy go get github.com/mholt/caddy/caddy

## 第一个例子

```go
package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	// plug in plugins here, for example:
	// _ "import/path/here"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}
```

- go mod init
- go mod tidy
- go run hello.go
- 需要注意，如果包含的库路径写错了，会报错，且不容易查出来。
- 之后可以通过 http://localhost:2015来访问，目前没有任何资源文件，所以访问是404

