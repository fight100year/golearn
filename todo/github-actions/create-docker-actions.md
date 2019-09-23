# 创建一个docker action

## 准备

- 需要在github创建一个仓库
- clone到本地

## 创建一个Dockerfile

```Dockerfile
# Container image that runs your code
FROM alpine:3.10

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
```

## entrypoint.sh

- 这里面就是action code
- 输入输出语法,可以参看[这儿](https://help.github.com/en/articles/development-tools-for-github-actions)
- 需要有可执行权限，chmod +x entrypoint.sh

```shell
#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo ::set-output name=time::$time
```

## 元信息

action.yml

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
    description: 'The time we greeted you'
runs:
  using: 'docker'
  image: 'Dockerfile'
  args:
    - ${{ inputs.who-to-greet }}
```

## 最后

测试和README和js action是一样的
