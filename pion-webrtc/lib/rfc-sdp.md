# SDP 

- session description protocol
- 会话描述协议
- 常与sip(session initiation protocol 会话初始协议)一起使用
- 2006年进行了标准化，rfc号是4566

## 介绍

需求：
- 多媒体电话会议/ip语音呼叫/视频流/其他会话
- 发起以上活动时，需给其他参与者发送自己的信息
- 信息包括：媒体细节/传输地址/其他描述性的元信息

实现：
- sdp正好用来解决上面的问题，另外sdp只管这些信息的表达，而不管信息的传输
- sdp是一个纯框架，专门定义了如何描述会话，这其中不包括sdp信息本身的传输
- sdp的传输可以按需指定：sap/sip/rtsp/带mime扩展的电子邮件/http协议

定位：
- sdp只描述会话，简单专一带来的是网络上的广泛使用
- sdp不包括会话内容的协商/不包括媒体格式的协商

## 术语

- 会议： 多媒体会议
    - 有2个或多个参与者,利用软件进行交流
- 会话： 一个多媒体会话
    - 有一系列多媒体发送者和接收者，数据流从发送者到接收者
    - 一个会议就是一个典型的多媒体会话
- 会话描述： 一种定义明确的格式
    - 用于向(发现者/参与者)传达足够的信息

## sdp用法例子

- 会话初始化
    - sip协议是应用层的控制协议，用于创建/修改/结束会话
    - sip可用于网络多媒体会话/网络电话呼叫/多媒体分发
    - 基于sip可以进行一个媒体协商，而sip创建的会话，她的描述格式就是sdp
    - sip+sdp 可用于offer/answer模式协商，这种模式限制性框架
- 流媒体
    - rtsp是一个应用层协议，传输数据时，带有一个实时属性
    - rtsp提供了一个可扩展的框架，易于控制，按需提供实时数据(音频/视频)
    - rtsp的服务端和客户端就需要协商很多信息(媒体传输信息)
    - 这些传输信息，部分是用sdp实现的
- 电子邮件和万维网
    - 通过电子邮件和www也可以进行sdp的传输
    - 她们可以通过application/sdp来支持sdp的传输
- 多播会话通知
    - 多播多媒体会议，或是多播会话，为了和未来参与者进行交流，可能需要一个分布式会话目录
    - 这样分布式会话目录需要定期给已知的多播组发送包含会话描述的包
    - 这样未来参与者就可以利用会话描述进行下一步工作
    - sap(session announcement protocol 会话通知协议)，就服务于这个分布式会话目录
    - 而里面描述会话的，就是sdp协议

## 要求和建议

- sdp的目的是在多媒体会话中，传达媒体流的信息，以便让会话描述中的接受者参与到会话中 
- 主要用在互联网，因为特别适合描述不同网络情况的会议
- 媒体流可以是多对多
- 会话不需要保活

目前互联网会议有两种：
- 基于多播会话
- 会议中的接收方可以加入会话(除非会话加密了)

相应的sdp有两个目的：
- 和已存在的会话交互
- 传达足够的信息，以便可以加入到会话

如果是单播环境，sdp只有第二个目的

sdp应该包含以下信息：
- 会话名 会话目的
- 会话激活时间
- 构成会话的媒体
- 接收这些媒体的相关信息(地址 端口 格式等)

考虑到资源不是无限的，下列信息也需要指定：
- 会话的带宽
- 会话负责人的联系信息

总结：sdp一般用于"传达足够信息以加入会话(非加密)";给待参与者通知已使用资源(eg：分布式会话目录)，这个是基于多播的

### 媒体和传输信息

媒体信息包括：
- 媒体类型，eg：音频 视频 其他
- 传输协议，eg：rtp/udp/ip，h.320等
- 媒体格式，eg：h.261 video, mpeg, 等

除了上面的，还需要地址和端口，对于ip多播会话，还包括：
- 多播组地址
- 传输端口 (这个多播的ip和端口，是目标地址和目标端口，无论是发送还是接收) 

对于非多播会话：
- 是远端地址
- 是远端端口 (就是数据发送者的ip和端口)

此处的媒体类型尽量不要重复定义，会增加实现的复杂度

### 时间信息

