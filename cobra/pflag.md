# pflag

相比标准库中的flag,pflag添加了对posix/gun风格的 --flags,
说白了就是多支持了一种写法,就是长参数写法.

相比标准库中的flag,只是做了一些功能的新增,是flag的一个简单的代替品.

从flag换到pflag,只有一个需要注意的:Flag创建实例时,多来一个成员变量Shorthand


