# coredns 内置插件

coredns是一个dns server，通过插件来扩展功能，
内置插件都是针对各种场景而提供的一些插件，也有扩展插件可用，
下面过一下内置插件，了解为主，重点学习和cncf项目相关的插件

## whoami

- 返回本地ip地址，端口和传输方式
- 指令：插件不管查询名，只要是查询a记录或aaaa记录，whoami插件就会响应
- a记录是ipv4 的ip记录  aaaa是ipv6的
- 默认都会启用这个插件，这个插件可用于查询coredns是否正在响应查询请求
- 在生产环境中，用途有限
- 配置 . { whoami }
- 使用 dig @172.168.10.119 -p 1053 a example.org
- whomai的响应，会作为附加块出现在请求的响应中
- 如下所示，ip、端口、传输方式，分别出现在两条记录中
  - a记录，展示了ip
  - srv记录，展示了传输方式，和端口

```txt
;; ADDITIONAL SECTION:
example.org.		0	IN	A	172.168.10.134
_udp.example.org.	0	IN	SRV	0 0 39127 .
```
说明了客户端的ip/port/传输方式是：172.168.10.134/39127/udp

- a记录，也称主机记录，使用最广，写的就是一个域名对应一个ip
- srv记录 作用是说明一个服务器能够提供什么样的服务
  - ldap._tcp.contoso.com 600 IN SRV 0 100 389 NS.contoso.com
  - ldap：提供ldap服务
  - "_tcp/_udp" 表示传输方式
  - contoso.com：此记录所值的域名
  - 600： 此记录默认生存时间（秒）
  - IN： 标准DNS Internet类
  - SRV：将这条记录标识为SRV记录
  - 0： 优先级，如果相同的服务有多条SRV记录，用户会尝试先连接优先级最低的记录
  - 100：负载平衡机制，多条SRV并且优先级也相同，那么用户会先尝试连接权重高的记录
  - 389：此服务使用的端口
  - NS.contoso.com:提供此服务的主机

## trace

opentracing是一个cncf项目，用于分布式跟踪，这个插件就是跟opentracing对接。

- 作用是跟踪dns请求经过的插件链
- 使用这个插件，就表示开启跟踪功能
- 指令 trace [endpoint-type] [endpoint]
  - endpoint-type, 跟踪目的地类型，支持zipkin和datadog，默认是zipkin
  - endpoint 跟踪的目的地，默认是localhost:9411
- 指令还可以指定跟踪细节：隔多长查询跟踪一个、跟踪指定服务、开启opentracing的某些特征
- 如果要把跟踪数据丢给zipkin，那么要部署一个zipkin服务
- docker run -d -p 9411:9411 openzipkin/zipkin
- 配置 trace zipkin 172.168.10.134:9411     因为docker容器在本机运行，coredns在虚拟机中运行
- 使用 dig @172.168.10.119 -p 1053 a example.org
- 查看结果 zipkin的结果在浏览器上  localhost:9411/zipkin 即可查看刚才请求的详细跟踪数据

zipkin:
- apache基金的一个孵化项目，分布式跟踪系统
- 用于微服务链路跟踪

datago：
- 一个商业的监控分析软件，用于云端应用监控
- 搜集数据，可视化，监控

## tls

之前也提到了，coredns支持3中交互方式：普通的纯dns、tls、grpc, 其中tls和grpc都可以配置tls。

tls插件就可以用到tls和grpc交互方式中。

- coredns利用tls插件，可以让查询请求加密，或是通过grpc交互时加密
- 普通的dns请求是没有加密的
- 指令 tls cert key [ca]
  - ca 可选，如果不写，会使用系统ca

## template

- 允许基于请求，动态响应
- 指令
  - class 查询类，IN或ANY
  - type 查询的dns记录类型
  - zone 区
  - regex 正则
  - rr 资源记录
  - rcode 响应码
  - upstream 解析cname(别名记录)
  - fallthrough

```txt
template CLASS TYPE [ZONE...] {
    match REGEX...
    answer RR
    additional RR
    authority RR
    rcode CODE
    upstream
    fallthrough [ZONE...]
}
```

