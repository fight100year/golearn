# 构建action

- 可以编写自定义代码来创建actions
- 这个action可以以你喜欢的方式来和仓库交互
    - 通过整合github api或其他公开可用的第三方api来实现
    - eg: 一个action可以发布npm，在紧急issues创建时发送短信，或发布可用于生产的代码
- 可以自己写action，也可以直接使用github.com/actions下的
- 要创建action，需要提供一个元文件，叫action.yml
    - 这个元文件包含输入/输出和主要入口
- action的类型有两种：
    - docker容器actions
        - 利用github actions 代码来打包环境
        - 这样的好处是更加一致和可靠，使用者也不用担心工具和依赖的不一致
        - docker容器的actions，要放在github的linux环境中运行
        - 在容器中可指定操作系统版本/依赖/工具/代码，而action运行正需要指定这些信息，简直是绝配
        - 因为容器的创建和检索，docker容器action要必js action慢一点
    - js actions 
        - 这类的action是可以直接在github提供的虚拟环境中直接运行的
        - action 代码和运行这个代码的环境是分开的
        - 好处也是有的：简单一点，执行也比docker容器action快一点

创建action：
- 推荐使用一个单独的仓库来存放action，而不是将action和应用程序代码放一起
- 单独仓库存放的好处是：可以打版本，可以跟踪，可以发布
- 另外一个好处是分享方面，其他人如果对这个action感兴趣，不代表她想看到应用程序代码
- 当然，如果非要将action和应用程序代码放一块，action放在.github/actions/xxx

action的版本：
- 版本可以用sha/分支/tag来标识
- 当然，github推荐使用更加有语义的版本，这样使用体验也会好很多
    - 使用v1.0.9来发布一个release
    - 用主版本v1 v2等tag来指向当前发布的release
    - 对于破坏性更改，放在新的主版本tag中，eg：接口参数个数变化

action的readme：
- 为了方便认识，最好添加readme，里面应包含：
- action做了什么，详细描述
- 必选的输入输出参数
- 可选的输入输出参数
- 和安全相关的设置
- 使用了哪些环境变量
- 一个如何在工作流中使用action的例子



