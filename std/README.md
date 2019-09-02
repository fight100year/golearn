# 查看标准库的依赖关系

参考了[xx](https://www.cnblogs.com/ios122/p/7639478.html),预览结果是[xx](https://gallery.echartsjs.com/editor.html?c=xSyJNqh8nW)

自己写程序的好处是:想到任何稀奇古怪的需求,都可以自己添加,我添加的如下:
- r 重新显示最底层库
- searchxx 搜索xx开头的包
- 输入完整的包名 显示此包的调用关系
- q 结束程序

运行: go run std.go,数据在info里,也可以通过go list -json std 获取,再转成json格式即可


## package 分类

目前接触到的分类有：
- std 就是标准库里的包 go list -json std
- builtin 语言内置包，主要是声明一些标识符 go list -json builtin
