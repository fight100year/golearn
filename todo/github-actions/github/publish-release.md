# elgohr/Github-Release-Action

- 通过action发布一个release
- 这是正式分析的第一个非官方action
- 这是一个功能性的action
- action带有测试

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM ubuntu:18.04
LABEL "com.github.actions.name"="Github Release"
LABEL "com.github.actions.description"="Publish Github releases in an action"
LABEL "com.github.actions.icon"="git-branch"
LABEL "com.github.actions.color"="gray-dark"

LABEL "repository"="https://github.com/elgohr/Github-Release-Action"
LABEL "maintainer"="Lars Gohr"

RUN apt-get update \
  && apt-get install software-properties-common -y --no-install-recommends \
  && add-apt-repository ppa:cpick/hub \
  && apt-get update \
  && apt-get install hub -y --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

- 容器基于ubuntu18.04,安装了一些基本软件
  - software-properties-common 里面包含了add-apt-repository命令
  - 通过add-apt-repository命令来安装hub
  - hub是github推出的一套和github交互的命令行工具,[用法在这儿](https://hub.github.com/hub-release.1.html)

## action 分析

entrypoint.sh的内容如下：

```shell
#!/bin/bash

MESSAGE=$*
hub release create -m ${MESSAGE} $(date +%Y%m%d%H%M%S)
```

- 整个action的code非常简单，就是publish一个release

    hub release create [-dpoc] [-a FILE] [-m MESSAGE|-F FILE] [-t TARGET] TAG

- 可以看出，入参有一个，就是消息(用空白行分割，前面是release的标题，后面是release的描述)
- TAG是用时间来处理的，而且时间精度是秒

## 测试workflow

```yaml
name: Test
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Build the Docker image
      run: docker build .
```

- 测试仅仅测试docker image 是否能构建成功


## 使用

- 这个action是支持老的workflow写法，也支持新的workflow写法
- 官方推荐的写法是yaml

```yaml
name: Publish Release
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Create a Release
      uses: elgohr/Github-Release-Action@master
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      with:
        args: MyReleaseMessage
```

## 总结

- 通过action来发布一个release，是进一步将开发生命周期的活动纳入workflow的标志
- 实现原理也颇为巧妙，通过hub工具来实现
- 可惜，还差那么一点，就是没有将tag打成v1.0.2类似的格式

扩展思路：
- 如果有一个action可以获取release的版本号，那此次分析的action就可以做扩展了
- 每次tag就可以选取v1.1.1或是v.1.2.1.3
- 本系列的任务'action熟悉'完成之后,可以尝试写一下,并给作者推一个pr
