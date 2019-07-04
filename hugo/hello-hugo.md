# hugo

- 口号是 构建站点最快的框架
- 开源/流行的静态站点生产工具之一
- 特点是速度快/灵活性高
- 有趣也是一方面

## 快速使用

- hugo安装之后,使用hugo version 查看版本
- 创建一个新的站点(hello) hugo new site hello
- 新站点创建之后,只要完成以下3步就可以跑起来了
    - 下载一个主题 
    - 为站点添加一些内容 content
    - 启动hugo server, hugo server
- 添加主题
    - git init
    - git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
    - 设置主题, echo 'theme = "ananke"' >> config.toml
- 添加站点内容
    - 添加博客分类的文章 hugo new posts/hello.md
- 启动hugo server: hugo server -D
- 构建静态站点: hugo即可,会在public目录生成静态资源

## 安装

源码编译或是直接用包管理器安装

## 基本用法

hugo cli 是一个命令行.

- 查看help文件 hugo help
- 生成静态站点 在站点目录执行 hugo命令,生成的文件在public目录
- 在内容content的front matter中,可设置draft/publishdate/expirydate,
分别是 草稿/发布日期/到期日期,这几个参数都可以影响content的发布状态,
本地部署或是hugo命令参数可以覆盖上面3个参数.
- 实时加载, 一个hugo服务运行时,修改文件,会被hugo捕获.好处是修改文件,不用重启服务
- 实时加载,服务不用重启,页面如果要接收通知,需要添加js脚本(这个脚本使用ws连接hugo服务)
- 关闭实时加载: hugo server --watch=false 或是 hugo server --disableLiveReload,
或是在config配置中指定 disableLiveReload = true
- 部署正确的姿势: 删除public目录 ,执行hugo生成站点
- hugo生成站点时不会主动删除public目录,所以在开发/生产环境,最好使用不同的输出目录
    - 开发环境 hugo server -wDs ~/hello -d dev
    - 生产环境 hugo -s ~/hello  默认输出到public
    - 参数解析:
        - w 表示实时加载
        - D 表示草稿也发布  开发环境正好需要这个,而生产环境,没有特殊要求就不用了
        - s 表示站点的目录
        - d 表示输出目录    开发输出到dev 生产输出到public

## 目录结构

```
root@7bfe83861e15:~/hello# tree .
.
|-- archetypes
|   `-- default.md
|-- config.toml
|-- content
|-- data
|-- layouts
|-- static
`-- themes

```
使用hugo new site创建一个新的站点,结构就如上面输出的一样

- archetypes
    - 每次用hugo new创建一个新的content文件时,文件中默认会包含date title,
    - 这些数据就源自archetypes下的模板,也有地方称这些模板为 文章模板
- assets
    - hugo pipe要处理的文件,都存在assets
    - 这个目录不是默认生产的
    - 目录下的文件,只有带有 .Permalink .RelPermalink 属性的才会被发布到public目录
- config
    - hugo有大量指令,可以在json/yaml/toml文件中配置指令
    - 指令可出现在启动参数/环境变量/配置文件中,写在配置文件中最省事
    - 站点根目录下 config.toml文件配置指令是最佳实践
- content
    - 站点的内容都放在这个目录下
    - 这个目录下会有多个子目录,每个子目录表示一个content分类
    - 子目录, eg: content/blog, content/articles, content/tutorials
- data
    - hugo生产站点时的配置文件, yaml/json/toml格式
- layout
    - .html文件,这些文件表明了content如何渲染展示
- static
    - 静态content,eg:图片/css/js
- themes
    - 主题目录,也称为皮肤

## 配置

- 指定配置文件

    通过启动命令指定
    --config string      config file (default is path/config.yaml|json|toml)
    可同时指定多个配置文件,用逗号分隔
    默认使用根目录下的config文件

- 配置目录

    针对多个配置文件,指定一个配置目录,便于维护
    --cacheDir string    filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/
    每个配置文件都表示一个根对象的配置, eg:params menus languages
    每个子目录都表示一组特定的环境指令配置
    配置文件可以本地化为指定语言
    这个配置目录可以起很大作用,eg:
    有个开发目录存开发环境的配置
    有个集成测试目录存测试环境的配置
    有个生产目录存生产环境的配置
    有个_default目录,存默认配置
    hugo --environment staging 表示使用测试环境配置,会使用staging和_default下的配置

`一般开发环境使用hugo server部署,生产环境使用hugo即可`

