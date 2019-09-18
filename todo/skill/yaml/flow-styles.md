# 流式

- 流式风格，涵盖了json的扩展
- 为了增加可读性，也添加了很多小规则
- node的tag是用于控制原生数据结构的构造
- node的别名是用于减少构造实例

## 别名node

- 别名node，序列化过程中，以前出现过node后面又出现
- 第一次出现是anchor，用&标记，后面出现就是别名，用\*标记
- 别名引用的是最近一个anchor
- 别名node不能指定其他属性和内容

```yaml
First occurrence: &anchor Foo
Second occurrence: *anchor
Override anchor: &anchor Bar
Reuse anchor: *anchor

%YAML 1.2
---
!!map {
  ? !!str "First occurrence"
  : &A !!str "Foo",
  ? !!str "Override anchor"
  : &B !!str "Bar",
  ? !!str "Second occurrence"
  : *A,
  ? !!str "Reuse anchor"
  : *B,
}
```

## 空node

- 空node，就是node省略内容的情况
- 如果出现空node，一般出现在流式 字面量plain风格中，node的内容会解析为null值

```yaml
# 下面有两个map[string]string
# 第一个值是空的，第二个key是空的
{
  foo : !!str°,
  !!str° : bar,
}

%YAML 1.2
---
!!map {
  ? !!str "foo" : !!str "",
  ? !!str ""    : !!str "bar",
}
```

- node的属性(tag/anchor)或node的内容都是可以忽略的
- 属性和内容全部省略的叫完全空node，出现的时候需明确指定出来
  - 完全空node，因为没有tag，也就没有类型，默认值也就是null了

```yaml
{
  ? foo :°,
  °: bar,
}

%YAML 1.2
---
!!map {
  ? !!str "foo" : !!null "",
  ? !!null ""   : !!str "bar",
}
```

## 流式字面量风格

- 流式字面量风格有3种：plain风格(非引号风格)，quoted风格(包括单引号和双引号风格)

双引号风格：
- 格式是 "xxx" 
- 也是唯一可以带转义字符的风格，转义字符是 \

```yaml
"implicit block key" : [
  "implicit flow key" : value,
]

%YAML 1.2
---
!!map {
  ? !!str "implicit block key"
  : !!seq [
    !!map {
      ? !!str "implicit flow key"
      : !!str "value",
    }
  ]
}
```

- 双引号风格如果存在多行，换行规则符合流式行折叠规则
  - 行尾的空白字符会忽略
  - 换行也是可以转义的，不过转义之后，换行和后面的空白字符就属于内容了
  - 普通的换行会解析成空格
  - 利用\和空白字符，可以在任意地方换行，出现\之后，她们前后空白字符不是全部忽略的，具体可看下面的例子

```yaml
"folded·↓           # 普通替换，替换成空格
to a space,→↓       # 普通替换，替换成一个换行，空白字符会忽略
·↓
to a line feed, or·→\↓    # 遇到转义字符\， 所以\之前的都是内容(使用这条规则，就不会使用换行的普通替换，替换成空格了)
·\·→non-content"          # 遇到转义字符\, 后面的空格和tab都属于内容

%YAML 1.2
---
!!str "folded to a space,\n\
      to a line feed, \
      or \t \tnon-content"
```

- 多行中，除了第一行的开头和最后一行的结尾，其他所有行的开头和结尾的空白字符都不属于内容
- 空行需要考虑到行折叠规则(第一个空行会被忽略掉) 

```yaml
"·1st non-empty↓
↓
·2nd non-empty·
→3rd non-empty·"

%YAML 1.2
---
!!str " 1st non-empty\n\
      2nd non-empty \
      3rd non-empty "
```

单引号风格：
- 格式 'xx'
- 可以使用 \ 和 ", 因为单引号不会进行转义
- 单引号只能出现可见字符
- 多行中，除了第一行的开头和最后一行的结尾，其他所有行的开头和结尾的空白字符都不属于内容

```yaml
# 在map的key中如果出现单引号/双引号字面量，只能是单行

'implicit block key' : [
  'implicit flow key' : value,
]

%YAML 1.2
---
!!map {
  ? !!str "implicit block key"
  : !!seq [
    !!map {
      ? !!str "implicit flow key"
      : !!str "value",
    }
  ]
}

# 多行，除了首行开头和尾行结尾的空白字符，其他行首行尾的空白字符会被忽略掉

'·1st non-empty↓
↓
·2nd non-empty·
→3rd non-empty·'

%YAML 1.2
---
!!str " 1st non-empty\n\
      2nd non-empty \
      3rd non-empty "
```

plain风格，也叫简洁风格，也叫非引号风格
- 没有指示符和转义，所以可读性是最高的
- 这种限制最多，很依赖上下文
  - 首先是字符集需要指定
  - 字面量非空，首尾无空白字符
  - 唯一断行的是空白字符左右都有非空白字符时
- 为了避免二义性，非引号字面量是不能包含 #/:
- 同理，在map的key中，如果出现非引号字面量，也不能出现[]{},字符
- 在map的key中如果出现单引号/双引号字面量，只能是单行

```yaml
# 可以看出，非引号风格字面量的可读性是最高的

# Outside flow collection:
- ::vector
- ": - ()"
- Up, up, and away!
- -123
- http://example.com/foo#bar
# Inside flow collection:
- [ ::vector,
  ": - ()",
  "Up, up and away!",
  -123,
  http://example.com/foo#bar ]

%YAML 1.2
---
!!seq [
  !!str "::vector",
  !!str ": - ()",
  !!str "Up, up, and away!",
  !!int "-123",
  !!str "http://example.com/foo#bar",
  !!seq [
    !!str "::vector",
    !!str ": - ()",
    !!str "Up, up, and away!",
    !!int "-123",
    !!str "http://example.com/foo#bar",
  ],
]

# map的key，字面量只能是单行

implicit block key : [
  implicit flow key : value,
]

%YAML 1.2
---
!!map {
  ? !!str "implicit block key"
  : !!seq [
    !!map {
      ? !!str "implicit flow key"
      : !!str "value",
    }
  ]
}

# 多行规则和引号风格是一致的

1st non-empty↓
↓
·2nd non-empty·
→3rd non-empty

%YAML 1.2
---
!!str "1st non-empty\n\
      2nd non-empty \
      3rd non-empty"
```

