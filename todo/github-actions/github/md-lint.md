# ocular-d/md-linkcheck-action

- docker action
- md lint

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine:3.10
LABEL "com.github.actions.name"="md-linkcheck"
LABEL "com.github.actions.description"="Validate links in markdown files."
LABEL "com.github.actions.icon"="link"
LABEL "com.github.actions.color"="green"
LABEL "repository"="https://github.com/testthedocs/md-linkcheck-action.git"
LABEL "homepage"="https://github.com/testthedocs/md-linkcheck-action"
LABEL maintainer="svx <sven@testthedocs.org>"

# Version of markdown-link-check
ENV MD_LINKCHECK 3.7.3

# Install
# hadolint ignore=DL3018
RUN apk add --no-cache \
        bash \
        nodejs \
        npm \
    && apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing \
        fd \
    && npm install --no-cache -g markdown-link-check@${MD_LINKCHECK}

WORKDIR /srv
COPY entrypoint.sh entrypoint.sh

RUN chmod +x entrypoint.sh
#ENTRYPOINT ["bash"]
ENTRYPOINT ["/srv/entrypoint.sh"]
```

## action 分析

- 环境配置都是最简的 alpine + 基本工具 (包括action的主角：npm安装markdown-link-check)

```shell
#!/usr/bin/env bash
#set -eu
set -eou pipefail

# Vars
ESC_SEQ="\\x1b["
RESET=$ESC_SEQ"39;49;00m"
YELLOW=$ESC_SEQ"33;01m"
RED=$ESC_SEQ"31;01m"

echo -e "${YELLOW}==> Checking Links <==${RESET}"

fd -e md -x markdown-link-check {} \; 2> error.txt
#exec markdown-link-check {} \; 2> error.txt


if [ -e error.txt ] ; then
  if grep -q "ERROR:" error.txt; then
    echo -e "${RED}Please check the log${RESET}"
    exit 1
  fi
fi
```

- 设置了输出的颜色
- 检查之后，如果遇到error级别的日志，就报错

## 使用

```yaml
# 常规写作
# 每次在push时才进行检查
# 下面的写法，适用于将action的代码克隆下来，再执行
# 可以使用通用方法，直接使用 - uses: ocular-d/md-linkcheck-action@1.0.1

on:
  push:
    paths:
    - '*.md'
    - '/docs/*'
name: Testing linkcheck
jobs:
  markdown-link-check:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: md-linkcheck
      uses: ./

# 定时调度
# 15分钟调度一次

on:
  schedule:
    # * is a special character in YAML so you have to quote this string
    - cron:  '*/15 * * * *'
    paths:
    - '*.md'
    - '/docs/*'
name: Testing linkcheck
jobs:
  markdown-link-check:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: md-linkcheck
      uses: ./
```

## 总结

- 可以利用这个action来进行md的检查
