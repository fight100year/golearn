# 基础结构

## 缩进

- 块式，都是用缩进来表示结构的
- 缩进就是行首有0个或多个空格
- 为了更好的移植性，缩进中不能有tab
- 块式，一个块的构造遇到什么才结束：如果下一行的缩进必块的缩进小，就表示块构造结束
- node的缩进要比她的父node的缩进要更进一步;兄弟node的缩进是一样的;node的内容是独立缩进的

```yaml
# 注释行前面的空格不是缩进也不是内容
# 流式 风格前的空格不是缩进也不是内容

··# Leading comment line spaces are
···# neither content nor indentation.
····
Not indented:
·By one space: |
····By four
······spaces
·Flow style: [    # Leading spaces
···By two,        # in flow style
··Also by two,    # are neither
··→Still by two   # content nor
····]             # indentation.

%YAML 1.2
- - -
!!map {
  ? !!str "Not indented"
  : !!map {
      ? !!str "By one space"
      : !!str "By four\n  spaces\n",
      ? !!str "Flow style"
      : !!seq [
          !!str "By two",
          !!str "Also by two",
          !!str "Still by two",
        ]
    }
}
```

- -?: 分别用于数组条目，map的key和value分割
- 块式中，这三个也被认为是缩进的一部分

```yaml
# yaml中的写法非常灵活
# 下面这种写法实际情况应该很少有
# 下面的结构这样解读：map[string]array
# array里第一个元素是字符串b，第二个元素是数组，数组里有c和d

?·a
:·-→b
··-··-→c
·····-·d

%YAML 1.2
---
!!map {
  ? !!str "a"
  : !!seq [
    !!str "b",
    !!seq [ !!str "c", !!str "d" ]
  ],
}
```

## 行内分割

- 除了缩进和字面量内容，一行中的token用空白字符(空格和tab)来分割

```yaml
# 空白字符用于行内分割，不术语node的内容

-·foo:→·bar
- -·baz
  -→baz

%YAML 1.2
---
!!seq [
  !!map {
    ? !!str "foo" : !!str "bar",
  },
  !!seq [ !!str "baz", !!str "baz" ],
]
```

## 行前缀

- 在字面量内容里，每行都会有一个行前缀，这个前缀不属于内容
- 一般这个前缀是缩进，可能是空格，也可能是tab

```yaml
# 字面量 
# 流式 plain风格，会去掉前缀
#      引用风格，只去掉前缀，其他属于内容
# 块式 折叠风格，只去掉前缀，其他属于内容
plain: text
··lines
quoted: "text
··→lines"
block: |
··text
···→lines

%YAML 1.2
---
!!map {
  ? !!str "plain"
  : !!str "text lines",
  ? !!str "quoted"
  : !!str "text lines",
  ? !!str "block"
  : !!str "text\n·→lines\n",
}
```

## 空行

- 空行就是行前缀后跟了个换行符
- 这里的空行是值字面量内容里的空行

```yaml
Folding:
  "Empty line
···→
  as a line feed"
Chomping: |
  Clipped empty lines
·

%YAML 1.2
---
!!map {
  ? !!str "Folding"
  : !!str "Empty line\nas a line feed",
  ? !!str "Chomping"
  : !!str "Clipped empty lines\n",
}
```

## 行折叠

- 为了提高长句的可读性，可通过行折叠将一行断成多行
- 换行后面跟了个空行，空行会被省略
  - 连续多个空行的情况：只有第一个空行被省略，剩下作为内容
- 其他情况，换行会被转换成一个空格，来连接两个行内容

```yaml
# 块式折叠字面量
>-
  trimmed↓   # 后面是空行，所以这个换行符会被忽略
··↓          # 第一个空行会被省略
·↓           # 第二个空行作为内容
↓            # 第三个空行作为内容
  as↓        # as前面有一个换行内容，后面的换行替换成空格
  space

%YAML 1.2
---
!!str "trimmed\n\n\nas space"
```

- 前面这个行折叠适合块式折叠风格和流式字面量风格
- 其中也有一些区别，块式折叠风格还会考虑非空行中的空白字符;流式引用风格，会将空白字符进行忽略处理

## 注释

- # 开头就是注释，注释不属于内容
- map的key和value后面都可以有注释，和其他编程语言类似

## 分割行

- 多行注释也能出现在map的key和value中

## 指令

- 指令是告诉yaml处理要怎么做，目前只有YAML和TAG两种指令
- 指令只影响最后字符流的呈现，不影响内容信息
- 指令都是%开头

YAML指令：
- 用于指明yaml文档的版本，目前都是1.2,默认也是1.2
- 如果yaml处理器只能处理1.2版本，来了个1.3会报警，来了个2.0会拒绝(发生高版本情况：小版本号警告，主版本号拒绝)
- yaml1.2是兼容1.1的，来了个1.1也是能处理的
- 不能同时用YAML指定多个版本

