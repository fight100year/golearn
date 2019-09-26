# go的ci

具体地址看[这儿](https://github.com/actions/starter-workflows/blob/master/ci/go.yml)

```yaml
name: Go
on: [push]
jobs:

  # 构建job
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:

    # 安装go环境
    - name: Set up Go 1.12
      uses: actions/setup-go@v1
      with:
        go-version: 1.12
      id: go

    # 克隆代码
    - name: Check out code into the Go module directory
      uses: actions/checkout@v1

    # dep安装依赖，不过现在可以直接使用go module
    - name: Get dependencies
      run: |
        go get -v -t -d ./...
        if [ -f Gopkg.toml ]; then
            curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
            dep ensure
        fi

    # 项目构建
    - name: Build
      run: go build -v .
```

- 基于上面的workflow，是可也做扩展的，ci test / cd等

```yaml
# 扩展之后的ci test

name: ci-test
on: 
  - push:
    paths:
      '*.go'
  - pull_request

jobs:
  build:
    name: Build
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-18.04,windows-2016,macOS-10.14]
        go: ['1.10.x','1.12.9','1.13']
      max-parallel: 3

    steps:
    - name: Set up Go 1.10+
      uses: actions/setup-go@v1
      with:
        go-version: ${{ matrix.go }}
      id: go

    - name: Check out code 
      uses: actions/checkout@v1

    - name: Build
      run: go build -v .

    - name: Test
      run: |
        go test -v
        go test -v ./...
```

`我已经新建了一个模版，用来做ci 测试的`

[地址](https://github.com/fight100year/go-ci-test-workflows-template)
