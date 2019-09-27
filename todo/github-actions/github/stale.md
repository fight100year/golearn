# action/checkout

- 将久无进展的issues/pr进行标记，再过一段时间，就会关闭

## 仓库分析

- 一个标准的ts action
- action缺少测试 workflow

```yaml
# action.yml

name: 'Close Stale Issues'
description: 'Action to close stale issues'
author: 'GitHub'
inputs:
  repo-token:
    description: 'Token for the repo. Can be passed in using {{ secrets.GITHUB_TOKEN }}'
    required: true
  stale-issue-message:
    description: 'The message to post on the issue when tagging it. If none provided, will not mark issues stale.'
  stale-pr-message:
    description: 'The message to post on the pr when tagging it. If none provided, will not mark pull requests stale.'
  days-before-stale:
    description: 'The number of days old an issue can be before marking it stale'
    default: 60
  days-before-close:
    description: 'The number of days to wait to close an issue or pull request after it being marked stale'
    default: 7
  stale-issue-label:
    description: 'The label to apply when an issue is stale'
    default: 'Stale'
  exempt-issue-label:
    description: 'The label to apply when an issue is exempt from being marked stale'
  stale-pr-label:
    description: 'The label to apply when a pull request is stale'
    default: 'Stale'
  exempt-pr-label:
    description: 'The label to apply when a pull request is exempt from being marked stale'
  operations-per-run:
    description: 'The maximum number of operations per run, used to control rate limiting'
    default: 30
runs:
  using: 'node12'
  main: 'lib/main.js'
```

## action 分析

- 上面是action的整个源文件
- 执行的是标准的ts源码
- 入参分析：
  - token
  - 标记issues时，发送的信息，没有这个，action不会标记issues
  - 标记pr时，发送的信息，没有这个，action不会标记pr
  - 标记时长，定义老旧具体有多长，默认60天
  - 关闭时长，标记之后，多长时间close，默认是7天
  - 标记issues的标签
  - 达到条件后，未被标记的issues，打上的标签,也叫豁免标签
  - 标记pr的标签
  - 达到条件后，未被标记的pr，打上的标签
  - 每次运行最多操作的数量，默认30个
- 因为是ts action，所以具体细节不分析了，主要看下用法

## 使用

```yaml
# 常规用法
# 调度事件触发workflow
# 每个小时的整点，开始触发workflow
# cron，语法是5个数[分/时/天/月/周]，其中还有多值/范围/步长的表示，最后取5数集合

name: "Close stale issues"
on:
  schedule:
  - cron: "0 * * * *"

jobs:
  stale:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/stale@v1
      with:
        repo-token: ${{ secrets.GITHUB_TOKEN }}
        stale-issue-message: 'Message to comment on stale issues. If none provided, will not mark issues stale'
        stale-pr-message: 'Message to comment on stale PRs. If none provided, will not mark PRs stale'

# 第二场景，自定义标记和关闭时间(30天标记，5天关闭)
# 也是每小时的整点触发
# 只处理issues，不处理pr

name: "Close stale issues"
on:
  schedule:
  - cron: "0 * * * *"

jobs:
  stale:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/stale@v1
      with:
        repo-token: ${{ secrets.GITHUB_TOKEN }}
        stale-issue-message: 'This issue is stale because it has been open 30 days with no activity. Remove stale label or comment or this will be closed in 5 days'
        days-before-stale: 30
        days-before-close: 5

# 第三场景，自定义标记标签，以及豁免标签

name: "Close stale issues"
on:
  schedule:
  - cron: "0 * * * *"

jobs:
  stale:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/stale@v1
      with:
        repo-token: ${{ secrets.GITHUB_TOKEN }}
        stale-issue-message: 'Stale issue message'
        stale-pr-message: 'Stale issue message'
        stale-issue-label: 'no-issue-activity'
        exempt-issue-label: 'awaiting-approval'
        stale-pr-label: 'no-pr-activity'
        exempt-pr-label: 'awaiting-approval'
```

## 总结

- action 缺少了一个测试workflow
- 上面3种场景的定制，一个比一个高，也更加适合项目的多样性选择
- 这个action是我见到过的第一个调度性action，也是由调度事件触发的action，之前看到的都是常规github事件触发

