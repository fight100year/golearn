# encoding

提供接口给其他包使用，将字节序和文本进行转换

现在包括gob json xml

标准库中,time.Time net.IP都实现了这些接口

encoding包包含了4个接口:
- BinaryMarshaler 这个接口类型,是将对象序列化成二进制
- BinaryUnmarshaler 将二进制反序列化成对象
- TextMarshaler 序列化成文本
- TextUnmarshaler 从文本放序列化成对象

## encoding/json

- Marshal函数:
    - 作用是将对象序列化成json字符串
    - func Marshal(v interface{}) ([]byte, error)
    - 这个函数会递归遍历v
    - 遇到值后,如果值不是nil,且实现了MarshalJSON方法,就调用,产生json数据
    - 如果遇到的值没有MarshlJSON方法,但有TextMarshaler,就会调用后者
    - 遇到内置类型,会进行相应的转换
        - 布尔 -- json的boolean
        - 浮点/整数/数值 -- json的数值
        - string -- json的utf-8编码的字符串
        - \< \> --进行转义
        - 数组和切片 -- json数组(输出的是基于base64编码的字符串)
        - 结构体 -- json的对象 (可导出的字段会作为json对象的成员)
        - 自定义json的key:
            - omitempty表示是空值(false/0/nil),在json对象中忽略
            - 如果含有一个短横线,表示永久忽略
            - 特殊情况:如果含有一个短横线+一个逗号,json对象的key就是一个短横线
            - 如果含有的是string,表明结构体存的值已经是一个json格式的数据
        - map -- json的对象
        - 指针 -- .A
        - 接口 -- json中用null表示
        - 通道/复数/函数 就不转化为json
- Unmarshal 反序列化,将json转换成接口对象
    - 
