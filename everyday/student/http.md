# net/http

说明:
- 标准库提供的，可实现http client，也可以实现http server
- 支持get head post postfrom 这些http请求
- client在处理完响应之后，需要关闭响应，濡染会有协程泄露
- 如果要控制client的请求头，eg：重定向策略等，需要创建一个http.Clinet对象
- 如果要控制代理、tls配置、保活、压缩等，要创建一个http.Transport对象
- Client和Transport对象是可以被多个协程并发使用，效率也高
- ListenAndServe函数，会开始一个http服务，第一个参数中会指定ip和端口
- ListenAndServe函数的第二个参数会指定使用哪个复用器，如果为nil，就使用默认复用器
- 默认复用器是DefaultServeMux，Handle和HandleFunc函数会将处理器添加到默认复用器
- 服务端的控制，需要创建一个http.Server对象
- go1.6开始，在使用https时默认使用http2协议。也可以手动关闭http2，方式是调用指定函数或环境变量
- Server和Transport对象都是自动启用了http2的，这是默认的配置，也是一个简单配置
- 如果要控制更复杂的配置，或是想试试最先版的http2，或是启用更多特征，可试试golang.org/x/net/http2


零散的说明：
- 请求头的最大字节数默认是1M，可以修改

