# 论坛

整个设计过程分以下部分：
- 了解需求(论坛和用法)
- 应用设计
- 数据模型
- 请求的接收和处理
- html的产生(响应)
- 附加主题
- 测试
- 总结

## 了解需求

互联网交流产品之一，
老一辈的是bbs，usenet，电子邮件等，
现在google group等都在利用论坛来聚集大批用户

论坛有几个概念：
- thread，主题
- post，帖子

用户发一个帖子，叫主题(帖子的一种)，
后面用户的回复，叫post

很多论坛都是按类分层，不同的话题有不同的子论坛，
部分人会有特权，称为版主

注册用户可以进行回贴，非注册只能看，不能回帖

## 应用设计

设计成一个典型的web app，收请求，然后返回一个响应

- 设计request的格式
    http://server-ip/handler-name?params
    handler(web app的两大核心之一，用于接收和处理request，
    另一个是模板引擎，用户产生html页面), 也分分层的,
    阅览主题的handler，可以取名为thread/read
    参数可取主题id

- 设计web app流程
    - client 发送request
    - server(web app) 把请求丢到多路复用检查器
    - 将请求分发给不同的handler
    - 处理完调用模板引擎产生html，返回response

## 数据模型

使用关系数据库 PostgreSQl

- 用户信息
- 会话信息
- thread信息
- post信息

## request的接收和处理

http包里有一个对象叫 复用器， 用处是将请求重定向到指定handler

http包还可以提供文件服务，静态文件的

http template:
    嵌入特定指定的html文件，称为actions，
    换句话说 actions就是代码有注解的html，注解格式：{{注解}}
    每一个模板文件都会定义个template

