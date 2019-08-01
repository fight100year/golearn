# 源码阅读

## cobra.Command结构体

说明:
- Command对应的就是命令,一个Command就是一条命令
- git pull, pull就是一个命令
- 这个结构体,需要定义好用法/描述,保证可读行

cobra所有的操作都是围绕着Command来做的,其他的都是和其他工具的兼容,
核心部分,除了标准库,还用到了一个github.com/spf13/pflag库,这是对flag的扩展.