- hugo定义的变量,()里是默认值
    - archetypeDir (“archetypes”) archetypes目录
    - assetDir (“assets”) assets目录
    - baseURL 站点目录名,一般是一个网址
    - blackfriday 是一个markdown渲染引擎 go md库,名字还是黑色星期五
    - buildDrafts (false) 构建站点时,默认不发布草稿
    - buildExpired (false) 默认不发布过期文章
    - buildFuture (false)  默认不发布未来文章(就是设置的发布日期还未到的文章)
    - caches 更好的缓存设置
    - canonifyURLs (false) 默认不将相对路径转换成绝对路径
    - contentDir (“content”) content目录
    - dataDir (“data”) data目录
    - defaultContentLanguage (“en”) content的默认语言是英语
    - defaultContentLanguageInSubdir (false) 默认不使用子目录表示语言
    - disableAliases (false) 默认不禁止使用别名
    - disableHugoGeneratorInject (false) 默认hugo会注入一些数据,仅限于首页html头中注入meta tag 
    - disableKinds ([]) 禁用某个分类的所有页面
    - disableLiveReload (false) 默认启用实时加载
    - disablePathToLower (false) 默认将url转换成小写
    - enableEmoji (false) 默认不启用表情符号
    - enableGitInfo (false) 默认不启用git信息
    - enableInlineShortcodes 是否启用shortcode,shortcode就是将重复的html封装了,减少重复代码
    - enableMissingTranslationPlaceholders (false) 缺少转换时,默认显示空字符串或默认值,而不是占位符
    - enableRobotsTXT (false) 默认不生成robots.txt
    - frontmatter 用来配置元信息的,放在文件头
    - footnoteAnchorPrefix (“”)  默认脚注锚点前缀是空
    - footnoteReturnLinkContents (“”)  脚注返回链接上显示的文字默认是空
    - googleAnalytics (“”) google 分析跟踪id
    - hasCJKLanguage (false) 默认不使用cjk检测,cjk是中日韩语言自动检测
    - imaging 图片处理配置
    - languages 语言配置
    - languageCode (“”) 配置站点的语言code
    - languageName (“”) 配置站点的语言名
    - disableLanguages 不启用某种语言
    - layoutDir (“layouts”) layout目录
    - log (false) 默认不启用日志
    - logFile (“”) 设置日志路径
    - menu 菜单设置
    - metaDataFormat (“toml”) Front matter meta-data format 只有3种:toml/yaml/json
    - newContentEditor (“”) 创建一个新content文章时,使用的编辑器
    - noChmod (false) 默认同步文件的权限模式
    - noTimes (false) 默认同步文件的修改时间
    - paginate (10)   默认分页数量是10
    - paginatePath (“page”) 分页时的分页元素是page
    - permalinks 固定链接 
    - pluralizeListTitles (true) 列表中启用多标题
    - publishDir (“public”) 发布目录
    - pygmentsCodeFencesGuessSyntax (false) 未指定语言的情况下,根据代码来猜测语法
    - pygmentsStyle (“monokai”) 语法高亮的主题  也就是颜色主题
    - pygmentsUseClasses (false) 默认不使用额外的css来补充颜色主题
    - related hugo有默认的,可以微调
    - relativeURLs (false) 默认不启用,相对路径是相对于content根目录说的,这个设置不影响绝对路径
    - refLinksErrorLevel (“ERROR”) 默认ref/relref错误时的日志打印级别是error
    - refLinksNotFoundURL ref/relref找不到指定page时,显示的占位符
    - rssLimit (unlimited) 默认不限制rss订阅数量
    - sectionPagesMenu (“”)  菜单相关
    - sitemap 
    - staticDir (“static”) static目录
    - summaryLength (70) .Summary文件最大字数长度是70 也就是摘要
    - taxonomies 用户自定义分类
    - theme (“”) 使用哪个主题 也就是哪个皮肤
    - themesDir (“themes”) 主题目录
    - timeout (10000) 生产page content的超时时间 单位毫秒,也就是10s超时, 这是为了防止递归生产
    - title (“”) 站点标题
    - uglyURLs (false) 默认不启用 丑陋的url(带后缀,就是被认为是丑)
    - verbose (false) 默认不启用详细输出
    - verboseLog (false) 默认不启用详细日志记录
    - watch (false) 默认不启用实时加载

命令行查看配置: hugo config | grep watch

## 配置环境变量

