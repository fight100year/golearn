# pflag

相比标准库中的flag,pflag添加了对posix/gun风格的 --flags,
说白了就是多支持了一种写法,就是长参数写法.

相比标准库中的flag,只是做了一些功能的新增,是flag的一个简单的代替品.

从flag换到pflag,只有一个需要注意的:Flag创建实例时,多来一个成员变量Shorthand

## 分析

- 查看源码目录下的文件看,值类型都单独出来了,值类型也做了扩展
- 额外的就是flag.go 和golangflag.go两个文件

类图整理

源码分析:
- pflag针对flag解析失败就退出的情况,做了一个处理:忽略出错的,继续解析
- FlagSet也扩展了,

值类型的变更:
- pflag对所有值类型的实现都独立出来,成为独立的源文件
- 所有flag中的值类型,在pflag中全部重新实现了,对外暴露的接口和flag保持一致,这样就兼容了flag
- 值类型除了flag中对外暴露的接口,还多了几个实现: VarP()系列
- pflag中长flag都用--指定, -c都认为是短flag,flag包中暴露的值类型接口,在pflag中认为是长flag
- VarP()系列就是短flag
- 值类型的种类较flag包,多了很多种,甚至基于flag包的封装格式,也可以当作是pflag值类型的一种

从源码文件名和大致浏览代码,可知falg.go文件是所有核心处理,其他都是值类型,或值类型相关的测试

继续分析:
- pflag重新定义了FlagSet和Flag数据结构,也定义了值类型接口Value,用于自定义符合pflag的值类型
- 下面只需要理清楚pflag所有对外暴露的方法的流程即可

parse():
- 遍历命令行参数,按长短flag分别处理

值类型的定义文件,也就是那些源文件,同时了指定了预期flag的定义,parse做了flag的解析,
下面看看还暴露了哪些方法,看了一下,和flag暴露的方法基本一致:
- Args() 命令行参数除了flag参数,剩下的就是程序参数,至少说明是这样的non-flag command-line arguments
- Arg() 读取第几个程序参数,就是除了flag参数之外的参数
- NArg() 返回程序参数的个数
- NFlag() 返回通过命令行设置的flag数量
- ParseAll() 解析命令行参数,之后对解析的flag执行指定函数,而Parse()执行的存储
- PrintDefaults() 打印Usage信息 和flag的有些许差异
- SetInterspersed() 是否在命令行参数中支持程序参数和flag参数混搭, 默认是开启混搭
- UnquoteUsage() 和flag一样,去掉第一对\`\`
- Var()指定长falg,没有短flag
- VarP()可以选择长短同时指定,也可以只指定长flag一种,`不能只指定短flag`
- Visit() VisitAll() 和flag的一致,分别是对已设置的flag/全部预期的flag,执行某个操作

## pflag的flag规则

- --a=b 最常见的长flag格式
- --a 适用于有默认值的flag
- --a b 适用于没有默认值的flag
- -a=b 最常见的短格式
- -a 适合有默认值的flag
- -abc 等价于 -a=bc
- -a b 适合没有默认值的flag

和flag的规则还是有很大不同的,flag格式:
- -a=b 常见格式
- -a b 适合非bool型
- -a 适合bool 型

可以看出 pflag的适用范围更大一些.还有一点:
- flag遇到 -- 会跳过,接着解析下一个falg
- pflag遇到 --, 会终止解析

issuse:
- FlagSet.Parse() L1119 "if len(arguments) < 0 {" 是否有意义,不应该是判断是否有参数吗,应该是 < 1才对啊
- #209 ,flag包中判断的是 == 0

## 总结

- 除了flag定义的规则有不同外,实现上功能上也略有不同
- 不过pflag提供了一个golangflag的值类型,这样pflag就能完全兼容flag了
- 说白了,不兼容的适用golanflag,要兼容gun,就适用pflag新增的功能
- 其次,pflag还扩展了很多值类型,总之就是强大

[pflag类图和主要函数的流程图](https://www.draw.io/?mode=github#H63isOK%2Fconference_graph%2Fmaster%2Fhugo%2Fflag)第二页