- 会话可能有时间限制，也可能没有。
- 会话可以指定何时激活
    - 可以指定任意开始/结束的时间，指定了，就是表示会话有时间限制
    - 可以重复指定 eg：每周一下午一点开始，为期1小时
- 这个时间是一致的，无论是哪个时区

## 私有会话

- sdp本身是不区分公共会话和私有会话
- 这里的私有会话是指在分发中传达的信息是加密的
- 一般是利用sip/sap协议进行加密


### 获取会话的更多信息

- sdp里应包含足够的信息，以便决定是否参加这个会话
- 所以sdp里可能会包含一些uris，这样就可以表示更多的信息

### 分类

- sdp可以由sap进行分发，也可以由其他机制分发，sdp里可以进行过滤
- 这样可以控制sdp支持哪些机制的分发
- a=cat:

### 国际化

- sdp 推荐使用iso 10646字符集，编码使用utf-8


## sdp specification 规格

- sdp是纯文本，iso10646字符集，utf-8编码
- sdp不含传输部分，所以不像其他协议有包头什么的
- sdp的字段和属性值不能直接和utf-8的比较，因为标准并未强制要求utf-8,只是建议
- 文本格式，而不是二进制等其他格式，好处是便于携带，可以兼容更多传输协议
- 文本格式，生成sdp和处理sdp也更加方便
- 考虑到有些环境对sdp的大小有限制，所以sdp采用紧凑性写法
- 考虑到底层传输协议和中转服务缓存的丢失，所以sdp的编码使用严格顺序和格式规则的设计
- 好处是更容易发现sdp是否出现异常
- 还允许接收方在没有相应解码key的时候，快速丢弃加密的会话通知
- sdp每行的格式都是<type>=<value>
    - type 要是单个大小写敏感的字符
    - value就是一个结构化的文本，一般用空格隔开，或者是字符串格式，=前后没有空格
- sdp分两部分：会话级别的部分和媒体级别的部分
    - 会话级别从"v="开始，直到遇到第一个媒体级别的行
    - 每一个媒体级别部分从"m="开始
    - 一般会话级别对所有媒体都起作用，除非是媒体级别进行了覆盖
- sdp的信息，有部分是必选，有部分是可选

```
sdp 格式，包含了顺序和每行表达的意思
Session description
   v=  (protocol version) 协议版本
   o=  (originator and session identifier) 发起者和会话标识
   s=  (session name) 会话名
   i=* (session information) 会话信息
   u=* (URI of description) 会话描述的uri,也就是前面提到过的，为了携带更多信息
   e=* (email address) 发起者的联系方式
   p=* (phone number) 发起者的联系方式
   c=* (connection information -- not required if included in all media) 连接信息，如果媒体级会话有指定，这条就可省略
   b=* (zero or more bandwidth information lines) 带宽信息
   One or more time descriptions ("t=" and "r=" lines; see below)
   z=* (time zone adjustments) 时区信息
   k=* (encryption key) 加密信息
   a=* (zero or more session attribute lines) 会话的属性信息
   Zero or more media descriptions

Time description
   t=  (time the session is active) 会话激活时间
   r=* (zero or more repeat times) 重复时间，就像先前的：每周一10点开始，为期1小时的会话

Media description, if present
   m=  (media name and transport address) 媒体名和传输地址
   i=* (media title) 媒体标题
   c=* (connection information -- optional if included at session level) 连接信息，这个是媒体级别的
   b=* (zero or more bandwidth information lines) 带宽信息
   k=* (encryption key) 加密信息
   a=* (zero or more media attribute lines) 媒体属性
```

