# repetitive/actions/auto-pull-request

- docker action
- 新建分支，推送到github时，会自动创建一个pr
- 好处是减少ui操作

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine

LABEL name="auto-pull-request"
LABEL version="1.0.0"
LABEL repository="http://github.com/repetitive/actions"
LABEL homepage="http://github.com/repetitive/actions"

LABEL maintainer="Anton Podviaznikov <anton@podviaznikov.com>"
LABEL com.github.actions.name="GitHub Action for automatically creating Pull Request"
LABEL com.github.actions.description="Create Pull Request when new branch is pushed"
LABEL com.github.actions.icon="git-pull-request"
LABEL com.github.actions.color="purple"
COPY LICENSE README.md /

RUN apk --no-cache add jq bash curl

COPY "entrypoint.sh" "/entrypoint.sh"
ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- 在alpine的基础上，安装了jq bash curl等基础工具

```shell
#!/bin/bash

set -e

if [[ -z "$GITHUB_TOKEN" ]]; then
	echo "Set the GITHUB_TOKEN env variable."
	exit 1
fi

if [[ "$(jq -r ".created" "$GITHUB_EVENT_PATH")" != true ]]; then
	echo "This is not a create push branch!"
	exit 78
fi

if [[ "$(jq -r ".head_commit" "$GITHUB_EVENT_PATH")" == "null" ]]; then
	echo "This push has not commits!"
	exit 78
fi

commit_message="$(jq -r ".head_commit.message" "$GITHUB_EVENT_PATH")"

echo "Commit message:"
echo "$commit_message"

REPO_FULLNAME=$(jq -r ".repository.full_name" "$GITHUB_EVENT_PATH")

DEFAULT_BRANCH=$(jq -r ".repository.default_branch" "$GITHUB_EVENT_PATH")
echo "Creating new PR for $REPO_FULLNAME..."

URI=https://api.github.com
PULLS_URI="${URI}/repos/$REPO_FULLNAME/pulls"
API_HEADER="Accept: application/vnd.github.shadow-cat-preview"
AUTH_HEADER="Authorization: token $GITHUB_TOKEN"

new_pr_resp=$(curl --data "{\"title\":\"$commit_message\", \"head\": \"$GITHUB_REF\", \"draft\": true, \"base\": \"$DEFAULT_BRANCH\"}" -X POST -s -H "${AUTH_HEADER}" -H "${API_HEADER}" ${PULLS_URI})

echo "$new_pr_resp"
echo "created pull request"
```

- 入参有三个
  - token
- 为本仓库的默认分支提交了一个草稿pr

## 使用

- 只需传入token即可

## 总结

- action自身没有测试workflow
- 非常符合功能特征分支的开发模式
- 不管是针对个人项目还是多人协作项目，这个action都非常有作用
