# 在unix上可以做什么

本章主要是系统编程，
毕竟就是因为go语言的发明者不满系统软件不好开发而做的，
所以说系统编程，在go中占有很大的成分

本章包含以下主题：
- 进程
- flag包
- io.Reader io.Writer接口的使用
- 处理信号
- 支持pipe
- 读文本文件
- 读csv文件
- 写文件
- bytes包
- syscall的高级使用
- 目录遍历
- 文件权限


## 进程

进程就是一个执行环境，包含指令，用户数据，系统数据，其他运行时资源

进程分3种：
- 用户进程 运行在用户空间，无特殊访问权限
- 守护进程 运行在用户空间，不依赖终端，运行在后台
- 内核进程 运行在内核空间，能访问所有内核数据

go并不像c提供了创建子进程的fork(),而是提供goroutine

## flag

flag 包，用于解析命令行参数

如果写的程序是带命令行参数的，使用flag包没错

## io.Reader io.Writer

要不要使用缓冲，都可以

写文件的方式很多：
- ftm.Fprintf
- os.File.WriteString()
- bufio
- ioutil.WriteFile
- io.WriteString

## 信号signal

os/signal包

利用signal包可以捕获信号，
这个包比flag用的更加广泛，比较每个服务程序都需要有这么一个信号处理

## pipe

unix哲学：一个程序只做一件事，做好。

多个程序之间的交互常用pipe，在pipe中，一个程序的输出是另一个程序的输入，
这样就可以组合多个程序，来实现更复杂更自定义的需求

pipe的限制：
- 两进程有一个共同的先祖(有名管道去掉了这一限制)
- 单向

## 目录遍历

## eBPF

enhanced Berkeley Packet Filter, 增强伯克利包过滤,
内核态的虚拟机，可以和linux 内核交互，用于linux跟踪

前面说道的strace和dtrace都是跟踪，eBPF是linux通用的。

要在linux中使用eBPF, 内核编译时需要带CONFIG_BPF_SYSCALL选项，
ubuntu是自动开启的。

eBPF要求的linux内核版本较高，macOS上不适用。

go中可以使用https://github.com/iovisor/gobpf

标准库syscall中也可以跟踪一些信息，eg：寄存器

syscall的一些跟踪用法还是需要进一步深入之后再分析。

画图可以选用glot： https://github.com/Arafatk/glot


