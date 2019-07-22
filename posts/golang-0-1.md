# golang的0到1

一篇基础教程

很多语言针对一个问题,会提出很多种解决问题的方法,golang因为特征少,
正确解决问题的方式往往只有一个,这就节省了时间,让维护大项目更加简单.
这就是golang推崇的简单实用

## 入门

golang以包为基础,编译器知道包是生成可执行,还是库文件.package main是特殊的.

### 工作区

- GOROOT, 一般是go安装的目录
- GOPATH, go定义的环境变量,一般用于定义工作区
- fmt 是内置包,实现格式化i/o的功能
- import func 都是关键字,各有各的用处

### 变量

主要是各种声明的写法:
- var a int
- var a = 1
- a := 1
- var b, c int = 2, 3

### 数据类型

内置的有:
- 数值
- 字符串
- 布尔型

引用类型:
- array 数组
- slices 切片
- map 映射

类型转换需要显示

### 条件语句

- if else
- switch case
- for

### 指针

\*,& 声明和取址

指针常用于:
- 传递结构体参数时
- 已存在类型,添加方法时,常用指针接收者

### 函数

- func 声明
- 多返回值

### 结构体/方法/接口

成为了golang编程的基石

### 包

- main 可执行
- fmt 内置包
- go get 会将包安装在$GOPATH/pkg下
- 包名跟着目录名走
- godoc 包名 Description 显示文档
- godoc -http=":8000" 通过http服务来显示文档
- 常用的内置包
    - fmt 格式化i/o函数
    - json json格式的序列化/反序列化

### 错误处理

- 通过返回值的errors判断
- 可在函数中返回自定义错误: return 123, errors.New("diy error")

### 异常 painc

- painc是无法处理的情况
- painc不是处理错误的理想情况, 错误用errors处理,而异常,是应该停止程序
- painc被捕获后,应该使用recover来恢复运行

### defer 延时

确保退出作用域时, 会执行某个操作

## 并发

内置支持go routine

- go routine的作用:让函数可以和其他函数并发/并行执行
- 协程, 轻量,创建上千个,占用的资源都不大
- 协程通过channel通信
- channel 有双向/单向的
- 使用select来管理多个channel
- 也可以使用带缓冲的channel来实现一些东西

golang成功的根本原因:简单



