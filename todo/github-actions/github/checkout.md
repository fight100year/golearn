# action/checkout

- checkout a git repository
- 将仓库代码克隆到$GITHUB_WORKSPACE，这样workflow就能访问仓库内容了
- 可以指定版本，这个action的相当于git fetch 和 git checkout $GTIHUB_SHA

## 仓库分析

- license/readme是常规操作
- 剩下的action.yml是元文件，也是action必须要有的文件

```yaml
# action.yml

name: 'Checkout'
description: 'Checkout a Git repository.'   # 描述，也是这个action的功能
inputs:
  repository:
    description: 'Repository name'
  ref:
    description: 'Ref to checkout (SHA, branch, tag)'
  token:
    description: 'Access token for clone repository'
  clean:
    description: 'If true, execute `execute git clean -ffdx && git reset --hard HEAD` before fetching'
    default: true
  submodules:
    description: 'Whether to recursively clone submodules; defaults to false'
  lfs:
    description: 'Whether to download Git-LFS files; defaults to false'
  fetch-depth:
    description: 'The depth of commits to ask Git to fetch; defaults to no limit'
  path:
    description: 'Optional path to check out source code'
runs:
  # Plugins live on the runner and are only available to a certain set of first party actions.
  plugin: 'checkout'
```

## action 分析

- 上面是action的整个源文件
- 运行调用的是github action服务中的runner 运行实例中的插件：checkout
- 具体runner中的checkout插件是如何做的，我们不用太关心(无非是调用git命令，并参考输入参数)
- 入参分析：
  - 仓库名
  - 引用，提交号/分支/tag
  - github 访问token
  - clean 在执行chekcout action之前，是否将目录清理干净
  - 是否递归克隆子项目
  - 是否下载lfs文件(大文件)
  - fetch深度，默认是全部，在大项目中可以设置为1
  - 指定源码目录，可选
- 用法分析
  - 需要仓库，一般在源码编译上用的比较多

## 使用

```yaml
# 这是一个典型的下载源码/安装编译环境/测试的过程
# 简单点说，如果workflow触发是push，那这个用法就是ci 测试
# 而且是最常规的用法

steps:
- uses: actions/checkout@master
- uses: actions/setup-node@master
  with:
    node-version: 10.x 
- run: npm install
- run: npm test
```

```yaml
# 再用的比较多的就是，克隆不同的分支，这也是第二使用场景
# 特别是触发workflow的事件是其他分支或tag

- uses: actions/checkout@master
  with:
    ref: some-branch
```
