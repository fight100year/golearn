# hmarr/auto-approve-action

- docker action
- pr的自动审阅

## 仓库分析

- readme是常规操作
- 缺少一个license
- 是一个ts action

```yaml
# action.yml

name: 'Auto Approve'
description: 'Automatically approve pull requests'
branding:
  icon: 'check-circle'
  color: 'green'
inputs:
  github-token:
    description: 'The GITHUB_TOKEN secret'
    required: true
runs:
  using: 'node12'
  main: 'dist/index.js'
```

## action 分析

- 这个index.js有1w多行，而且是ts的，所以不分析了
- action只有一个入参，token，必选

## 使用

```yaml
# 常规用法
# 所有的审阅都自动通过

name: Auto approve
on: [pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: hmarr/auto-approve-action@v2.0.0
      with:
        github-token: "${{ secrets.GITHUB_TOKEN }}"

# 可以进行过滤，对指定的某些人才自动审阅通过
steps:
- uses: hmarr/auto-approve-action@v2.0.0
  if: github.actor == 'dependabot[bot]' || github.actor == 'dependabot-preview[bot]'
  with:
    github-token: "${{ secrets.GITHUB_TOKEN }}"
```

## 总结

- 在某些场景下，会有出场机会