下面是一个firefox抓的一个webrtc sdp
```
会话级信息开始

版本
v=0

发起者和会话标识
o=mozilla...THIS_IS_SDPARTA-68.0.2 2913518868991814669 0 IN IP4 0.0.0.0

会话名
s=-

会话开始时间
t=0 0

会话附加属性
a=sendrecv
a=fingerprint:sha-256 EE:B0:51:71:0B:0F:8F:84:0D:2E:04:23:F4:34:C2:F9:24:66:AF:B9:55:56:01:D8:A0:33:1A:C9:66:B0:EA:4E
a=group:BUNDLE 0 1
a=ice-options:trickle
a=msid-semantic:WMS *

媒体级信息标识
m=audio 43280 UDP/TLS/RTP/SAVPF 109 9 0 8 101

连接信息
c=IN IP4 172.168.10.200

媒体属性
a=candidate:0 1 UDP 2122252543 172.168.10.200 43280 typ host
a=candidate:1 1 UDP 2122187007 172.17.0.1 42073 typ host
a=candidate:2 1 TCP 2105524479 172.168.10.200 9 typ host tcptype active
a=candidate:3 1 TCP 2105458943 172.17.0.1 9 typ host tcptype active
a=candidate:0 2 UDP 2122252542 172.168.10.200 60735 typ host
a=candidate:1 2 UDP 2122187006 172.17.0.1 42024 typ host
a=candidate:2 2 TCP 2105524478 172.168.10.200 9 typ host tcptype active
a=candidate:3 2 TCP 2105458942 172.17.0.1 9 typ host tcptype active
a=sendonly
a=end-of-candidates
a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level
a=extmap:2/recvonly urn:ietf:params:rtp-hdrext:csrc-audio-level
a=extmap:3 urn:ietf:params:rtp-hdrext:sdes:mid
a=fmtp:109 maxplaybackrate=48000;stereo=1;useinbandfec=1
a=fmtp:101 0-15
a=ice-pwd:a907fd71fa56fc0be5586a3e59657662
a=ice-ufrag:efe37ebd
a=mid:0
a=msid:{c5e88a33-ebf0-4da6-91fc-c6cf47f692d6} {b059efa8-381c-411f-b74e-edabace30766}
a=rtcp:60735 IN IP4 172.168.10.200
a=rtcp-mux
a=rtpmap:109 opus/48000/2
a=rtpmap:9 G722/8000/1
a=rtpmap:0 PCMU/8000
a=rtpmap:8 PCMA/8000
a=rtpmap:101 telephone-event/8000
a=setup:actpass
a=ssrc:3225358466 cname:{b4be1496-139c-4030-aed0-cbf1d7704f71}
m=video 33040 UDP/TLS/RTP/SAVPF 120 121 126 97
c=IN IP4 172.168.10.200
a=candidate:0 1 UDP 2122252543 172.168.10.200 33040 typ host
a=candidate:1 1 UDP 2122187007 172.17.0.1 55259 typ host
a=candidate:2 1 TCP 2105524479 172.168.10.200 9 typ host tcptype active
a=candidate:3 1 TCP 2105458943 172.17.0.1 9 typ host tcptype active
a=candidate:0 2 UDP 2122252542 172.168.10.200 56040 typ host
a=candidate:1 2 UDP 2122187006 172.17.0.1 49538 typ host
a=candidate:2 2 TCP 2105524478 172.168.10.200 9 typ host tcptype active
a=candidate:3 2 TCP 2105458942 172.17.0.1 9 typ host tcptype active
a=sendonly
a=end-of-candidates
a=extmap:3 urn:ietf:params:rtp-hdrext:sdes:mid
a=extmap:4 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time
a=extmap:5 urn:ietf:params:rtp-hdrext:toffset
a=fmtp:126 profile-level-id=42e01f;level-asymmetry-allowed=1;packetization-mode=1
a=fmtp:97 profile-level-id=42e01f;level-asymmetry-allowed=1
a=fmtp:120 max-fs=12288;max-fr=60
a=fmtp:121 max-fs=12288;max-fr=60
a=ice-pwd:a907fd71fa56fc0be5586a3e59657662
a=ice-ufrag:efe37ebd
a=mid:1
a=msid:{c5e88a33-ebf0-4da6-91fc-c6cf47f692d6} {6abe6872-778f-4beb-9b7a-2617d6c97b92}
a=rtcp:56040 IN IP4 172.168.10.200
a=rtcp-fb:120 nack
a=rtcp-fb:120 nack pli
a=rtcp-fb:120 ccm fir
a=rtcp-fb:120 goog-remb
a=rtcp-fb:121 nack
a=rtcp-fb:121 nack pli
a=rtcp-fb:121 ccm fir
a=rtcp-fb:121 goog-remb
a=rtcp-fb:126 nack
a=rtcp-fb:126 nack pli
a=rtcp-fb:126 ccm fir
a=rtcp-fb:126 goog-remb
a=rtcp-fb:97 nack
a=rtcp-fb:97 nack pli
a=rtcp-fb:97 ccm fir
a=rtcp-fb:97 goog-remb
a=rtcp-mux
a=rtpmap:120 VP8/90000
a=rtpmap:121 VP9/90000
a=rtpmap:126 H264/90000
a=rtpmap:97 H264/90000
a=setup:actpass
a=ssrc:3668201432 cname:{b4be1496-139c-4030-aed0-cbf1d7704f71}
```