HUGO_NUMWORKERMULTIPLIER 设置并行处理的数量,默认是逻辑cpu数

## 配置查找的顺序

- config.toml
- config.yaml
- config.json

## 环境变量来做配置

env HUGO_ + 大写的配置key = 值 hugo

env表示后面接环境变量  表现形式是k=v ,后面接 hugo命令

## 渲染时忽略部分文件

直接在配置中添加 ignoreFiles = [ "\\.foo$", "\\.boo$" ] 忽略foo和boo结尾的文件

## front matter

frontmatter
: 元信息,一般在文件开头,用3个短横线开头和结尾,中间的就是元信息

- 配置日期
    - 就是配置在page content中如何访问date
    - 方法就是在config.toml中添加frontmatter节 (toml和ini文件的节 键 值 类似)
    - 配置中,日期可能有多个候选者,也可以自己加日期候选,每次使用日期都会选用第一个可以匹配上的
    - 日期有很多:当前日期 发布日期 最后修改日期 过期日期等
    - 候选者不带冒号前缀的,就是普通候选者,一般都是hugo内部定义的一些日期别名
    - 带冒号的候选,更多的是一种特殊的日期处理,这些候选都不是从frontmatter(content的元信息)中取到的
    - :fileModTime 从content文件最后修改的时间戳中取日期
    - :filename 从content文件名中取日期
    - :git 从content在git上最后一次提交中取日期,这个需要开启 enableGitInfo=true配置
- 配置 Blackfriday
    - 前面也说了,这是hugo实现makrdown解析的引擎
    - 如果要配置,也在blackfriday节
    - blackfriday的选项蛮多, 暂时不列出.
    - blackfriday的扩展有很多是有用的: 大部分支持和github的md写法类似,下面是需要提一下的:
    - 前后都被两个波浪号包围的文字,显示将删除(是不显示还是显示删除线: ~~~123~)
    - "[^1]" 脚注写法,就和论文的脚注类似,都是对某句话的进一步说明
    - 和github的md写法类似, 块上面都需要插入一个空行
    - 可以通过"{#id}"指定header id, 标题的header id是自动生成的
    - 行位的反斜杠会自动转换成换行符
    - 定义一个术语: 术语名单独一行,下一行用冒号开头,后跟术语定义
    - 默认会删除新行,并加入一行. 可避免过多的空白行

## hugo 输出其他格式

默认就是一个静态站点,如果要输出其他格式,就需要配置

## 配置文件缓存

cacheDir可配置缓存目录,cache节还可以具体配置缓存细节

## 主题 theme

- 将重复的提取出来,减少重复工作,提出来的就是主题
- hugo没有默认主题,需要通过命令行下载主题
- 下载全部主题 git clone --depth 1 --recursive https://github.com/gohugoio/hugoThemes.git themes
- 使用主题,可通过配置;可通过启动参数 -t,指定主题名即可
- 可以配置多个主题,组合起来使用,也可以嵌套,也可以自定义主题

## content

- 站点的生成,也就是hugo,需要保持可伸缩性/可管理性,那么除了front matter和一系列模板,还需要其他扩展
- hugo的受众除了开发人员,还有content管理者,作者.从这3个角度才能看到完整的hugo

### content的组织

content organization
: 一般假设,content 源文件的结构和站点渲染的结构类似

- content源文件的组织方式,应和站点渲染一一映射
- content目录可以有目录,想windows下的文件系统一样,可以任意嵌套
- 一般,content目录下的子目录,用于content类型的分类, 也决定了布局

下面看看组织结构和站点渲染映射关系:
- _index.md
    - hugo中的一个特殊文件,可以添加front matter和content到list模板
    - 这些list模板可以是 section模板/Taxonomy模板/Taxonomy术语模板/首页模板
    - permalink = baseurl + section(url)
    - 对于_index.md,最后在站点渲染时,转换成index.html
- section中的单个page
    - 单页模板会将page渲染成站点的页
    - 链接是 section(url) + slug(一般取至md文件的文件名)

组织结构对应站点渲染过程,涉及的路径术语:

section
: content的类型由section决定,最后会出现在渲染的路径中,不能在front matter中配置

slug
: 要么是 "文件名.后缀", 要么是 "文件名/", 由content文件名指定;也可由front matter配置

path
: content的路径由sectin的路径决定, 是完整的baseUrl + section, 不包含slug

url
: 相对路径, 基于content的路径,或是在front matter配置中指定

