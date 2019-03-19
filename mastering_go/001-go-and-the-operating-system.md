# go与操作系统

本书第一句话说出了真诚："go主题的介绍，可能会有轻微天真和幼稚"，
言下之意是我们可能把牛皮吹大了一点

和实践学科类似，动手是理解最好的方式

本章会介绍以下主题

- go历史
- go的优势
- 编译、执行go代码
- 下载使用已有的go包
- unix的标准输入输出和错误输出
- 打印
- 捕获输入
- log文件处理
- 错误处理

本书会分3个部分：
- 介绍go中的概念
- 代码组织、项目设计、go中的高级特征
- 实际运用手段，系统编程、并发编程，测试 优化 分析，网络编程

go支持过程、编发、分布式编程

默认是静态链接，一次编译，多处运行，不用担心依赖库

go的缺点是：
- 不直接支持oop，但是有代替的概念
- go不会取代c，而目标是java和c++
- 对unix系统来说，c是无法取代的，因为unix是c写的

言下之意是go还准备取代c

go发布时带有很多工具包，有个叫godoc的工具，可以直接查看
函数或包的信息，而不需要联网查，可以在命令行当命令查看，
也可以在一个命令行程序中使用，这样通过web浏览器查看结果
- 命令行中，充当命令
    godoc fmt Printf 查看fmt包中Printf函数的信息
    godoc fmt 查看fmt包的信息
- 运行一个web服务，用浏览器查看
    godoc -http=:8081 通过网页访问，可以看到和官网一样的网站

go不太注重源码的文件名，main函数是入口，一个项目只能有一个

go build file  编译  
go run 运行，生成的可执行文件随后会删除掉

go的两条规则：
- 要么使用package，要么不包含package
- 大括号的格式化，只有一种方式

除了标准库，还可以使用github上的库，eg：  
go get -v github.com/mactsouk/go/xxx,
下载的项目在$GOPATH/src/github.com/下面,
go get还会对项目进行编译，结果放在 $GOPATH/pkg/linux_amd64/github.com/

go clean -i -v -x github.com/mactsouk/go/xxx 清理不需要的项目

unix文件描述符：
- 0 /dev/stdin    os.Stdio
- 1 /dev/stdout   os.Stdout
- 2 /dev/stderr   os.Stderr

fmt有很多打印函数：
- fmt.Print() 一般打印
- fmt.Println() 多了一个换行，参数之间多了加一个空格
- fmt.Printf() 格式化打印
 
除此之外，还有S家族打印函数,fmt.Sprint(),fmt.Sprintln(),
fmt.Sprintf(),用于组成一个字符串
 
F家族打印函数,fmt.Fprint(),fmt.Fprintln(),fmt.Fprintf(),
用于向一个文件写字符串

标准输入输出和打印一族不是同一个概念，打印fmt负责格式化输入输出，
标准输入输出是在os package中

输入的捕获：
- 程序的命令行参数
- 用户输入
- 从文件读

var 用于声明变量，常用于声明无初始化值的全局变量

标准输出和标准错误输出，有时候很难区分，因为她们都打印在屏幕上，
在少数bash上是有颜色上的区分，再重定向的时候，有区分是非常有用的。

日志的好处：
- 持久化，信息不会丢失
- unix上很多工具可以对日志信息做二次处理，如果是终端打印，就很不方便了

日志对调试也很有帮助，特别是对go程序

错误error在go中有很重的分量，错误条件是一件事，怎么处理这个错误是另一件事。
错误处理在go中很重要，基本上绝大部分函数都会返回一个错误类型的对象

panic，报异常的工具函数，如果使用太多，表示代码需要重构。

