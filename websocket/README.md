# websocket 库

基于rfc6455实现

约定:
- Conn类型表示ws连接
- 在http的请求处理函数中,调用Upgrader.Upgrade,会返回一个Conn对象
- Conn.WriteMessage conn.ReadMessage会收发消息,消息类型可以是文本/二进制
- 也可以使用io.WriteCloser io.Reader来收发消息

消息:
- 文本(utf-8)/二进制
- 是否是有效的utf-8格式,需要应用程序去判断

控制消息:
- 协议中定义了3种控制消息:close/ping/pong
- 调用以下方法时会发送一个控制消息到对端:WriteControl/WriteMessage/NextWriter
- 收到close消息,默认会返回一个close消息给对端
- 收到ping消息,默认会防护及一个pong消息给对端
- 收到pong消息,默认啥都不做

并发:
- 支持并发,使用一个读对象/写对象即可

缓冲:
- 为了减少系统调用,增加了一个网络缓冲
- 一个websocket帧,每次通过网络传输数据时,都会添加一个websocket帧头
- 减少网络缓冲大小,会导致帧的消耗
- ReadBufferSize/WriteBufferSize可以指定缓冲大小
- 默认缓冲是4k
- 缓冲大小并不影响传递消息的字节数
- 缓冲的生命周期是由连接决定的
- 应用程序可以在内存和性能方面作出平衡,做法就是调节缓冲大小
- 缓冲越大,内存越大,网络读写的系统调用次数越小
- 设置缓冲大小的参考:
    - 缓冲大小最好比消息的预期长度大一点(小了会导致写帧头的次数增加)
    - 缓冲大小比消息的预期长度略少一点(会大幅减少内存消耗,后果是略微影响性能)
    - 具体还要看消息长度的分布状态,最好是大于90%的消息长度,同时可大幅减少内存使用

实验性的压缩:
- 基于rfc 7692
- 可用参数启动此功能

## 常量

- DefaultDialer: 全都是默认值的dialer(Dialer是连接ws服务时的选项集合)
- ErrBadHandshake: 服务端响应时,发现握手时无效时,返回这个
- ErrCloseSent: 应用程序向已关闭的ws发送消息时,会收到这个
- ErrReadLimit: 读的消息长度超过限制时,返回这个

## api中的约束

- 一个连接(Conn),最多只能有一个打开的reader(接收者)
- Conn.NextReader,会读取下一个message,如果前一个消息没有被消费,就丢掉
- 一旦Conn.NextReader返回错误,就需要退出循环,就是这么设计的
- Conn.NextWriter,最多只能有一个打开的writer,前一个如果没写完,会被终结
- 可设置超时,一旦读写超时了,连接就怪掉了

## 服务端使用api流程

- 启动一个http服务,等待升级成websocket的http请求
- 升级成websocket连接之后,创建一个cleint对象来跟踪
- 一个client表示一个ws连接,会有两个协程负责读写
- 之后还有一个管理所有客户端的对象,叫hub,添加上业务就是完整的ws服务

## 客户端使用api流程

- 相对更加简单,以ws url作为参数,利用库创建一个ws连接
- 之后就是处理收发

