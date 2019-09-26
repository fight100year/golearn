# jessfraz/branch-cleanup-action

- 一个标准的docker action
- 在pr被合并之后，删除分支

## 仓库分析

- license/readme是常规操作
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine:latest
LABEL maintainer="Jessica Frazelle <jess@linux.com>"

LABEL "com.github.actions.name"="Branch Cleanup"
LABEL "com.github.actions.description"="Delete the branch after a pull request has been merged"
LABEL "com.github.actions.icon"="activity"
LABEL "com.github.actions.color"="red"

RUN	apk add --no-cache \
	bash \
	ca-certificates \
	curl \
	jq

COPY cleanup-pr-branch /usr/bin/cleanup-pr-branch

ENTRYPOINT ["cleanup-pr-branch"]
```

## action 分析

- 在alpine中安装了一些基本工具
- 其中就包含了curl，pr大部分操作都是利用curl完成(倒是可以考虑使用hub来代替)


```shell
#!/bin/bash
set -e
set -o pipefail

if [[ -n "$TOKEN" ]]; then
	GITHUB_TOKEN=$TOKEN
fi

if [[ -z "$GITHUB_TOKEN" ]]; then
	echo "Set the GITHUB_TOKEN env variable."
	exit 1
fi

URI=https://api.github.com
API_VERSION=v3
API_HEADER="Accept: application/vnd.github.${API_VERSION}+json"
AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"

# Github Actions uses either status code 0 for success or any other code for failure.
# Docs: https://help.github.com/en/articles/virtual-environments-for-github-actions#exit-codes-and-statuses
NO_BRANCH_DELETED_EXIT_CODE=${NO_BRANCH_DELETED_EXIT_CODE:-0}

main(){
	action=$(jq --raw-output .action "$GITHUB_EVENT_PATH")
	merged=$(jq --raw-output .pull_request.merged "$GITHUB_EVENT_PATH")

	echo "DEBUG -> action: $action merged: $merged"

	if [[ "$action" != "closed" ]] || [[ "$merged" != "true" ]]; then
	    exit "$NO_BRANCH_DELETED_EXIT_CODE"
	fi

	# delete the branch.
	ref=$(jq --raw-output .pull_request.head.ref "$GITHUB_EVENT_PATH")
	owner=$(jq --raw-output .pull_request.head.repo.owner.login "$GITHUB_EVENT_PATH")
	repo=$(jq --raw-output .pull_request.head.repo.name "$GITHUB_EVENT_PATH")
	default_branch=$(
		curl -XGET -fsSL \
			-H "${AUTH_HEADER}" \
 			-H "${API_HEADER}" \
			"${URI}/repos/${owner}/${repo}" | jq .default_branch
		)

	if [[ "$ref" == "$default_branch" ]]; then
		# Never delete the default branch.
		echo "Will not delete default branch (${default_branch}) for ${owner}/${repo}, exiting."
		exit 0
	fi
	is_protected=$(
		curl -XGET -fsSL \
		-H "${AUTH_HEADER}" \
		-H "${API_HEADER}" \
		"${URI}/repos/${owner}/${repo}/branches/${ref}" | jq .protected
	)

	if [[ "$is_protected" == "true" ]]; then
		# Never delete protected branches
		echo "Will not delete protected branch (${ref}) for ${owner}/${repo}, exiting."
		exit 0
	fi

	pulls_with_ref_as_base=$(
		curl -XGET -fsSL \
			-H "${AUTH_HEADER}" \
			-H "${API_HEADER}" \
			"${URI}/repos/${owner}/${repo}/pulls?state=open&base=${ref}"
	)
	has_pulls_with_ref_as_base=$(echo "$pulls_with_ref_as_base" | jq 'has(0)')


	if [[ "$has_pulls_with_ref_as_base" != false ]]; then
		# Do not delete if the branch is a base branch of another pull request
		pr=$(echo "$pulls_with_ref_as_base" | jq '.[0].number')
		echo "${ref} is the base branch of PR #${pr} for ${owner}/${repo}, exiting."
		exit 0
	fi

	echo "Deleting branch ref $ref for owner ${owner}/${repo}..."
	response=$(
		curl -XDELETE -sSL \
			-H "${AUTH_HEADER}" \
			-H "${API_HEADER}" \
                        --output /dev/null \
			--write-out "%{http_code}" \
			"${URI}/repos/${owner}/${repo}/git/refs/heads/${ref}"
	)

	if [[ ${response} -eq 422 ]]; then
		echo "The branch is already gone!"
	elif [[ ${response} -eq 204 ]]; then
		echo "Branch delete success!"
 	else
		echo "Something unexpected happened!"
		exit "$NO_BRANCH_DELETED_EXIT_CODE"
	fi

	exit 0
}

main "$@"
```

- 必选入参只有一个： token
- 可选入参也有一个：如果删除分支出现错误，action的退出码，默认是0(不影响job执行)，也可自定义
- action的流程分析：
  - action执行的场景是：已经合并，且pr已经close
  - 不删除默认分支
  - 不删除保护分支
  - 如果执行了git rebase，就正常退出
  - 分支已经没了，也是正常退出
  - 如果分支删除分支出错，就返回预置的返回码(可通过入参指定)

## 使用

- hcl语法，还未转换成yaml

## 总结

- 在整个自动化流程中，这个action可以使用，也是一个非常方便的action
- 因为review和megre两个操作都需要人来参与，所以这个action够用了
