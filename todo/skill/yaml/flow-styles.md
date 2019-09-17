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

- node的属性或node的内容都是可以忽略的
- 


## 流式字面量风格
## 流式数组风格
## 流式map风格
## 流式Node
