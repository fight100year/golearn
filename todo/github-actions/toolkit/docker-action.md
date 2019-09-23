# docker action

模版在[这里](https://github.com/actions/container-template)

## 仓库分析

- license/readme 是常规操作
- Dockerfile/entrypoint.sh是docker action所要求的
- action.yml是每个action都要有的

```Dockerfile
FROM alpine:3.10

COPY LICENSE README.md /

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

entrypoint.sh如下：

```shell
#!/bin/sh -l

echo "hello $1"
```

- 之前也提到过了，docker action需要修改Dockerfile的ENTRYPOINT，替换成自己的entrypoint.sh
- 现在，这个模版中的action做了一件事，就是echo

```yaml
# action.yml

name: 'Container Action Template'
description: 'Get started with Container actions'
author: 'GitHub'
inputs: 
  myInput:
    description: 'Input to use'
    default: 'world'
runs:
  using: 'docker'
  image: 'Dockerfile'
  args:
    - ${{ inputs.myInput }}
```

- 在action中，就是指定了Dockerfile要用docker来执行
- 输入参数也指定了

## 使用

```yaml
steps:
  - using: actions/setup-node@master
  - using: actions/container-template@master
    with: yb
```

## 总结

- action是一个操作，并不局限于一段编程的代码，或是几个shell命令
- action的本质还是一个操作，这个操作可以用编码来描述，也将操作可以丢在docker中
- 这时就可以看出js action和docker action的区别了：
  - 操作的颗粒度小，就使用js action通过nodejs编码的方式来描述
  - 颗粒度大，就可以使用docker来处理
  - js action的操作和环境是分离的，docker的操作和环境是一致的

总之，action的本质还是"一个操作"