## secondary

- 从主服务器提供服务，即当前服务是辅助服务器

## root

- 作用是确定在哪个目录查找zone文件
- 默认root是当前工作目录，使用root插件，就可以修改目录

## rewrite

- 支持重写内部消息
- 重写消息，客户端是不知道的
- 简单重写非常fast，复杂重写速度肯定slower
- 指令 rewrite [continue|stop] field [from to|from ttl]
  - field 指明 请求、响应中的哪部分需要重写
    - type 请求的type域要重写，field 后跟的dns记录类型 eg: rewrite type any hinfo 查询记录类型重any改为hinfo
    - class 重写消息的class， field后跟 dns class类型 eg: rewrite class ch in  将ch改为in
    - name 重写请求的查询名
    - answer name，修改响应中的查询名
    - edns0 请求中附加的edns0选项
    - ttl 重写响应中的ttl值
  - from 名字或类型，在匹配时用
  - to 修改后的值
  - ttl 修改后的ttl

## prometheus

- prometheus 普罗米修斯
- cncf下的一个毕业项目，作用是监控
- 这个插件就是和普罗米修斯对接
- 只要使用这个插件，就可以在coredns或任何插件中获取要监控的指标
- 指令 prometheus [ADDRESS] 
- 查看数据 默认是localhost:9153/metrics看数据

## loadbalance

- 提供一个负载均衡的策略
- 指令 loadbalance [policy] 目前只有一个策略 就是 round_robin 轮询调度
- 这就是dns解析中的负载均衡，支持a记录 aaaa记录 mx记录

## kubernetes

- 对接cncf第一个毕业项目 k8s
- 这个插件给k8s提供了基于dns服务发现的解决方案，已经成功替代了k8s自身的服务发现组件
- 取代k8s集群中的kube-dns，成为k8s内置dns服务发现组件
- 指令 kubernets [zones ...]
- 插件可以处理查询请求的同时，还可以连接k8s集群
- k8s service不提供rtp记录查询；k8s pod不提供a记录查询

详细配置：

```txt
kubernetes [ZONES...] {
    resyncperiod DURATION  指定k8s 数据api的期限，默认是0，表示不同步
    endpoint URL           k8s远端api端点，如果忽略，就使用集群服务账号
    tls CERT KEY CACERT    tls证书
    kubeconfig KUBECONFIG CONTEXT   连接k8s集群时的认证信息，支持tls/用户密码/token等方式，是可选的
    namespaces NAMESPACE... 指定k8s的namesapce，如果忽略，那就能访问k8s的所有namespace
    labels EXPRESSION      用标签选择器来匹配k8s对象资源
    pods POD-MODE          处理基于ip的pod的a记录
    endpoint_pod_names     处理a记录时，用端点名来触发pod名
    upstream [ADDRESS...]  定义上游解析器，一般上游解析器是外部主机，eg：cname记录
    ttl TTL                定义响应的ttl
    noendpoints            关闭一些模式
    transfer to ADDRESS... 转发
    fallthrough [ZONES...] 
    ignore empty_service
}
```

## import

和log一样，都是一些功能性插件，给其他插件用。

- 使用在corefile中

## etcd

- etcd 是cncf的一个孵化项目，这个插件就是和etcd对接
- 作用是 从etcd v3实例中读取zone数据
- etcd插件使用forward插件去转发和查询其他服务
- 指令 etcd [zones ...]
- 因为etcd v3 的存储结构采用了扁平设计，所以插件处理时需要注意以前的目录式改为现在的前缀式

```txt
etcd [ZONES...] {
    fallthrough [ZONES...]
    path PATH                 在etcd实例中，key是"/skydns"
    endpoint ENDPOINT...      etcd端点，默认是 http://localhost:2379
    credentials USERNAME PASSWORD
    upstream [ADDRESS...]
    tls CERT KEY CACERT
}
```

etcd中可存a记录 aaaa记录 srv记录 txt记录 等等



