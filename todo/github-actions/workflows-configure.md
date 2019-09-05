# github actions 工作流配置

## 创建一个 workflow

- 前提：有指定仓库的写权限或管理权限
- 就可以创建/查看/编辑 wrokflow
- 一个仓库可以有多个工作流 在仓库的.github/worklflows/下
- workflow至少要有一个job
- job包含多个step，每个step执行一个独立的小任务，step可以是运行命令，或使用一个action
- action可以自己创建，也可以去社区找
- 配置的工作流的启动可以通过github事件触发，通过调度触发，或通过外部事件触发
- 配置文件用yaml格式，并将配置文件保存到仓库里

创建一个配置文件：
- 在.github/workflows/下创建一个.yml文件
- 按workflow语法构建一个配置文件
- 将配置文件提到到仓库的指定分支，如果想这个分支上运行工作流。


## 管理 workflow run

## wrokflow 语法

## 事件

## 虚拟环境

## 虚拟环境中的软件

## 上下文和表达式语法
