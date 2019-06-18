# etcd v3 api

etcd v3 版本的发布标志着数据模型和api的正式稳定。

- kv api支持了mini事务
- 提供了一套v3的api，v2和v3的api共享同一套raft协议代码，v2和v3可以看成raft协议代码行的两个应用，相互独立
- 效率、可靠性、性能上进行了优化

v3基于v2的改进和优化：
- grpc + pb 取代 http + json
- lease租约，更轻量级的自动过期机制，取代了key的ttl自动过期机制
- watcher 观察者机制重新设计，v2的机制基于http长连接事件驱动机制，v3机制基于http2的server push，对事件有多路复用优化
- v3数据模型改进，v2是一个kv内存数据库，v3支持事务和多版本并发控制的磁盘数据库。

## 重要特性

grpc
- google开源的高性能、跨语言的rpc框架
- 基于http2协议实现
- 使用bp进行序列化，数据模型和rpc接口都基于pb

pb
- 效率高于json，差不多两倍多

减少tcp连接
- 利用http2的多路复用，减少通信

租约机制 lease
- 多个过期时间相同的key，绑定一个租约，客户端维护这个租约即可，不用维护每一个key

观察者模式
- 服务端通知客户端，而不是客户端去轮询(zookeeper/consul就是轮询)
- 轮询：一次性watch，会导致用户无法感知两次watch之间发生的事件。etcd可以通过索引轮询和流式观察来解决这个问题
- v2的索引轮询，每次轮询都有一个http长连接，v3对一个客户端进行多路复用，减少服务端压力

数据存储模型
- v2只保存key最新的value，历史记录放在缓冲中(只有1000个，而且是全部key共享的)
- v3放弃了v2这种滑动窗口的设计，引入了mvcc(多版本并发控制)，采用了从历史记录为主索引的存储结构
- v3放弃了v2的目录式层级化设计，而使用了一个扁平化的设计
- v3采用了mvcc，保存了k-v的历史版本，数据大了很多，就从内存数据库转为磁盘数据库，底层存储引擎是 BoltDB,最后采用的是coreos维护的。

mini事务
- v3除了提供kv api，还提供了 事务api，v2提供了cas/cad来对单个key进行更新
- v2的cas/cad操作依赖提供具体版本号或当前值，一旦条件不满足，操作就会失败，而且多个key变更时，v2力有不逮
- v3引入mini事务来解决分布式锁和事务。v3 v2的差异是，v2的操作适用于单个key，v3适用于多个key

快照
- v2是内存数据库，最多支持数十万级级别的key，不是说不能支持更多，而是raft一致性是基于日志，日志不能无限增长，到一定程度就要存快照，持久化到磁盘
- v3对raft和存储系统进行了重构，支持增量快照和传输较大的快照，目前v3支持百万到千万级别的key

大规模watch
- v2中每个watch都占用一个tcp资源和一个go协程资源，差不多30k-40k
- v3使用http2的多路复用，同一个用户不同watch共用一个go协程，减少了服务器资源消耗

## grpc服务

现在api是通过grpc来提供的，所以在分类上有了细化：
- kv 键值
- cluster 集群
- maintenance 维护
- auth 认证
- watch 观察
- lease 租约

其中auth、cluster、maintenance也称为管理集群api，kv、watch、lease被称为键值空间api

## 请求和响应

v3中的rpc方法都是一个格式：rcp Range(RangeRequest) returns (RangeResponse) {},
一个入参一个返回值，请求参数中尽可能覆盖了更多的场景，响应有固定格式：有个响应头。
- 响应头有4个字段，分别是
- 集群id
- 成员id
- 版本号
- raft任期

v2 更多的是在http响应头中带有3个信息：etcd节点索引，raft索引和raft任期。v3更多是用版本来表示。


## kv api

kv键值对是kv api能处理的最小单位 mvccpb.KeyValue
```golang
type KeyValue struct {
    // key is the key in bytes. An empty key is not allowed.
    Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    // create_revision is the revision of last creation on this key.
    CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
    // mod_revision is the revision of last modification on this key.
    ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
    // version is the version of the key. A deletion resets
    // the version to zero and any modification of the key
    // increases its version.
    Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
    // value is the value held by the key, in bytes.
    Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
    // lease is the ID of the lease that attached to key.
    // When the attached lease expires, the key will be deleted.
    // If lease is 0, then no lease is attached to the key.
    Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}
```
除了kv值，还有创建版本、最后修改版本、当前版本、租约号

