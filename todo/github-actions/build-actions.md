# 构建action

- 一个可用的action，要包含元文件和action代码文件
- 元文件用于定义输入输出，叫action.yml，(不是必需的，有些action的输入比较简单，可直接通过环境变量来搞定)
- action代码文件，主要定义action的运行环境和执行命令

## action 的说明

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
        - 因为容器的创建和检索，docker容器action要比js action慢一点
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

## 元文件语法

上面也提到过了，在仓库中创建一个action来执行某个任务，需要一个元文件

docker容器型action和js型的action都需要一个元文件，action.yml,
这个文件定义了输入/输出/和action的入口，格式是yaml

语法：
- name：action的名称，用于github显示，标识不同的actions
- author：可选，action的作者
- description： 一个简短的说明
- inputs：可选，输入参数，这个输入参数会以环境变量的形式转入，会被转成小写
    - 可选参数，一般会有一个默认值
    - 在workflow file中使用这个action，需要用with关键字来指定输入参数
    - 对于action的入参，对应的环境变量是 INPUT\_变量名,自动转大写，空格用\_代替
    - input.input_id: 必选，字符串，用于标识入参
        - description：入参的说明
        - required： 是否是必选参数，false表示可选参数
        - default： 默认值，这个可以不指定 
- ouputs: 可选，输出参数
    - 输出参数可以给工作流后续的action使用
    - 自动转小写
    - output_id: 必选，字符串，用于标识出参
    - description： 说明，必选
- runs： 必选，action要执行的命令
    - using：指明，用来执行main中code的程序，eg：docker，node
    - env: 可选，一个kv对，会设置成虚拟环境中的环境变量 
    - main：对于js action，就是action code，配合using使用
    - image：对于docker容器action，配额和using使用
        - 这个镜像可以是docker hub的镜像，仓库里的Dockerfile文件，其他注册中心的镜像
    - entrypoint：docker容器action用于覆盖Dockerfile的ENTRYPOINT
        - 如果action中未指定runs关键字，那entrypoint指定的命令就会执行
    - args：docker容器action，这个参数表示入参数组，这里面都是一些硬编码
        - 这些参数会传给容器的ENTRYPOINT
        - args就是用于替换Dockerfile中的CMD指令
            - README中的必选参数，可在CMD中忽略
            - 没有指定args，就使用默认值
            - 如果action暴露了类似 --help标记，在文档中使用默认值
- branding：品牌
    - 可将创建的action作为一系列品牌发布出去
    - 可以指定颜色和图标

## 创建一个js action

- 创建一个action，需要使用一个官方的套件，在actions/tookkit

## 创建一个docker action

## github actions 开发工具

利用github actions的node.js开发套件来创建js action非常快

- 使用node.js工具套件

还有使用日志命令 具体可查看文档
