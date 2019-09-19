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

# 这个build矩阵要测试14和18下node版本为6/8/10 各种情况的组合
```

- 有很多官方的action也是可以直接使用的
    - uses: actions/checkout@v1 就是github官方维护的github.com/actions/checkout仓库，v1表示使用v1版本
    - 针对clone，也可以只克隆最新代码

```yaml
- uses: actions/checkout@v1
  with:
    fetch-depth: 1    # 只拉取最新代码
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
# 引用公共仓库的action
jobs:
  my_first_job:
    name: My Job Name
      steps:
        - uses: actions/setup-node@v1 # 引用github.com/actions/setup-node v1就是tag
          with:
            node-version: 10.x

# 引用同一仓库的action
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      # This step checks out a copy of your repository.
      - uses: actions/checkout@v1                   # 引用公共仓库的
      # This step references the directory that contains the action.
      - uses: ./.github/actions/hello-world-action  # 引用同一仓库的action

# 引用docker hub的容器
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
        - 构建矩阵的每个条件都可以对应多个版本等，这样会创建多个job
        - 可以用include来指定某个条件下的新约束
        - 可以用exclude来排除某个条件下的约束
        - 设置fail-fast 为true表示如果某个job出错，会取消所有矩阵对应的job，默认是true
        - max-parallel 可指定矩阵运行的job的最大并发数，github 会默认执行所有的矩阵
    - jobs.job_id.container
        - 在job中一个容器包含任意个step
        - 如果指定了，这个job的所有step都跑在这个容器中
          - 如果step也指定了容器,同时还有js脚本action，那么step指定的容器和job的容器是兄弟容器，一起运行
          - 并这些容器共享网络和存储
        - 如果不设置容器，那所有的step都运行在虚拟环境
        - 容器可指定相关参数，如果不指定，可省略image关键字
        - image：docker镜像，生成的容器用于运行action
        - env：容器的环境变量，可用数组传递
        - ports：容器暴露端口
        - volumes：容器存储
        - options：附加选项
    - jobs.job_id.services
        - 工作流中一个job的附加容器
        - eg：创建db服务，创建缓存redis服务，虚拟环境实例自动管理网络和这些服务容器的生命周期
        - 这种服务的ip是动态变化的
        - 端口由service_name.ports指定
        - image： docker镜像
        - env： 环境变量
        - ports: 就是上面说到的端口
        - volumens：存储
        - options：附加选项

## 事件

- 事件，可配置什么事件发生时触发什么样的workflow
- 当然除了github事件，还有调度和外部事件都能触发wrokflow，这些都可以算成事件

限制：
- workflow中的action并不能触发另一个新的workflow

webhook事件：
- 可利用一个或多个webhook事件触发workflow
- eg: on:push 或是 on:[push,pull_request]
- 在一个workflow run中，GITHUB_SHA, GITHUB_REF都作为环境变量存在虚拟环境中
- github事件有多种触发类型
    - release，当有一个release版本执行发布/下架/创建/编辑/删除/预发布时触发
- 也有部分事件只有一个触发类型
    - 这时，action关键字会直接使用webhook payload
    - eg: commit_commnet 只在提交comment时触发事件,push事件也是类似的
    - eg：多触发类型： pull_request 里面有分配 取消分配 标签 取消标签 打开 编辑 repoen 加锁等多个触发类型
