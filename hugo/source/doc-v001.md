# 第二次提交对应的文档

也是第一份初始文档, 同时也是本次代码说对应的site例子

## 站点资源文件的组织结构

```
|-- config.json  // 这是site配置文件
|-- content
|   `-- doc      // content type只有一个:doc, 下面都是一个md内容文件
|       |-- configuration.md
|       |-- contributing.md
|       |-- contributors.md
|       |-- example.md
|       |-- front-matter.md
|       |-- installing.md
|       |-- license.md
|       |-- organization.md
|       |-- release-notes.md
|       |-- roadmap.md
|       |-- shortcodes.md
|       |-- source-directory.md
|       |-- templates.md
|       |-- usage.md
|       `-- variables.md
|-- layouts      // 布局,也可以看成内容的渲染方式,与content独立出来是有好处的
|   |-- chrome   // 具体的布局细节
|   |   |-- footer.html
|   |   |-- header.html
|   |   |-- includes.html
|   |   `-- menu.html
|   |-- doc      // 和content同名,默认使用这个布局去渲染doc下的content
|   |   `-- single.html
|   `-- index.html  // 特殊文件, 这里是首页, 特指一个常规的page,而不是像content下的page资源
`-- public       // 公共资源,被布局资源所引用
    `-- static
        `-- css
            |-- bootstrap-responsive.css
            `-- bootstrap.min.css
```

## 文档内容

### home 首页

- hugo的特点: 生产静态站点速度快/灵活/有趣

### getting started 入门

- 安装
    - 一个可执行,能访问即可
- 用法
    - 从用法参数上基本可以看出下面几个参数是很重要的
    - c 配置文件, p 相对目录, w 实时加载, port 端口, s 生产站点之后,启一个http服务
    - 还有一些是功能上的扩展,或是在配置文件中可以配置的
- 配置
    - 目录结构很重要,模板(layout下的文件),两者决定了site的大部分配置
    - 配置文件更多的是对目录结构和模板的一个补充和细节的扩展
    - hugo有些默认约定,符合最佳实践,没有特别理由,不需要修改:
    - content是源文件目录,也就是md文件的目录,里面存的是site想展示的资源
    - layouts是模板目录,里面更多的是告诉content,按什么样的方式去展示
    - public是生产站点,输出目录
    - tags 默认是分类和tag,具体意义不明,待后面补充,表示content是按分类索引还是标签索引
    - baseurl, url名,和最后的站点生产有关
- 组织结构
    - 目录的组织很重要,hugo的入参就是一个具体目录
    - hugo要用最小的配置,完成最高级的定制,秘密就在于模板
    - 一个site,可能有多个content type. spf13就有博客/演讲/自我介绍/相关项目
    - content可使用两种索引:分类和标签. csdn的博客,分类可能是c++,标签可能是linux和网络
    - content有3种不同的展示方式: 列表/摘要/完整的页面. list可用于导航/摘要是补充

### layout 布局

- 模板
    - go中的html/template是模板引擎,特点是轻量,hugo中使用她,是因为正好合适
    - hugo有5种类型的模板:
        - index.html 必须存在, 渲染之后,就是site的homepage
        - rss.xml 如果有rss文档,这个文件就必须存在
        - indexes 一种特殊文件,用于展示多个content, 博客中按标签分类时,会用到
        - content type, hugo支持多种类型的content, eg: 博客/音乐/等等
        - chrome, 一种简单装饰,对site来说可有可无,更多的作用是方便写模板
    - 模板和content type是对应的,下面看看如何新增一个content type:
        - layouts下创建一个目录,注意:目录名是单数,目录名和content type对应
        - 目录下创建一个single.html 文件,这个是用来渲染一个page的
        - 创建layouts/index/目录名.html 
        - 支持content的多种显示方式:只需在layouts/目录名/创建一个模板即可
- 变量
    - .Title, content的标题
    - .Description, content的描述
    - .Keywords, content的元数据 keyword, 这个对搜索引擎优化有帮助
    - .Date, 发布日期
    - .Indexes, 多索引名
    - .Permalink, page的固定链接
    - .Fuzzyordount, content中的字数,不一定绝对正确
    - .RSSLink, indexes中的rss链接
    - 还有一些其他变量
    
### content 

- 组织结构
    - hugo中在md文件的最前面,用特定标签标记的部分称为 front matter
    - front matter也可以定义一些配置,不过,推荐还是用组织结构来表示配置
    - 因为组织结构和site渲染有对应关系,而且简单明了
- front matter
    - 可以为content添加元数据, 这个版本的元数据格式主要是json,后面支持多种格式
    - hugo预定义了一些变量,用户也可以自定义变量,在模板中通过.Params来访问自定义变量
    - 必须要的变量: Title,Description,Pubdate,Indexes
    - 可选变量: Draft,Type,Slug,Url
    
### 附加
    
- shortcodes
    - 简单的代码块
    - markdown可以内嵌html
    - shortcodes,就是将这些html代码全部封装起来,这样md里只需要关注内容即可
    - hugo用模板去渲染content,如果遇到shortcodes,替换,最后将完整的md丢给黑色星期五(md解析引擎)
    - {{%  %}} 就是shortcode, 空格结束, 第一个是shortcode名,后面跟的是参数,最佳实践是单参数
    - 创建一个shortcode步骤:
        - 在layouts/shortcodes/下创建一个模板,模板名就是shortcode名
        - 模板中获取参数 
            - 按参数位置获取 {{ index .Params 0 }}
            - 按参数名获取 {{ index .Params "class" }}
        - 模板中检查参数有没有传过来
            - {{ if isset .Params "class"}} class="{{ index .Params "class"}}" {{ end }}

    选择md,是因为md简单
    md也有很多事是不支持的,且不能在写md时每次写很多html
    直到后来md支持了嵌入html,但是问题是每个md都需要写很多html代码
    此时,shortcodes解决了这个问题 

### 文档的其他信息

路线图 roadmap(不分先后):
- 分页
- 支持其他顶级page,现在只有homepage,后面还有about等
- 支持系列
- 语法高亮
- 页面的 前一页/后一页
- 相关post
- 支持toml的front matter
- 适当支持yaml的front matter
- 支持其他格式

