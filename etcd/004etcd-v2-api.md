# etcd v2 api

v2版本的api更多的是提供读写
- 读：范围查、watch
- 写：更新、删除

一个api完成的的标志是：
- 已通过raft一致性提交
- 并执行了(持久化到etcd存储引擎了)

客户端收到服务器的响应，意味着操作已经完成。

api的特点：
- 原子性，api要么执行完，要么不执行
- 一致性，所有的api都保证了顺序一致性。
- 隔离性，可串行化隔离，读操作不会读到任何中间数据
- 持久性，任何完成的操作都是持久的，读操作不会返回未持久化的数据

v2 api走的http+json格式，可以使用curl来测试

## 集群管理api

    查看etcd版本 curl http://localhost:2379/version | jq
    节点健康价差 curl http://localhost:2379/health | jq

## kv api

etcd使用类似文件系统的树型结构来表示kv，根节点用/表示，
里面可以存储两种内容：key和目录。
- key，存储字符串
- 目录，存储一些key和目录
看起来和windows的注册表很像

v2的key一般存储在/v2/keys/下

    增 curl http://localhost:2379/v2/keys/abc -XPUT -d value="123" | jq
    删 curl -v -XDELETE http://localhost:2379/v2/keys/abc | jq
    改 curl http://localhost:2379/v2/keys/abc -v -XPUT -d value="789" | jq
    查 curl -v http://localhost:2379/v2/keys/abc | jq

curl的-X表示使用http的其他动词，默认curl使用的GET，
PUT表示修改服务器数据，DELETE表示删除数据，后面都用得上，
-d 表示数据。  
jq是一个命令行json格式化工具，可用apt install jq -y 安装

## key的ttl

v2中，key的过期是用ttl来实现的，在v3中引入了lease租约的概念。

只有领袖才能主动让某个key过期，如果群众和领袖的网络断掉了，那么key就永远不会过期，知道重新连接上领袖

    设置ttl curl -v -XPUT http://localhost:2379/v2/keys/abc -d value=abc -d ttl=5 | jq
    刷新ttl curl -v -XPUT http://localhost:2379/v2/keys/abc -d value=456 -d ttl=5  -d prevExist=true | jq


ttl的应用场景：agent代理创建一个缓存key，定时刷新ttl，etcd如果发现key删除了，就判断agent掉线了。

