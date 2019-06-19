# v1.5.0官方手册

## 什么是coredns

- dns 服务
- go编写
- 链式插件，所以灵活
- 插件可单机运行，可多机提供dns功能


什么是dns功能：
- 实现了coredns插件api的软件。
- 只要实现了插件api，插件的功能就可以任意发挥。eg：dshow插件机制和dshow的各种filter
- 有的插件接受数据不会进行响应，eg：监控、缓存
- 有的插件只是提供了一些额外的附加功能
- 有的插件可以从文件或数据库中读取数据

coredns定了了插件api，收获的是链式插件，这就是灵活的来源：
- 目前有30个内置插件，还有额外的扩展插件，功能很强大
- 扩展插件需要重新编译coredns
- 写新插件有两点要注意：一是需要用go语言；另一个是需要了解dns内部工作细节

## 安装 coredns

下载源码安装
```
    git clone git@github.com:coredns/coredns
    cd coredns
    make CHECKS= all
```

安装时，使用代理下载包
```
    # Enable the go modules feature
    export GO111MODULE=on
    # Set the GOPROXY environment variable
    export GOPROXY=https://goproxy.io 
```

测试
```
    ./coredns -dns.port=1053
    # 另一个终端
    dig @localhost -p 1053 a www.baidu.com
```

## plugins 插件

coredns运行后就解析配置，并对外提供服务。
每个服务都需要指定端口号和服务的区。每一个服务都有自己的插件链。

服务的区：zone，eg：有的服务负责处理baidu.com的请求，有的负责处理mp3.baidu.com的请求，
其中baidu.com和mp3.baidu.com就称为不同的区。

dns服务，说白了，对外提供的是域名解析服务，coredns对每一个查询，都会有以下步骤执行：
- 如果有多个服务都利用同一个端口对外提供服务，至于用哪个服务来处理查询请求，就看哪个区最匹配
  - 有两个服务 baidu.com和 map.baidu.com
  - 如果查询的是 a.map.baidu.com 那么匹配的是map.baidu.com,这称为longest suffix match，最长后缀匹配
- 一旦处理请求的服务确定了，请求会都给插件链去处理
  - 每一个服务都有自己的插件链，定义在plugin.cfg
  - 插件链的顺序是非常重要的
- 插件拿到请求后，根据请求中的查询名和其他属性，有以下几种处理：
  - 处理
  - 不处理
  - 处理，如果带有fallthrough关键字，请求还会丢给下一个插件进行处理
  - 处理，还会丢给下一个插件处理

插件处理请求，意味着插件会返回一个响应给客户端。

前面也说道了插件可以提供任意想实现的功能，但目前coredns的插件对请求的处理都是上面4种处理策略之一。

下面具体了解一下插件处理请求的细节：
- 处理
  - 当前插件处理请求，返回一个响应给客户端，请求结束
  - 插件链后面的插件不会处理这个请求，因为不会传过去
- 不处理
  - 当前插件不处理，丢个插件链上的下一个插件
  - 如果最后一个插件也不处理，coredns就返回一个SERVFAIL给客户端
- 处理，带fallthrough
  - 当前插件会处理这个请求，如果fallthrough被指定了，下一个插件也会处理
  - 首先会在/etc/hosts中查找，如果找到结果，就返回，如果没有找到结果，就丢给下一个插件
  - 说白了，就是当前插件没有结果返回，就将请求传递给下一个插件
- 处理，带hint
  - 当前插件会处理，也会丢给下一个插件
  - 带hint，是提供一个入口去查看写到客户端中的数据。监控就非常合适用这点查看每个处理阶段的结果

还有一些插件不需要注册：
- 这类插件不处理任何dns数据，但会影响coredns的行为
- eg： bind插件，决定coredns可以绑定哪些接口
- eg： root插件，设置coredns的插件目录
- eg： health插件，请用http的健康检查

插件的套路，或者说细节：
- 每个插件都有启动、注册、处理部分。有些不需要注册，eg：如上
- 启动：解析配置和插件指令
- 处理：处理请求，实现所有逻辑
- 注册：将插件注册到coredns，这个是在编译coredns源码时实现的，服务可以用到所有注册过的插件，至于在运行时具体使用哪些插件，由corefile配置文件决定