TAG指令：
- 创建node tags的一个简便操作(说白了：将tag的公用部分放在TAG指令，这样在创建tag时就简便很多)
- TAG指令格式 TAG tag处理 tag前缀 (tag处理匹配上，就将tag前缀和tag简写组合成一个完整的tag)
- 一个document最多只能有一个相同的tag处理

```yaml
# !yaml! 就是tag处理
%TAG !yaml! tag:yaml.org,2002:
---
!yaml!str "foo"

%YAML 1.2
---
!!str "foo"
```

tag处理和tag前缀是有匹配关系的，yaml有3种tag处理类型：
- 主要处理
  - 主要处理只有一个!，为了更加紧凑
  - 默认前缀匹配的处理是!
  - tag简写用这个，好处是简写的写法和本地tag的写法一致
- 次要处理
  - 次要处理前面带两个!
  - yaml tag仓库用这个
- 基于名字的处理
  - 写法类似于 !名字!
  - 这种写法一般不用再tag简写上，除非TAG指令明确指出

```yaml
# 主要处理
# tag处理是!,所以tag前缀都是匹配的
# 查看顺序，TAG指令对第二个document起作用
# 将tag前缀和tag简写组合成一个完整的tag

# Private
!foo "bar"
...
# Global
%TAG ! tag:example.com,2000:app/
---
!foo "bar"

%YAML 1.2
---
!<!foo> "bar"   # 局部tag，也叫私有tag，前面带一个!
...
---
!<tag:example.com,2000:app/foo> "bar"

# 次要处理

%TAG !! tag:example.com,2000:app/
---
!!int 1 - 3 # Interval, not integer

%YAML 1.2
---
!<tag:example.com,2000:app/int> "1 - 3"

# 基于名字的处理

%TAG !e! tag:example.com,2000:app/
---
!e!foo "bar"

%YAML 1.2
---
!<tag:example.com,2000:app/foo> "bar"
```

tag前缀有两种：
- 本地tag前缀
  - 本地tag前缀前面是有一个!
  - 最后组合而成的tag也是一个本地的
  - 本地tag前缀并不要求是一个有效的uri
- 全局tag前缀
  - 前面没有!
  - 全局tag前缀必须是一个有效的uri

```yaml
# 本地tag前缀
%TAG !m! !my-
--- # Bulb here
!m!light fluorescent
...
%TAG !m! !my-
--- # Color here
!m!light green

%YAML 1.2
---
!<!my-light> "fluorescent"
...
%YAML 1.2
---
!<!my-light> "green"

# 全局tag前缀
%TAG !e! tag:example.com,2000:app/
---
- !e!foo "bar"

%YAML 1.2
---
!<tag:example.com,2000:app/foo> "bar"
```

## Node属性

- 模型上也提到了，Node的属性只有两个：tag和anchor
- 这两个属性可以忽略

```yaml
!!str &a1 "foo":
  !!str bar
&a2 baz : *a1

%YAML 1.2
---
!!map {
  ? &B1 !!str "foo"
  : !!str "bar",
  ? !!str "baz"
  : *B1,
}
```

Node的tag
- tag标识着node的原生数据结构的类型
- tag用!来标识，这个和TAG指令中的!不是一个东西，
  - 可以看到TAG指令最后组合成完整tag，tag前面都有一个!
- tag有3个属性
  - verbatim tag
    - 写法是是tag被\<\>包围
    - yaml处理器会原样(verbatim 逐字)
  - 简写tag
    - 前面讲TAG指令时也提到过 !xx/!!xx/!xx! 三种tag简写对应主要处理/次要处理/基于名字的处理
  - non-specific tag
    - 未指定tag，前面出现过很多次
    - 如果一个node没有tag属性，就会被分配一个未指定tag
    - 对于非plain字面量node，tag就是!
    - 对于其他node，tag就是?
    - 惯例，会依据node的类型来设置如下几种tag
      - tag:yaml.org,2002:seq
      - tag:yaml.org,2002:map
      - tag:yaml.org,2002:str

```yaml
# verbatim tag
# 通过逐字tag来指定类型，
# !<!bar> 第一个!表示是一个tag，里面那个!表示是本地tag

!<tag:yaml.org,2002:str> foo :
  !<!bar> baz

%YAML 1.2
---
!!map {
  ? !<tag:yaml.org,2002:str> "foo"
  : !<!bar> "baz",
}

# 未指定tag
# Assuming conventional resolution:
- "12"
- 12
- ! 12  # 这个的tag会被设置为str

%YAML 1.2
---
!!seq [
  !<tag:yaml.org,2002:str> "12",
  !<tag:yaml.org,2002:int> "12",
  !<tag:yaml.org,2002:str> "12",
]
```

Node的anchor
- anchor 锚，用&开头，后面可以引用
- 这是yaml信息处理第二阶段(输出事件树)的信息
- 不能包含 []{}, 原因是避免流式集合的二义性
- 别名机制上面也详细提到过了，这里就直接上例子了

```yaml
First occurrence: &anchor Value
Second occurrence: *anchor

%YAML 1.2
---
!!map {
  ? !!str "First occurrence"
  : &A !!str "Value",
  ? !!str "Second occurrence"
  : *A,
}
```
