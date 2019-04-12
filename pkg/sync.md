# sync

提供了一些同步原语，eg：互斥锁

也提供了once和WaitGroup类型，这两个类型常用于底层库

上层业务的同步最好使用协程和通道来完成

## sync.Mutex

互斥锁，默认值是unlocked状态

- Lock() 
- Unlock()

互斥量的加锁解锁

## sync.RWMutex

读写互斥锁，锁可以被多个读或单个写所持有

- RLock()
- RUnlock()

还包括Mutex的加锁解锁(用于写)

## WaitGroup

用于等待协程结束
- Add() 主协程设置增量 cnt
- Done() 子协程结束时调用， cnt减一
- Wait() 如果cnt不为0，就阻塞
