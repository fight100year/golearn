# 创建一个js action

本节关注于如何使用action开发套件来创建一个 js action

前一节的开发套件介绍，主要熟悉了创建action的几种模版，下面关注于如何diy

## 前期准备

- 安装node.js 12.x 其中已经包含了npm
- 在github创建一个新的仓库
  - 要么从模版仓库中创建
  - 要么手动创建，并添加相关的文件
- clone到本地
- 切换到本地仓库，执行npm init -y,这是为了初始化package.json文件

## 创建一个元信息文件 action.yml

```yaml
name: 'Hello World'
description: 'Greet someone and record the time'
inputs:
  who-to-greet:  # id of input
    description: 'Who to greet'
    required: true
    default: 'World'
outputs:
  time: # id of output
    description: 'The time we greeted you'
runs:
  using: 'node12'
  main: 'index.js'
```

- 这是一个简爱的js actin元信息
- 包含了输入输出，执行信息

## nodejs相关

- 使用开发套件的package
- 写action code

## README

- 这个readme需要详细描述：
  - 描述
  - 输入输出(可选和必选)
  - 安全信息
  - 环境变量
  - 如何使用，一个小例子

## 测试

- 一般会创建一个测试workflow来测试当前action
- 那几个模版中都有测试例子

