# maddox/actions/wait-for-200

- 检查一个http的状态是否是200
- 可附加失败条件：重试多少次，每次多长时间,查看代码之后，发现这些都是必选项

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM debian:stable-slim

LABEL "maintainer"="maddox <jon@jonmaddox.com>"
LABEL "repository"="https://github.com/maddox/actions"
LABEL "version"="1.0.1"

LABEL "com.github.actions.name"="Wait for 200"
LABEL "com.github.actions.description"="Poll a URL until it returns a 200 HTTP status code."
LABEL "com.github.actions.icon"="refresh-cw"
LABEL "com.github.actions.color"="blue"

RUN apt-get update && apt-get install -y curl

ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- 和sleep的环境有些不一样，sleep的环境更加简单一点
- 这个Dockerfile指定的是debian，还安装了curl工具
- 一个标准的Docker action写法，Dockerfile中指定了自定义entrypoint.sh

entrypoint.sh:

```shell
#!/bin/sh

set -e

while [ $MAX_TRIES -gt 0 ]
do
  STATUS=$(curl -L --max-time 1 -s -o /dev/null -w '%{http_code}' $URL)
  if [ $STATUS -eq 200 ]; then
    exit 0
  else
    MAX_TRIES=$((MAX_TRIES - 1))
  fi
  sleep $SECONDS_BETWEEN_CHECKS
done

exit 1
```

- 利用shell来实现的一个功能
- 结合curl只取http状态码，实现了一个简洁的查询功能


## 使用

```yaml
steps:
- name: sleep x second
  uses: maddox/actions/wait-for-200@master
  with:
    url: www.baidu.com          # 要查询的url地址
    seconds_between_checks: 2   # 每次查询的间隔
    max_tries: 10               # 最大重试次数
```

## 总结

- action没有测试 workflow
- 查询url的状态码，还添加了重试功能
- 在某些场景下会非常有用