- 对于每一个事件，可以设置一个触发类型，如果不设置，就会有一个默认设置
    - 具体的默认设置可以查看[传送们](https://help.github.com/en/articles/events-that-trigger-workflows)
    - pull_request事件中，默认只打开了少数几个触发类型
- 目前github提供的api v3 有50个rust api,对应到github actions的事件有28个
- fork 一个仓库，并提交pr
    - 只会触发上游仓库的pull_request事件
- 由于一些GITHUB_TOKEN机制，如果像让pr触发ci测试，要开启监听push事件

调度事件：
- 调度工作流基于默认分支或上游分支的最后一次提交
- 使用的是posix corn语法
- on.schedule.corn
    - corn有5个由空格间隔的部分
    - 5个部分分别是：分0-59/时0-23/每月的第几号1-31/月1-12/每周第几天0-6
    - 这5个部分的交集就是指定时间
    - \*表示匹配任何值
    - ,表示当前字段是一个数组，可选多个值
    - - 表示是一个范围值
    - / 步长 字段分上的20/15 表示20 35 50 是条件

外部事件：
- 可以使用github 的api来触发一个webhook事件(repository_dispatch)
- 这是在github外部触发工作流的事件
- 如果要触发这个外部事件，需要用POST请求等，具体可查看文档

## 虚拟环境

- 这个虚拟环境由github提供，里面包含了一些工具/包/可设置的action
- 这些虚拟环境跑在微软的azure(一个云服务)，具体是Standard_DS2_v2机器，毕竟github给微软收购了
- windows和linux跑在Standard_DS2_v2虚拟机，macos跑在macstadium
- 可以直接跑在虚拟环境中，或泡在docker容器中
- 每个job都放在一个独立的虚拟环境实例中
    - job中的step都在这个虚拟环境实例中运行
    - 所有action就能利用文件系统共享书信
    - 当然，也可以配置成job都跑在同一个环境中
- 每个虚拟环境可用资源是 双核cpu/7G内存/14G ssd磁盘空间

直接跑在虚拟环境中，部分文件系统路径：
- docker容器运行的action，是在/github目录下，而脚本的的目录并不是静态的
- github actions中可配置 home/workspace/workflow目录
    - home： 用户数据，环境变量对应HOME - workspace： actions执行目录，环境变量对应 GITHUB_WORKSAPCE
    - workflow/event.json： 触发工作流的webhook事件的POST payload，环境变量对应GITHUB_EVENT_PATH

跑在docker容器，部分文件系统路径：
- /github/home  /github/workspace  /github/workflow
- 上面是默认路径，github推荐不要修改默认的环境变量 

关于环境变量：
- workflow run中的每个step都可以访问环境变量
- 一般在配置文件 workflow file中指定环境变量
- 可以通过jobs.job_id.steps.env 来为一个spte指定环境变量
- github会自动将环境变量转成大写，所以在配置中写环境变量时，大小写均是支持的
- 推荐在actions中使用环境变量来访问文件系统，而不是硬编码(直接将文件路径卸载actins中)
- 虚拟环境已经提供了一些默认的环境变量，具体可查看文档
- 自定义的环境变量来表示文件路径，推荐使用\_PATH作为后缀

安全：
- 加密，环境变量也可以进行加密来保证安全
- 这些加密信息只能用于github actions
- 这些安全加密的信息要在github 仓库页去设置，然后通过配置文件，通过环境变量传递给action
- 具体的安全信息的配置和使用，可查看文档

结束代码和状态：
- 可以用返回码来表示action的执行状态
    - 0 表示成功
    - 非0 表示失败

## 虚拟环境中的软件

这里列出来虚拟环境中可用的一些软件,针对于ubuntu16 18/ windows server 2016 2019/ macOS10.14

## github actions 上下文和表达式语法

表达式：
- 可以通过表示来设置workflow file中的值，并访问上下文
- 表达式可以是 文字值 上下文引用 函数 进行组合，组合的方式是通过操作符
- 配置中是可以使用if语句的，if语句后使用表达式，是表达式最常使用场景
- 使用if 后接表达式， 此时表达式不用使用$, 因为此时github会自动当成表达式去计算
- 表达式第二个场景是用在环境变量设置时 env_var:${{<表达式>}}

上下文：
- 上下文是访问信息的一种方式,这些信息包括：工作流实例，虚拟环境，job，step
- 书写方式是： ${{<上下文>}},和表达式有点类似
- 具体上下文对象有哪些：
    - github： 表示wrokflow run 就是工作流运行实例
    - job：当前执行的job对象
    - steps： 当前运行job中包含的step
    - runner： 当前job运行实例
    - secrets： 开启访问github 安全设置
    - strategy： 策略参数，包括快速失败 最大并发 最大job数 job索引
    - matrix： 矩阵 
- 每个对象都有很多属性可以访问

字面量：
- boolean： 值有true和false
- null：值是null
- number： json支持的数值格式
- string：单引号包围的字符串，里面如果含有单引号，是需要转以的
- 这些字面量可以直接用于组合表达式

还有一些操作符，具体可以查看文档：
- 松散式比较
    - 类型不同就转换成成number再比较
    - NaN和NaN的比较结果不一定是true
    - 比较字符串时，不区分大小写
    - 对象和数组只有是在都是相同实例才算是相等
    - 字面来转number具体看文档

函数：
- contains 字符串包含
- startsWith 字符串前缀比较
- endsWith 比较后缀
- format 字符串替换
- join 数组转字符串
- toJson 转json
- job状态检查函数：
    - success 是否成功
    - always 强制认为条件是成功的，就是直接返回true
    - cancelled 是否取消了
    - failure 是否失败了

对象过滤：支持通配符
