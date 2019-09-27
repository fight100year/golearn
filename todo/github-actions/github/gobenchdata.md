# bobheadxi/gobenchdata

- docker action

## 仓库分析

- licence/readme是常规操作
- 这是一个标准的docker action

```Dockerfile
FROM golang:latest

LABEL maintainer="Robert Lin <robert@bobheadxi.dev>"
LABEL repository="https://go.bobheadxi.dev/gobenchdata"
LABEL homepage="https://bobheadxi.dev/r/gobenchdata"

# version label is used for triggering dockerfile rebuilds for the demo, or on
# release
ENV VERSION=master
LABEL version=${VERSION}

RUN apt-get update && apt-get install -y --no-install-recommends git && rm -rf /var/lib/apt/lists/*
ENV GO111MODULE=on
RUN go get -u go.bobheadxi.dev/gobenchdata@${VERSION}

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

最后调用的shell脚本

```shell
#!/bin/bash
set -e

# generate some defaults
INPUT_SUBDIRECTORY="${INPUT_SUBDIRECTORY:-"."}"
INPUT_PRUNE_COUNT="${INPUT_PRUNE_COUNT:-"0"}"
INPUT_BENCHMARKS_OUT="${INPUT_BENCHMARKS_OUT:-"benchmarks.json"}"
INPUT_GO_TEST_PKGS="${INPUT_GO_TEST_PKGS:-"./..."}"
INPUT_GO_BENCHMARKS="${INPUT_GO_BENCHMARKS:-"."}"
INPUT_GIT_COMMIT_MESSAGE="${INPUT_GIT_COMMIT_MESSAGE:-"add benchmark run for ${GITHUB_SHA}"}"

# output build data
echo '========================'
command -v gobenchdata
gobenchdata version
env | grep 'INPUT_'
echo "GITHUB_ACTOR=${GITHUB_ACTOR}"
echo "GITHUB_WORKSPACE=${GITHUB_WORKSPACE}"
echo "GITHUB_REPOSITORY=${GITHUB_REPOSITORY}"
echo "GITHUB_SHA=${GITHUB_SHA}"
echo "GITHUB_REF=${GITHUB_REF}"
echo '========================'

# setup
mkdir -p /tmp/{gobenchdata,build}
git config --global user.email "${GITHUB_ACTOR}@users.noreply.github.com"
git config --global user.name "${GITHUB_ACTOR}"

# run benchmarks from configured directory
echo
echo '📊 Running benchmarks...'
RUN_OUTPUT="/tmp/gobenchdata/benchmarks.json"
cd "${GITHUB_WORKSPACE}"
cd "${INPUT_SUBDIRECTORY}"
go test \
  -bench "${INPUT_GO_BENCHMARKS}" \
  -benchmem \
  ${INPUT_GO_TEST_FLAGS} \
  ${INPUT_GO_TEST_PKGS} \
  | gobenchdata --json "${RUN_OUTPUT}" -v "${GITHUB_SHA}" -t "ref=${GITHUB_REF}"
cd "${GITHUB_WORKSPACE}"

# fetch github pages branch
echo
echo '📚 Checking out gh-pages...'
cd /tmp/build
git clone https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git .
git checkout gh-pages

# generate output
echo
echo '☝️ Updating results...'
if [[ -f "${INPUT_BENCHMARKS_OUT}" ]]; then
  echo '📈 Existing report found - merging...'
  gobenchdata merge "${RUN_OUTPUT}" "${INPUT_BENCHMARKS_OUT}" \
    --flat \
    --prune "${INPUT_PRUNE_COUNT}" \
    --json "${INPUT_BENCHMARKS_OUT}"
else
  cp "${RUN_OUTPUT}" "${INPUT_BENCHMARKS_OUT}"
fi

# publish results
echo
echo '📷 Committing and pushing new benchmark data...'
git add .
git commit -m "${INPUT_GIT_COMMIT_MESSAGE}"
git push -f origin gh-pages

echo
echo '🚀 Done!'
```

- 可以看出，这个action是执行指定的基准测试，让后将结果发布到github page上

## 使用

```yaml
steps:
- name: checkout
  uses: actions/checkout@v1
  with:
    fetch-depth: 1
- name: gobenchdata to gh-pages
  uses: ./
  with:
    PRUNE_COUNT: 30
    GO_TEST_FLAGS: -cpu 1,2
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## 总结

- 这个action和具体的项目代码集合在一起，不是太好
- 其次一个action包含了ci/cd，中间能扩展的地方太少了
- 这也是见到过的第一个包含cd的action, 可以作为后面研究的案例
