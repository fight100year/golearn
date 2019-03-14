---
part1: introduce about go web app
---
# go

- simple
- efficient
- back end system

for web app and *-as-a-server system

large-scale web app features:
- scalable
  - vertical scaling: increasing cpu
    goroutine
  - horizontal scaling: increasing machines
    layer: add a proxy layer
- modular
  easy to add/remove/modify feature
  easy to reusing modular components
- maintainable
- high performance


## web app and web server

app: software program

web app: 
* return html,(client render html and display to user)
* transport data by http

web service:
* not return html
* server for other program

## http

app-level communication protocol

* stateless
* text-based
* request-response
* c/s

### request

a request consists of a few line:
- request-line
- zero or more request headers
- an empty line
- the message body (optional)

GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(empty line)
abc

request-line = request method + uri + http version
request headers = a:b c:d

__request method__  
- get: Tells the server to return the specified resource.
- head: The same as GET except that the server must not return a message body.
- post: Tells the server that the data in the message body should be passed to the resource identified by the URI .
- put: Tells the server that the data in the message body should be the resource at the given URI . 
- delete: Tells the server to remove the resource identified by the URI .
- trace: Tells the server to return the request
- options: Tells the server to return a list of HTTP methods that the server supports.
- connect: Tells the server to set up a network connection with the client.
- patch: Tells the server that the data in the message body modifies the resource identified by the URI .

__idempotent__  
the result of 1 call = the result of 1000 times call

__request_header__  

request header consists:
- info of request
- info of client

if request have a message, Content-Length and Transfer-Encoding is need

common http request headers:
- Accept: Content types that are acceptable by the client as part of the HTTP response.
- Accept-Charset: The character sets required from the server.
- Authorization: This is used to send Basic Authentication credentials to the server.
- Cookie: The client should send back cookies that were set by the calling server.
- Content-Length: The length of the request body in octets.
- Content-Type: The content type of the request body.
- Host: The name of the server, along with the port number.
- Referrer: The address of the previous page that linked to the requested page.
- User-Agent: Describes the calling client.

### response

a response consists of a few line:
- a status line
- zero ro more response headers
- an empty line
- the message body (optional)

status line = status code + reason phrase

__status_code__  
- 1xx: Informational. This tells the client that the server has already received the request and is processing it.
- 2xx: Success. This is what clients want; the server has received the request and has processed it successfully. The standard response in this class is 200 OK.
- 3xx: Redirection. This tells the client that the request is received and processed but the client needs to do more to complete the action.
- 4xx: Client Error. This tells the client that there’s something wrong with the request.
- 5xx: Server Error. This tells the client that there’s something wrong with the request but it’s the server’s fault.

__response_header__  

common http response headers:
- Allow: Tells the client which request methods are supported by the server.
- Content-Length: The length of the response body in octets
- Content-Type: The content type of the response body
- Date: Tells the current time
- Location: This header is used with redirection, to tell the client where to request the next URL.
- Server: Domain name of the server that’s returning the response.
- Set-Cookie: Sets a cookie at the client.
- WWW-Authenticate: Tells header the client what type of authorization clients should supply in their Authorization request header.

---
- rui: name of resource
- rul: location of resource

