# 介绍

coredns称为2019年首个毕业项目 2019.1.25
- 也是继k8s、prometheux、envoy之后的第4个cncf毕业项目
- 项目成立于2016.3，google的一个工程师建立
- 2017年加入cncf，称为沙箱项目，2018.2称为孵化项目
- 插件式的架构设计，易于集成和扩展

CoreNDS提供了两大功能：
- DNS
- service discovery

CoreDNS现状：
- v1.5.0 发布于2019.4.6
- k8s默认集群dns服务

coredns是什么：
- 一个go编写的dns服务
- 相比其他dns服务来说，优势是灵活：在很多环境都很适用：k8s、以及其他混合云
- 开源
- cncf还有一个公共的聊天交流频道 https://cloud-native.slack.com 
- 好了 coredns的主职是dns服务，同时可以利用dns的svr记录进行服务发现
- 插件式，容易被其他项目集成
- 简单 本身就有一套默认配置
- 服务发现，副业，配合etcd完成服务发现，或是集成到k8s提供服务发现
- 快速 灵活，这也是设计目标，也是项目发起的初衷
- cncf毕业项目

## 官方历史

- coredns项目起源于caddy项目，起初是给caddy项目做一个coredns插件，名字还取了很多，有叫daddy 也有叫caddy-dns, 2016.5.10 
- 整个项目的定位有了初步思考后，取名为caddy-dns, 2016.5.14
- 以caddy-dns为名，正式开工， 2016.5.17
- 官方宣布使用 coredns，不用caddy-dns， 2016.5.18

起于caddy，也继承了caddy的优良血统：链式插件。所以灵活性足够。
