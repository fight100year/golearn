# 代码测试 优化 分析

- 优化：部分程序跑的更快，效率更高，资源更少
- 测试：确保程序按我们的意图在执行，边开发边写测试代码
- 分析：明晰程序代码的工作细节

本章主题如下：
- 代码分析
- go tool pprof工具
- web接口
- 测试
- go test
- go tool trace
- 代码基准 benchmarking go code
- 交叉编译
- 从代码生成文档
- 生成例子函数
- 找出未使用代码

## 优化

前提：代码中没有bug

第一点要先找出影响性能的点

## 分析

第一使用命令行接口，第二使用web接口

需要导入runtime/pprof包,net/http/pprof(web编程中导入)

go中支持cpu分析和内存分析，两种同时启用会相互影响，一次最好启用一种。

代码中插桩之后，会输出一个分析文件，用go tool pprof x.out来分析，
top 显示资源占用最高的几个函数，list显示具体代码，
pdf生成pdf报告，svg生成svg报告。

如果觉得麻烦，可以直接使用github.com/pkg/profile包来生成分析文件

go tool pprof -http=[host]:[port] profile 可以支持web接口，
前提是安装Graphviz，在浏览器中可以看到各种分析图。一目了然。

到目前为止，一个竞争分析，一个分析工具 成了每个程序写完之后必备的工具，
相比c++来说，go做的确实很不错。

### Graphviz

用途是画图，里面包含了很多工具，用于维护双向/无向图，和生成图层

画图要用到DOT语言，不过各个语言都有提供对DOT的映射。

开发的话，如果要画图，使用Graphviz是非常不错的。

### go提供的跟踪工具

go tool trace 调优工具,和go tool pprof类似，trace是用来分析跟踪文件的。

跟踪文件有3中生成方式：
- runtime/trace
- net/http/pprof
- go test -trace

go tool trace -http=172.17.0.2:8080 trace.out ,查看web接口

pprof和trace结合使用，让性能更高。

## 测试

测试分很多类很多种，下面只聊到go中的自动化测试。

自动化测试函数的结果只有两种：PASS和FAIL

约定：
- 测试函数在go源码中，并以_test.go结尾
- 测试函数以Test开头，后面的和函数名类似
- 导入testing包，并执行go test子命令

如果要测试非导出函数，在test文件中，函数名也要大写

    go test testMe.go testMe_test.go 执行所有测试项
    go test testMe.go testMe_test.go -v 带信息
    go test testMe.go testMe_test.go -run='F2' -v 运行指定测试项

## benchmarking

基准测试，是几个大公司共同指定的一些测试规则，工具不限，主要用于性能测试，
bechmark测试用于挖掘整个系统的性能(压测)，profile工具是呈现系统运行时的性能指标，
通常使用这两个工具一起做性能测试。

benchmark测试，不要在unix机器比较繁忙时做，和单元测试一样，有一下规则：
- 基准测试函数由Benchmark开头
- 测试文件导入testing包，文件以_test.go结尾，测试命令还是go test

go test -bench=. benchmarkMe.go benchmarkMe_test.go,
其中-bench=.表示哪个测试项要进行测试， 正则表达式.表示匹配所有测试项，
结果如下：
    goos: linux
    goarch: amd64
    Benchmark30f1-4   	     200	   6138675 ns/op
    Benchmark10f1-4   	 3000000	       402 ns/op
    Benchmark30f2-4   	     200	   6614956 ns/op
    Benchmark10f2-4   	 3000000	       436 ns/op
    PASS
    ok  	command-line-arguments	7.220s

Benchmark30f1-4,后面的-4表示启用了4个协程，GOMAXPROCS指定的，
同时也可以看到是linux系统，64位，第二列的时间是函数执行次数，
第三列是单次执行的平均时间，如何要看下内存申请情况，加入 -benchmem,
加入内存后，一列显示函数执行一次申请的内存，另一列是申请了多少次

基准测试可以很明显看到哪种测试项好一些，对比很明显。

写缓冲的长度，会影响到最终效果，用基准测试，可以一目了然地看出对比效果

## 如何找出未使用的代码

在go的思想中，未使用到的代码，是逻辑错误,
这里的未使用，是指永远无法执行到的代码(eg：写在return之后的代码)

go tool vet xx.go 可以检查出来，除了未使用，还可以检测出printf格式、冗余代码、死循环、漏报、
混杂错误、性能等问题，是一个优雅的、项目定期使用的工具


## 交叉编译

go中的交叉编译是指编译出在不同平台使用的可执行二进制文件。

好处是：写代码可以在任意平台。

交叉编译，go已经内置支持

要实现交叉编译，只需要指定平台和arch即可：
env GOOS=windows GOARCH=386 go build xx.go 即可编译出windows上32位的程序，
[go现在支持的参数](https://golang.google.cn/doc/install/source#environment)

如果结合docker，就可在一台机器上实现在多平台开发需求，这就不需要折腾虚拟机和重其他系统了。

## 例子函数

测试源码中包含这一行  // Output: 

这样导出文档的时候，也可以看到这个函数的使用例子

前面提到了自动化测试和基准测试，这也是中特殊测试：

- 源文件已_test.go结尾
- 导入testing包，用go test命令执行
- 例子函数以Example开头
- 例子函数没有入参和返回

## 文档生成

godoc工具来生成文档，
在源码中尽量记录一切信息，很明显的信息不需要记录。
不要有“我创建一个int变量”，而是“这个变量使用来干嘛的”，
最后，好的代码不太需要文档。

文档记录的几个约定：
- 声明前几行用// 来描述，适用于函数、变量、常量、包
- 包的第一行文档，会在生成的文档中描述包，所以第一行文档必须是描述包的文字，而且要完整
- BUG(xxx)的注释，会出现的生成文档的bug章节，(这个不一定要写在申明的前几行)
- 其他注释，在生成文档时会被忽略掉

使用go install xx 先编译，在用godoc -http="172.17.0.2:8080"，就可以找到xx的包信息，
非常方便，例子函数也有了。

go确实提供了很多友好的工具，让开发的注意力更多的放在codeing上，而不是周边。

## go提供的必备工具

- godoc -http=":8080" 经常使用 构建项目文档
- go tool vet 定期使用 检查易遗漏的细小问题
- benchmark test 按需使用 按实际需要调整策略或调整参数
- 自动化测试 经常使用 代码覆盖率要达到一定比例
- go tool pprof 写一个package使用一次 测试cpu和内存 性能分析
- go tool trace 写一个package使用一次 调优工具
- go build/run -race 经常使用 竞争检查 
