# kentaro-m/auto-assign

- docker action
- 里面执行的是ts程序
- 自动添加审阅人和负责人

## 仓库分析

- license/readme是常规操作
- 是一个docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM node:10

LABEL "com.github.actions.name"="Auto Assign"
LABEL "com.github.actions.description"="Add reviewers/assignees to pull requests when pull requests are opened."
LABEL "com.github.actions.icon"="user-plus"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="https://github.com/kentaro-m/auto-assign"
LABEL "homepage"="https://probot.github.io/apps/auto-assign/"
LABEL "maintainer"="Kentaro Matsushita"

ENV PATH=$PATH:/app/node_modules/.bin
WORKDIR /app
COPY . .
RUN npm install --production && npm run build

ENTRYPOINT ["probot", "receive"]
CMD ["/app/lib/index.js"]
```

## action 分析

- 在node环境下，使用npm来构建程序，然后执行
- 执行的是一个ts程序


## 使用

- 会有一个配置文件，里面存了审阅人列表和负责人列表
- 会有以下几种规则：
  - 这个action会自动给新建pr添加审阅人/负责人
  - 配置里指定了多个审阅人，action会自动随机分配一个
  - 自动随机分配，也可以按组分
  - 如果pr中包含了某些关键字，action就不会自动分配
- 这个配置文件放在 .github/auto_assign.yml

```yaml
# 给pr只分配一个审阅人/负责人

# Set to true to add reviewers to pull requests
addReviewers: true

# Set to true to add assignees to pull requests
addAssignees: false

# A list of reviewers to be added to pull requests (GitHub user name)
reviewers:
  - reviewerA
  - reviewerB
  - reviewerC

# A number of reviewers added to the pull request
# Set 0 to add all the reviewers (default: 0)
numberOfReviewers: 0

# A list of assignees, overrides reviewers if set
# assignees:
#   - assigneeA

# A number of assignees to add to the pull request
# Set to 0 to add all of the assignees.
# Uses numberOfReviewers if unset.
# numberOfAssignees: 2

# A list of keywords to be skipped the process that add reviewers if pull requests include it
# skipKeywords:
#   - wip

# 给pr分配多个审阅人/负责人

# Set to true to add reviewers to pull requests
addReviewers: true

# Set to true to add assignees to pull requests
addAssignees: false

# A number of reviewers added to the pull request
# Set 0 to add all the reviewers (default: 0)
numberOfReviewers: 1

# A number of assignees to add to the pull request
# Set to 0 to add all of the assignees.
# Uses numberOfReviewers if unset.
# numberOfAssignees: 2

# Set to true to add reviewers from different groups to pull requests
useReviewGroups: true

# A list of reviewers, split into different groups, to be added to pull requests (GitHub user name)
reviewGroups:
  groupA:
    - reviewerA
    - reviewerB
    - reviewerC
  groupB:
    - reviewerD
    - reviewerE
    - reviewerF

# Set to true to add assignees from different groups to pull requests
useAssigneeGroups: false

# A list of assignees, split into different froups, to be added to pull requests (GitHub user name)
# assigneeGroups:
#   groupA:
#     - assigneeA
#     - assigneeB
#     - assigneeC
#   groupB:
#     - assigneeD
#     - assigneeE
#     - assigneeF

# A list of keywords to be skipped the process that add reviewers if pull requests include it
# skipKeywords:
#   - wip
```

- 如果想每次给pr都分配同一个审阅人，可利用code owners来实现
- 如果需要随机，就可以使用当前action
- 这个action只有一个参数，token

## 总结

- 这个action是非常有用的，特别是多人协作，一个审阅人忙不过来时，这个就是非常好的
- 特别是随机分组
- 不过，对于有code owners的项目(一般是由某个组织的项目维护的项目)，owners机制就非常强大了
