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
deferred | 命令延时执行
disassemble (disass) | 汇编
down | 当前栈 下移
edit (ed) | 打开源码
exit (quit/q) | 退出调试
frame | 设置当前栈针，或在某一栈针执行命令
funcs | 显示函数列表
goroutine | 显示或更改当前协程
goroutines | 显示程序所有协程
help (h) | 帮助
list (ls/l) | 显示当前源码
locals | 显示局部变量
next (n) | 源码下一步
on | 设置命中断点执行的命令
print (p) | 打印表达式
regs | 打印cpu寄存器
restart (r) | 重新执行，起点可以是checkpoint,也可以是事件
rewind (rw) | 继续执行
set | 修改变量值
source | 执行一个文件中的所有delve命令
sources | 显示所有的source 文件
statck (bt) | 显示栈
step (s) | 单步,会进入到子函数
step-instruction (si) | 汇编单步
stepout | 跳出当前函数
thread (tr) | 切换线程
threads | 显示所有跟踪线程
trace (t) | 设置一个跟踪点,断点的一种，不过不会中断，只会显示一些信息
types | 显示类型,支持正则匹配，类似funcs
up | 栈帧上移
vars | 显示package变量
whatis | 显示表达式类型





## gdb调试