如何在front matter中,重新指定路径:
- filename,这个就是文件名,不带后缀,无法在front matter中定义
- slug,可以在front matter中重新定义, 最后会在站点渲染的路径中体现出来
- section, 不能再front matter中定义
- type, 默认也是由section决定,但和section不一样的是:type可以在front matter中定义
- url, 一个完整的url,可以被重新定义,但也要基于baseURL,此时会忽略 --uglyURLs设置

type是非常有用的,修改type,可指定不同的layout模板,这样就可以随意修改md文件渲染的方式.

### page导航 最佳实践

- content的组织是利用了page导航的
- page导航就是将一些page资源进行了分组

page导航本质上是一个导航树,共分两类:
- leaf bundle    叶子 没有子节点了
- branch bundle  枝干 一般是首页/section/taxonomy术语/taxonomy列表

| | leaf | branch |
|--|--|--|
|用法| 单页内容和附件的集合 | section页的附件集合|
|索引文件名| index.md | _index.md|
|可使用的资源| page页面和非page资源(pdf/图片等)| 只能是非page资源|
|可以放在哪| 任何叶子导航目录下都可以|只能在枝干导航目录,也就是说要和_index.md在同级目录|
|layout分类| single|list|
|嵌套|不能嵌套|可以嵌套(枝干下可以是枝干,也可以是叶子)|
|例子|content/posts/2019/index.md|content/posts/_index.md|
|非索引的其他content文件|只能被当成page资源访问|只能被当成常规page访问|

leaf bundle 叶子节点:
- 在content目录下
- 包含一个index.md文件
- 对叶子导航节点来说,深度是无所谓的,只要叶子不是在另一个叶子下就行
- 有一种特殊的叶子导航节点,叫headless bundle
    - 并不会被发布
    - 没有Permalink, 也不会渲染成html到输出目录
    - 不是.Site.RegularPages, leaf bundle本就不是常规page
    - 可通过.Site.GetPage获取
    - 一个leaf要指定为headless,只需在front matter中指定 headless = true即可
    - headless应用场景: 共享媒体库/重用页面内容的片段

branch bundle 枝干导航节点:
- 在content目录下
- 至少包含一个 _index.md文件
- content目录下也可以有一个 _index.md文件

### hugo支持的content格式

.md是最主要的一个, 其次其他的也会支持,eg: .maark .org .html .htm

hugo中使用黑色星期五来做md文件的解析引擎,目前版本,tab要是4个空格,如果是2个会出错

blackfriday的选项可参考https://gohugo.io/content-management/formats/#blackfriday-extensions

blackfriday的扩展可以查看文章上面的章节

blackfriday支持的md相对于标准来说,做了如下扩展:
- 任务列表,类似github的 todo list, 这个功能是默认启用的
- 支持表情,在content内容中直接使用表情,这个功能需要在配置中启用
- shortcode,类似代码片段,不过这个是原始的html代码片段,好处是扩展md更便捷,hugo重要特性之一
- code block代码块,类似github的md,也支持代码块的高亮,用3个反勾号即可,高亮引擎默认使用chroma,也可以使用其他的

blackfriday功能上的扩展:
- mmark,黑色星期五自己实现的md, 感觉没必要
- 使用mmark,需要在front matter中指明:   markup: mmark
- mathjax,一个js库,可以直接显示数学公式,使用简单,但也有些限制,具体可以查看文档

hugo对外部工具的支持: 可查看文档