- sdp每行都是小写字母开头，而且不准备扩展，识别不了的都应该忽略
- a= 这种附加属性，是主要扩展sdp的方式，她和具体应用或媒体有强依赖关系
- a= 这种，有小部分是公共约定，大部分还是跟具体环境走，eg：h264 rfc也会定义使用sdp时的字段
- 额外信息：
    - 前面提到过 u=外部uris信息，k=加密信息，a=属性都可能包含一个uri
- 信息的覆盖：
    - 会话级别和媒体级别都有c=连接信息，a=属性，媒体级别信息会覆盖会话级别信息
    - 如果不通过a=指定字符集和编码格式，就默认使用标准的
- 域名，出现的域名也需要遵循一定的规则

### 协议版本 v=0 必选

- 指sdp协议版本，这份rfc定义的是0,表示没有小版本号

### 发起者和会话标识 o= 必选

    o=<username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
    
    o=mozilla...THIS_IS_SDPARTA-68.0.2 8450451625470945144 0 IN IP4 0.0.0.0

字段说明:
- username: 用户在主机上的登录名
    - 如果是-表示用户主机不支持用户id的概念
    - 用户名不能带空格
- sess-id: 会话的唯一标识
    - 是一个数字字符串
    - 是由工具生成，不过一般建议由ntp格式的时间戳来确保唯一
- sess-version: 会话描述的版本，这个和sdp协议版本不是一个意思
    - 每次修改一次会话描述，这个版本就变一次
    - 推荐使用ntp格式时间戳来表示
- nettype: 网络类型
    - 目前只有一种，用IN表示互联网
    - 未来可能做扩展
- addrtype: 地址类型
    - IP4 表示ipv4， IP6表示ipv6
- unicast-address: 单播地址
    - 创建会话的机器地址，也可以理解为单播发送者的地址

约束说明：
- o=一般用于标识唯一的会话，记录会话描述的版本
- 有时为了隐私安全，用户名和ip会被做一定处理，就像上面的例子，可以填入任意值

### 会话名 s= 必选

    s=<session name>

    s=

- 会话名，只能有一个
- 非空，且满足字符集和编码格式
- 如果没有一个有意义的值，可直接使用s=就行，就像上面的例子

### 会话信息 i= 可选

    i=<session description>

- 一个会话最多一个会话信息
- 要符合指定的字符集和编码格式
- i=会话信息 可以适用于所有媒体
- 在媒体级别的信息中也有一个 i=媒体标题，这个i就媒体流的一个标签，一个会话有多个媒体流就有i=的用武之地
- i= 可提供格式自由，易于人读的描述
- 一般这行不做自动解析

### URI资源 u= 可选

    u=<uri>

- 在万维网客户端使用，用于确定一个资源标识
- 要在媒体级信息之前出现，顺序不能错
- 每个sdp最多只能有一个

### 联系方式 e= p= 可选

    e=<email-address>
    p=<phone-number>

    p=+1 617 555-6011
    e=j.doe@example.com (Jane Doe)
    e=Jane Doe <j.doe@example.com>

- 一个是邮箱地址，一个是电话号码
- 这个都是会话负责人的信息，并不一定是会话发起者
- sdp协议，前一个版本这两个值是必选的，这个版本开始是可选的
- 必须在媒体级信息前面，顺序不能错
- 可添加多个邮件地址和电话
- 电话号码要符合itu-t标准，要以+开头，中间可添加空格增加可读性
- 邮件后面可接一个人名，人名要放在括号中，就如上面的例子那样,或者先写名字后跟邮件地址
- 要符合指定字符集和编码格式

### 连接数据 c= 可选

    c=<nettype> <addrtype> <connection-address>

    单播
    c=IN IP4 172.168.10.200

    多播，带ttl的
    c=IN IP4 224.2.36.42/127

    多播svc，接收者指定接收具体的层
    c=IN IP4 224.2.1.1/127/3

