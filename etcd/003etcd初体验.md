# etcd 初次接触

etcd是go实现的，下载下来可直接运行，加入环境变量即可。

- 启动： ./etcd
- 查看版本： etcd --version
- 默认端口是 2379/2380
- 使用etcdctl工具可与etcd服务交互

- 读写：
    - ETCDCTL_API=3 etcdctl --endpoints=localhost:2379 put foo abc
    - ETCDCTL_API=3 etcdctl --endpoints=localhost:2379 get foo
    - ETCDCTL_API=3表示使用etcd v3，--endpoints表示etcd服务地址， put/get表示读写，后面表示key和value

## 单机部署多etcd实例

使用goreman来管理单机上的多进程，前提是需要go环境。
- 下载go的二进制包
- 配置go环境，下载goreman

```shell
export GOPATH=$HOME/go-ws
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin:$HOME/etcd/etcd-v3.3.13
export ETCDCTL_API=3

go get github.com/mattn/goreman
```

etcd的下载按官网下载安装即可，上面的环境配置中也添加了etcd的目录，接下来用goreman来管理etcd集群：

```shell
# 单机多实例

# etcd1  
etcd1: etcd --name infra1 --listen-client-urls http://127.0.0.1:12379 --advertise-client-urls http://127.0.0.1:12379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new

# etcd2
etcd2: etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new

# etcd3
etcd3: etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new
```

命令参数说明:  

    --name infra1 表示etcd节点名,不重复即可
    --listen-client-urls 客户端连接地址，可监听多个
    --advertise-client-urls 推荐的客户端地址，这个地址要告诉其他成员，用于请求重定向,proxy模式时要注意避免死循环
    --listen-peer-urls 节点间通信的地址
    --initial-advertise-peer-urls 推荐使用的节点间通信地址
    --initial-cluster-token 集群启动时初始化的一个集群令牌，多集群中标识自己
    --initial-cluster 初始化集群配置，这个配置里包含了所有成员之间的通信地址
    --initial-cluster-state 新建集群的标识
上面goreman创建了一个单机集群，节点间的访问地址是12380/22380/32380，客户端访问地址是12379/22379/32379.

显示集群成员：etcdctl --endpoints=http://localhost:12379 member list  

## 多节点集群化部署

有两种启动方式：
- 静态配置
- 服务发现

静态配置：
- 适用于线下环境，有两个条件：
    - 节点个数已知
    - 节点地址已知，在--initial-cluster中指定
- 需要注意以下错误
    - 节点地址枚举要有，如果没有就不满足静态配置的两个条件
    - 节点间的地址，要和枚举的一致，不然无法匹配
    - 为了确定具体某一集群，最好带集群令牌

服务发现：
- 适用于不知道集群成员地址的情况，eg：使用dhcp网络或使用云提供商提供的环境时
- 服务的自发现有两种模式，etcd自发现模式和dns自发现模式

etcd自发现模式：

```shell
    curl 'https://discovery.etcd.io/new?size=3'
    https://discovery.etcd.io/21ccbd1a72655d22659bc5ee3bc2e05a

    # etcd服务自发现

# etcd1
etcd1: etcd --name infra1 --listen-client-urls http://172.168.10.119:2379,http://127.0.0.1:2379 --advertise-client-urls http://172.168.10.119:2379 --listen-peer-urls http://172.168.10.119:2380 --initial-advertise-peer-urls http://172.168.10.119:2380 --discovery https://discovery.etcd.io/21ccbd1a72655d22659bc5ee3bc2e05a

```

相比静态配置，etcd自发现不需要指定集群令牌、集群状态、枚举信息

dns自发现模式：
- 利用dns的srv记录来进行服务发现，需要在dns服务器上配置，暂时不实验了。

## etcdctl 常用命令行

- 使用v3的api，export ETCDCTL_API=3，我们已经在环境变量中配置了。
- 写 etcdctl put foo 123
- 缓存写 etcdctl put foo 123 --lease=abc 利用租约来设置老化时间 abc引用的是x分钟，过期查询会返回100，表示key不存在
- 读 etcdctl get foo
    - --print-value-only 只打印值
    - etcdctl get foo foo3 会打印foo到foo3范围内的key，半开区间[foo,foo3)
    - etcdctl get --prefix foo 遍历前缀是foo的key
    - etcdctl get --prefix --limit=2 foo 限制输出数量为2
