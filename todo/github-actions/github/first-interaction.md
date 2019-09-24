# actions/first-interaction

当贡献者的第一次创建pr/issues时，发送一条指定消息

## 仓库分析

- ts action ? docker action
  - 分析可得出，这个是docker action
- action没有对应的测试

## action分析

```yaml
# action.yml

name: 'First interaction'
description: 'Get started with Container actions'
author: 'GitHub'
inputs:
  repo-token:
    description: 'Token for the repo. Can be passed in using {{ secrets.GITHUB_TOKEN }}'
    required: true
  issue-message:
    description: 'Comment to post on an individuals first issue'
  pr-message:
    description: 'Comment to post on an individuals first pull request'
runs:
  using: 'docker'
  image: 'Dockerfile'
```

- actions/first-interaction 是一个docker action
- 整个action的动作就是启动一个容器
- 这个action有3个入参：
  - 仓库的token
  - 第一次issue的消息
  - 第一次pr的消息
  - 这些消息都是要进行传递的，具体是发给贡献者还是通知维护者，就看下面的分析

```Dockerfile
FROM node:slim

COPY . .

RUN npm install --production

ENTRYPOINT ["node", "/lib/main.js"]
```

- docker配置里，就是执行一个nodejs程序
- nodejs里是ts程序，所以才会误以为是ts actions
- 因为ts完全看不懂，所以下面看下这个action如何使用，看看能否看清楚这个action的作用

## 使用

```yaml
steps:
- uses: actions/first-interaction@v1
  with:
    repo-token: ${{ secrets.GITHUB_TOKEN }}
    issue-message: '# Message with markdown.\nThis is the message that will be displayed on users' first issue.'
    pr-message: 'Message that will be displayed on users' first pr. Look, a `code block` for markdown.'
```

- 看样子，如果是贡献者第一次issues/pr，就会给用户发一些消息，这是个有用的功能

```yaml
name: Greetings

on: [pull_request, issues]

jobs:
  greeting:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/first-interaction@v1
      with:
        repo-token: ${{ secrets.GITHUB_TOKEN }}
        issue-message: 'thank you very much for your first issue, i will promote this as soon as possiblie(非常感谢您贡献的issue，我会尽快推进处理)'
        pr-message: 'thank you very much for your first pr, i will promote this as soon as possiblie(非常感谢您贡献的pr，我会尽快推进处理)'
```

## 总结

- 这是一个交互性的action，主要针对第一次issues/pr
- 因为不熟悉ts，所以对issues/pr的扩展就不聊太多
- 其次，这个action的写法是docker action，但执行的是TypeScript代码
- 第三，这是第一个用到token的action,未来会遇到更多
