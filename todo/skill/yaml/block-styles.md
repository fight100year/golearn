# 块式

- 块式采用缩进来表示数据结构，而不是标记符
- 好处是更好的可读性
- 坏处是不那么紧凑

## 块式字面量风格

- 块式字面量有两种风格：literal和folded，文字风格和折叠风格

块式字面量头部：
- 块式字面量的内容前面会有一些标记，叫头部，用于控制内容
- 这个头部后面跟的是非内容的换行符，也允许有注释的
- 这种情况下，注释后面不能有其他注释

```yaml
- | # Empty header↓
 literal
- >1 # Indentation indicator↓
 ·folded
- |+ # Chomping indicator↓
 keep

- >1- # Both indicators↓
 ·strip

%YAML 1.2
---
!!seq [
  !!str "literal\n",
  !!str "·folded\n",
  !!str "keep\n\n",
  !!str "·strip",
]
```

- 头部有很多写法，下面来看一下
- 块式 可在字面量头部指定相对父级的缩进等级

```yaml
# 这种写法，连最后的换行都算在内容里了

- |°        # 普通头部，文字风格，无注释
·detected
- >°        # 普通头部，折叠风格，第一个空行要忽略
·
··
··# detected
- |1        # 1表示缩进一个空格
··explicit
- >°        # 折叠模式，\t是内容
·→
·detected

%YAML 1.2
---
!!seq [
  !!str "detected\n",
  !!str "\n\n# detected\n",
  !!str "·explicit\n",
  !!str "\t·detected\n",
]
```

- 头部中还有一种写法：chomping标识符
  - 用于控制最后一个换行和空行
  - - 表示最后的换行和结尾的空行不属于字面来内容
  - + 表示最后的换行和结尾的空行属于字面来内容
  - 如果不指定，就是默认行为(clipping),最后的换行属于内容，但结尾空行不属于内容

```yaml
strip: |-
  text↓
clip: |
  text↓
keep: |+
  text↓

%YAML 1.2
---
!!map {
  ? !!str "strip"
  : !!str "text",
  ? !!str "clip"
  : !!str "text\n",
  ? !!str "keep"
  : !!str "text\n",
}
```

- 为了避免二义性，如果结尾空行后面有注释，注释的缩进不能小于块式字面量内容的缩进
- 只有在这种情况下，有注释限制，这是唯一注释有限制的地方

```yaml
# Strip
  # Comments:
strip: |-
  # text↓
··⇓
·# Clip
··# comments:
↓
clip: |
  # text↓
·↓
·# Keep
··# comments:
↓
keep: |+
  # text↓
↓
·# Trail
··# comments.

%YAML 1.2
---
!!map {
  ? !!str "strip"
  : !!str "# text",
  ? !!str "clip"
  : !!str "# text\n",
  ? !!str "keep"
  : !!str "# text\n",
}
```

- 如果一个块式字面量的内容只有空行，那这些就是结尾空行算成了内容，也受chomping影响

```yaml
strip: >-
↓
clip: > # 折叠风格，第一个空行忽略
↓
keep: |+
↓

%YAML 1.2
---
!!map {
  ? !!str "strip"
  : !!str "",
  ? !!str "clip"
  : !!str "",
  ? !!str "keep"
  : !!str "\n",
}
```

literal(文字)风格：
- 标记符是 |
- 这种写法 是块式字面量可读性最高的，也是最简单的，同样是限制最多的

```yaml
|↓
·literal↓
·→text↓
↓

%YAML 1.2
---
!!str "literal\n\ttext\n"
```

- 文字字面量，所有的缩进都被认为是内容，包括空白字符
- 换行是正常处理，空行不是折叠
- 不是引号风格，所以也不存在转义。同样没有断行一说

```yaml
|
·
··
··literal↓
···↓
··
··text↓
↓
·# Comment

%YAML 1.2
---
!!str "\n\nliteral\n·\n\ntext\n"
```

folded(折叠)风格：
- 标记符是 >
- 和文字风格类似，都是比较简单的写法
- 折叠风格是遵循行折叠规则
  - 换行转换成空格
  - 任意断行，断行点要在空格处，且左右要有非空白字符
- 多缩进的行，不使用行折叠规则
- 换行和空行分割的折叠行，不使用行折叠规则
- 最后一个换行和结尾的空行，由chomping控制，不使用折叠规则

```yaml
# 块式 折叠字面量

>↓
·folded↓
·text↓
↓

%YAML 1.2
---
!!str "folded text\n"

# 多缩进和之间的空行不使用折叠规则
# 断行在空格处

>

·folded↓
·line↓
↓
·next
·line↓
   * bullet

   * list
   * lines

·last↓
·line↓

# Comment

%YAML 1.2
---
!!str "\n\
      folded line\n\
      next line\n\
      \  * bullet\n
      \n\
      \  * list\n\
      \  * lines\n\
      \n\
      last line\n"
```

## 块式 集合风格

- 为了可读性，块式集合风格中没有标记符

块式数组：
- 数组的每个元素都是一个node，每个node前都有一个 "-"
- "-"后面要跟一个空格，如果没有空格，就是非引用字面量了

```yaml
block sequence:
··- one↓
  - two : three↓

%YAML 1.2
---
!!map {
  ? !!str "block sequence"
  : !!seq [
    !!str "one",
    !!map {
      ? !!str "two"
      : !!str "three"
    },
  ],
}
```

- 每个条目node(数组元素)，可以完全是空，或者嵌套其他node，或者用紧凑写法
- 这些node无法指明属性

```yaml
-° # Empty
- |
 block node
-·- one # Compact
··- two # sequence
- one: two # Compact mapping

%YAML 1.2
---
!!seq [
  !!null "",
  !!str "block node\n",
  !!seq [
    !!str "one"
    !!str "two",
  ],
  !!map {
    ? !!str "one"
    : !!str "two",
  },
]
```

块式map：
- 基本和前几章提到过的内容一样

块式Node：
- 流式node可以嵌入到块式node中，反之不行
- 块式node嵌套流式node，流式node的缩进至少要多出一个空格

```yaml
-↓
··"flow in block"↓
-·>
 Block scalar↓
-·!!map # Block collection
  foo : bar↓

%YAML 1.2
---
!!seq [
  !!str "flow in block",
  !!str "Block scalar\n",
  !!map {
    ? !!str "foo"
    : !!str "bar",
  },
]
```

- 块式node的属性可能跨越几行，这些属性的缩进至少要多出一个空格

```yaml
literal: |2
··value
folded:↓
···!foo
··>1
·value

%YAML 1.2
---
!!map {
  ? !!str "literal"
  : !!str "value",
  ? !!str "folded"
  : !<!foo> "value",
}
```

- 人对"-"的的感知，如果存在嵌套，嵌套至少要多出一个空格

```yaml
sequence: !!seq
- entry
- !!seq
 - nested   # 多一个空格，就表示是嵌套结构
mapping: !!map
 foo: bar

%YAML 1.2
---
!!map {
  ? !!str "sequence"
  : !!seq [
    !!str "entry",
    !!seq [ !!str "nested" ],
  ],
  ? !!str "mapping"
  : !!map {
    ? !!str "foo" : !!str "bar",
  },
}
 ```
