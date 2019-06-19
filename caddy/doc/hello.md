# caddy 介绍

- 2015年4月开始的一个开源项目，截止目前为止，star数 2w+，与目前如日中天的etcd的star差不多
- 主要由个人维护，v1.0.0已发布，目前主要以稳定和解决bug为主，更新放缓
- 一个通用的 http2 web server，特征是默认使用https
- 基本支持全平台
- cncf 毕业项目coredns 底层就使用了caddy，说明在go实现的webserver中，还是有一定优势的

## 特征

- 易于配置，直接使用caddyfile配置即可
- 默认自动支持https，通过let's encrypt 自动申请证书，自动续期
- 默认使用http2
- 支持virtual host，这样多个站点也可以很好工作
- quic支持，实现性质的


quic ：
- cutting-edge technology 尖端传输技术，尖端表示最新最前沿的技术
- Quick UDP Internet Connection 谷歌制定的一种基于UDP的低时延的互联网传输层协议
- 传输层除了tcp/udp, 新贵就是quic，2016.11 ietf开始了标准化之路
- 牛在哪：融合了tcp/tls/http2等特性，基于udp，在效率和部署上都有很大的优势
- 2018.10.28，标准化组织ietf，正式更名为http/3协议，再一次贴了一个标签:牛
-