- 读取老版本的key etcdctl get --prefix --limit=2 --rev=5 foo
- 按key的字段序来读取 etcdctl get --from-key foo 取大于等于foo字节值的key
- 删除一个key或一个key范围
    - etcdctl del foo
    - etcdctl del foo foo1
    - etcdctl del foo --prev-kv  删除key的同时，返回响应的value
    - etcdctl del --prefix foo 按前缀删除
    - etcdctl del --from-key foo 删除大于等于foo的所有key
- 观察一个值，除非遇到退出信号量，否则一致等待
    - etcdctl watch foo
    - etcdctl watch foo foo100
    - etcdctl watch --prefix foo
    - etcdctl watch -i 支持交互模式，可watch多个key
    - etcdctl watch --rev=2 foo 查看从版本2之后的所有变化
    - etcdctl watch --prev-kv foo 观察的同时，显示最近一个版本的value
- key压缩，随着历史版本越来越多，可以将不会用到的历史进行压缩，压缩之后，版本之前的key和value都不可用
    - etcdctl compact 100   压缩版本100
    - etcdctl get foo -w=json获取当前etcd节点的版本号

## 租约 lease

这是etcd v3的特性，比超时的概念更丰富一些。
- lease 租约，有一个ttl，time to live，实际上的ttl会大于用于授予的ttl，这个由etcd决定
- key可以和lease绑定，lease的ttl超时了，表示租约过期了，与之绑定的所有key都会被自动删除。


    etcdctl lease grant 10 创建一个10s的租约
    etcdctl put --lease=02f06b452397aa0a b 1 将key和lease绑定
    etcdctl lease revoke 02f06b452397aa0e 撤销租约，与lease绑定的key全部删除
    etcdctl lease keep-alive 02f06b452397aa13 续租，将lease的ttl值重置为最初授予的值
    etcdctl lease timetolive 02f06b452397aa19 获取租约信息
    etcdctl lease timetolive 02f06b452397aa19 --keys 获取租约上绑定的key

## etcd 常用配置参数

etcd的启动参数通过命令行参数和环境变量来配置，
- ETCD_MY_FLAH
- --my-flag

参数：
- member
    - 节点名
    - 数据目录
    - wal目录
    - 触发快照的提交事务次数，默认10w次
    - 领袖心跳间隔，默认100ms
    - 选举超时事件，默认1s
    - 集群节点间通信地址
    - 客户端请求的地址
    - etcd保存最大快照文件数，默认5个，0表示无限
    - wal最大文件数，默认5个
    - 跨域资源共享白名单CORS，逗号分割
- cluster
    - 以 --initial为前缀 适用于member最初的启动过程和运行时，重启会忽略掉
    - 以 --discovery为前缀 适用于服务发现
    - 集群内部交互数据的地址
    - 初始启动的集群配置，也就是上面提到过的枚举 --initial-cluster
    - 集群初始化状态，默认是new
    - 集群token
    - 创建集群的服务发现URL，etcd自发现模式中的服务发现URL；dns自发现模式中的dns srv域名
    - 服务发现失败时的处理行为，proxy或exit，默认是proxy
    - 服务发现使用的http代理
    - 拒绝重新配置，前提是法定人数信息丢失
    - mvcc kv存储自动压缩间隔，单位小时，0表示不自动压缩
    - 是否接受etcd v2的请求
- proxy
    - 以 --proxy为前缀 适用于etcd运行在proxy模式下，反向代理模式只支持v2 api
    - 是否启用反向代理模式
    - 后端发生错误时，反向代理间隔多长时间再次使用该后端，默认5s
    - 后端刷新间隔，默认30s
    - 与后端连接超时时间，默认1s
    - 写后端超时，默认5s
    - 读后端超时，默认0，表示没有超时，一直等
- 安全
    - 各种证书、认证
- 日志
    - 将etcd所有子项目的日志等级调到debug等级，默认是info
    - 设置某个子项目的日志等级
- 不安全参数
    - 强制创建只有一个节点的集群，会移除集群内现存的节点，会破坏raft一致性
- 统计
    - 包括运行时性能分析和监控数据
    - 启用收集，并通过http服务对外暴露
    - 设置导出数据的详细程度
- 认证
