# cobra项目

cli: command-line interface 命令行界面

了解:
- 翻译 眼镜蛇
- 她是一个go库,用于创建更加强大的cli程序
- 也是一个命令行文件
- 很多流行的go项目都使用到了cobra
- cobra是一个库,可以让应用程序创建一个强大的cli,至少可以达到git/go的命令
- cobra是一个程序,是快速开发基于cobra程序的脚手架

特点:
- 基于子命令的cli
- 兼容posix格式(支持长命令/短命令)
- 支持嵌套子命令
- 支持全局标识/本地标识/级联标识
- cobra init appname; cobra add cmdname 
- 智能提示,eg:git输错命令后,会有一个智能提示
- help命令的自动生成
- bash自动补全
- man page自动生成
- 支持命令别名
- 灵活的自定义help/usage
- 可选,可与viper(spf13的另一个项目)一起,为12-factor程序提供功能

cobra中的概念:
- cobra中分命令/参数/标识
- 命令表示动作
- 参数表示某些东西
- 标识表示动作的一些变更

cobra追求的目标:
- 整个命令看起来,就应该像句子一样,易读
- 易懂,无歧义

匹配模式:
- APPNAME VERB NOUN --ADJECTIVE
- APPNAME COMMAND ARG --FLAG

命令 command:
- 应用程序重要的一部分,每一次交互,都可以考虑做成一条命令
- 命令时可以有子命令的,也可以选择执行一个动作

标识 flags:
- 修改命令行为的一种方式
- 兼容posix标识,兼容go的flag包

## 用法

一般使用cobra的项目,在main.go的同级目录会有一个单独目录来存放相关源文件,
一般这个目录叫cmd,或是commands

在main函数中,一般是cmd.Execute() 用于初始化cobra环境

使用cobra生成工具:  
如果程序已经写好了,这时使用cobra程序,可自由添加命令到程序,
可以堪称cobra+已有程序

使用cobra库:  
就像上面提到过的,main函数中初始化cobra环境,然后在指定目录添加源文件即可

简单就是:
- 在app/cmd/xx.go 中 构造一个cobra.Command结构体
- 然后在Execute()中调用cobra.Command.Execute()即可,这样就添加了一个app xx 命令

flag用法:
- flag分两类:
    - 永久flag, 一个命令的flag,她的子命令都可以用这个flag
    - 局部flag, 只有特定的命令能用这个flag
- flag默认都是可选的,可通过api指定flag为必带的,不带就报错

## 源码阅读

- [cobra库源码](/cobra/source.md) 可以作为第三方库,集成到应用程序上,让cli更加强大
- [cobra可执行源码](/cobra/source-app.md) 作为一个基于cobra的应用程序,可以让已存在的应用程序,快速添加强大的cli
- [pflag](/cobra/pflag.md) 对标准库flag做扩展的库, [标准库的flag](/cobra/flag.md)
