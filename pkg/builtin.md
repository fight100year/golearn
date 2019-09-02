# builtin

语言内置包，作用是给一些标识符做声明

```
go list -json builtin
{
	"Dir": "/usr/local/go/src/builtin",
	"ImportPath": "builtin",
	"Name": "builtin",
	"Doc": "Package builtin provides documentation for Go's predeclared identifiers.",
	"Target": "/usr/local/go/pkg/linux_amd64/builtin.a",
	"Root": "/usr/local/go",
	"Match": [
		"builtin"
	],
	"Goroot": true,
	"Standard": true,
	"GoFiles": [
		"builtin.go"
	]
}

```

## 说明

这个包是特殊的

- 源码中的源文件，只是给godoc用的，语言实现内部，还是有一个对应builtin包
- 源码中更多的是声明，没有相关的实现
- 源码中用 type bool bool 或是 type int int 来声明类型


## 源码分析

具体源码就不贴在这儿了

- bool类型 true和false的声明
- int系列类型
- float32和float64
- complex64和complex128
- string
- uintptr 指针类型，所有其他的指针类型，都可以存放
- nil 变量，支持指针/channel/func/interface/map/slice类型
- 别名：byte和rune
- 常量 iota 无类型的int，初始值是0
- 接口类型 type error interface{ Error() string }
- 内置方法：
    - append 适用于slice，在切片后面新增一个元素
        - 用法简单：第一个参数是切片，第二个是元素，支持可变参，返回新切片
    - copy 将切片拷贝到目的地址
        - 可将string的byte拷贝到[]byte
        - 拷贝可能会发生重叠现象
        - 返回拷贝byte的数目(这个是src和dst的最小值) 
    - delete 适用于map，删除map中的一个kv对
        - 第一个参数是map对象，第二个是key值
        - 如果map对象为nil，或map中找不到key，就什么都不做
    - len 计算对象的长度
        - array： 数组的元素个数，和容量cap不是一个概念
        - array的指针，那就计算*array的元素个数，指针可能为nil
        - slice/map： 元素个数，如果slice/map为nil，len()返回0
        - string: byte的个数
        - channel： 缓冲区中未读的元素个数，如果channel是nil，len()返回0
    - make 申请并初始化对象，类型可以是slice/map/chan
        - 第一个参数是类型，第二个是个数,支持可变参
        - 返回值就是类型的对象，new返回的就是对象的指针
        - 针对slice，第三个参数表示容量cap，
            - 可以不指定，如果不指定，容量和长度一样
            - 如果指定了，容量是需要大于等于长度的
        - 针对map，后面的长度可以省略
        - 针对channel，后面的长度表示带缓冲的channel，长度为0或省略表示非缓冲channel
    - new 申请对象，返回指针，参数就是类型
    - complex 构造一个复数对象
    - real 从一个复数对象中取实数部分
    - imag 从一个复数对象中取虚数部分
    - close 关闭一个channel，通道要么是双向的，要么是只发送的
        - 只能由发送端执行，执行关闭的时机是：最后一个发送的元素被接收到了，就执行关闭
        - 从一个已关闭的通道接收数据，不会阻塞，会立马返回，操作返回false
    - panic 结束当前协程的正常执行
        - 一个函数中调用了panic，这个函数的defer函数会正常执行
        - 调用panic会返回到调用方，就是这个函数的调用者(也是一个函数)
        - 直到 - 这个协程上的第一个函数
        - 此时，会结束程序，报告错误信息，包括引起异常的值
        - 上面被称为异常的执行序列
        - 异常可以被内置函数recover控制
    - recover 允许程序接管一个发生异常的协程
        - 将recover放在defer函数中，不放在defer函数中，将无法接管
        - recover会结束异常序列，恢复正常执行
    - print/prinln 打印信息到stderr，用于bootstrapping和调试
        - 后面不一定有了，可能会被遗弃掉

## 总结

- 感觉这个包有点大杂烩，类型接口，常量变量，函数都有
- 这个包的主要作用还是提前声明，具体实现也不在这，想看也只能通过汇编来分析
