# 预览

## 集合

- 前面说到了，yaml三大原语其中有两个就是集合：map和序列
- block collections(集合块) 包含了缩进和每行中的条目
- block sequences(集合序列,可理解为数组) 每行中的条目，都是以 "- "开头的
- map 用冒号风格key和value
- 注释用#开头

```yaml
字面量的数组
- Mark McGwire
- Sammy Sosa
- Ken Griffey

字面量到字面量的映射，用go翻译就是：map[字面量]字面量
hr:  65    # Home runs
avg: 0.278 # Batting average
rbi: 147   # Runs Batted In

字面量到数组的映射，用go翻译就是 map[字面量]数组
american:
  - Boston Red Sox
  - Detroit Tigers
  - New York Yankees
national:
  - New York Mets
  - Chicago Cubs
  - Atlanta Braves

map的数组，用go翻译：[]map[字面量]字面量
-
  name: Mark McGwire
  hr:   65
  avg:  0.278
-
  name: Sammy Sosa
  hr:   63
  avg:  0.288
```

- yaml除了用缩进来表示，也可以用标记来表示，类似json，这种风格称为flow风格
- 数组用[]来表示，数组元素用逗号分割
- map用大括号表示

```yaml
数组
- [name        , hr, avg  ]
- [Mark McGwire, 65, 0.278]
- [Sammy Sosa  , 63, 0.288]

map
Mark McGwire: {hr: 65, avg: 0.278}
Sammy Sosa: {
  hr: 63,
  avg: 0.288
}
```

## 结构

- yaml中用---(3个-)来分割yaml指令和文档内容
- yaml中用...(3个.)来表示文档结束

```yaml
一个stream中的两个文档，每个文档都是以一个注释开头
# Ranking of 1998 home runs
---
- Mark McGwire
- Sammy Sosa
- Ken Griffey

# Team ranking
---
- Chicago Cubs
- St Louis Cardinals

```

- 重复的nodes(或对象)，第一次标识出现时可在前面带一个&，后面通过别名来引用(前面带一个星号)

```yaml
单文档带两个注释
---
hr: # 1998 hr ranking
  - Mark McGwire
  - Sammy Sosa
rbi:
  # 1998 rbi ranking
  - Sammy Sosa
  - Ken Griffey

一个node出现两次，使用别名来引用
---
hr:
  - Mark McGwire
  # Following node labeled SS
  - &SS Sammy Sosa
rbi:
  - *SS # Subsequent occurrence
  - Ken Griffey
```

- 一个"? "表示一个复杂的key
- 在一个块集合中，map的k:v对可以直接从 短横线/冒号/问号开始

```yaml
map[数组]数组
? - Detroit Tigers
  - Chicago cubs
:
  - 2001-07-23

? [ New York Yankees,
  Atlanta Braves ]
: [ 2001-07-02, 2001-08-12,
  2001-08-14 ]


```

## 字面量

- 在block风格中，文字风格的|用于表示换行
- 另外，折叠风格中,换行用空格代替，直到遇到一个空行或更加缩进的行

```yaml
文字风格中，换行用|代替
# ASCII Art
--- |
  \//||\/||
  // ||  ||__

折叠风格中，换行会替换成空格
--- >
  Mark McGwire's
  year was crippled
  by a knee injury.

折叠风格中，遇到更加缩进行或空行
>
 Sammy Sosa completed another
 fine season with great stats.

   63 Home Runs
   0.288 Batting Average

 What a year!

缩进确定范围
name: Mark McGwire
accomplishment: >
  Mark set a major league
  home run record in 1998.
stats: |
  65 Home Runs
  0.278 Batting Average

```

- flow风格中，字面量包含简单风格(不含引号)和两种引号风格
- 双引号风格支持转义,不需要转义就使用单引号风格
- 所有的flow风格都支持多行

```yaml
引号风格
unicode: "Sosa did fine.\u263A"
control: "\b1998\t1999\t2000\n"
hex esc: "\x0d\x0a is \r\n"

single: '"Howdy!" he cried.'
quoted: ' # Not a ''comment''.'
tie-fighter: '|\-*-/|'

多行flow风格字面量
plain:
  This unquoted scalar
  spans many lines.

quoted: "So does this
  quoted scalar.\n"
```

## tags


