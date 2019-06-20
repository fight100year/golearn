# 对caddy的使用和配置

## 完全新手

这部分针对 以前没玩过webserver的人

安装部分参考官网，可选择二进制，也可以选择编译

配置部分，部署的机器要暴露80和443端口

caddy -host example.com 为域名开启https访问，中间会遇到错误，需要给权限 sudo setcap cap_net_bind_service=+ep $(which caddy)

不过如果要开启https，还需一个真实域名

使用Caddyfile来配置站点，caddy会默认去加载Caddyfile， 也可以用-conf 指定Caddyfile的地址

## Caddyfile 入门

- Caddyfile 是一个纯文本文件
- 第一行是站点地址  eg：localhost:8000, 跨机器访问可写 ip:port
- 之后的行里面都是指令，常用的会在下面提到
- gzip 开启gzip压缩: 即消息体的gzip压缩
- log access.log  表明记录访问日志
- 指令的参数配置灵活，eg：gzip无参数、log后面带一个参数
- 注释使用#
- 多个站点都可以配置在一个文件里，每块用 {} 包裹，而且还有严格的格式要求
- 如果多个站点的配置一样，那多个站点可共用一个 {} 
- 站点名可做一定的匹配，eg 可以用 * 来做占位符
- 配置文件里可以使用环境变量  写法固定  {$PORT}

## todo

- Caddyfile 写法规则
- Caddyfile 插件
- 域名使用和sdk开发
