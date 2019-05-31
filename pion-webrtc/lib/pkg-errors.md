# errors包

因为内置的error包，在打印信息时并没有带上上下文信息，
errors做了一个扩展，在error的基础上添加了上下文信息。

errors的未来，因为go2中支持了更加强大的错误处理，
errors并不进行演进了，而保持在1.0版本中。


## Frame

```go
    type Frame uintptr
```

表示堆栈中的一条，封装的功能有：
- 取函数名
- 取源文件全路径
- 取函数对应的行号
- 支持转换成字节流

## StackTrace

```go
    type StackTrace []Frame
```

表示的是一个堆栈,封装的功能有：
- 格式化，将整个堆栈都转换成字节流

## stack

维护的是一个程序计数器，用于runtime堆栈信息和 StackTrace 之间的一个转换

# 说明

errors 主要是在error的基础上添加了堆栈信息。