- 连接信息，发起者的地址
- 要么在会话级别指定一个，要么在每个媒体级都指定一个
- 也可以出现sdp有一个会话级和多个媒体级的连接信息，当然媒体级会覆盖会话级的信息

字段分析：
- nettype: 网络类型
    - 目前只有一种，用IN表示互联网
    - 未来可能做扩展
- addrtype: 地址类型
    - IP4 表示ipv4， IP6表示ipv6
- connection-address: 连接地址
    - 可选

约束：
- 如果addrtype是IP4或IP6，connection-address需要遵循以下规则：
    - 如果是多播，连接地址就是IP多播组的地址
    - 如果是单播，那地址就是数据源的ip地址，也可以是中继的ip地址
    - 虽然没有禁止多播会话中出现单播地址，但最好不要出现
    - 如果是ipv4多播会话，一定要附加一个到多播地址的ttl值(存在时间)
    - 这个ttl的是0-255的, ttl和地址之间用/分割，具体如上面的例子
- IP6 多播是不需要ttl的

分层编码方案：
- 将一个媒体数据源编码成多个层，h264中比较出名的svc，也叫柔性编码
- svc会让媒体服务变成真正的传输分发服务，而不是媒体处理服务
- 在sdp中，接收者可按需选择接收不同质量的流(目前最主要受限于带宽)，做法是订阅不同的层
- 在连接地址中可以选择不同的层，格式：\<base multicast address>[/\<ttl>]/\<number of addresses>
- 上面指定层的写法出现在接收者，可以看上面示例

### 带宽 b= 可选

    b=<bwtype>:<bandwidth>

- 可在会话级和媒体级指定，当然，媒体级会覆盖会话级
- b=CT:... 带宽是一个范围内的
    - CT: conference total 会议整个带宽限制
    - 好处是让两个或多个会议正常进行，因为针对每个会议都做了限制
    - 如果是基于rtp，ct就是所有rtp会话的带宽总和
- b=AS   表明是某个程序的最大带宽
    - 如果是基于rtp，此时as表示的某个rtp会话的带宽
- ct是限制整个会议的带宽，as是限制单个媒体流的带宽
- bandwidth的单位是kb/s,每秒千字节

### 时间 t= 必选

    t=<start-time> <stop-time>

    t=0 0

- 指定了会话的开始时间和结束时间
- 可有出现多次，因为会议有可能在多个不规则间隔的时间段里进行
- 如果间隔时间是有规则的，就用r=代替
- 开始时间和结束时间都是十进制，ntp时间值，从1900开始的秒数
- 结束时间设置为0,表示会议没有时间限制
- 如果开始时间设置为0,表示这个会话是永久性的，写法可参看上面的例子

### 重复时间 r= 可选

    r=<repeat interval> <active duration> <offsets from start-time>

    每周二10点-11点的会议，为期3个月
    t=3034423619 3042462419
    r=604800 3600 0 90000

    下面的写法是一样的
    r=7d 1h 0 25h

- 需配合t= 使用，t= 指定开始时间，r= 指定间隔，结束时间从t=中获取

### 时区设置 z= 可选

    z=<adjustment time> <offset> <adjustment time> <offset> ....

- 调整时区

### 加密key k= 可选

    k=<method>
    k=<method>:<encryption key>

- 传输加密
- 虽然这个k= 的目的是为了兼容其他实现，但是并不推荐使用

### 属性 a= 可选

    a=<attribute>
    a=<attribute>:<value>

    开关类
    a=recvonly

- 附加属性，作为对sdp的扩展手段，在会话级和媒体级都有,媒体级覆盖会话级
- 媒体级描述可以有任意个属性，这是跟具体的媒体相关的
- 属性有两种格式
    - a=flag 这种用于描述开关类属性，一般这种写法更多的是会话属性
    - a=attribute:value 这种有属性名和属性值
- 属性如何解释，要看使用的媒体工具
- 属性名要符合字符集和编码格式
- 一般属性值会遵循字符集和编码格式，不受charset属性的影响，具体要查看标准，很复杂

