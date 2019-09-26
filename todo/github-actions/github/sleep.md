# maddox/actions/sleep
t
- sleep 几秒
- 一个docker action

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM busybox:latest

LABEL "maintainer"="maddox <jon@jonmaddox.com>"
LABEL "repository"="https://github.com/maddox/actions"
LABEL "version"="1.0.1"

LABEL "com.github.actions.name"="Sleep"
LABEL "com.github.actions.description"="Stall execution for N seconds"
LABEL "com.github.actions.icon"="moon"
LABEL "com.github.actions.color"="yellow"

ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- 一个标准的Docker action写法，Dockerfile中指定了自定义entrypoint.sh

entrypoint.sh:

```shell
#!/bin/sh 
set -e

sleep "$*"
```

- 具体实现的功能非常简单：
  - 输入一个数字，然后调用sleep命令来休眠x秒


## 使用

```yaml
steps:
- name: sleep x second
  uses: maddox/actions/sleep@master
  with:
    args: 15
```

## 总结

- action没有测试 workflow
- 这个sleep功能确实非常简洁
