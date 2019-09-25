# 分析已有action

前期主要以github官方的action为主，后期以action市场为主

官方action列表：
- [actions/checkout](/todo/github-actions/github/checkout.md) 完成项目的fetch 和 checkout
  - 说明：将项目克隆到指定目录，并将git头指针指向指定版本
- [actions/setup-go](/todo/github-actions/github/setup-go.md) 安装go环境
  - 说明：非常适合使用矩阵构建，测试多平台的开发
- [一个推荐的go ci](/todo/github-actions/github/go-ci.md)
  - 基于ci/cd的，可以基于这个workflow进行扩展
- [actions/first-interaction](/todo/github-actions/github/first-interaction.md)
  - 贡献者第一次创建issues和pr时，发送一条指定信息
- [actions/labeler](/todo/github-actions/github/labeler.md)
  - 根据pr修改的文件路径自动为pr打标签
- [actions/stale](/todo/github-actions/github/stale.md)
  - 标记并关闭一段时间内未更新的issues/pr
