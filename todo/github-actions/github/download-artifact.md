# action/download-artifact

- 将构建的输出文件下载下来

## 仓库分析

- license/readme是常规操作
- 剩下的action.yml是元文件，也是action必须要有的文件
- 非常常规的一个仓库，这是第二次见到这类action
- 第一次是actions/checkout,这个是第二次，都是基于github 服务中的功能插件
- 这类action和js action/ts action/docker action是有些许不同的

```yaml
# action.yml

name: 'Download artifact'
description: 'Download workflow artifacts'
author: 'GitHub'
inputs:
  name:
    description: 'Artifact name'
    required: true
  path:
    description: 'Destination path'
runs:
  # Plugins live on the runner and are only available to a certain set of first party actions.
  plugin: 'download'
```

## action 分析

- 通过download插件来下载输出文件
- 入参分析：
  - 输出文件名字
  - 要下载的目的目录

## 使用

```yaml
# 基础用法
# 没有指定目录，表示上传当前目录
# 上面那句是官方说法，不过这里有个疑问
# 这是一个下载action，难道是通过一个可选参数来决定上传或者下载吗？
# 其次，如果是上传，是上传当前目录下的my-artifact文件吗？上传到哪儿？
# 所以，要么是我理解还不够，要么是文档描述错误，导致我有了疑问
# 这个问题可以在后面接触的cd action中，应该可以得到解答

steps:
- uses: actions/checkout@master

- uses: actions/download-artifact@master
  with:
    name: my-artifact

- run: cat my-artifact

# 第二使用场景
# 带有path，表示下载输出文件

steps:
- uses: actions/checkout@master

- uses: actions/download-artifact@master
  with:
    name: my-artifact
    path: path/to/artifact

- run: cat path/to/artifact
```

## 总结

- 和checkout一样，缺少了一个action 测试环节
- 这个action的功能描述很具体：下载输出文件
- 但是文档中关于用法的描述，很让人疑惑，这是第一个让人不明用法的action
- 在后续的 ci/cd流程中继续关注这个action的相关问题
