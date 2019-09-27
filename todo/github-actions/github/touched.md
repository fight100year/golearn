# cds-snc/github-actions/touched

- docker action
- 主要看push中是否有符合匹配的文件
- 如果不匹配，action就失败，如果匹配，action就通过

## 仓库分析

- readme是常规操作
- 没有license文件 
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine:3.9 as builder
RUN apk add --no-cache crystal=0.27.0-r0 shards=0.8.1-r0 libc-dev=0.7.1-r0 && rm -rf /var/cache/apk/*
WORKDIR /src
COPY . .
RUN crystal build --release --static src/run.cr -o /src/run

FROM scratch

LABEL "name"="touched"
LABEL "maintainer"="Max Neuvians <max.neuvians@cds-snc.ca>"
LABEL "version"="1.0.0"

LABEL "com.github.actions.name"="File touched"
LABEL "com.github.actions.description"="Uses a pattern to see if any files in that pattern have been touched"
LABEL "com.github.actions.icon"="file"
LABEL "com.github.actions.color"="orange"

WORKDIR /app
COPY --from=builder /src/run /app/run
ENTRYPOINT ["/app/run"]
```

## action 分析

- Dockerfile里的入口并不是shell命令
- 而是通过crystal语言编译出的可执行
- 所以具体逻辑也就不用分析了，只关注使用即可

## 使用

```yaml
steps:
- uses: cds-snc/github-actions/touched@master
  with:
    args: "**jpg,**png"
```


## 总结

- action自身没有测试workflow
- docker的入口也是自己编译出的程序
- 这个功能在某些场景也会很有作用，不包含某些文件action就不会通过
