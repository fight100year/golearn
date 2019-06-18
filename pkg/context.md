# context

上下文

这个包里包含了一个Context接口类型，支持取消操作，
最主要的也是支持取消

Context类型用于跨api边界和进程，传输`截止时间`、`取消信号`、`其他请求范围的值`

为了保证满足跨包时的接口一致性和静态分析工具检查context时的传递性，要遵循以下规则：
- context不要放在结构体里
- context应作为函数的第一个参数进行传递，一般命名为ctx
- 不要传递nil context，不确定可已传递context.TODO
- 只有跨进程、跨api，访问请求范围的值，才使用context值。如果要传可选参数，可以利用api参数实现。
- context是协程安全的

## 什么时候适合用context

Context 上下文，是go并发模式的一种。2014年提出的。

作为服务端，每来一个请求，我们都会用一个单独的协程去处理，
更多的时候，会开额外的协程去处理后台数据访问、rpc服务访问。
这是常态，用户的一个请求，需要多个协程来完成以下事情：
用户校验、认证token、请求的deadline。如果这个请求取消，
或超时了，那跟这个请求相关的协程都必须快速终止，释放资源。
这时，context模式就很适用。

Context是一个接口类型，在跨api边界时，可以携带一些值，这些值种类很多，
但都只有一个目的：让协程快速结束

```golang
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```
上面的Done()返回一个channel，这个channel表示取消信号，channel关闭意味着请求取消了，
这时需要协程停止工作，并return  
Err()描述了取消的原由  
一个操作若需要多个子操作来完成一些事情，可以将子操作放在协程里，有一点需要明确：
子操作里不能取消父操作。  
Deadline()可以设置超时时间，i/o操作时比较适用  
Value()可以带一个请求相关的值，这个值要可以被多个协程使用

## Context的派生

go中有个Context树，如果其中一个Context被取消了，那这个Context派生的其他Context也被取消了。
只有根节点不会被取消

```golang
// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context

// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// WithValue returns a copy of parent whose Value method returns val for key.
func WithValue(parent Context, key interface{}, val interface{}) Context
```
上面的Background()可以创建一个根节点  
接下来的Withxxx()都是从第一个参数的Context进行派生一个新的Context

当请求返回时，和这个请求相关的Context都应该取消。
WitchCancel()对于取消冗余请求非常有用，eg：多个协程在获取数据(冗余处理)，
只要有一个返回，就取消其他的。WithTimeout非常适合客户端放给服务端的场景。