etcd版本的版本是集群内的一个64位计数器，也可以称为全局逻辑时钟，键空间的变化，会导致计数器的增加。单调递增。

键区间，[key1,keyn) 扁平key空间，每一个key都有一个索引。不像v2时代的目录式。
v3的键区间，也称为区间，支持单键查找、前缀查找。

range api,表示的一组kv
```golang
type RangeRequest struct {
    // key is the first key for the range. If range_end is not given, the request only looks up key.
    Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    // range_end is the upper bound on the requested range [key, range_end).
    // If range_end is '\0', the range is all keys >= key.
    // If range_end is key plus one (e.g., "aa"+1 == "ab", "a\xff"+1 == "b"),
    // then the range request gets all keys prefixed with key.
    // If both key and range_end are '\0', then the range request returns all keys.
    RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
    // limit is a limit on the number of keys returned for the request. When limit is set to 0,
    // it is treated as no limit.
    Limit int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
    // revision is the point-in-time of the key-value store to use for the range.
    // If revision is less or equal to zero, the range is over the newest key-value store.
    // If the revision has been compacted, ErrCompacted is returned as a response.
    Revision int64 `protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
    // sort_order is the order for returned sorted results.
    SortOrder RangeRequest_SortOrder `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=etcdserverpb.RangeRequest_SortOrder" json:"sort_order,omitempty"`
    // sort_target is the key-value field to use for sorting.
    SortTarget RangeRequest_SortTarget `protobuf:"varint,6,opt,name=sort_target,json=sortTarget,proto3,enum=etcdserverpb.RangeRequest_SortTarget" json:"sort_target,omitempty"`
    // serializable sets the range request to use serializable member-local reads.
    // Range requests are linearizable by default; linearizable requests have higher
    // latency and lower throughput than serializable requests but reflect the current
    // consensus of the cluster. For better performance, in exchange for possible stale reads,
    // a serializable range request is served locally without needing to reach consensus
    // with other nodes in the cluster.
    Serializable bool `protobuf:"varint,7,opt,name=serializable,proto3" json:"serializable,omitempty"`
    // keys_only when set returns only the keys and not the values.
    KeysOnly bool `protobuf:"varint,8,opt,name=keys_only,json=keysOnly,proto3" json:"keys_only,omitempty"`
    // count_only when set returns only the count of the keys in the range.
    CountOnly bool `protobuf:"varint,9,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
    // min_mod_revision is the lower bound for returned key mod revisions; all keys with
    // lesser mod revisions will be filtered away.
    MinModRevision int64 `protobuf:"varint,10,opt,name=min_mod_revision,json=minModRevision,proto3" json:"min_mod_revision,omitempty"`
    // max_mod_revision is the upper bound for returned key mod revisions; all keys with
    // greater mod revisions will be filtered away.
    MaxModRevision int64 `protobuf:"varint,11,opt,name=max_mod_revision,json=maxModRevision,proto3" json:"max_mod_revision,omitempty"`
    // min_create_revision is the lower bound for returned key create revisions; all keys with
    // lesser create revisions will be filtered away.
    MinCreateRevision int64 `protobuf:"varint,12,opt,name=min_create_revision,json=minCreateRevision,proto3" json:"min_create_revision,omitempty"`
    // max_create_revision is the upper bound for returned key create revisions; all keys with
    // greater create revisions will be filtered away.
    MaxCreateRevision int64 `protobuf:"varint,13,opt,name=max_create_revision,json=maxCreateRevision,proto3" json:"max_create_revision,omitempty"`
}
```
RangeRequest定义了排序因子、排序方式，还依次定义了以下元素：
- key的区间，左右两端
- 一次请求能返回的最大key数量
- 版本
- 表示是否通过可串行化读取数据
  - 可串行化读：直接读服务节点存储的数据(这个服务节点不一定是领袖节点，可能读之间有写操作，但领袖节点并未同步到当前群众节点)
  - 可线性化读：强调的是raft保证一致性，当前读的结果一定是之前写操作的结果。
  - 可串行化读的好处：更高的性能和可用性。损失的是 有一定概率读到旧数据
  - 可线性化读的好处：数据一致。 损失的是 对读也做了强一致性，所以读操作的性能降到和写操作差不多了。
