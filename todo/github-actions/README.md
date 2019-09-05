# github actions

github actions的定位：
- 在gtihub仓库中自定义'软件开发生命周期'的工作流

## 介绍

- 灵活自定义工作流
- 每一个单独的任务称为一个action(操作)，多个action组合起来，就是自定义的工作流
- 这些工作流是自动执行的，可用于构建/测试/打包/发布/部署
- github actions 也可用于端到端的ci/cd,github actions内部集成了ci功能
- 工作流可运行在linux/macos/windows/容器里，而创建的action可以放在仓库，放在github公共仓库，放在docker容器里
- action也可以和github社区分享
- 工作流可配置成执行某些特殊的事件

## github actions中的核心概念

- workflow 工作流
    - 在仓库中可设置的一个"可配置的自动化流程"
    - 可在github仓库中对任何项目进行编译/测试/打包/发布/部署
    - workflow可分成多个job，由事件调度或激活
- workflow run 工作流运行实例
    - 当配置的事件发生时，是workflow的一个运行实例，
    - 有job实例/action实例/log实例/status实例
- workflow file 配置文件
    - 最新job的工作流配置文件，格式是yaml
    - 这个文件放在github仓库，目录是.github/workflows目录
- job 任务
    - 由step组成的一个任务
    - 每个job都在虚拟环境的新实例中运行
    - 在workflow file中定义job运行的依赖规则
    - 多个job可以并行运行，或者依赖上一个job的状态并按顺序执行
    - eg：一个工作流可以包含两个任务：编译/测试。测试就依赖编译的状态，编译失败，测试也就不用测试了
- step 步骤
    - 一个job会执行多个步骤，每一个步骤都被称为step
    - 一个job的step，她们都在相同的虚拟环境中进行执行
    - 步骤中的action是通过文件系统共享信息的
    - 步骤中可运行命令，也可运行actions
- action 操作
    - 一个单独的小任务，用于组合成步骤，步骤是为了job服务
    - action是workflow中最小的可移植构建块
    - action 可以通过github社区进行分享，可自定义公共的action
    - action必须包裹在step中才能使用
- ci 持续集成
    - 在软件开发中，像共享仓库经常性提交小的代码改动，这一实践称为持续集成
    - 通过github actions，可自定义工作流来自动构建和自动测试代码
    - 在仓库中，可以直接查看代码更改状态和每次action的日志
    - ci通过一个快速的反馈来节省开发者的时间
- cd 持续部署
    - cd是基于ci的
    - ci的测试通过后，代码会自动部署到生产环境(也可以是其他环境)
    - 通过github actions，可以自动部署到任何云/自有服务/平台等
    - cd通过自动部署来节省时间，并更快地为客户提供"经过测试的/稳定的"代码更改
- virtual environment 虚拟环境
    - github提供了linux/macos/windows的虚拟环境，用以执行工作流
- runner 运行者
    - 一个github服务
    - 表示的是一个虚拟环境中的github服务，用于执行可能出现的job
    - 运行者会执行job的action，并报告执行的进度/日志/和最后结果
    - 一个运行者同一时间只能跑一个job
- event 事件
    - 一个具体的活动，用来触发工作流运行实例(workflow run)
    - 这个活动(事件)可能来之github(可以是提交/issue的创建/pr的创建)
    - 事件也可以来之github外部，eg：这里面要用到仓库的webhook
- artifact 输出
    - 构建和测试时生成的文件，eg：二进制或包文件，测试规则，屏幕截图，日志文件
    - artifact和workflow run一般是一起用的
    - artifact 这些文件可以用于另一个job，或是用于部署

理解：
- action好像一个单独的功能函数，里面包含一个独立的小功能(打开本地文件)
- step是一个小业务块，由多个action组合而成(打开abc文件，如果打开失败就提示)
- job像是一个线程，里面放了多个业务块，为外面提供一个业务功能(读取abc，处理，显示结果)
- workflow像一个进程，里面包含多个job，为外面提供一个完整的业务功能(读取abc文件，无论是本地还是网络的)
- workflow run就是具体的进程实例
- workflow file就是配置文件

## workflow run的通知

- github actions 可启用邮件或web通知
- workflow run 运行完之后就会有通知，通知的状态包括成功/失败，中立/取消等
- 当然也可以直接去github仓库页的actions tab查看

## 在github社区获取actions

- 自己分享：在公共仓库开源一些actions
- 找别人的：取github.com/actions找

## github actions 支持的语法

- 支持yaml
- hcl 语法已经明确表示放弃了

## 使用限制

- 超出限制，可能会导致各种错误
- 限制可能未来会改变，目前的限制是：
    - 一个仓库最多20个工作并发
    - 一个仓库所有的actions，请求数不能超过1000条/每小时
    - 工作流中的每个job，最多执行6个小时
    - 一个仓库所有的工作流，最多能有20个job并发
- 还有一些就是不能违反github利益的

## 加入github actions 受限公共版

- 目前github action的功能还在变更中，不推荐将高价值工作迁移到github actions


