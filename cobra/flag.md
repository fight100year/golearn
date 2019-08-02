# 标准库flag

标准库的实现只用到标准库中的库,
flag的作用是解析命令行参数.


## 分析

- 差不多1/3的代码定义了一种值类型接口,实现这个接口的有bool/int/uint/int64/uint64/float64/string/duration
- 这个值类型接口定义了对值的读和写,以及字符串读
- 实现值类型的接口,有各自的构造函数
- 剩下的篇幅,全部是围绕FlagSet展开的
- FlagSet:一个结构体,代表的是一组定义好的Flag,每一个FlagSet都有一个名字
- flag包默认有一个全局的FlagSet对象,叫CommandLine, 在init()

## 源码分析

值类型分析:
- 具体看类图

flag.Parse():
- 过程很简单,命令行参数重头到尾一个个分析,是不是预定义的flag
- 如果是,设置相应的值,然后分析下一个命令行参数

flag.IntVar()
flag.BoolVar()
flag.UintVar()
flag.Int64Var()
flag.Uint64Var()
flag.StringVar()
flag.Float64Var()
flag.DurationVar()
- 最后调到falg.Var()
- 根据传入的参数的k 查看flag是否存在
- 存在就报错,不存在就实例化一个flag存储起来

flag.Visit() 系列:
- 便利flag集合,对每个flag都执行指定函数
- flag集合有两种:一种是程序能接收的,也就是代码中指定的可接收flag集合
- 另一种是运行时,通过命令行参数指定的,可能是上面集合的一个子集

flag.Lookup():
- 查找通过flag的k,查找flag

flag.Set()
- 通过flag的k,设置flag的值(手动设置,使用场景在定义预期flag之后)

flag.PrintDefaults():
- 遍历所有的flag,打印 -flag名 值类型 usage 默认值 换行

flag.NFlag()
- 返回通过命令行参数设置的flag个数

flag.Arg()
- 返回第几个参数, 这个参数是命令行参数

flag.Args()
- 返回完整的命令行参数

flag.Int()
flag.Int64()
flag.Uint()
flag.Uint64()
flag.String()
flag.Float64()
flag.Duration()
flag.Bool()
- flag.BoolVar()的一个封装

flag.Parsed()
- 判断是否已经解析过了

## 支持的格式

源码分析:
- -a=b 写法 支持所有数据类型
- -a 如果是bool型 默认a是true
- -a b 如果不是bool型 等同于a=b, 如果a是bool型,那么b就无法识别为一个flag了

换句话说:
- -a=b 最中庸的写法,是ok的
- -a 写法 只适用于bool
- -a b 写法 只适用于非bool型
- 对标准flag库来说 前缀是-和--都是同样的效果

官方文档上写的是:

    -flag
    -flag=x
    -flag x  // non-boolean flags only
    前缀是一个破折号或两个破折号是等价的

注:-flag写法只适用于bool型,非bool型会报一个flag needs an argument的错误,源码之下,了无秘密

## 分析完源码之后,查看官方文档

flag: 实现了命令行标识的解析(命令行标识就是flag,一个命令行参数可能包含多个flag)

用法:
- 定义程序预期的flag
    - 这个定义的动作反在解析之前即可,一般推荐反在init()中
    - 定义的方式有多种,flag包提供了IntVar()系列/Int()系列
    - 除了定义规定的数据类型,还可以通过Var()定义自定义数据类型的flag
- 解析flag
    - flag.Parse()
- 使用通过命令行参数传进来的值(这也是这个包最大的意义)
    - 在定义预期flag时,就指定了解析之后,值存放在哪个变量上,使用时直接用这个变量即可

除了核心的上面三步外,还提供了以下辅助功能:
- 取命令行参数/取第一个命令行参数/取命令行参数个数
- 取命令行参数设置flag的个数/
- 打印所有预期flag信息: flag名 数据类型 usage 默认值
- 手动修改flag的值
- 针对所有的flag或已设置的flag,执行某个函数


文档的额外提示:
- 多个flag可以绑定到一个变量,这样可以实现长flag/短flag,坏处是如果命令行参数同时指定长短flag,只能以后出现的为准
- 未指定的flag的默认值,都是go语言的零值

## 总结

flag在大部分场景都是够用的,以下是实现时未考虑到的:
- 长短flag,不够精细
- 解析:遇到一个未解析的,后面的就不再解析

[flag类图和主要函数的流程图](https://www.draw.io/?mode=github#H63isOK%2Fconference_graph%2Fmaster%2Fhugo%2Fflag)
