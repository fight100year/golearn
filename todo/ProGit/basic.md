# git 基础

本章记录基础命令,包括仓库初始化/文件跟踪/暂存/提交/撤销/忽略/历史/推送/拉取/差异

- git init 初始化一个仓库
- git add 添加跟踪, 开始跟踪一个文件
- git rm 取消跟踪,如果文件已添加到暂存区,可以用-f参数强制删除
- git commit -m "xx" 提交
- git commit -a 表示已跟踪过的文件自动暂存,一起提交,可以少一步git add步骤
- git clone url 目录, 克隆一个项目
- git支持多种数据传输协议:https://; git://; ssh协议,这几种协议的差异后面讨论
- git status 查看文件状态(未跟踪/未修改/已修改/已暂存)
- 默认分支名是master
- .gitignore 忽略文件
- git diff 查看工作区和暂存区的差异
- git diff --staged 查看暂存区和仓库的差异
- git mv file1 file2 修改名字,不过实际上 git mv = mv + git rm + git add
- git log 查看提交历史
    - git log -p -3 显示每次提交差异,只显示最近两次
    - git log --stat 查看统计信息
    - git log --pretty=oneline 提交数很大时,可将提交放在一行显示
    - git log --pretty=format:"" 指定格式显示
    - 一个变更的作者(将成果完成的人)和提交者(将成果提交到仓库的人)是有区别的
    - git log --graph 形象地显示分支合并历史
    - git log -<n> 显示最近n条提交
    - git log --since --until 按时间筛选
- git reset HEAD file 取消某文件的暂存
- git checkout -- file 取消工作区的修改(和最后一次提交同步)


