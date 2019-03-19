# go调试

## delve调试

下文只讨论delve在golang中进行源码级调试的过程

安装： go get -u github.com/go-delve/delve/cmd/dlv

源码级调试可以选择delve

- 使用-- 来传递flags参数

### 调试命令

cmd | 说明
---|---
arg | 显示函数参数
break (b) | 设置断点
breakpoints (bp) | 打印断点
call | 注入一个函数调用，并恢复处理
check (checkpoint) | 当前位置创建一个checkpoint
checkpoints | 打印已存在的checkpoint
clear | 删除breakpoint
chear-checkpoint (clearcheck) | 删除checkpoint
clearall | 删除多个breakpoint
condition (cond)| 设置条件断点、跟踪点
config | 修改配置
    config -list | 显示所有配置参数
    config -save | 将配置持久化到磁盘
    config <k> <v> | 修改配置的值
    config substitute-path <from> <to> | 新增替换规则
    config substitute-path <from> | 移除替换规则
    config alias <cmd> <alias> | 定义一个别名
    config alias <alias> | 删除一个别名
continue (c) | 继续执行


## gdb调试
