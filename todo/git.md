# git常用命令

- 查看日志
    - git whatchanged --stat --graph --pretty=oneline 查看日志,顺带显示每次提交的文件变更
    - git log  --graph --pretty=oneline 更好的以图的方式看日志

## git学习

为啥要学git
- 文件防丢失
- 最流行的版本控制工具
- 让合作和协作更加便捷

## git flow

git可以将项目中每一个变动都跟踪起来,git的工作原理是将
每一次的变动都记录下来,存储起来,并提供随时查看.

- git init 初始化git环境(准备好跟踪变动的各种工具),这样就有了一个git项目
    - git项目有3部分:
        - 工作区:我们的工作目录,可以对文件进行增删改和组织
        - 暂存区:对工作区的修改,都可以通过暂存区用列表显示出来
        - 仓库:永久存储这些变更地方
    - git flow的规则:
        - 在工作区编辑文件,在暂存区添加文件,在仓库保存变更
        - git add 将变更文件添加到暂存区
        - git commit 将变更提交到永久仓库,提交日志格式:引号/现在进行时/50字左右
        - git status 查看工作区的变更
        - git diff 查看工作区和暂存区之间的差别
        - git log 查看日志
- git and github
    - git 是版本控制工具,可以存草稿(就是暂存区),一个git管理的项目称一个git仓库
    - github 提供了git仓库的托管服务,通过github可以对文件进行存储/共享/协作
    - 简单点讲,gitub是一个使用git的一套工具.
    - 在本地建立一个仓库后,可以将仓库托管一份到github
    - git remote add origin https://github.com/63isOK/git_practice.git
    - 之后可以用git push -u origin master 将本地仓库同步到远端仓库
- git 对变更进行undo
    - HEAD 一般指最后一次commit, git show HEAD 会显示最后一次提交的信息
    - 如果工作区做了修改,但是我们不想要这部分修改,还原指定文件到最后一次提交:
    - git checkout HEAD filenme 针对工作区
    - 放在暂存区的变更,提交到仓库中时,使用的是同一commit号
    - 上面是工作区的变更undo,下面是unstaged暂存区的变更
    - git reset HEAD filename 针对暂存区
    - 下面是丢弃仓库中的某次提交
    - git reset e49f8d97e   仓库的HEAD会指向e49f8d97e, e4之后的提交不再属于仓库
    - 丢弃仓库中的提交,本质上是历史的倒带,要谨慎使用
- 分支
    - 创建分支的目的之一: 可以在分支版本中进行实验性操作
    - 分支之间是互不影响的,除非在分支合并时才有交集
    - git branch 查看当前分支,带有星号的分支是当前分支
    - git branch 分支名 用于创建一个新分支
    - git chekcout 分支名 切换到分支
    - 分支合并: git merge 分支名, 将指定分支合并到当前分支
    - 分支合并最大的问题就是冲突,在合并的时候会提示哪个文件有冲突,解决之后commit就ok了
    - 分支解决完一个问题后,就结束了,结束的标志就是合并操作,之后就可以删除掉了
    - git branch -d 分支名 删除指定分支
    - 如果特征分支(就是新创建,用于解决特定问题的分支)没有合并就要删除, 使用 -D
- 协作
    - 目标:每个参与者都有自己本地仓库,可以跟踪其他人的变更,可以访问最终版本
    - 实现:用一个远端仓库来和本地仓库同步,这样每个参与者都可以独立去做事
    - 从远端克隆一个仓库: git clone remote_location clone_name
    - 默认,我们clone远端仓库时,远端地址名叫 origin
    - git remote -v 显示远端信息
    - 当其他人更新之后,我们若想自己的本地仓库保持同步,要执行 git fetch
    - git fetch 只是将远端仓库同步到本地的remotes/origin/master,并没有和本地master作合并操作
    - git fetch之后,若想对其他人提交的变更进行查看和修改,就需要先合并到本地分支
    - 协作5步:
        - 同步远端分支到本地分支 fetch/merge 
        - 从本地分支创建一个新分支,用于开发新特征
        - 开发新特征,并提交工作(最后合并到本地分支,删除新建的分支)
        - 重要: 再次同步远端分支到本地分支(为了防止冲突,和svn一个道理)
        - 将本地分支同步到远端分支 push
        - git push origin 分支名 将本地分支同步到远端仓库

## flow

目前有3种flow: git flow/ github flow/ gitlab flow,
目的都是为了让多人协作得更加流畅,使项目井井有条地发展下去.

上面3种工作流的共同点是:fdd,功能驱动式开发.
即:需求是开发的起点,先有需求,后有功能分支(特征分支),
功能开发完成,分支合并到主干,然后删除分支.

github flow,是git flow的简化版,适合"持续发布":
- 只有master一个长期分支
- 不区分功能(特征)分支和补丁(bug)分支
- master分支的更新和产品发布是一致的,master就是线上稳定的代码
- 有些功能开发后,并不需要立马发布,这就需要一个production分支跟踪线上版本

github flow 以部署为主,git flow以发布的release为主
