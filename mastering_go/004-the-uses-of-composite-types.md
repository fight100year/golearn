# 复杂类型的使用

前一章提到了array、slice、map、pointer、const、for、for range，
以及time date package，
本章主题包括tuple和string，switch语句，struct和正则

具体展开的主题有以下几个：
- struct
- tuple
- 字符串 string/runes/byte slice/string literals
- 正则
- 匹配模式
- switch
- string package
- 计算圆周率π
- kv存储


复杂类型包括了两方面：
- struct 随意组合复杂类型
- tuple 函数可以返回多个值

## struct

new 关键字，返回的是地址

make和new的区别
- make 会适当初始化，new不会初始化
- make返回不是指针，new返回指针
- make只适用于slice channel map

## tuple

元组，以前常说的一元组，二元组，三元组等，
元组是由多个部分组成的有限有序列表。

重点是有限个数，有序

go中使用tuple的地方，大多数在于函数返回一个多元组

可以用_来指出，某一个元组的值被忽略了

## regex

匹配模式和正则表达式，是go中的重要一部分，
主要用于字符串的搜索。regexp 包

## strings

go中的string就是一个只读的byte slice

rune 是int32的别名，可以用来表示unicode code point，
也就是单个unicode字符，一个rune字面量就是'单个字符'

import 的时候可以给package取一个别名

## switch

使用switch的唯一原因是因为：switch支持正则
