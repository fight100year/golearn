# 字符

## 字符集

- 为了增加可读性，yaml字符流中只能出现可打印的unicode字符集
- 不可见的控制字符中，允许 tab 回车 换行 del，不可见的字符需要进行转义
- 为了和json兼容，引用字面量中可以包含所有非控制字符，不可见字符要进行转义

## 字符编码

- yaml支持utf-8 utf-16,为了兼容json，也支持utf-32
- 如果yaml字节流开头用一个字节来表示字节序，那字符编码也会尊需这个设置
- 为了方便地组合字节流，每个document的开头都可以设置字节序标志。一个流中的所有document都应该用同样的字节序
- 为了和json兼容，引号字面量都可以指定字节序，为了可读性，输出的时候要转义
- 推荐使用utf-8
- 字节序需要出现在document开头，不然就是无效或错误的

## 指示字符

- 一类有特殊语义的字符
  - -连字符，短橫线，表示数组的条目
  - ?，表示map的key
  - :, 表示map的value

```yaml
# 下面有两个map，key都是string，一个value是数组，一个value是map[string]string

# 块式 next-line风格
sequence:
- one
- two
mapping:
  ? sky
  : blue
  sea : green

# 流式
%YAML 1.2
---
!!map {
  ? !!str "sequence"
  : !!seq [ !!str "one", !!str "two" ],
  ? !!str "mapping"
  : !!map {
    ? !!str "sky" : !!str "blue",
    ? !!str "sea" : !!str "green",
  },
}
```

- , 流式集合(数组和map)条目的结尾
- [和] 流式数组的开始和结尾
- {和} 流式map的开始和结尾

```yaml
# 块式写法 in-line风格
sequence: [ one, two, ]
mapping: { sky: blue, sea: green }
```

- 注释用#开头
- &表示node的anchor属性，也就是后面会用占位符来引用这个，节点图上没这个玩意
- \*表示别名node
- !,用于tag指令和tag属性的处理，也可用于本地标签，或者非plain风格的字面量的非指定tag

```yaml
# 块式 map 别名
anchored: !local &anchor value
alias: *anchor

# 流式
%YAML 1.2
---
!!map {
  ? !!str "anchored"
  : !local &A1 "value",
  ? !!str "alias"
  : *A1,
}
```

- | 结尾，表示文字风格的字面量块
- > 结尾，表示折叠风格的字面量块

```yaml
# 块式 字面量 的两种风格：文字风格和折叠风格
literal: |
  some
  text
folded: >
  some
  text

# 流式 plain风格
%YAML 1.2
---
!!map {
  ? !!str "literal"
  : !!str "some\ntext\n",
  ? !!str "folded"
  : !!str "some text\n",
}
```

- 引号，流式字面量

```yaml
# 流式 字面量 引号风格
single: 'text'
double: "text"

%YAML 1.2
---
!!map {
  ? !!str "single"
  : !!str "text",
  ? !!str "double"
  : !!str "text",
}
```

- %表示是指令行

```yaml
%YAML 1.2
--- text

#等同于
%YAML 1.2
---
!!str "text"
```

- @和`是保留符号
- 上面提到过的这些指示符号，[]{},这5个只用于流式集合中

## 换行符

- ascii的换行符有两种：回车 换行
- 非ascii换行符也有一些，但json不支持，为了和json兼容，yaml不认为非ascii换行符是换行
- 字面量里的换行符会被当成常规处理

```yaml
|
  Line break (no glyph)
  Line break (glyphed)↓

%YAML 1.2
---
!!str "line break (no glyph)\n\
      line break (glyphed)\n"
```

## 空白字符

- 空白字符包含了空格和tab
- 其他非打印的空白字符在yaml不算是空白字符

```yaml
# Tabs and spaces
quoted:·"Quoted →"
block:→|
··void main() {
··→printf("Hello, world!\n");
··}

%YAML 1.2
---
!!map {
  ? !!str "quoted"
  : "Quoted \t",
  ? !!str "block"
  : "void main() {\n\
    \tprintf(\"Hello, world!\\n\");\n\
    }\n",
}
```

## 杂项字符

- 数字字符 0-9
- 16进制转义字符 a-f A-F
- ascii的阿拉伯字母 a-z A-Z
- 单词字符： 包含 数字 字母 和-(连字符)
- tag的uri字符，包括[]和ipv6的相关字符
  - 除了可打印的ascii字符，uri出现的其他字符，需要用utf-8编码，并用%进行转义
  - yaml不会扩展这些字符
  - tag字符是需要出现在yaml字符流中的，不能有转义处理的
- !可用于标记tag处理的结尾，所以在tag缩写中是有限制的
  - 为了避免二义性，不能包含[]{},

## 转义字符

- 不可打印的字符是需要转义的
- 只有双引号字面量中才会出现转义
- \ 就是转义字符的起始，[具体可转义的不可见字符包括](https://yaml.org/spec/1.2/spec.html#ns-esc-null)


