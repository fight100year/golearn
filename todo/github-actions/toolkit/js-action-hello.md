# 一个js action的简单例子

对应的仓库在[这里](https://github.com/actions/hello-world-javascript-action)

## 介绍

- 这个action的作用是打招呼
  - hello world
  - hello xxx
- 这是一个简单，标准的js action，整个构建过程符合"创建js action"
  - 因为"创建 js action"文章会在下一篇分析，所以现在省略掉，先关注我们能看到的

## 源码分析

- 整个仓库来看，该有的全有
  - action.yml 定义action
  - README.md 描述，包括了(作用/输入/输出/usage例子)

```yaml
# action.yml

name: 'Hello World'
description: 'Greet someone and record the time'
inputs:
  who-to-greet:  # id of input
    description: 'Who to greet'
    required: true
    default: 'World'
outputs:
  time: # id of output
    description: 'The time we we greeted you'
runs:
  using: 'node12'
  main: 'index.js'

# 分析
# 每级缩进是2个空格
# 这这个js action太标准，输入输出，还有运行code都指定了
# runs中，是用node12来执行index.js
```

- 仓库中还有一个index.js,就是js action要执行的code

```yaml
const core = require('@actions/core');
const github = require('@actions/github');

try {
  // `who-to-greet` input defined in action metadata file
  const nameToGreet = core.getInput('who-to-greet');
  console.log(`Hello ${nameToGreet}!`);
  const time = (new Date()).toTimeString();
  core.setOutput("time", time);
  // Get the JSON webhook payload for the event that triggered the workflow
  const payload = JSON.stringify(github.context.payload, undefined, 2)
  console.log(`The event payload: ${payload}`);
} catch (error) {
  core.setFailed(error.message);
}

# 这就是node要运行的js代码
# 从代码中也看到，已经引用了@actions/core和@actions/github两个package
```

- 现在看到的东西太少，还是得看完js action的构建流程再回头来理解这个例子
