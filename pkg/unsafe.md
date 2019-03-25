# unsafe

类型安全的一些操作

unsafe.Pointer 类型的转换会覆盖go类型系统，
说白了就是unsafe.Pointer的转换优先级更高，
所以用法上一定要保证正确


unsafe这个包由go编译器实现，被很多底层package使用，
eg：runtime，syscall，os等

使用unsafe之后，一些无效内存地址的访问，编译器是不会报错的

