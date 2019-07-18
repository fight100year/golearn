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

远程仓库的管理:
- git remote 查看已配置的远程仓库,每个仓库会带一个简写,origin是clone时的默认缩写
- git remote -v 显示远程仓库缩写,及对应的读写url
- 在协作中,一般会拥有多个远程仓库
- git remote add 简写 url 添加远程仓库
- git fetch 简写 从远程仓库拉取数据(将会拥有远程仓库所有分支的引用,此时并未合并)
- git push origin master 推送到远程仓库,前提是先拉取合并(push之前若有其他人先push,那本次push会失败)
- git remote show 简写  查看某一远程仓库的详细信息
- git remote rename 老简写 新简写 远程仓库简写的重命名
- git remote rm 简写 在配置中移除一个远程仓库

标签:
- 给历史的某次提交打上标签,以示重要. 最常用的用法是标记发布节点
- git tag 列出已有标签
- git tag -l 'xxx\*' 列出xxx开头的标签
- git tag -a v0.8.0 -m '0.8.0版本' 创建附注标签, -a创建 -m注释
    - git tag -a v1.0.0 -m '1.0版本' 提交号  为指定提交号创建标签
- git show 标签名 查看标签信息和对应的提交信息
- git tag v0.9.0 创建轻量标签 不带-a -s -m参数
- git push origin v0.8.0 显示将标签推送到远程仓库
    - git push origin --tags 将所有标签推送到远程仓库
- git tag -d 标签名 删除本地标签
- git push 简写 :refs/tags/标签名 删除远程仓库的标签
- git checkout 标签名 切换到标签指向的文件版本, 会导致detacthed head
- git checkout -b 分支名 标签名 从指定标签创建一个新分支

别名:
- 在git config文件中配置,主要是配置一些命令的简写

