# action类型

类型有两种:
- js action
  - 运行在github的host机器上
  - 运行的code和环境是分离的
- docker action
  - 容器包含了运行的code和环境，以及依赖

这两种action都可以访问"工作区间"和"github事件的payload和上下文"

## 何时选择docker action

- 要运行的code和环境都在docker里
- 好处：
  - docker带来更好的一致性，每次运行的环境都是确定的
  - 更好的可靠性，每次action运行时，不用担心工具集和依赖问题
- 限制：
  - 目前只能用在linux上

## 何时选择host action

这里为什么用host action：
- host action表示action运行在github host机器上(而不是host机器上的容器中)
- host action后面会包含多种写法
  - 大类是js action，其中也包括TypeScript action
  - 说不定后面还有其他种类的actin

js action不像docker action将运行code和环境绑在一起
- js action的运行code和环境是分离的
- js action直接运行在github的host机器或虚拟机上

```yaml
# 一个基于构建矩阵的workflow
# 非常适合用js action，因为运行code(action)和环境有多种组合

on: push

jobs:
  build:
    strategy: 
      matrix:
        node: [8.x, 10.x]
        os: [ubuntu-16.04, windows-2019]
    runs-on: ${{matrix.os}}
    actions:
    - uses: actions/setup-node@master
      with:
        version: ${{matrix.node}}
    - run: | 
        npm install
    - run: |
        npm test
    - uses: actions/custom-action@master
```

- js action，只要不超出host机器的运行范围，都可以很方便地使用
  - host机器支持的工具集，就是js action很好的基石，就像上面例子中的setup-node,就是在host机器上安装node
  - 官方文档有支持的[工具集](https://help.github.com/en/articles/software-in-virtual-environments-for-github-actions)
  - 如果action中需要用到哪些工具，可以像上面使用setup-\* 的action来安装，也可以直接通过命令安装(就如上例的npm)
