# banyan/auto-label

- 根据pr修改的文件路径，给pr打上标签
- github pr的标签，主要作用是为了分类


## 仓库分析

- 一个标准的docker action
- 自身不带测试
- 和官方的类似，都是docker里运行的是ts程序

```Dockerfile
FROM node:alpine

LABEL "com.github.actions.name"="autolabel"
LABEL "com.github.actions.description"="Add labels to Pull Request based on matched file patterns"
LABEL "com.github.actions.icon"="flag"
LABEL "com.github.actions.color"="gray-dark"

COPY . .
RUN yarn install
RUN apk --no-cache add git
ENTRYPOINT ["node", "/dist/entrypoint.js"]
```

## action 分析

- 用node运行的ts文件
- ts文件就不继续深入分析了

```json
{
  "rules": {
    "frontend": ["*.js", "*.css", "*.html"],
    "backend": ["app/", "*.rb"],
    "ci": ".circleci"
  }
}
```

## 使用

- 使用之前需要创建一个标签文件，放在.github/auto-label.json
- 入参只有一个，token

## 总结

- 缺少了一个action 测试环节
- 使用这个action，需要一个额外的标签文件
- 对比官方的自动标签action，区别是配置文件格式，一个是yaml，一个是json
