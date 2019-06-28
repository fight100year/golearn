# database/sql

针对标准sql或类sql提供了通用的接口

- 使用sql包，必须和数据库驱动包driver一起。
- 驱动包driver不只是context的取消操作，要等查询完成后才能返回

使用的步骤：
- 创建一个db处理对象 sql.Open
- 检查db对象是否可以进行查询 db.PingContext
- db对象使用完之后，需要Close

查询步骤：
- ExecContext函数，适用于无返回的sql语句，eg：增删改
- QueryContext函数，适用于select 查
- QueryRowContext函数，适用于单行返回的场景
- PrepareContext函数，适用于预编译语句

事务：
- 上面的查询，都可以放在一个事务里
- 开始事务; BeginTx
- 事务的结束： Commit 或是 Rollback

空处理：
- 如果表中的某个字段可以为空 null，在Scan函数中，就需要支持null的类型
- sql包中，只实现了如下支持null的类型，其他支持null的类型，需要放在驱动包里
    - NullBool
    - NullFloat64
    - NullInt64
    - NullString

## api分析

- 驱动：注册和查询
- 字段类型：字段类型 字段名 长度等
- db的单连接：查询、事务都可以做
- db的连接池：功能比单连接更加丰富
- 连接池的统计信息：也就是状态
- 事务的隔离级别：从最低的级别到串行化，最高的线性化
- 命名参数：查询语句中，用于替换其中参数而存在的
- 4种支持null的类型
- 存储过程输出类型 Out
- 内存数据，eg：[]byte, 可用于从结果集中读取每一row
- 执行的结果集(增删改的)：sql命令的返回结果
- 结果集的一行：还可从中取具体某一个字段
- query的结果集(select的)：支持遍历，取字段
- 支持自定义取字段的接口
- 预处理类型
- 事务类型
- 事务的设置选项

标准的sql包，带来了很大的好处，将所有db操作抽象出来封成sql包，
而针对不同的db，采用驱动包来对接，这样无疑提高了开发效率，go再一次诠释了“实用”二字。


