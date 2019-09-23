# js action 的创建

仓库在[这里](https://github.com/actions/javascript-action)

- 这个仓库是个模板，用于创建一个js action
- 做测试/静态检测/工作流/发布/打版本都可以参考这个模板
- 如果要利用这个模版来快速构建一个自己的js action，直接引用这个模版(use this template)

## 模版分析

- 整个仓库目录分析：
  - 核心的还是 action.yml(里面是js action的详细描述/功能/输入输出等信息)
  - action.yml最后还是指定了这个action就是执行一段nodejs，执行的文件就是index.js
  - 具体的执行文件是index.js 她用node执行时的一些依赖就是package.json
  - index.js中还依赖了本地js：wait.js，且还使用了github提供的一些库：@actions/core
  - git忽略文件/readme/license都是常规操作
  - 还有这个仓库也有一个workflow，用来测试js action
    - .github/workflows下有一个test.yml
    - 配合index.test.js 在发布时做测试workflow
- 核心文件分析：
  - action.yml
    - 这里的例子只是一个简单的nodejs例子
    - 输入一个数字，输出一个数字，nodejs程序
  - index.js
    - 这里也是简单例子，
      - debug打印当前时间
      - 等待一段时间(调用的是wait.js, wait里也使用js里的超时来实现的)
      - debug打印当前时间
      - 将当前时间输出
    - 写法是nodejs的写法
- 测试分析
  - 这个action也有工作流在发布时做测试
  - 也可以在提交之前，手动测试(npm install,npm test)

接下来就是针对需求，修改action

## diy action

修改action.yml来自定义action：
- action的输入输出
- action的名称，描述等
- 对同步异步操作的支持

diy之后，就要发布一个release版本：
- 其实就是新增了一个releases的分支

## 其他项目的使用

在某个step中：

```yaml
- uses: 63isOK/hello-js@releases/v1
  with:
    milliseconds: 5000
```

## 总结分析

- js actions做了一件事：定义好输入输出，执行一个nodejs程序
- 除了js actions ，其他语言实现的actions都是ok的，
- 说回来，actions的本质还是一段可执行程序，workflow就是将一个个操作组合成一个自动化集合
