# ci: 持续集成

可以在github仓库，使用github action创建自定义持续集成工作流

## 持续集成

- ci是一种软件实践，频繁提交代码到共享仓库的实践
- ci的好处是快速反馈，出错了立马就知道，知道了就可以立马调试
- 频繁提交，步子小，所以合并方便，方便支持多人协作
- ci的目的是让开发人员专注于代码开发，减少调试错误和解决合并冲突
- ci里指定测试，会让提交马上得到反馈，
    - 代码静态检测
    - 安全检测
    - 代码覆盖率
    - 函数的单元测试
    - 自定义检测
- ci中的构建和测试都需要服务器，提交到github之前肯定需要在本地服务器上运行通过
    - 以前有travis-ci做测试，现在使用github action可以实现更多的功能

## github actions 中的ci

- github actions的工作流虽然不局限于ci/cd，但ci是一个刚需
- github提供了仓库代码的构建服务，提供了虚拟环境用于执行测试
- workflows run 用的是云服务的虚拟环境，会clone仓库
- 可以配置ci工作流，让github事件出现时，触发ci工作流,也可以通过调度或外部事件来触发
- 在github actions中，除了ci，workflow可以在整个软件开发的声明周期中都用得上
- workflow 支持多种语言的ci，[具体可以看](https://github.com/actions/starter-workflows/tree/master/ci)

## workflow 在ci中的一些概念和术语

- workflow
    - 可配置的自动化流程
    - 用于build/test/package/release/deploy
    - workflow由多个job组成
    - 由事件激活
- workflow run
    - 一个流程的运行实例
    - 已经预先配置好了事件触发
    - 可以看到每个运行实例的job/actions/logs/statues
- workflow file
    - yaml文件
    - 术语workflow的配置文件，里面最少配置了一个job
    - 这个配置文件在 .github/workflows
- job
    - 一个job由多个steps(步骤)组成
    - 每个job都有一个干净的虚拟环境来运行
    - job都是并发，也可以配置依赖规则，来让job顺序执行
- step
    - 步骤
    - 每个step都在同一个虚拟环境下执行
    - 通过文件系统进行信息共享
    - step(步骤)包括commands 或是 actions
- action
    - 组合成steps，steps组合成job，job组合成workflow
    - action是workflow中最小的可移植构建块
    - 可以自己编写action，也可以从github.com/actions等社区获取
    - action只能出现在steps中
- ci
    - continuous integration 持续集成
    - 频繁提交到共享仓库的一种实践
    - 好处是快速反馈和合并方便
- cd
    - continuous deployment 持续部署
    - 基于持续集成
    - 好处是让变更更快的呈现给客户
- virtual environment
    - 虚拟环境
    - github提供的
- runner
    - github提供的一种服务
    - 在每个虚拟环境中都有
    - runner就是等待job，选择job之后就执行，一次只能执行一个job
    - 她用于执行job的action，向github反馈进度/日志和最后的结果
- event
    - 触发workflow运行的事件
- artifact
    - build/test时的输出
    - 二进制/包文件/测试结果/屏幕截图/日志文件
    - 输出是作为下一个job的输入，或是用于部署

## 通知

web和邮件

## 状态徽章

- 用于表明当前工作流是成功还是失败
- 在README.md中添加，一般在master中显示

## github actions 中设置ci


