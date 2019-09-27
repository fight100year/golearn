# pullreminders/assignee-to-reviewer-action

- docker action
- 对于pr，基于负责人来自动分配审阅人

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM debian:9.6-slim

LABEL "com.github.actions.name"="Assignee to reviewer"
LABEL "com.github.actions.description"="Automatically create review requests based on assignees"
LABEL "com.github.actions.icon"="arrow-up-right"
LABEL "com.github.actions.color"="gray-dark"

LABEL version="1.0.4"
LABEL repository="http://github.com/pullreminders/assignee-to-reviewer-action"
LABEL homepage="http://github.com/pullreminders/assignee-to-reviewer-action"
LABEL maintainer="Abi Noda <abi@pullreminders.com>"

RUN apt-get update && apt-get install -y \
    curl \
    jq

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- 标准的docker action 写法
- 在配置环境时，除了debian，还安装了curl和jq

```shell
#!/bin/bash
set -eu

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo "Set the GITHUB_TOKEN env variable."
  exit 1
fi

if [[ -z "$GITHUB_EVENT_NAME" ]]; then
  echo "Set the GITHUB_REPOSITORY env variable."
  exit 1
fi

if [[ -z "$GITHUB_EVENT_PATH" ]]; then
  echo "Set the GITHUB_EVENT_PATH env variable."
  exit 1
fi

API_HEADER="Accept: application/vnd.github.v3+json; application/vnd.github.antiope-preview+json"
AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"

action=$(jq --raw-output .action "$GITHUB_EVENT_PATH")
number=$(jq --raw-output .pull_request.number "$GITHUB_EVENT_PATH")
assignee=$(jq --raw-output .assignee.login "$GITHUB_EVENT_PATH")

# Github Actions will mark a check as "neutral" (neither failed/succeeded) when you exit with code 78
# But this will terminate any other Actions running in parallel in the same workflow.
# Configuring this Environment Variable `REVIEWERS_UNMODIFIED_EXIT_CODE=0` if no reviewer was added or deleted.
# Docs: https://developer.github.com/actions/creating-github-actions/accessing-the-runtime-environment/#exit-codes-and-statuses
REVIEWERS_UNMODIFIED_EXIT_CODE=${REVIEWERS_UNMODIFIED_EXIT_CODE:-78}

update_review_request() {
  curl -sSL \
    -H "Content-Type: application/json" \
    -H "${AUTH_HEADER}" \
    -H "${API_HEADER}" \
    -X $1 \
    -d "{\"reviewers\":[\"${assignee}\"]}" \
    "https://api.github.com/repos/${GITHUB_REPOSITORY}/pulls/${number}/requested_reviewers"
}

if [[ "$action" == "assigned" ]]; then
  update_review_request 'POST'
elif [[ "$action" == "unassigned" ]]; then
  update_review_request 'DELETE'
else
  echo "Ignoring action ${action}"
  exit "$REVIEWERS_UNMODIFIED_EXIT_CODE"
fi
```

- 这个action是，给pr指定负责人时，就将负责人同时设置为审阅者，取消的时候也一样
- 入参有3个：
  - token
  - 仓库
  - 触发事件

## 使用

- 在使用时，可用下列事件来限制一下覆盖范围：

```yaml
on:
  pull_request:
    types: [assigned,unassigned]
```

- 好处是不用管其他事件的处理了，action对这种情况，默认返回78
- 作者也提供了一个方法，通过环境变量，将这个78改为0
- 这么做的是为了不将并发的job取消掉

## 总结

- action自身没有测试workflow
- 一个普通的docker action
- 一些注释没写的那么清晰，导致yaml写法的3个入参需要进一步研究
