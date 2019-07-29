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

