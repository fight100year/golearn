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
    节点健康检查 curl http://localhost:2379/health | jq

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
接收请求的可能是群众，也可能是领袖，如果群众和领袖的时钟不一样，会影响key的ttl，所以群众和领袖的时钟差异超过1s，etcd会发出提示。

    可以只刷新ttl，不更新值 curl -v -XPUT http://localhost:2379/v2/keys/abc -d value=abc -d ttl=5 | jq

## watch

等待变化通知。etcd的watch是通过轮询机制实现。

    观察一个key curl 'http://localhost:2379/v2/keys/abc?wait=true' 这个是一次性watch
    带索引的watch  curl 'http://localhost:2379/v2/keys/abc?wait=true&waitIndex=20' | jq 有就直接返回，没有就挂起
    流式watch curl 'http://localhost:2379/v2/keys/abc?wait=true&recursive=true&stream=true'  非一次性的


etcd v2版本的 历史事件的缓冲只有1000个事件，超过就会覆盖，如果watch收到响应，最好立马返回。
如果watch index时，缓冲区满了导致想等的事件被etcd丢弃了，那么watch会返回一个400的http状态码。

流式watch相比一次性watch，好处是可靠，不会丢事件；坏处是历史事件监听不到。  
流式带索引的watch：
- waitIndex参数大于etcd当前索引，参数无效，行为等同于流式watch
- waitIndex参数小于等于etcd当前索引，只会获取waitIndex到当前索引之间的事件，后续不再接收事件，也不退出

v2 版本最佳watch套路：
- 流式watch
- 带索引的一次性watch

这个套路适用于获取当前开始后面所有的变动，要找历史只能用索引，考虑到缓冲时刻都在被覆盖，尽量不要把业务基于历史事件实现。

watch 可以观察一个目录，目录下所有的变化事件都能被观察到，recursive=true 只适用于目录，对key无效。

etcd还支持exec-watch功能，为事件挂载一个处理函数。

## 自动创建有序key

http的put和post请求都是更改资源，不过put是幂等性，post不是。所以put多用于修改，post多用于新增。
eg：新增一张信用卡，适合用post，post 5次表示加5张信用卡；修改余额为100，put 1000次，余额还是100.

创建有序key就是一个post请求，(创建具体某个key，使用put即可)

put创建的key，post创建的有序key

自动创建有序key，就是说创建key时，不指明具体的key，只指定目录和value，etcd会自动创建key，这个key和etcd的修改索引有关，
能保证单调自增，但不保证连续，查询时，可以排序，这样会得到一个有序的key。

    增加4个key 
    curl http://localhost:2379/v2/keys/d1 -XPOST -d value="job1" | jq
    curl http://localhost:2379/v2/keys/d1 -XPOST -d value="job2" | jq
    curl http://localhost:2379/v2/keys/d1 -XPOST -d value="job3" | jq
    curl http://localhost:2379/v2/keys/d1 -XPOST -d value="job4" | jq
    查询
    curl 'http://localhost:2379/v2/keys/d1?recursive=true&sorted=true'  | jq

## 目录的ttl

目录的ttl过期后，下面所有的key都会被删除

    创建目录时指定ttl curl http://localhost:2379/v2/keys/d2  -XPUT -d ttl=30 -d dir=true | jq
    刷新已存在目录的ttl curl http://localhost:2379/v2/keys/d1  -XPUT -d ttl=30 -d dir=true -d prevExist=true | jq


## 原子的cas

cas： compare and swap 先比较，如果不一样，就交换。  
- cas是分布式锁服务的一个基本操作。
- cas不适用于目录

cas支持的比较条件：
- prevValue: 检查key之前的value
- prevIndex: 检查key之前的modifiedIndex
- prevExist: 检查key是否存在，存在就是更新，不存在就是新建

    设置一个值
    curl http://localhost:2379/v2/keys/f2 -XPUT -d value=123 
    值不存在就新建
    curl http://localhost:2379/v2/keys/f2?prevExist=false -XPUT -d value=456 | jq
    值为123就更新为456
    curl http://localhost:2379/v2/keys/f2?prevValue=123 -XPUT -d value=456 | jq

## 原子的cad

cad: compare and delete 先比较，如果一样，就删除
- cas不适用于目录

cas支持的比较条件：
- prevValue: 检查key之前的value
- prevIndex: 检查key之前的modifiedIndex

    curl 'http://localhost:2379/v2/keys/f2?prevValue=123' -XDELETE

## 目录操作

大多数下是自动创建

    创建目录
    curl http://localhost:2379/v2/keys/d1/d2/d3 -XPUT -d dir=true
    创建d3下的key
    curl http://localhost:2379/v2/keys/d1/d2/d3/k1 -XPUT -d value=123 | jq
    获取目录d1
    curl http://localhost:2379/v2/keys/d1 | jq
    递归展开d1
    curl 'http://localhost:2379/v2/keys/d1?recursive=true' | jq
    删除一个空目录
    curl 'http://localhost:2379/v2/keys/d1/d2/d3/d4?dir=true' -XDELETE
    删除一个目录及目录下的所有内容
    curl 'http://localhost:2379/v2/keys/d1/d2/d3/d4?recursive=true' -XDELETE | jq

## 隐藏节点

以_为前缀的key或目录称为隐藏key、隐藏目录。
默认情况下，http get请求不会返回隐藏节点。

## 存储小配置文件

    echo "Hello\nWorld” > afile.txt
    curl http://127.0.0.1:2379/v2/keys/afile -XPUT --data-urlencode value@afile.txt

## 线性读

发生在写之后的读一定能读到数据  
quorum=true,读性能会下降到和写性能差不多。

    curl 'http://localhost:2379/v2/keys/k1?quorum=true' | jq

## 统计

集群运行时，etcd会收集一些数据，客户端可通过api来获取这些数据。eg：请求延时、数据带宽、运行时长等

领袖的数据：
- 集群中每个节点的延时
- 失败/成功的raft rpc请求次数

领袖的数据保存在 /v2/stats/leader中

    获取领袖数据，只能在领袖节点运行
    curl http://localhost:2379/v2/stats/leader | jq

节点自身的数据：
```json
{
  "name": "infra3", // 节点名
  "id": "12d46ddcd3b78dca", // 节点唯一标识符
  "state": "StateLeader", // raft协议里的角色
  "startTime": "2019-06-12T09:58:32.86635184+08:00", // etcd节点启动时间
  "leaderInfo": { // 当前领袖信息
    "leader": "12d46ddcd3b78dca", // 领袖节点id
    "uptime": "4h34m3.524383174s", // 领袖节点启动事件
    "startTime": "2019-06-12T09:58:33.967011264+08:00"
  },
  "recvAppendRequestCnt": 0, // 已处理的append请求数
  "sendAppendRequestCnt": 4631 // 已发送的append请求数
}
```
节点自身数据保存在 /v2/stats/self, 领袖和群众的数据会各不相同

    curl http://localhost:2379/v2/stats/self | jq

还有一些数据放在 /v2/stats/store 

## member api

通过成员api可管理节点，包括增删改查。
- list member: curl http://localhost:2379/v2/members | jq
- add member
- curl http://localhost:2379/v2/members/12d46ddcd3b78dca -XDELETE | jq
- 修改member的peer url
