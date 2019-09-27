# action/labeler

- 根据pr修改的文件路径，给pr打上标签
- github pr的标签，主要作用是为了分类


## 仓库分析

- 一个标准的ts action
- 自身是带有测试的?
  - 咋一看，是有ts测试的，可惜是手工测试
  - 缺少一个自动测试的workflow

## action 分析

```yaml
# action.yml

name: 'Pull Request Labeler'
description: 'Label pull requests by files altered'
author: 'GitHub'
inputs:
  repo-token:
    description: 'The GITHUB_TOKEN secret'
  configuration-path:
    description: 'The path for the label configurations'
    default: '.github/labeler.yml'
runs:
  using: 'node12'
  main: 'lib/main.js'
```

- 一个标准的js action写法
- 有两个入参：
  - token
  - 标签文件的路径(标签文件里记录了源码目录对应的标签，后期pr就是根据这个来匹配)

切换tag到v2版本，就可以看到lib/main.js


## 使用

- 使用之前需要创建一个标签文件，放在.github/labeler.yml,或是其他地方(其他地方就需要使用action时指定路径)
- 标签文件的写法也是有[规则](https://github.com/isaacs/minimatch)

```yaml
# 常规用法

name: "Pull Request Labeler"
on:
- pull-request

jobs:
  triage:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/labeler@v2
      with:
        repo-token: "${{ secrets.GITHUB_TOKEN }}"
```

## 总结

- 缺少了一个action 测试环节
- 使用这个action，需要一个额外的标签文件
- 目前还没有测试这个action，不过在[awesome-actions](https://github.com/sdras/awesome-actions)中还是有不少使用这个action的
- 回头再在使用中分析这个action
