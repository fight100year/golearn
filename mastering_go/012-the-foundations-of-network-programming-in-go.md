# 网络编程的基础

前一章提到了很多实用工具：
- 基准测试 自动化测试 例子函数
- 交叉编译
- 分析 (pprof trace -race)
- 文档

本章主题：
- tcp/ip
- ipv4 ipv6
- netcat 命令行工具
- dns解析
- net/http包
- web例子
- 网站
- wireshark和tshark
- http超时


## net包

net包，处理网络，包括http、rpc、mail等，下一章节具体说，
http包，提供了http和https请求。

http.RoundTripper是一个接口，表示处理http事务的能里,
简单讲，就是有接受request的能力，有返回response的能力。

- http.Response 表示响应
- http.Request 表示请求
- http.Transport 是传输，实现了http.RoundTripper,支持http和https，还支持http代理

## tcp/ip

互联网中著名的协议族，取名于tcp协议和ip协议

tcp：transimission control protocol 传输控制协议

两个机器用tcp传输数据时，是一段一段的传，每一小段称为一个tcp包，
主要特征是可靠，意味着开发者不需要额外的代码来保证tcp包的传递。
如果没有证据表明tcp包已经被传递了，那么这个tcp包就会被重传。

一个tcp包可用于：
- 建立连接
- 传输数据
- 确认发送
- 关闭连接

建立连接后，一个全双工的通道就在两端建立了，这样就可以收发数据，
如果连接断掉了，也会有一套机制告诉相关的程序。

ip：internet protocol 网际协议,有了网际协议，就可以进行网络寻址。
ip协议不是可靠的，无连接的。tcp/ip来保证数据传递的可靠性。

udp：user datagram protocol, 用户数据报协议，不可靠，但速度快

ipv4:32位二进制，192.168.100.1；
ipv6:128位二进制，3fce:1706:4523:3:150:f8ff:fe21:56cf

## netcat

nc工具,被称为瑞士军刀，
功能很多，可以进行端口扫描、传文件、聊天、shell等

- 作为tcp/udp客户端
    nc ip port // tcp方式
    nc -u ip port // udp方式
- 模拟服务端
    nc -l 
- 查网络问题
    nc -v 或是 nc -vv

## 网卡接口信息

数据链路层的以太网协议，规定每个接入网络的设备都要有一个网卡接口，
和之前说的一样，一个网卡接口信息包括4方面：
- ip
- 子网掩码
- 网关
- dns

本地一般有一个叫lo的网卡(下面所有的网卡接口network interface简称网卡)，
这时一个特殊的虚拟网卡，用于本地回环，就是本机内部和自己交互。

## web编程

go中的web编程在效率和安全都做了很多事，但是，
如果要支持模块、多站点、虚拟主机等强大功能时，推荐使用apache和nginx。

所以go在web编程做做小事还是ok的，其他还不行。

apache的ab压力测试工具，用于测试web编程

抓包工具可以使用tcpdump、wireshark(tshark是命令行版本)

## hugo

go实现的静态站点生成器，和jekyll类似，spf13在google主导这个项目