- 是否只返回key，不要value
- 是否只返回key的数量
- 用修改版本的范围值做为一个参数，过滤掉一部分数据
- 用创建版本的范围值做为一个参数，过滤掉一部分数据

RangeRequest对应的响应：
```golang
type RangeResponse struct {
    Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
    // kvs is the list of key-value pairs matched by the range request.
    // kvs is empty when count is requested.
    Kvs []*mvccpb.KeyValue `protobuf:"bytes,2,rep,name=kvs" json:"kvs,omitempty"`
    // more indicates if there are more keys to return in the requested range.
    More bool `protobuf:"varint,3,opt,name=more,proto3" json:"more,omitempty"`
    // count is set to the number of keys within the range when requested.
    Count int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}
```
其中count表示满足请求的key数量，more表示是否还有数据没写到响应结果中。

DeleteRangeRequest是范围删除
```golang
type DeleteRangeRequest struct {
    // key is the first key to delete in the range.
    Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    // range_end is the key following the last key to delete for the range [key, range_end).
    // If range_end is not given, the range is defined to contain only the key argument.
    // If range_end is one bit larger than the given key, then the range is all the keys
    // with the prefix (the given key).
    // If range_end is '\0', the range is all keys greater than or equal to the key argument.
    RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
    // If prev_kv is set, etcd gets the previous key-value pairs before deleting it.
    // The previous key-value pairs will be returned in the delete response.
    PrevKv bool `protobuf:"varint,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
}
```
PrevKv，表示是否要返回删除的键值对，对应的，响应如下
```golang
type DeleteRangeResponse struct {
    Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
    // deleted is the number of keys deleted by the delete range request.
    Deleted int64 `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
    // if prev_kv is set in the request, the previous key-value pairs will be returned.
    PrevKvs []*mvccpb.KeyValue `protobuf:"bytes,3,rep,name=prev_kvs,json=prevKvs" json:"prev_kvs,omitempty"`
}
```
其中包含了删除key的数量

上面提到了范围查，范围删，下面看下单个key的新增或修改 PutRequest
```golang
type PutRequest struct {
    // key is the key, in bytes, to put into the key-value store.
    Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    // value is the value, in bytes, to associate with the key in the key-value store.
    Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
    // lease is the lease ID to associate with the key in the key-value store. A lease
    // value of 0 indicates no lease.
    Lease int64 `protobuf:"varint,3,opt,name=lease,proto3" json:"lease,omitempty"`
    // If prev_kv is set, etcd gets the previous key-value pair before changing it.
    // The previous key-value pair will be returned in the put response.
    PrevKv bool `protobuf:"varint,4,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
    // If ignore_value is set, etcd updates the key using its current value.
    // Returns an error if the key does not exist.
    IgnoreValue bool `protobuf:"varint,5,opt,name=ignore_value,json=ignoreValue,proto3" json:"ignore_value,omitempty"`
    // If ignore_lease is set, etcd updates the key using its current lease.
    // Returns an error if the key does not exist.
    IgnoreLease bool `protobuf:"varint,6,opt,name=ignore_lease,json=ignoreLease,proto3" json:"ignore_lease,omitempty"`
}
```
里面包含了key-value、租约、是否返回修改前的key-value值、是否只修改key、是否不修改租约。  
响应PutResponse：
```golang
type PutResponse struct {
    Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
    // if prev_kv is set in the request, the previous key-value pair will be returned.
    PrevKv *mvccpb.KeyValue `protobuf:"bytes,2,opt,name=prev_kv,json=prevKv" json:"prev_kv,omitempty"`
}
```

事务，一个事务可以有多个请求，同一事务里产生的事件都有同一个版本号，一个事务里禁止多次修改同一个key。
    事务类似于：
    if con1 {
        事务1
    }
    if con2 {
        事务2
    }
    if con3 {
        事务3
    }
    利用比较条件来保证集群中事务的不受其他的干扰。
    一般比较条件可以选择是否等于某个值，或版本号
```golang
type Compare struct {
    // result is logical comparison operation for this comparison.
    Result Compare_CompareResult `protobuf:"varint,1,opt,name=result,proto3,enum=etcdserverpb.Compare_CompareResult" json:"result,omitempty"`
    // target is the key-value field to inspect for the comparison.
    Target Compare_CompareTarget `protobuf:"varint,2,opt,name=target,proto3,enum=etcdserverpb.Compare_CompareTarget" json:"target,omitempty"`
    // key is the subject key for the comparison operation.
    Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
    // Types that are valid to be assigned to TargetUnion:
    //	*Compare_Version
    //	*Compare_CreateRevision
    //	*Compare_ModRevision
    //	*Compare_Value
    //	*Compare_Lease
    TargetUnion isCompare_TargetUnion `protobuf_oneof:"target_union"`
    // range_end compares the given target to all keys in the range [key, range_end).
    // See RangeRequest for more details on key ranges.
    RangeEnd []byte `protobuf:"bytes,64,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
}
```
里面包含了比较因子、比较方式、待比较的key、以及用于比较的数据

事务里的比较条件成功之后，会执行请求块里的操作
```golang
type RequestOp struct {
    // request is a union of request types accepted by a transaction.
    //
    // Types that are valid to be assigned to Request:
    //	*RequestOp_RequestRange
    //	*RequestOp_RequestPut
    //	*RequestOp_RequestDeleteRange
    //	*RequestOp_RequestTxn
    Request isRequestOp_Request `protobuf_oneof:"request"`
}
```
操作里面包含了范围查、范围删、key更新。删和更新的key在一个事务里是不重复的。

etcd中，一个事务就是一个txn api调用。
```golang
type TxnRequest struct {
    // compare is a list of predicates representing a conjunction of terms.
    // If the comparisons succeed, then the success requests will be processed in order,
    // and the response will contain their respective responses in order.
    // If the comparisons fail, then the failure requests will be processed in order,
    // and the response will contain their respective responses in order.
    Compare []*Compare `protobuf:"bytes,1,rep,name=compare" json:"compare,omitempty"`
    // success is a list of requests which will be applied when compare evaluates to true.
    Success []*RequestOp `protobuf:"bytes,2,rep,name=success" json:"success,omitempty"`
    // failure is a list of requests which will be applied when compare evaluates to false.
    Failure []*RequestOp `protobuf:"bytes,3,rep,name=failure" json:"failure,omitempty"`
}
```
先比较，成功就执行成功操作，失败就执行失败操作。响应如下：

```golang
type TxnResponse struct {
    Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
    // succeeded is set to true if the compare evaluated to true or false otherwise.
    Succeeded bool `protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
    // responses is a list of responses corresponding to the results from applying
    // success if succeeded is true or failure if succeeded is false.
    Responses []*ResponseOp `protobuf:"bytes,3,rep,name=responses" json:"responses,omitempty"`
}
```
里面包含了'成功操作'的结果，下面看下结果对应的ResponseOp
```golang
type ResponseOp struct {
    // response is a union of response types returned by a transaction.
    //
    // Types that are valid to be assigned to Response:
    //	*ResponseOp_ResponseRange
    //	*ResponseOp_ResponsePut
    //	*ResponseOp_ResponseDeleteRange
    //	*ResponseOp_ResponseTxn
    Response isResponseOp_Response `protobuf_oneof:"response"`
}
```
里面包含了范围查、范围删、单个key的更新，以及事务操作。事务里也能包含事务

还有一个压缩 Compact
```golang
  Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error)