### 媒体描述 m= 必选

    m=<media> <port> <proto> <fmt> ...

    单播，svc，媒体描述的格式
    m=<media> <port>/<number of ports> <proto> <fmt> ...

    单播svc 例子
    m=video 49170/2 RTP/AVP 31   
    视频，rtp/avp传输协议，端口49170,2个rtp会话,
    第一个rtp会话的rtp/rtcp端口是49170/49171
    第二个rtp会话的rtp/rtcp端口是49172/49173
    格式是31

    多播，number of ports和多播地址有一个对象关系
    c=IN IP4 224.2.1.1/127/2
    m=video 49170/2 RTP/AVP 31 同样的媒体描述
    这里有一些隐含的规则：
    第一个rtp会话的rtp/rtcp的是224.2.1.1的49170/49171
    第二个rtp会话的rtp/rtcp的是224.2.1.2的49172/49173


- 会指定媒体名和传输地址
- 一个sdp会包含多个媒体描述
- 一个sdp的多个m= 指定同样的地址，行为是未定义的，所以不要重复

字段解释：
- media： 媒体类型
    - audio/video/text/application/message 
- port：媒体流发送的端口
    - 这里的端口和c=连接信息中的ip和m=媒体描述中的传输协议，三者配合使用
    - 有些端口是放在a=属性里的，eg：rtp的rtcp端口等(a=rtcp:)
    - 如果有一些非配置的端口使用到了(不在m=中指定的),一定要在属性中指定出来，就像上面的rtcp端口
    - rtcp的端口，不是rtp端口减1,没有这个规定，所以通过属性rtcp指定是有必要的
    - 单播，如果使用了分层编码(svc),可能需要指定多个端口
        - m=格式就如上面写的那样
        - 端口也是依赖于传输协议的，eg：rtp就需要指定偶数端口，number of ports是指rtp会话个数     
    - 多播，会有一些隐含的约束，可参看上面的例子
- proto：传输协议
    - 和c= 连接属性的ip，上面的port端口，是强关联
    - 具体有3中协议：
        - udp，基于udp，并未指定具体哪个协议(应用成协议)
        - rtp/avp, 基于udp，使用rtp来控制音视频
        - rtp/savp，基于udp，使用rtsp协议
- fmt： 媒体格式描述
    - fmt 以及后面的字段都属于媒体格式描述
    - 媒体格式和前面的传输协议很相关
        - 如果协议是rtp/avp或rtp/savp
            - fmt里需要指定rtp payload类型的值，就如上面例子中的31
            - 这个rtp payload类型值可以是一个列表，第一个是默认的
            - 如果rtp payload类型是动态的，需要用a=rtpmap:属性来指定
            - a=rtpmap: 是用来指明payload类型值和具体编码的映射关系
        - 如果协议是udp
            - fmt需要指明具体的媒体类型
            - 这些媒体类型是audio/video/text/application/message类型下的子类型

## sdp属性

下面这些属性都是sdp规定了的，除此之外还有很多属性，不过这些属性都是依赖具体的媒体

- a=cat:分类 上面也提到过,是给接收者进行过滤的，会话级属性，不依赖字符集
- a=keywds:和上面的相反，这个是接收者主动告知要接收哪些分类，会话级属性，依赖字符集
- a=tool:创建sdp信息的工具名和工具的版本，会话级属性，不依赖字符集
- a=ptime:一个包里的媒体时间是多少毫秒，只对音频有效，好处是不用解码就知道pts，媒体级属性，不依赖字符集
- a=maxptime:一个包里最大的媒体时间，只对音频有效，好处是不用解码就知道pts，媒体级属性，不依赖字符集
- a=rtpmap:\<payload type> \<encoding name>/\<clock rate> [/\<encoding parameters>]
    - 使用非常广泛一个属性
    - 将rtp payload类型值和实际编码做一个映射关系
    - a=rtpmap:96 opus/48000/2
    - 96映射的是opus 时钟频率是48k，后面的编码参数是双声道
    - 媒体级属性，不依赖字符集
    - 虽然在m= 媒体描述中也可以指定具体的payload类型，但是是静态的，动态的还是用a=属性映射好
    - 使用动态映射payload类型，更加常用一些
    - 对每种媒体格式都可以映射一个，在m=媒体描述中就更加简洁了
    - 对于音频，编码参数表示声道数，编码参数是可选的，默认是单声道
    - 对于视频，编码参数没有意义
    - 编码相关的参数应该用a=fmtp:属性来指定
- a=recvonly:


