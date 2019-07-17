# 历史

- git 存储的是快照,而不是变更列表
- git大部分操作,只需要访问本地资源和文件,所以速度快
- git保证了完整性
- 命令行可以体验所有功能,gui只能体验部分,所以推荐使用命令行
- git config 可以配置git环境
    - .git/config 针对仓库
    - ~/.gitconfig 针对当前用户
    - /etc/gitconfig 针对当前系统用户
    - 从上到下,优先级越低
- git安装后第一件事是配置用户名和邮箱

```
➜  ~ git config --global -l
user.email=1876180681@qq.com
user.name=63isOK
core.editor=vim
```

- git config -l 显示所有配置信息
- git help log 显示git log 命令帮助