[md链接](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

### front matter

前面一直提到front matter 指的是元数据, 也就是文件头中,用三个短横线包裹的内容,
下面具体介绍一下front matter

front matter可以添加在:
- 配置文件config
- content文件

关于front matter:
- 翻译: 前端的事项, 是不是指文件头部指定的信息
- 作用是:可以在content类型的实例中附加元数据,并嵌入到实例中
- 是hugo的一个重要的能力组件
- 在yaml格式的配置中,用---作为识别代码
- 在toml格式的配置中,用+++作为识别代码
- 在json格式的配置中,用{新行}作为识别代码
 
front matter 中的变量:
- 预定义变量,在使用时用 .变量名,还分大小写
    - aliases,发布时,生产的别名,一个或多个,一般用于有修改content文件在渲染时改名的场景
    - audio,一组和page相关的,音频文件路径,被opengraph内部模板使用
    - date,page的日期,一般从front matter的date中获取,可配置
    - description, content的描述
    - draft, 是否发布草稿.当然,如果启动参数指明不发布草稿,那这个设置就不起作用
    - expiryDate,过期时间,过期后hugo就不发布这个content,也受启动参数影响
    - headless, 是否是特殊的leaf bundle
    - images, 一组和page相关的,图片文件路径,被内部模板使用
    - isCJKLanguage, 是否是中日韩语言, 会导致字数统计有差异
    - keywords, 关键字
    - layout, 布局选择器,和k8s的选择类似,只要选择器的名字一样,就使用指定的布局来套用当前content
    - lastmod, content最后修改时间
    - linkTitle, 创建一个链接,指向content,
    - markup, 实现性功能, 指明当前content的格式. 一般不使用,因为用的主要是md
    - outputs, 输出格式,种类多多
    - publishDate, 发布日期,受启动参数影响
    - resources, 用于配置页面导航资源
    - series, 一系列page属于哪个系列,可以理解为分类,被opengraph内部模板使用
    - slug, 输出url的尾部, 看是否需要修改url中的文件名部分
    - summary, 文本, 也就是摘要
    - title, content的标题
    - type, content的类型,自动从目录(section)中继承而来,不能被front matter指定,但可以使用
    - url, 完整的url路径,当然是渲染之后,可访问的
    - videos,一组和page相关的,视频文件路径,被opengraph内部模板使用
    - weight, 加载显示的优先级, 值越小,优先级越高,越优先显示
    - <taxonomies>, 分类, 可有多个,hugo内置了一些,用户也可以自定义分类,两者工作原理不一样
- 用户自定义变量,在使用时用 .Params.变量名

front matter中的weight变量,可用于content的排序,分类中也能用到 

### page 资源

page资源,一般包括图片,其他page,文档,还有和page相关的url和元素数据等

- 属性
    - ResourceType 资源主类型 eg:图片, 图片下还有png 和jpg等
    - Name 资源名,默认就是文件名,可在front matter中设置
    - Title 值和资源名一样,可在front matter中设置
    - Permalink 绝对url, page类型的资源是没有这个值的, page的固定链接是拼出来的
    - RelPermalink 相对url,
    - Content 资源内容,大部分资源都是将内容放到字符串里返回
    - MediaType 资源的mime类型,eg: image/jpg
    - MediaType.MainType 资源的主mime类型
    - MediaType.SubType 子mime类型
    - MediaType.Suffixes mime类型中的后缀
- 方法
    - ByType 根据类型获取page资源
    - Match  根据资源名匹配结果来获取page资源,匹配区分大小写
    - GetMatch 只返回第一个匹配的page资源
- 元数据
    - page资源的元数据由page的front matter和叫resources的参数管理
    - name 可指定Name
    - title 可指定Title
    - params 可自定义k-v对
    - 占位符 :counter 可用在name/title中, 设置之后,Name和Title才会各自取各自的值

### image处理

image page资源可以 resized and cropped, 重置大小/裁剪

页面导航中,获取所有image


```
{{ with .Resources.ByType "image" }}
{{ end }}
``` 

下面的image处理方法,并不适用于/static目录下的图片:
- Resize 重新定义宽高,可能做了裁剪,可能做了空白填充
- Fit 缩放到指定宽高
- Fill 等比缩放到指定宽高,可能会有裁剪

```
{{ $image := $resource.Resize "600x" }}
{{ $image := $resource.Resize "x400" }}
{{ $image := $resource.Resize "600x400" }}
{{ $image := $resource.Fit "600x400" }}
{{ $image := $resource.Fill "600x400" }}
```

目前hugo还不支持image的exif信息,因为go下的库还不支持

处理方法中的一些选项:
- JPEG Quality, jpeg质量,范围1-100, 默认75
- Rotate, 旋转
- Anchor, 坐标点, 只用在等比填充,指明以哪儿为初始点
- Resample Filter, 重采样filter

最佳实践,将image处理的放在shortcode中

config中配置image处理,在imaging节.就算不指定,hugo也有默认处理

### shortcode

代码段,在content中使用,或在自定义模板中使用,一般里面是html数据

- 在hugo中,md是偏爱的,特别适合content格式,不足之处就衍生出了shortcode
- md + shortcode + template 保证3个方面都可以独立发展,在生产站点时由hugo来保证衔接
- 在content中,调用shortcode的方式是 {{% shortcodename params %}}
- 参数由空格分隔,参数包含空格,就需要用引号包裹
- 其中的%和<>是一样的,都是终结符,和mysql存储过程的结束符是一个意思

```
html123
content
html456

当html123和html456被抽象成shortcode abc之后,
content中的写法应该是 {{% abc %}} md内容 {{% /abc %}}

为了照顾到html, 包含shortcode的写法也是成对的,
当然,如果shortcode里的html是完整的,那就不用成对出现了.
``` 

shortcode 可以嵌套, 在子shortcode中获取父shortcode,可用.Parent

hugo提供了一些内置的shortcode:
- figure, 处理image
    - src, 要显示imge的url
    - link, 要显示image的超链接,目的地址
    - target, 可选,如果link指定了,这个设置就有意义
    - rel, 可选,link指定了才有意义
    - 等等
- gist, 适用于blog,适用于在content中引用代码片段
- highlight, 代码高亮
- param, 在front matter中查找当前page参数的值
- ref/relref, 通过相对路径或逻辑名,或固定链接或相对固定链接找page

### content 相关

怎么知道page中包含哪些相关content, 在content的front matter中标记着

两个page相关:
- 共享相同的数据
- 共享keyword参数

```
// 页面集合中相关方法是.RegularPages
// $related := .Site.RegularPages.Related . 返回和当前页相关的集合
{{ $related := .Site.RegularPages.Related . | first 5 }}  
{{ with $related }}
<h3>See Also</h3>
<ul>
	{{ range . }}
	<li><a href="{{ .RelPermalink }}">{{ .Title }}</a></li>
	{{ end }}
</ul>
{{ end }}
```

fornt matter中定义related:
- 顶层配置
    - threshlod, 0-100, 值越小,匹配的越多,相关新越小
    - includeNewer, 
- 每个索引的配置,可选

### content sections

section
: 基于content 组织结构, 位于content/目录下, 一些page的集合

- content目录下的第一层子目录, 都形成了各自的section, 这些都称为顶层的节
- 如果要定义一个层数更深的section,那么指定目录下至少要有一个 _index.md文件
- section 只能被组织结构定义, front matter 元数据参数无法定义section
- section可以嵌套,除了顶层section,其他section的目录下保证有 _index.md文件即可
- section 就像一棵树, 可用于导航, 最底层的section至少需要一个content文件(eg: _index.md)

有了内容,再加上模板,就可以构成样式各异的站点,所以,下面常用到  section + 模板,
这里的section就是内容,一般指层section,如果子section要指定某个模板,
需要通过front matter 元数据中的type 或是 layout指定

### content type

section和type比较类似, 默认两者取值是一样的,如果在section下指定了archetype,那两者就不一样了

section更多的是站点渲染后的分类, content type更多是源码组织结构

content type:
- 可以有唯一一个元数据集合(eg: 元数据就是front matter)
- 可以自定义模板
- 可以通过hugo new命令创建,期间可以带archetypes

一个content
- 可能是一个相片/一篇博客,
- 每一种,都带有不同的元数据集合和不同的可视化渲染方式

hugo中的约定:
- content type和section是对应的, 一个类型用一种配置,
- 如果所有的content放在一起,就需要为每个content配置,显然很麻烦
- 用目录来表示content type,比较简洁
- 也可以在front matter中用type指定content type,都可以

新建一个content type:
- 在content下创建一个子目录, eg: content/events/
- 在events下的content文件的front matter如下配置即可

```
type = "event"
layout = "birthday"
```

所以新建一个content type 非常方便, 渲染指定birthday layout

新建一个layout type:
- 在/layouts下创建一个event目录即可
- content type创建时, 目录是events, type是event,目录是复数,带s
- layout type创建时,目录是单数,不能是复数,也就是不能带s

hugo中的views:
- content渲染有多种方式
- 比较特殊的有: 显示section列表的单页面,还要显示摘要
- 这时就需要在layout type目录下,有一个模板来附加views

自定义content type的模板查找顺序:
- layouts/event/birthday.html
- layouts/event/single.html
- layouts/events/single.html
- layouts/_default/single.html

如果要创建一个archetype,就要在/archetypes/下创建

### archetypes

- 是一种模板,创建新content时会用到
- 预先配置了front matter, 使用hugo new 时就会用到

hugo new posts/hello.md, 模板查找顺序:
- archetypes/posts.md
- archetypes/default.md
- themes/my-theme/archetypes/posts.md
- themes/my-theme/archetypes/default.md
- 顺序是先找本地,再找主题,先找content type完全匹配的,再找缺省的






