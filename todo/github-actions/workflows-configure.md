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
    - 引用同一个仓库中的action，写法是{owner}/{repo}@{ref} 或是./path/to/dir
    - 引用docker hub中的容器，写法是docker://{image}:{tag}

```yaml
引用公共仓库的action
jobs:
  my_first_job:
    name: My Job Name
      steps:
        - uses: actions/setup-node@v1 引用github.com/actions/setup-node v1就是tag
          with:
            node-version: 10.x

引用同一仓库的action
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      # This step checks out a copy of your repository.
      - uses: actions/checkout@v1  引用公共仓库的
      # This step references the directory that contains the action.
      - uses: ./.github/actions/hello-world-action  引用同一仓库的action

引用docker hub的容器
jobs:
  my_first_job:
    steps:
      - name: My first step
        uses: docker://alpine:3.8

```

## 管理 workflow run

我们做以下的事：
- 查看workflow中每个step的状态和结果
- 取消阻塞的workflow
- 调试或重新运行一个失败的workflow
- 查找或下载日志
- 下载输出

github actions利用github的api，让管理工作流(就是上面那些事)成为了可能，
也可以浏览工作流的历史，可以重新运行workflow run，可以取消正在运行的workflow run，
可以查看每个step的详细运行信息，可以在运行信息中进行搜索，以及下载

## wrokflow 语法

说明：
- yaml语法，文件扩展名是.yml或是.yaml
- 配置文件放在.github/workflows/下
- 限制
    - 一个仓库最多20个workflow并发
    - 一个仓库总的api请求，不能超过1000次/小时
    - workflow中的job运行不能超过6小时
    - 一个仓库的所有workflow，不能超过20个job的并发

语法：
- name
    - 工作流的名称
    - 如果省略，就取配置文件名
- on
    - 必选
    - 用string来表示单个事件
        - on: push 单个事件触发工作流
    - 用array来表示多个事件
        - on: [push, pull_request] 多个事件
    - 用map来配置一个工作流调度
        - on.schedule
        - 通过 posix corn 定时任务来触发
    - 用指定文件/指定tag/指定branch来限制工作流的执行
        - on.<push|pull_request>.<tags|branches>
        - tag branch 都可以使用通配符来进行匹配
        - 如果只指定分支或tag中的一个，那其他类型的都不会触发工作流
        - 多个匹配过程可能会发生交集，以最后一个匹配为准
        - on.<push|pull_request>.paths
        - 这个就是通过文件目录来进行匹配
- jobs
    - workflow就是由多个job组成
    - job默认都是并发执行，也可以配置成顺序执行
    - job通过runs-on指定一个新的虚拟环境
    - jobs.job_id 每个job都有一个id关联,job_id以字母和\_开头,包含字母数组-\_
    - jobs.job_id.name job名用于github上显示
    - jobs.job_id.needs 指明这个job的前置job
    - jobs.job_id.runs-on 指定新的虚拟环境，目前(20190906)可用虚拟环境是
        - ubuntu-latest, ubuntu-18.04, or ubuntu-16.04
        - windows-latest, windows-2019, or windows-2016
        - macOS-latest or macOS-10.14
    - jobs.job_id.steps job的步骤叫step
        - step可运行命令，可设置任务，可运行action
        - 多个step之间可共享工作区间和文件系统
        - 但step中修改的环境变量不会影响到其他step，她们都有各自的进程
        - jobs.<job_id>.steps.id 标识step
        - jobs.<job_id>.steps.if 可利用if来设置条件，满足就执行job
        - jobs.<job_id>.steps.name step名称
        - jobs.<job_id>.steps.uses 选择使用哪些action，action是可重用的最小单元
        - 推荐：在使用action时带上git 引用或sha或docker的tag
            - 使用提交的sha，最安全可靠
            - 使用action的版本，有可能是分支名或tag，在一定范围内是可以进行修复，维护需要注意兼容性
            - 使用master分支是最方便的，如果发布了一个主版本，就可能导致工作流失败
        - 部分action是需要with关键字的，是否需要，需要查看action的README
        - action要么是js脚本，要么是docker容器,docker容器要运行在linux环境中
        - jobs.job_id.steps.run 操作系统shell执行的命令
            - 执行shell会在虚拟环境中的新进程内执行
            - 可以通过shell指定命令执行的shell，可以是bash sh cmd python等
        - jobs.<job_id>.steps.with 参数，是一个map，kv对
            - 这些参数是作为环境变量传进去的
            - 这些参数对应的环境变量是加上前缀 INPUT_， 并转换成大写
        - jobs.<job_id>.steps.with.args 参数，这个是给docker容器的参数
        - jobs.<job_id>.steps.with.entrypoint 覆盖docker容器的entrypoint
        - jobs.<job_id>.steps.env 设置虚拟环境的环境变量
        - jobs.<job_id>.steps.working-directory 设置job的工作目录
        - jobs.<job_id>.steps.continue-on-error 设置为true表示step失败后可以继续执行后面的step
        - jobs.<job_id>.steps.timeout-minutes job运行的最大时长，超过就kill
    - jobs.job_id.timeout-minutes 工作流运行的最大时长 默认是360=6小时
    - jobs.job_id.strategy 创建构建矩阵
    - jobs.job_id.strategy.matrix 具体的构建矩阵




## 事件

## 虚拟环境

## 虚拟环境中的软件

## 上下文和表达式语法
