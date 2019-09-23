# 打版本号

前面也提到过了，action的引用，是可以指定版本的：
- 提交号 sha
- 分支名
- tag名

前面也说到了，这3种各有优势，也有各自的缺点，
master分支是不推荐的，因为稳定版本master可能会升级，
推荐使用sha或是发布的对应的版本号

- 绑定主版本号，在小版本中可以进行一个修复工作，也有一个兼容性
- 不绑定master分支的主要原因是：如果一旦发生了主版本升级，那master就存在兼容性问题

## 建议

- node_modules不要放在master分支，最好忽略掉
  - js action
  - 一旦node_modules的文件过期，会导致action执行失败
- 每一个主版本都用一个分支来跟踪
  - 在向master添加新需求之前，先将变更推送到分支
  - 这个分支就是用来做alpha发布测试
- 当一个稳定版本要release时，打一个主版本的tag
  - eg： v1 v2
- 通过github ui，可给每个小版本或补丁发布一个release
  - eg： v1.2.3
- 兼容性断点
  - 给新的主版本创建一个分支releases/v3,一个tag v3,
  - 输入输出的变更，一点要走主版本变更

## 一个简单的git 流

- 当master有了一个稳定版本，创建分支 releases/v1
- master修复bug，合并之后，发布release 1.0.1,并添加tag v1
- master修复bug，合并之后，发布release 1.0.2,并移动tag v1指向最新的小版本1.0,2
- 遇到一个新的主版本升级，创建分支 releases/v2,然后从第一步继续循环
