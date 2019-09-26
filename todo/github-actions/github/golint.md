# ArangoGutierrez/GoLinty-Action

- docker action
- 在pr事件发生时，调用go lint进行检查

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM golang:1.13rc1-alpine
MAINTAINER Eduardo Arango <eduardo@sylabs.io>

LABEL version="1.1.0"
LABEL repository="https://github.com/ArangoGutierrez/GoLinty-Action"
LABEL maintainer="ArangoGutierrez"

LABEL com.github.actions.name="Go-Linty"
LABEL com.github.actions.description="Linty support an iterative process to clear out lints from go code"
LABEL com.github.actions.icon="activity"
LABEL com.github.actions.color="green"

RUN apk add --no-cache \
	bash git grep \
	ca-certificates \
	curl \
	jq

RUN go get -u golang.org/x/lint/golint

COPY linty/action /usr/bin/github_action

ENTRYPOINT ["github_action"]
```

## action 分析

- 在配置环境时，安装了一些常用工具
- 下载了官方的golint工具
- docker的入口是linty/action的文件，这个文件是一个shell文件

```shell
set -e
set -o pipefail

if [[ -z "$GITHUB_TOKEN" ]]; then
	echo "Set the GITHUB_TOKEN env variable."
	exit 1
fi

if [[ -z "$GITHUB_REPOSITORY" ]]; then
	echo "Set the GITHUB_REPOSITORY env variable."
	exit 1
fi

# Small hack to sanitize GOPATH
WORKDIR="/go/src/github.com/${GITHUB_REPOSITORY}"
# create go work dir
mkdir -p ${WORKDIR}
# copy all files from workspace to work dir
cp -R /github/workspace/* ${WORKDIR}
cp -R /github/workspace/.linty ${WORKDIR}
# cd into the work dir and run all commands from there
cd ${WORKDIR}

if [[ ! -f ".linty/linty.conf" ]]; then
	echo "No linty conf file found"
	exit 1
fi

URI=https://api.github.com
API_VERSION=v3
API_HEADER="Accept: application/vnd.github.${API_VERSION}+json;application/vnd.github.antiope-preview+json"
AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"

linty(){
for pkg in `go list ./...`; do
	# Check package for lint
	lint=$(golint -set_exit_status ${pkg} 2>/dev/null)
	has_lint=$?

	# Check if the package is expected to have lint
	if grep -Fxq $pkg ".linty/linty.conf"; then
		if [ "$has_lint" -eq 1 ]; then
			# Still has lint...
			lint_count=$(echo "$lint" | wc -l)
			printf " %5s | %s\n" "$lint_count" "$pkg" >> "$1"
		else
			# Lint free!
			echo -e  "*WARNING:*"
			echo -e  "The package $pkg is lint free, but is listed in the LINTY config.\n"
			echo -e  "Please remove it from 'linty.config'!\n"
			rm $1
			exit 0
		fi
	else
		if [ "$has_lint" -eq 1 ]; then
			# New lint...
			echo -e  "$lint\n\n"
			echo -e  "ERROR: package $pkg contains NEW lint. Please address the issues listed above!\n"
			rm $1
			exit 0
		fi
	fi
done

# Sort results by count
sort -nr $1 -o $1

# Print results table
echo -e "##  L I N T Y   W A L L   O F   S H A M E "
echo -e " Count	| Name of Linty Package\n"
while IFS= read -r var
	do
		echo -e  "$var\n"
	done < "$1"
echo -e "\\n"
echo -e "Help LINTY fight the good fight, golint today!\n"
}

main(){
	action=$(jq --raw-output .action "$GITHUB_EVENT_PATH")
	ref=$(jq --raw-output .pull_request.head.ref "$GITHUB_EVENT_PATH")
	REPO_OWNER=$(jq --raw-output .pull_request.head.repo.owner.login "$GITHUB_EVENT_PATH")
	REPO_NAME=$(jq --raw-output .pull_request.head.repo.name "$GITHUB_EVENT_PATH")
	ISSUE_NO=$(jq --raw-output .number "$GITHUB_EVENT_PATH")

# Temporary table file
tmp_file="/tmp/linty.${ref}"
touch $tmp_file
if [[ ! -f "/tmp/linty.${ref}" ]]; then
	echo "Failed to create /tmp/linty.${ref}"
	exit 1
fi

linty_out=$(linty $tmp_file)
JSON_STRING=$( jq -n \
				--arg bd "$linty_out" \
				'{body: $bd}' )

echo ""
echo "${JSON_STRING}"
echo ""

# the real action!
curl -sSL -H "${AUTH_HEADER}" -H "${API_HEADER}" \
 -H "Content-Type: application/json" \
 -X POST -d "${JSON_STRING}" \
 "${URI}/repos/${REPO_OWNER}/${REPO_NAME}/issues/${ISSUE_NO}/comments"

# Remove temporary file
if [[ -f "$tmp_file" ]]; then
	rm $tmp_file
fi
}

main "$@"
```

## 使用

- 这个action是利用一个配置文件来配置lint的
- 这个配置文件里列出了一个lint的表，哪些lint启用都可以配置
- 这个配置放在源码目录的 .linty/linty.conf

最后的使用如下：

```yaml
name: Lint
on:
  pull_request:
    branches:
    - master
jobs:

  Lint:
    name: GoLinty
    runs-on: ubuntu-latest
    steps:

    - name: Check out code into the Go module directory
      uses: actions/checkout@v1
    - name: Go-Linty
      uses: ArangoGutierrez/GoLinty-Action@go-1.13rc1-alpine
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

- 非常简单的一个使用
- 只针对master分支上的pr


## 总结

- action自身没有测试workflow
- 针对pr，确实有这个lint检查，会方便很多
