#  RTP Payload Format for H.264 Video

- h264 视频的rtp有效载荷格式
- 2011年进行了标准化，rfc号是6184





## payload 格式的参数


### sdp参数

接收方遇到本文未提到的参数，应该忽略

#### payload 参数到sdp参数的映射

针对h264的video，在sdp中涉及到的部分应该如下：
- m=video
- a=rtpmap 中要指定h264
- a=rtpmap 的时钟频率是90000

可选参数(以a=fmtp 开头，分号分割，格式是：参数=值)：
- profile-level-id
    - 16进制数，eg：profile-level-id=42001f
    - 3个字节，分别表示profile\_idc,profile-iop,level\_idc
    - profile\_idc,和profile-iop指定使哪种profile，4200表示基础profile，不一定遵循A.2所有条款
    - level\_idc表示一个等级，涵盖了帧率/码率/分辨率 
    - 针对176×144推荐使用 42000b
- max-recv-level
- max-mbps
- max-smbps
- max-fs
- max-cpb
- max-dpb
- max-br
- redundant-pic-cap
- use-level-src-parameter-sets
- in-band-parameter-sets
- level-asymmetry-allowed
- packetization-mode
    - 表示载荷类型
    - 0 表示nal
    - 1 表示 非交错
    - 2 表示 交错，隔行扫描
- sprop-interleaving-depth,sprop-deint-buf-req
- deint-buf-cap
- sprop-init-buf-time
- sprop-max-don-diff
- max-rcmd-nalu-size
- sar-understood
- and sar-supported

#### sdp offer/answer模式的使用

- level-asymmetry-allowed 
    - 表示是否开启 level asymmetry,这个是不对称等级，不开启，offer和answer的level要保持一致
    - 不管是offer和answer中，只要置为0或不设置，表示不开启level asymmetry
    - 如果offer和answer中都设置为1,才开启不对称等级，这是针对提议/应答模式的个设置

#### 在声明性会话描述中的用法
