# softprops/action-gh-release

- 通过action发布release

## 仓库分析

- license/readme是常规操作
- 是一个ts action

```yaml
# action.yml

# https://help.github.com/en/articles/metadata-syntax-for-github-actions
name: 'GH Release'
description: 'Github Action for creating Github Releases'
author: 'softprops'
inputs:
  body:
    description: 'Note-worthy description of changes in release'
    required: false
  body-path:
    description: 'Path to load note-worthy description of changes in release from'
    required: false
  name:
    description: 'Gives the release a custom name. Defaults to tag name'
    required: false
  draft:
    description: 'Creates a draft release. Defaults to false'
    required: false
  prerelease:
    description: 'Identify the release as a prerelease. Defaults to false'
    required: false
  files:
    description: 'Newline-delimited list of path globs for asset files to upload'
    required: false
env:
  'GITHUB_TOKEN': 'As provided by Github Actions'
runs:
  using: 'node12'
  main: 'lib/main.js'
branding:
  color: 'green'
  icon: 'package'
```

## action 分析

- 这个一个标准的action写法，token是通过环境变量传入
- 入参：
  - 变更描述，可选
  - 变更描述路径，可选
  - release的标题，可选，默认是tag名
  - 是否是release草稿，可选，默认是false
  - 是否是预发布，可选，默认是false
  - 文件列表，可选



## 使用

```yaml
# 常规用法
# 向git的某一个tag中push
# 其中在step中使用了if 

name: Main

on: push

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v1
      - name: Release
        uses: softprops/action-gh-release@v1
        if: startsWith(github.ref, 'refs/tags/')
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

# 也可以直接配置，只作用于tag

name: Main

on:
  push:
    tags:
      - 'v*.*.*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v1
      - name: Release
        uses: softprops/action-gh-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

# 还可以指定发布的文件列表

name: Main

on: push

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v1
      - name: Build
        run: echo ${{ github.sha }} > Release.txt
      - name: Test
        run: cat Release.txt
      - name: Release
        uses: softprops/action-gh-release@v1
        if: startsWith(github.ref, 'refs/tags/')
        with:
          files: Release.txt
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

# 除此之外，还可以添加额外的文件

name: Main

on: push

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v1
      - name: Build
        run: echo ${{ github.sha }} > Release.txt
      - name: Test
        run: cat Release.txt
      - name: Release
        uses: softprops/action-gh-release@v1
        if: startsWith(github.ref, 'refs/tags/')
        with:
          files: |
            Release.txt
            LICENSE
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

- 还可以将changelog等都放在action上

## 总结

- 这个通过action发布release的版本还不错，
- 这个action更多集中在tag上，而之前那个action每次push都会触发
- 还是没有找到一个完全自定义，且可控的action来完成这件事