RUI form:  
    < scheme name> : < hierarchical part> [ ? < query> ][ # < fragment> ]"

http://user:passwd@www.baidu.com/doc/file?name=test&ip=123#sum  

## http2

- focuse on performance
- based on spdy/2
- binary protocol (http1.x base on text)
- full multiplexed
- compress the header
- allow server to push response to client

in go1.6 and after, http2 is default

## define web app

- get a http request form client
- process the requet, do some work
- generate html and return it in an http response message

web app = handlers + template engine

### handlers

- receives and process the http request
- call the template engine to generate the html and something about response 

    mvc(model-view-controller pattern)
    divide a arogram into three parts: model, view, controller
    model - underlaying data
    view - visualization of the model for the user
    controller - use input(from user) to modify the model
    model change --> view update automatically

mvc is a good way for web app, but is not the only way

### template engine

come form ssi technology

there are two type of template with different design philosophies:
- static template or logic-less template
    use placeholder tokens
    no logic
- active template
    placehoder tokens + other programming language 
    eg: jsp asp erb

# example - forums
- [论坛设计](/chitchat/README.md)

---
part2: details in go web program
---
# handling requests

    货物崇拜编程：在不理解需求痛点的情况下，复制一份可运行的代码，
    对代码也不很了解，最后导致扩展很困难。
    换句话说就是使用了不理解的解决方案，导致无法明确预期。
    在编程中，可能是使用了一个强大的框架，但不知道正确的使用规则。

为什么client要持久化cookie，server要持久化session信息：  
因为http是无状态协议，而且每次请求，不会带上上次请求的相关信息

web app的框架首推标准库中的net/htpp + html/template,其次是其他三方库，
为了避免货物崇拜编程，需要了解一下标准库中的一些规则(就像上面的cookie和session)

```go
  func ListenAndServe(addr string, handler Handler) error
  // addr 网络地址，空字符串表示使用80端口
  // handler 为nil，默认处理器(handler)就是默认复用器，DefaultServeMux

  http.ListenAndServe(":80", nil);  // 这个server未做配置
  server := http.Server{Addr: ":80", Handler: nil,}
  server.ListenAndServe()  // 这个是通过Server结构体来配置server
```
监听tcp端口，处理request，

在web app中(golang中)，handler和handler function并不是一个，
handler是一个接口：
```go
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
```
任何实现这个这个方法的，都可以被称为一个handler

接下来看看DefaultServeMux：
```go
// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
// 用两个参数去初始化一个Server结构体，然后调用其中的方法

type Server struct {
    Addr    string  // TCP address to listen on, ":http" if empty
    Handler Handler // handler to invoke, http.DefaultServeMux if nil
    ...
}
// 如果Handler未指定，就默认取http.DefaultServeMux

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
// 结构是ServeMux

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
// 绕了一圈，DefaultServeMux是实现了Handler接口的
```
这个默认Handler的来历是理顺了，那她的作用是：  
根据不同的rul，将请求丢给不同的handler，说白了就是一个默认的复用器。

为什么需要一个复用器？  
如果不用复用器，请求全都会被一个handler处理，
显然从设计上讲，分层是必须的，至少扩展是非常方便的

尽量避免一个handler处理全部的request

```go
    http.Handle("/hello", &hello)
```
用这个来指定对固定uri的处理，除了/是匹配所有，其他的都是完全匹配，
有一点差别都会报404， /hello/ 多了一个/也是404,
为啥设计时不考虑将/hello 和/hello/兼容呢？ 最小惊讶原则

再总结一下Handler：  
- 是一个接口
- 有一个方法来实现这个接口
- http.Handle(uri, &handler)来指定如何处理

什么是handler function：
- 和handler的功能类似
- 不是方法，不是接口，只是一个函数，参数和ServeHTTP一致
- http.HandleFunc(uri, handler function)来指定如何处理

```go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
// 这个直接是使用了DefaultServeMux默认复用器

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
// 实际上调用了ServeMux的方法
// 这个HandlerFunc，是一个函数类似，T(a)这种写法是类型转换
// 用处是将handler function 转换成Handler

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	if pattern[len(pattern)-1] == '/' {
		mux.es = appendSorted(mux.es, e)
	}

	if pattern[0] != '/' {
		mux.hosts = true
	}
}
// 这个就是将uri和handler成对保存起来
// 说白了就是一个handler注册过程，后面遇到对应的uri，调用不同的handler来处理
```

handler 和 handler function, 一个是接口，一个是函数，
handler = HandleFunc(handler function),可使用类型转换直接转换

为什么要提供功能相同，只是写法上有所区别的两种概念：  
实际上使用的还是handler，但写法上，handler function简单很多，所以这就是原因

反过来，既然handler function很方便，为啥还要暴露出handler概念：  
设计上的需求，可以提高模块化(实际上是为了更好的兼容性)

go 不是一个函数性的语言，但函数形语言的一些基本特征还是包括的：  
函数类型，匿名函数， 闭包。

有个新的概念叫aop 和oop可以互补，aop 面向切面编程，属于设计模式的延伸，
其中有个概念叫cross-cutting concern，叫横切关注点，
映射到go中，就是日志handler，安全handler，等很多独立的功能handler，
只关注自己业务上的事，至于和其他业务模块联合起来，就是aop中的织网，
在go语言中，织网可以用chain(链式handler)来实现。

    链式
    假定我们的模块都已经做好了(各个功能性的handler已经完成，且高内聚低耦合)，
    下一步就是织网，在c++中，要考虑的是各个模块调用的参数问题，
    在go中使用统一的格式即可：
    func xxx(h http.HandlerFunc) http.HandlerFunc {}
    这样织网时就不需要考虑模块之间参数的问题，
    A(B(C)), A(C) 按业务进行组合即可

链式的handler，可以非常长，也被称为 pipeline processing
```go
    func xxx(h http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            ...  // 做模块自己的事
            h(w, r)
        }
    }
    // 最后一个链式节点，使用方法(实现Handler的)
```
链式的handler和链式的handler function是一样的

除了默认的复用器，也有第三方的，
ServeMux用来处理固定的uri是很合适的，处理动态uri就力有不怠

http2 需要使用https

# process request

前面提到过：  
RUI form:  
    < scheme name> : < hierarchical part> [ ? < query> ][ # < fragment> ]"

http://user:passwd@www.baidu.com/doc/file?name=test&ip=123#sum  

RUL form：  
scheme://[userinfo@]host/path[?query][#fragment]  
eg: http://www.example.com/post?id=123&thread_id=456

query是k-v键值对，fragment在浏览器发送请求之前就被丢弃了，所以服务端不考虑这个，
但是非浏览器工具发送请求时，有可能有这个fragment，就需要做特殊处理了

http头，在库中用Header类型来表示
```go
    type Header map[string][]string  // 是一个map类型

    // 主要4个基本方法，用于增删改查
    func (h Header) Add(key, value string)
    func (h Header) Del(key string)
    func (h Header) Get(key string) string
    func (h Header) Set(key, value string)

    // Add和Set的区别
    // Header的key是string， value是[]string 切片
    // Set是先创建一个空白切片，切片的第一个值就是value
    // Add是在切面后面追加
```

http消息体，用 Body io.ReadCloser来表示
```go
    Body io.ReadCloser

    type ReadCloser interface {
      Reader
      Closer
    }

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

    type Closer interface {
        Close() error
    }
```
消息体是一个接口，实际上用两个接口来表示：
- Reader：用于读body
- Close：

http GET 请求不带消息体，POST才带，浏览器地址栏的请求都是GET请求，
如果要发post请求，只能用其他工具

## 表单

http post一般会带一个表单，也就是form，
而request中message的组织有多种方式，都可由请求之前指定：
- 简单的文本，使用url编码格式
- 大数据(eg:file),使用multipart- MIME格式
- 二进制，使用base64编码格式

http get是没有消息体，也就是没有body的，她的参数都是加在url后面的

### new book
- [hands on software architecture with golang](/hands_on_golang/README.md)
- [mastering go](/mastering_go/README.md)

### package(pkg)
- [fmt](/pkg/fmt.md)
  
