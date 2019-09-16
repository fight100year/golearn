# the tag uri scheme

- tag 的uri方案
- 这个还属于草稿状态，并不是已发布的标准协议
- 2005年提出的，rfc号是4151

## 介绍

说明：
- tag uris一般被设计用于，人类进行某项全局跟踪活动事的唯一标签
- 不同于其他uris，tag uris没有认证机制
- tag uris只是一个纯实体标签
- 对于非http可访问资源来说，使用tag uris必http ruis要好用一些

设计需求:
- 标识在时间/空间都是唯一的，这是uri的要求
- 标识要让人类易于 创建/阅读/归类/记忆等
- 无需集中注册，创建标识的成本要小
- 标识不能依赖特定的解析方案

对于用户而言，可能更关心下面几点：
- 全局唯一
- 可跟踪
- 自己就可以创建，无需和其他人沟通
- 不要暗示我有默认解析机制

## tag 语法和规则

### tag语法和例子

    tagURI = "tag:" taggingEntity ":" specific [ "#" fragment ]

    taggingEntity = authorityName "," date
    authorityName = DNSname / emailAddress
    date = year ["-" month ["-" day]]
    year = 4DIGIT
    month = 2DIGIT
    day = 2DIGIT
    DNSname = DNScomp *( "."  DNScomp ) ; see RFC 1035 [3]
    DNScomp = alphaNum [*(alphaNum /"-") alphaNum]
    emailAddress = 1*(alphaNum /"-"/"."/"_") "@" DNSname
    alphaNum = DIGIT / ALPHA
    specific = *( pchar / "/" / "?" ) ; pchar from RFC 3986 [1]
    fragment = *( pchar / "/" / "?" ) ; same as RFC 3986 [1]
  
- 这是abnf语法，一种语法规范，和之前go中的ebnf来之同一个家族
- taggingEntity是uri的命名空间，为了避免冲突，添加权威名称(邮件或域名)
- 权威名称 推荐小写，因为区分大小写
- 为了简便，权威名称限定于域名或电子邮件

```
  tag:timothy@hpl.hp.com,2001:web/externalHome
  tag:sandro@w3.org,2004-05:Sandro
  tag:my-ids.com,2001-09-15:TimKindberg:presentations:UBath2004-05-19
  tag:blogger.com,1999:blog-555
  tag:yaml.org,2002:int
```

### 创建tag的规则

### tag的解析

### 如何判断tag是否是相同的
