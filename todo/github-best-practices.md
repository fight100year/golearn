# github 10大最佳实践

- 基于分支模型,master分支不允许直接提交,好处是保证master时刻都是可部署的版本
- commit时使用正确的邮件地址,好处是跟踪方便
- 启用code owners特征,将一部分人自动添加到reviewer列表中
- 源码中不要包含敏感信息(这里指安全信息:密码等),安全信息最好放在环境变量
- 源码中不要提交依赖,用其他方式(其他工具)来处理依赖
- 源码中不要提交配置文件
- 使用一个定义明确的git 忽略文件
- 不再维护的仓库,应该归档,变成只读
- 有些依赖定义文件,会定义依赖库的版本,最好指明具体版本,而不是latest
- 可以的话,使用一个package,尽量再保持主版本号不变的同时,多试试不同的小版本号

基于第三点:
- CODEOWNERS 文件,定了这个仓库代码的所有者/所有team
- 仓库的管理员或owner有权限添加CODEOWNERS文件
- 在文件中定义的owner,需要有对仓库的写权限
- 这么做的好处是:开启一个pr时,owner会自动添加review中,草稿pr会不自动添加
- 文件中定义的owner,可以批准pr
- 文件位置可以是根目录 docs/ .github/ 任何分支都可以添加
- 一个分支中的文件中定义了一个分支的所有者,所以可以给不同的分支定义不同的所有者,就是每个分支都搞这么一个CODEOWNERS

CODEOWNERS文件的语法:
- @username 或者 user@qq.com 前者是github的@系统,后者是邮箱,可以引用一个人
- @org/team-name 引用一个team
- 文件后面的规则会覆盖文件前面的规则

    @global-abc @global-bcd 开启一个pr时,abc/bcd会收到review请求
    顺序很重要,最后那个优先级最高

    .js @js-owner 如果pr只修改了js文件,那么只有js-owner会收到review请求
    globa不会收到review请求

    .go abc@qq.com 也可以通过邮箱指定owner

    /log/xx @abc abc负责review /log/xx下的所有文件

    docs/* @abc abc负责review docs下的所有文件(不管abc是不是在根目录), 不嵌套进入子目录

    docs/ @abc abc负责docs下的所有文件,包括子目录内的

    /docs/ @abc abc负责/docs下的所有文件的review


    

