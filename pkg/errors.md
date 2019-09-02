# errors

一些管理错误的函数

```golang
go list -json errors

{
	"Dir": "/usr/local/go/src/errors",
	"ImportPath": "errors",
	"Name": "errors",
	"Doc": "Package errors implements functions to manipulate errors.",
	"Target": "/usr/local/go/pkg/linux_amd64/errors.a",
	"Root": "/usr/local/go",
	"Match": [
		"errors"
	],
	"Goroot": true,
	"Standard": true,
	"GoFiles": [
		"errors.go"
	],
	"XTestGoFiles": [
		"errors_test.go",
		"example_test.go"
	],
	"XTestImports": [
		"errors",
		"fmt",
		"testing",
		"time"
	]
}

```

## 说明

这是一个应用非常广泛的包，同时，也是一个非常简单的包：
- 使用广泛: go1.12.6 154/193的使用，除了少数几个功能包，基本都用到了这个包
- 简单，等会看源码分析

作用：
- 对error类型附加了一个属性：错误说明
- error是一个接口类型，后面会具体分析到，在builtin包中

## 源码分析

```golang
// Package errors implements functions to manipulate errors.
package errors
 
// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}
 
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}
 
func (e *errorString) Error() string {
    return e.s
}
~        
```

- 实现了error接口(实现了Error方法)
- 提供了一个对象创建的方法New
- 对error的扩展是：扩展了一个错误信息
- 特点是：简洁

## 例子分析

    errors_test.go分析

- 隐含约束测试
    - 任意两次用New创建的对象都是不一样的，至少，errorString对象是不同的
    - 只有自己和自己做比较，才是true，其他的都是false
- 方法测试
    - New方法太简单了，并未单独做测试，标准库并未为了覆盖率的数字而追求一些不靠谱的东西
    - 在Error方法中，还是调用了New方法的
- example
    - 除了实现接口的方法外，主要就是添加新增方法的Example
- example的扩展
    - 使用本包的例子测试
    - 一般是测试另一个包的功能，这个功能又是利用本包实现的
    - 主要目的是告诉使用者，除了直接调用本包，还可以使用基于本包的封装包
    - 说白了就是：如果不想使用本包，也可以使用更加高级一点的包来代替
    - 这个例子说的是：可以用fmt.Errorf()来代替本包的New，来创建一个实例

    example_test.go分析

- 这个不是针对包中扩展方法如何使用的展示
- 而是更高一个层次的探索，和errors包同级，毕竟errors只是为error附加了一个错误信息
- 如果我们的场景需要更多信息，可自定义一个包，具体实现，就是example\_test.go想表达的
- 这类文件，主要是显示在文档中，给使用者一个思路上的提示

## 总结

- 集成测试和单元测试都有做
- 不单独为简单方法做测试，而是在其他方法测试时体现(不要为了数字上的好看，而忘了初心)
- 测试的函数名规则：
    - 集成测试： Test功能()
    - 单元测试： 测试方法： Test方法名Method(); 如果是函数，应该也会有对应的写法
- 除了普通测试(单元测试/集成测试/基准测试/例子等),还可以有和包同级的diy实现测试 