## coredns配置

coredns的插件配置是plugin.cfg,配置里决定了哪些插件是要编译进coredns的。  
coredns的运行时配置是Corefile，./coredns -conf 可以指定运行时配置，如果不指定，就是用当前目录的Corefile。

corefile：
- 运行时的配置
- 可以配置一个服务或多个服务，每一个服务可指定使用哪些插件
- 还可以进一步为插件配置指令
- corefile的插件顺序是无所谓的，插件链的顺序实在编译时由plugin.cfg决定
- corefile中的注释用#开头

## 环境变量

corefile配置文件中可以使用环境变量，会在解析配置时替换环境变量。

环境变量的写法是： {$ENV_VAR}

## corefile写法

除了上面提到过的环境变量和注释，corefile还包括以下元素

- import语法
- 服务块
- 协议
- 插件的详细设置
- 扩展插件的集成
- 可能的错误

import语法
- 常用的是代码片段，多个服务可能会使用相同的插件，这部分插件可已抽出来定义成公共的片段，后面导入即可
- 格式： (片段名){公共插件列表}
- 使用：. { import 片段名 }
- 好处：避免出错，简化写法

服务块
- 对应不同的区 zone，有不同的服务去对应，这里的服务就称为服务块
- 服务块定义在 {} 之间
- . { 插件列表 } 这就是一个服务块，.表示为.下面的所有区提供服务
- dns服务的端口默认是53，也可以指定另一个端口。当然如果用启动阐述 -dns.port 指定了，就不用在corefile中指定
- .:1053 { 插件列表 } 表示coredns会在1053端口提供.下面所有区的服务
- 服务的标识由区和端口组成，配置文件中不能重复

协议
- coredns v1.5.0版本支持3种协议：纯dns、基于tls的dns，基于grpc的dns
- 可以为每一个服务指定一个的协议，具体做法是在定义服务块时，在区前面添加下面的前缀来表示
- dns:// 默认方式，纯dns
- tls:// 基于tls的dns
- grpc:// 基于grpc的dns

服务块中插件的详细设置
-  最简单的，只包含插件名  . { chaos } 可以返回版本或作者的信息，从txt记录中获取信息。
-  除此之外还可以配置指令  . { chaos coredns-1.5.0 coredns.io } 
-  用chaos插件配套的指令来查询 dig @localhost -p 1053 ch version.bind txt
-  插件还可以包裹一层  . { plugin { chaos } } 
-  corefile可以有多个服务块 
```
coredns.io:5300 {
    file db.coredns.io
}

example.io:53 {
    log
    errors
    file db.example.io
}

example.net:53 {
    file db.example.net
}

.:53 {
    kubernetes
    forward . 8.8.8.8
    log
    errors
    cache
}
```
这个就是在两个端口上提供了4个zone服务

扩展插件
- 扩展插件默认没有编译进coredns，需要额外添加

可能的错误
- 只要corefile的写法不违背上面服务块和插件配置的写法，就不会出什么问题
- 具体每个插件都有自己的写法和规格，但所有原则不能违背上面的原则

## 启动

coredns常用的启动项有两个：
- dns.port  设置默认服务的端口
- -conf 设置配置文件路径，不设置就取当前目录下的Corefile

转发：
- 使用forward关键字可以进行转发消息
- google 公共dns服务 8.8.8.8
- ibm 公共dns服务 9.9.9.9
- . { forward . 8.8.8.8 9.9.9.9 } 将.下的所有请求都转发到其他dns服务
- 如果发给abc.com的请求要转发给8.8.8.8，剩下的发给/etc/resolv.conf 那么最好的做法是用两个服务块来描述

## 如何写一个插件

方便理解具体插件是如何玩的

- setup.go setup_test.go 负责实现解析配置corefile，一旦在配置文件中发现了插件名，就调用启动函数
- 插件名.go 包含处理请求的逻辑  插件名_test.go 基础的单元测试，用于检查插件是否正常工作
- README.md 显示插件如何配置
- license 许可文件，这个是coredns需要的类apl许可

## todo

接下来
- 了解插件的大致写法
- 看完所有内置插件的提供的功能
- 了解coredns + etcd如何提供服务发现

