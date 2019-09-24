# 安装go环境

- 仓库在[这里](https://github.com/actions/setup-go)
- 这是一个ts action
- action自身是带有测试的

action功能是安装go环境，附带了两个选项：
- 对指定版本的go，进行下载或是缓存，并添加到PATH
- 在错误输出中注册一个匹配器，用于匹配指定问题

## 仓库分析

- 仓库和ts action基本类似

```yaml
# action.yml

name: 'Setup Go environment'
description: 'Setup a Go environment and add it to the PATH, additionally providing proxy support'
author: 'GitHub'
inputs:
  go-version:
    description: 'The Go version to download (if necessary) and use. Example: 1.9.3'
    default: '1.10'
# Deprecated option, do not use. Will not be supported after October 1, 2019
  version:
    description: 'Deprecated. Use go-version instead. Will not be supported after October 1, 2019'
    deprecationMessage: 'The version property will not be supported after October 1, 2019. Use go-version instead'
runs:
  using: 'node12'
  main: 'lib/setup-go.js'
```

- 描述：安装一个go环境，并添加到PATH，也可以提供代理支持
- 输入参数是 go-version，默认是1.10版本
- 后面调用的是setup-go.js，这个是TypeScript写的，大致意思就是安装go
- ts 具体就不分析了，也就不分析github action开发套件的package了

## 使用

```yaml
# 常规用法，安装指定go版本
# 常与action/checkout配合，在go语言的ci 测试中使用

steps:
- uses: actions/checkout@master
- uses: actions/setup-go@v1
  with:
    go-version: '1.9.3' # The Go version to download (if necessary) and use.
- run: go run hello.go

# 矩阵测试
# 这里是测试代码在不同go版本中的测试
# 也可以配合做跨平台的测试

jobs:
  build:
    runs-on: ubuntu-16.04
    strategy:
      matrix:
        go: [ '1.8', '1.9.3', '1.10.x' ]
    name: Go ${{ matrix.go }} sample
    steps:
      - uses: actions/checkout@master
      - name: Setup go
        uses: actions/setup-go@v1
        with:
          go-version: ${{ matrix.go }}
      - run: go run hello.go
```

## action 的测试

- 常规的ts action测试

## 总结

- 配合checkout，常用在ci test
- 配合矩阵构建，可测试跨平台，或测试不同go版本之间的支持
