# maddox/actions/ssh

- docker action
- 打印环境变量和evnet payload

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine

LABEL "repository"="http://github.com/hmarr/debug-action"
LABEL "homepage"="http://github.com/hmarr/debug-action"
LABEL "maintainer"="Harry Marr <harry@hmarr.com>"

LABEL "com.github.actions.name"="Debug Action"
LABEL "com.github.actions.description"="Log the action's environment variables and event payload"
LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="yellow"

RUN apk --no-cache add jq

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- Dockderfile就配置了一个jq工具(格式化json的工具)

```shell
#!/bin/sh

set -e

echo
echo "-- Environment variables ----------------------------------------------"
env
echo "-----------------------------------------------------------------------"

echo
echo "-- Event JSON ---------------------------------------------------------"
cat "$GITHUB_EVENT_PATH" | jq -M .
echo "-----------------------------------------------------------------------"
echo
```

- 可以看出，打印了环境变量和事件信息

## 使用

```yaml
steps:
- uses: hmarr/debug-action@v1.0.0
```

## 总结

- 在开发调试时，会有一定助力
