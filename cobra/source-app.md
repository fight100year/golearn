# cobra 可执行

这个可执行也叫cobra生成器,用处是为已存在的程序添加cli

## 套路第一步 cobra init

这个命令很强悍,执行之后,会自动创建cmd目录,并添加一个root命令,
其次也将main函数改造了,基本三就是自动添加了一个cobra架子

## 套路第二步 cobra add

- cobra add xxx 添加xxx命令
- cobra add yyy -p 'xxxCmd' 在xxx下,添加子命令yyy

针对命令,自动生成一些文件,这个架子又好看了一些


## 套路第三步 配置cobra生成器

更加强大的架子

