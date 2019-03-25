# 了解go内部

前一章的日志，输入输出是基础，
本章关注于go内部的一些细节，其中包括了gc，
怎么调用c代码，被c调用(用到的次数会很少，但go也是提供了一条道路的)，
之后就是panic recover defer等如何使用的讨论

具体分以下几个主题来展开：
- go编译器
- go中的gc是如何运行的
- 如何检查gc的操作
- go中调用c
- c中调用go
- painc 和 recorer函数
- unsafe包
- defer关键字的使用
- linux效率工具 strace
- FreeBSD和macOS中的效率工具dtrace
- go环境中查找信息
- 节点树
- go汇编

## 编译器

编译一个源文件，可使用 go tool compile xx.go,
编译出一个xx.o文件, .o是对象文件，二进制文件。
file xx.o可以查看对象文件的文件类型。
对象文件里面是机器码，大多数情况下不能直接执行，
好处是链接时占用空间少。

go 1.11之后 用go tool compile编译出来的不是对象文件，
虽然也叫xx.o，但用file命令查看是一个归档文件。

归档文件，ar archive格式，二进制，包含一个或多个
文件，在go中，这种归档文件被称为ar。
可用ar t xx.a 来查看归档文件里的内容。

.o是对象文件，.a是什么？  
go tool compile -pack xx.go 生成的就是ar文件。
.a就是归档文件。

go tool compile还有一个常用的参数，是-race,竞争条件，用于并发。
go tool compile -S   这个让程序不容易被反编译

## gc 垃圾回收

go中的gc是并行处理的。

除了从runtime包中获取垃圾回收信息，还可以这样：
GODEUB=gctrace=1 go run gc.go

### 三色算法

go中的垃圾回收机制是基于三色算法，也称三色垃圾回收算法，
三色算法并不是go独有的，也可以用做排序，
在go中，这个算法还有个名称是'三色标记-删除算法'

black set: 不会有指针引用white set的对象

go中也可以手动创建一个gc，通过runtime.GC()，
手动创建的不是异步的，对象多了之后，很容易卡死。

gc最关注的是低延时

stw：stop the world，中断型(非并行)

- 传统的gc算法是 标记-删除，也是最简单的，stw型，增加了延时
- 三色标记-删除算法， 异步的，去掉了stw
- 除了这些，还有分代收集，有更低的延时，未来go可能会吸收进来

## unsafe

绕开类型安全和内存安全的代码，称为unsafe代码

## 和c代码互相调用

c很强力，有些场景使用c可以获得更好的效果：
eg：db，驱动等

如果使用和c互调，使用的次数太多，
就需要重新考虑问题的解决方案，或是换一种编程语言。

使用和c互调的三种场景：
- 为了性能
- 为了和其他语言交互
- go语言实现不了的情况

## defer 关键字

延时执行，执行顺序类似stack顺序，都是后进先出

## 异常和恢复

panic()函数是一个内建函数，结束当前执行，并开始异常处理(记录信息和退出)，
recover()也是一个内建函数，从调用panic的goroutinue中取回控制权。
这里的内建，说的是builtin包

类似c++中的抛出异常，捕获异常

## linux效能工具

查问题的工具,strace dtrace,

strace,用于跟踪系统调用和信号

strace -c xxx 可查看执行时长，系统调用，错误等

dtrace 是动态追踪技术的鼻祖，是系统性能调优的工具之一

## 环境信息

可以从runtime中获取信息，eg：系统信息 cpu等 内核版本 go版本

## go 汇编

## node trees

节点树，go中的node是一个结构体，有很多属性

通过go tool compile -W nodeTreeMore.go 可以查看更多源码编译解析过程

通过go build -x xx.go 可以查看更多构建时的细节，包括更多底层执行命令

## 几个实用技巧

- 如果函数中有error，要么记日志要么返回，没有好的理由，不要两者都做
- interfaces定义了行为，不包含数据和数据结构
- 使用io.Reader io.Writer让代码的扩展性更高
- 函数传值的时候，需要的时候传指针，其他时候传值
- error变量的类型是error，而不是字符串
- 产生环境不要测试go代码
- 使用go的新特性之前，需要自己测试一下，尤其是开发一些给其他人用的工具或app时
- 害怕犯错，只能止步眼前


