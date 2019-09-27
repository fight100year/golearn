# maddox/actions/ssh

- docker action
- 通过ssh登录服务器，并执行一些命令

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM debian:stable-slim

LABEL "maintainer"="maddox <jon@jonmaddox.com>"
LABEL "repository"="https://github.com/maddox/actions"
LABEL "version"="1.0.1"

LABEL "com.github.actions.name"="SSH"
LABEL "com.github.actions.description"="Run command via SSH"
LABEL "com.github.actions.icon"="server"
LABEL "com.github.actions.color"="orange"

RUN apt-get update && apt-get install -y \
  openssh-client && \
  rm -Rf /var/lib/apt/lists/*


ADD entrypoint.sh /entrypoint.sh


ENTRYPOINT ["/entrypoint.sh"]
```

## action 分析

- Dockerfile中定义了使用debian，并安装openssh-client
- 这样通过这个ssh 客户端就可以访问ssh服务了

```shell
#!/bin/sh

set -e

SSH_PATH="$HOME/.ssh"

mkdir -p "$SSH_PATH"
touch "$SSH_PATH/known_hosts"

echo "$PRIVATE_KEY" > "$SSH_PATH/deploy_key"

chmod 700 "$SSH_PATH"
chmod 600 "$SSH_PATH/known_hosts"
chmod 600 "$SSH_PATH/deploy_key"

eval $(ssh-agent)
ssh-add "$SSH_PATH/deploy_key"

ssh-keyscan -t rsa $HOST >> "$SSH_PATH/known_hosts"

ssh -o StrictHostKeyChecking=no -A -tt -p ${PORT:-22} $USER@$HOST "$*"
```

- 前面准备环境，后面用ssh命令来连接服务器，并执行命令
- 必选入参：
  - private_key，ssh的私有密钥
  - host，ssh要连接的主机
  - user，连接ssh的用户名，认证就用上面的私有密钥
  - 要执行的命令
- 可选入参：
  - port，默认是22

## 使用

- 私有密钥/host/user/port都属于安全信息
- 只有要执行的命令属于普通入参
- 安全信息还不太了解，在后面的action遇到了再一起分析

## 总结

- action自身没有测试workflow
- 这个action在cd上会很有作用
