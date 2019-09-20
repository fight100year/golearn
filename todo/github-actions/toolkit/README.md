# github action 开发工具的学习

[官方资料](https://github.com/actions/toolkit)

- 这个系列的焦点是如何开发github actions
- 这套工具的目的是让开发github actions更加容易，提高一致性
- 这套工具的做法就是提供一系列package
  - @acrions/core 获取输出/设置输出/设置结果/日志/安全/环境变量相关的核心函数
  - @actions/exec 命令行运行这些工具的必须函数
  - @actions/io cli文件系统的核心函数
  - @actions/tool-cache 下载/缓存这些工具的必须函数
  - @actions/github Octokit客户端和action运行的上下文(Octokit是一个github api客户端，13年发布的)

action要么运行在容器中，要么运行在github提供的host机器上
- [ ] [选择哪种action类型](/todo/github-actions/toolkit/action-type.md)
- [ ] [创建js action的简单示例](/todo/github-actions/toolkit/js-action-hello.md)
- [ ] [js action 例子](/todo/github-actions/toolkit/js-action.md)，test/lint/workflow/publish/version
- [ ] [TypeScript action 例子](/todo/github-actions/toolkit/ts-action.md)，compile/test/lint/workflow/publish/version
- [ ] [docker action 例子](/todo/github-actions/toolkit/docker-action.md)
- [ ] [使用octokit的docker action例子](/todo/github-actions/toolkit/docker-action-octokit.md)
- [ ] [action的版本](/todo/github-actions/toolkit/version.md)，打版本，发布，tag打标签