```
这是clientv3中kv服务的一个方法

## watch api

v3基于事件来监测key的变化

```golang
type Event struct {
    // type is the kind of event. If type is a PUT, it indicates
    // new data has been stored to the key. If type is a DELETE,
    // it indicates the key was deleted.
    Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=mvccpb.Event_EventType" json:"type,omitempty"`
    // kv holds the KeyValue for the event.
    // A PUT event contains current kv pair.
    // A PUT event with kv.Version=1 indicates the creation of a key.
    // A DELETE/EXPIRE event contains the deleted key with
    // its modification revision set to the revision of deletion.
    Kv  *KeyValue `protobuf:"bytes,2,opt,name=kv" json:"kv,omitempty"`
    // prev_kv holds the key-value pair before the event happens.
    PrevKv *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv" json:"prev_kv,omitempty"`
}
```
里面包含了事件类型(更新还是删除)、当前值和之前的值。

v3的watch机制对事件有以下保证：
- 有序，低版本事件先发生，高版本后发生
- 可靠，不会漏事件。v2就有可能漏掉
- 原子性，一个事件的多个key，不会通过多个事件发送，都在一个事件里

流式watch的发送和接收：
```golang
type WatchCreateRequest struct {
    // key is the key to register for watching.
    Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    // range_end is the end of the range [key, range_end) to watch. If range_end is not given,
    // only the key argument is watched. If range_end is equal to '\0', all keys greater than
    // or equal to the key argument are watched.
    // If the range_end is one bit larger than the given key,
    // then all keys with the prefix (the given key) will be watched.
    RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
    // start_revision is an optional revision to watch from (inclusive). No start_revision is "now".
    StartRevision int64 `protobuf:"varint,3,opt,name=start_revision,json=startRevision,proto3" json:"start_revision,omitempty"`
    // progress_notify is set so that the etcd server will periodically send a WatchResponse with
    // no events to the new watcher if there are no recent events. It is useful when clients
    // wish to recover a disconnected watcher starting from a recent known revision.
    // The etcd server may decide how often it will send notifications based on current load.
    ProgressNotify bool `protobuf:"varint,4,opt,name=progress_notify,json=progressNotify,proto3" json:"progress_notify,omitempty"`
    // filters filter the events at server side before it sends back to the watcher.
    Filters []WatchCreateRequest_FilterType `protobuf:"varint,5,rep,packed,name=filters,enum=etcdserverpb.WatchCreateRequest_FilterType" json:"filters,omitempty"`
    // If prev_kv is set, created watcher gets the previous KV before the event happens.
    // If the previous KV is already compacted, nothing will be returned.
    PrevKv bool `protobuf:"varint,6,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
    // If watch_id is provided and non-zero, it will be assigned to this watcher.
    // Since creating a watcher in etcd is not a synchronous operation,
    // this can be used ensure that ordering is correct when creating multiple
    // watchers on the same stream. Creating a watcher with an ID already in
    // use on the stream will cause an error to be returned.
    WatchId int64 `protobuf:"varint,7,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
    // fragment enables splitting large revisions into multiple watch responses.
    Fragment bool `protobuf:"varint,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
}


type WatchResponse struct {
    Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
    // watch_id is the ID of the watcher that corresponds to the response.
    WatchId int64 `protobuf:"varint,2,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
    // created is set to true if the response is for a create watch request.
    // The client should record the watch_id and expect to receive events for
    // the created watcher from the same stream.
    // All events sent to the created watcher will attach with the same watch_id.
    Created bool `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
    // canceled is set to true if the response is for a cancel watch request.
    // No further events will be sent to the canceled watcher.
    Canceled bool `protobuf:"varint,4,opt,name=canceled,proto3" json:"canceled,omitempty"`
    // compact_revision is set to the minimum index if a watcher tries to watch
    // at a compacted index.
    //
    // This happens when creating a watcher at a compacted revision or the watcher cannot
    // catch up with the progress of the key-value store.
    //
    // The client should treat the watcher as canceled and should not try to create any
    // watcher with the same start_revision again.
    CompactRevision int64 `protobuf:"varint,5,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
    // cancel_reason indicates the reason for canceling the watcher.
    CancelReason string `protobuf:"bytes,6,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
    // framgment is true if large watch response was split over multiple responses.
    Fragment bool            `protobuf:"varint,7,opt,name=fragment,proto3" json:"fragment,omitempty"`
    Events   []*mvccpb.Event `protobuf:"bytes,11,rep,name=events" json:"events,omitempty"`
}
```
请求中包含了：
- key的范围
- 起始版本号 可选
- 是否接收 无事件消息
- 事件类型过滤
- 是否显示事件之前的key-value数据

响应中包含了：
- watch响应id
- 是否是watch 创建
- 是否是watch 删除
- watch版本号已经被压缩了，就返回最小的版本号
- 对应watch id的event有序列表

## 租约 lease api

一个key最多关联一个租约。
```golang
// 创建一个租约
// ttl 单位秒
// id 是服务端生成
type LeaseGrantRequest struct {
    // TTL is the advisory time-to-live in seconds. Expired lease will return -1.
    TTL int64 `protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
    // ID is the requested ID for the lease. If ID is set to 0, the lessor chooses an ID.
    ID  int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
}
```
删除只需要传入租约号即可。LeaseRevokeRequest  
续租也只需要传入租约号，LeaseKeepAliveRequest  

v3 的api 走的是grpc，go测试可用官方客户端包clientv3来测试。




