# cds-snc/github-actions/auto-commit

- docker action
- 通过action来做commit

## 仓库分析

- readme是常规操作
- 没有license文件 
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM alpine/git:1.0.7

LABEL "name"="auto-commit"
LABEL "maintainer"="Max Neuvians <max.neuvians@cds-snc.ca>"
LABEL "version"="1.0.0"

LABEL "com.github.actions.name"="Auto-commit for GitHub Actions"
LABEL "com.github.actions.description"="Auto-commits and changes back to the branch"
LABEL "com.github.actions.icon"="git"
LABEL "com.github.actions.color"="orange"

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]
```

## action 分析

- docker环境是一个轻量的，alpine

```shell
#!/bin/sh
set -e

sh -c "git config --global user.name '${GITHUB_ACTOR}' \
      && git config --global user.email '${GITHUB_ACTOR}@users.noreply.github.com' \
      && git add -A && git commit -m '$*' --allow-empty \
      && git push -u origin HEAD"
```

- action中利用git命令来做提交
- 因为用户名邮箱写的太死，后面使用会受到影响

## 使用

- 有一个入参，可选，就是提交日志

## 总结

- action自身没有测试workflow
- 暂时没想到会有啥作用
- 而且邮箱配置不够灵活，如果要使用，需要修改
