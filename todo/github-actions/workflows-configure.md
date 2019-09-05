# github actions 工作流配置

## 创建一个 workflow

- 前提：有指定仓库的写权限或管理权限
- 就可以创建/查看/编辑 wrokflow
- 一个仓库可以有多个工作流 在仓库的.github/worklflows/下
- workflow至少要有一个job
- job包含多个step，每个step执行一个独立的小任务，step可以是运行命令，或使用一个action
- action可以自己创建，也可以去社区找
- 配置的工作流的启动可以通过github事件触发，通过调度触发，或通过外部事件触发
- 配置文件用yaml格式，并将配置文件保存到仓库里

创建一个配置文件：
- 在.github/workflows/下创建一个.yml文件
- 按workflow语法构建一个配置文件
- 将配置文件提到到仓库的指定分支，如果想这个分支上运行工作流。


```yaml
name: Greet Everyone
# This workflow is triggered on pushes to the repository.
on: [push]

jobs:
  build:
    # Job name is Greeting
    name: Greeting
    # This job runs on Linux
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v1
      # This step uses the hello-world-action stored in this workflow's repository.
      - uses: ./hello-world-action
        with:
          who-to-greet: Octocat
        id: hello
      # This step prints the time.
      - run: echo The time was ${{ steps.hello.outputs.time }}
```

- 通过事件触发工作流(目前有3种：github事件/调度/外部事件)
    - 工作流的名称对应着name:
    - 触发事件对应着on:,push表示任何分支的push操作都会触发这个工作流
    - 调度触发或外部事件触发，可查看[文档](https://help.github.com/en/articles/configuring-a-workflow)
- 也可以指定工作流只对部分分支起作用
- 可选择每个job的虚拟环境(操作系统/工具/包/设置)
    - 一个job的虚拟环境是不变的，job里的action通过文件系统共享信息
    - runs-on: 就可以指定具体的虚拟环境实例，目前可选的有ubuntu/linux/macos
- 配置build矩阵
    - 为了同一时间在多操作系统/多平台/多语言版本下测试通过，可配置一个build矩阵
    - 使用build矩阵，要使用strategy

```yaml
runs-on: ${{ matrix.os }}
strategy:
  matrix:
    os: [ubuntu-14.04, ubuntu-18.04]
    node: [6, 8, 10]

这个build矩阵要测试14和18下node版本为6/8/10 各种情况的组合
```

- 有很多标准的action也是可以直接使用的
    - uses: actions/checkout@v1 就是checkout 标准的action，v1表示使用稳定版本
    - 针对clone，也可以只克隆最新代码

```yaml
- uses: actions/checkout@v1
  with:
    fetch-depth: 1
```

- action的类型
    - docker容器 action
    - javascript action
- 在工作流中引用action
    - 工作流中的action可能来至公共仓库/同一个仓库的其他地方/docker hub的镜像
    - 如果是私有仓库，那workflow file和action要在同一个仓库
    - 不能引用另一个私有仓库的action，即使这个私有仓库和当前仓库属于同一个组织
    - 可以通过git或docker的tag来指定action的版本
    - 引用一个公共仓库的action，写法是{owner}/{repo}@{ref} 或是{owner}/{repo}/{path}@{ref}

```yaml
jobs:
  my_first_job:
    name: My Job Name
      steps:
        - uses: actions/setup-node@v1 引用github.com/actions/setup-node
          with:
            node-version: 10.x
```





## 管理 workflow run

## wrokflow 语法

## 事件

## 虚拟环境

## 虚拟环境中的软件

## 上下文和表达式语法
