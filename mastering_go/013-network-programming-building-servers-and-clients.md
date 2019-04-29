# 网络编程 服务端和客户端

上一章介绍了tcp/ip的基础知识，包括网卡接口、dns等相关处理，
也提供了http的服务端和客户端，http的超时，http，跟踪分析。

本章关注于tcp/udp服务端和客户端的编程。

主题包括如下：
- net包
- tcp
- tcp的并发服务端
- udp
- rpc

## net包

- net.Dial() 客户端连接服务器时调用
- net.Listen() 服务端监听

连接和监听时可以指定协议：
- tcp
- tcp4 ipv4
- tcp6 ipv6
- udp
- udp4 ipv4
- udp6 ipv6
- ip
- ip4
- ip6
- Unix (unix sockets)
- Unixgram
- UnixPacket

net包出了支持tcp和udp，还支持其他协议，包括自定义协议

## tcp客户端

    数据链路层以太网协议的首部是 源mac地址 + 目的mac地址 + 类型
    网络层ip协议的首部是 源ip 目的ip
    传输层tcp协议的首部是 源端口 目的端口

## rpc

在go中，rpc远程调用是一种基于tcp/ip的cs机制，服务于ipc 进程间通信

如果定义一个接口，服务端对象类型名要和接口名一样就可以运行，
为啥要一样具体不清楚，只能在后面的学习中解惑

除了这些，还可以实现自定义的的网络协议，比如说捕获发送icmp包，
太底层了，后面是个方向，现在不需要了解



