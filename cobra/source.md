# 源码阅读

## cobra.Command结构体

说明:
- Command对应的就是命令,一个Command就是一条命令
- git pull, pull就是一个命令
- 这个结构体,需要定义好用法/描述,保证可读行

cobra所有的操作都是围绕着Command来做的,其他的都是和其他工具的兼容,
核心部分,除了标准库,还用到了一个github.com/spf13/pflag库,这是对flag的扩展.

## 源码阅读

- cobra.Command 是一个树型结构,有父节点,也有指节点
- 初始化help命令,前提是没有定义help命令,且没有子命令
- 校验参数是否正确:只有当根命令无子命令,且还有参数时报错
- 解析flag:遇到--就意味着flag结束(不再将后面的数据解析为flag)


## 查看hugo中cobra的用法

todo: 继续补充类图

## 历史版本阅读

### v0.0.1

一个命令的flag分类:
- 一个全集,包含所有的flag
- persistent flag
- local flag
- inherited flag 继承的flag
- 父命令的 persistent flag

如果不考虑command对象的扩展功能,核心功能还是很简单的:
拿到一个命令行,识别命令,解析参数,执行命令,有错就报错.

实际用法也遵循了套路:
- 在init()中定义命令行的flag
- main()调用Execute(),里面调用了解析参数,指定指定命令

init()中的讨论:
- rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
- rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
- addCmd.Flags().StringVarP(&packageName, "package", "t", "", "target package name (e.g. github.com/spf13/hugo)") 
- addCmd.Flags().StringVarP(&parentName, "parent", "p", "RootCmd", "variable name of parent command for this command")

`其他版本也大多类似`,所以不再分析

