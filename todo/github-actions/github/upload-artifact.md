# action/download-artifact

- 将构建的输出文件进行上传

## 仓库分析

- license/readme是常规操作
- 剩下的action.yml是元文件，也是action必须要有的文件
- 非常常规的一个仓库，这是第三次见到这类action
- 下载是基于download插件，上传是基于publish插件

```yaml
# action.yml

name: 'Upload artifact'
description: 'Publish files as workflow artifacts'
author: 'GitHub'
inputs:
  name:
    description: 'Artifact name'
    required: true
  path:
    description: 'Directory containing files to upload'
    required: true
runs:
  # Plugins live on the runner and are only available to a certain set of first party actions.
  plugin: 'publish'
```

## action 分析

- 通过publish插件来上传输出文件
- 入参分析：
  - 输出文件名字
  - 文件的目录
  - 这两个参数都是必选(下载action的path是可选的)

## 使用

```yaml
# 基础用法
# 将一个文件，上传到workflow

steps:
- uses: actions/checkout@v1

- run: mkdir -p path/to/artifact

- run: echo hello > path/to/artifact/world.txt

- uses: actions/upload-artifact@master
  with:
    name: my-artifact
    path: path/to/artifact
```

文件上传到哪了：
- 上传到某一位置(github页面中点击下载按钮，就可以下载这个文件)
- 也就是说，我们要控制下载哪些东西，都是可以通过这个action来控制的

## 总结

- 和checkout一样，缺少了一个action 测试环节
- 相比下载action，这个action的功能和用法都清晰很多
- 在整个业务流程中，这个action就是控制能在github中下载哪些东西
