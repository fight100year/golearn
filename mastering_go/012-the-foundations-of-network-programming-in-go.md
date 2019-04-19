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

ip：internet protocol 网际协议