## 流式数组风格

- 流式集合是可以嵌套在块式集合中的，或者嵌套在流式集合中
- 流式集合也可以称为map的kye，或者map的value
- 流式集合的条目，用","标识结尾，当然最后一个条目的后面的","是可以忽略的
- 集合中条目的类型，不一定是一样的，这和编程语言有些差别，yaml的表达范围会更广一些

流式数组：
- 紧凑型风格是[a,b,c]，当然，每个条目都可单独成行

```yaml
- [ one, two, ]
- [three ,four]

%YAML 1.2
---
!!seq [
  !!seq [
    !!str "one",
    !!str "two",
  ],
  !!seq [
    !!str "three",
    !!str "four",
  ],
]

[
"double
 quoted", 'single
           quoted',
plain
 text, [ nested ],
single: pair,
]

%YAML 1.2
---
!!seq [
  !!str "double quoted",
  !!str "single quoted",
  !!str "plain text",
  !!seq [
    !!str "nested",
  ],
  !!map {
    ? !!str "single"
    : !!str "pair",
  },
]
```

## 流式map风格

- map的格式是{}
- 每个kv对都是用","分割，最后一对后的","可以省略
- ?表示map的key是一个复杂结构
- ?后的kv对都是node，都可以是完全空node，写法如下

```yaml
- { one : two , three: four , }
- {five: six,seven : eight}

%YAML 1.2
---
!!seq [
  !!map {
    ? !!str "one"   : !!str "two",
    ? !!str "three" : !!str "four",
  },
  !!map {
    ? !!str "five"  : !!str "six",
    ? !!str "seven" : !!str "eight",
  },
]

# ?表示key是复杂结构
# 还有kv都是完全空node，值都是null:null
{
? explicit: entry,
implicit: entry,
?°°
}

%YAML 1.2
---
!!map {
  ? !!str "explicit" : !!str "entry",
  ? !!str "implicit" : !!str "entry",
  ? !!null "" : !!null "",
}

```

- map中用":"来将kv分开，但是value之前一定有有个空格
- 基于这点，非引号字面量里是可以包含":"的，只要":"后面没有空格
  - 这点非常适合：非引号字面量 url和时间戳
  - 也可以支持a:b这类写法，a:b是非引号字面量，a: b就是map的kv对了

```yaml
{
unquoted·:·"separate",
http://foo.com,     # http后面的":"不属于map，实际上，这整个非引号字面量都是key，value和":"都被省略了
omitted value:°,    # value省略了
°:·omitted key,     # key省略了
}

%YAML 1.2
---
!!map {
  ? !!str "unquoted" : !!str "separate",
  ? !!str "http://foo.com" : !!null "",
  ? !!str "omitted value" : !!null "",
  ? !!null "" : !!str "omitted key",
}
```

- 为了兼容json格式，":"后面没有空格，可直接跟value
- 这个更多是依赖上下文的结构，如果检测到是json格式，yaml处理器会自动识别
- 这种写法会降低可读性

```yaml
{
"adjacent":value,
"readable":·value,
"empty":°
}

%YAML 1.2
---
!!map {
  ? !!str "adjacent" : !!str "value",
  ? !!str "readable" : !!str "value",
  ? !!str "empty"    : !!null "",
}
```

- 为了更加紧凑的写法，有时会省略{}
  - 前提条件是map的kv中的node没有指定任何属性(tag/anchor)
- 明确指定? 可减少二义性

```yaml
# 简写
[
foo: bar
]

%YAML 1.2
---
!!seq [
  !!map { ? !!str "foo" : !!str "bar" }
]

# 减少二义性

[
? foo
 bar : baz
]

%YAML 1.2
---
!!seq [
  !!map {
    ? !!str "foo bar"
    : !!str "baz",
  },
]

```

- ?是可以省略的，也是比较常见的
- 省略?，map中就依据":"来区分key和value，限制：1024个unicode字符内要出现":"
- 当然，key可以是任意node，所以1024长度的限制还是有意义的

```yaml
# 常用写法

- [ YAML·: separate ]
- [ °: empty key entry ]
- [ {JSON: like}:adjacent ]

%YAML 1.2
---
!!seq [
  !!seq [
    !!map {
      ? !!str "YAML"
      : !!str "separate"
    },
  ],
  !!seq [
    !!map {
      ? !!null ""
      : !!str "empty key entry"
    },
  ],
  !!seq [
    !!map {
      ? !!map {
        ? !!str "JSON"
        : !!str "like"
      } : "adjacent",
    },
  ],
]

```

## 流式Node

- 流式非引用字面量风格，只适合node没有指明属性的情况

```yaml
- [ a, b ]
- { a: b }
- "a"
- 'b'
- c

%YAML 1.2
---
!!seq [
  !!seq [ !!str "a", !!str "b" ],
  !!map { ? !!str "a" : !!str "b" },
  !!str "a",
  !!str "b",
  !!str "c",
]
```

- 流式node是可以有属性的

```yaml
- !!str "a"
- 'b'
- &anchor "c"
- *anchor
- !!str°

%YAML 1.2
---
!!seq [
  !!str "a",
  !!str "b",
  &A !!str "c",
  *A,
  !!str "",
]
```
