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
- [actions/download-artifact](/todo/github-actions/github/download-artifact.md)
  - 下载构建的输出文件
- [actions/upload-artifact](/todo/github-actions/github/upload-artifact.md)
  - 将workflow中的输出进行上传,在github页面上可点击下载按钮进行下载
- [elgohr/Github-Release-Action](/todo/github-actions/github/publish-release.md)
  - 通过action来发布一个release版本, 发布的tag是时间格式，后期可以修改一下
- [softprops/action-gh-release](/todo/github-actions/github/publish-release2.md)
  - 通过action来发布一个release版本，主要基于tag来触发
- [einaregilsson/build-number](/todo/github-actions/github/build-number.md)
  - 通过action生成一个带顺序的build号,action每执行一次，build号加1
- [maddox/actions/sleep](/todo/github-actions/github/sleep.md)
  - sleep n秒
- [maddox/ations/wait-for-200](/todo/github-actions/github/wait-for-200.md)
  - 检查http的状态是否是200,还添加重试功能
- [maddox/ations/ssh](/todo/github-actions/github/ssh.md)
  - 通过ssh连接上host，并运行一些东西
- [hmarr/debug-action](/todo/github-actions/github/debug.md)
  - 开发调试action的助手，可以打印环境变量和事件的payload
  - 安全信息会自动过滤
