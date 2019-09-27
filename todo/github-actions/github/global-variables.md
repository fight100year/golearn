# zweitag/github-actions/global-variables

- js action
- 将一个文件的kv读取，转换成环境变量，只在当前job有效

## 仓库分析

- js action

```Dockerfile
name: 'global-variables'
author: 'zweitag'
description: 'Set global vars for all following build steps'
inputs:
  file:
    description: 'file with global vars'
    required: false
    default: '.github/workflows/.env.global'
runs:
  using: 'node12'
  main: 'index.js'
branding:
  icon: 'settings'
  color: 'green'
```

## action 分析

- 只有一个输入，可选，默认文件路径是.github/workflows/.env.global

```js
const core = require('@actions/core');
const fs = require('fs');

try {
  var file_name = core.getInput('file');
  var file_content = fs.readFileSync(file_name, 'utf8')
  file_rows = file_content.split("\n")

  file_rows.forEach(function (row, index) {
    row = row.trim();
    if (row && !row.startsWith('#')) {
      global_var = row.split("=")
      core.exportVariable(global_var[0], global_var[1]);
    }
  });
} catch (error) {
  core.setFailed(error.message);
}
```

- js大致意思：
  - 读取文件
  - 按行分割
  - 对于每行的kv对，调用package里的功能，设置为环境变量
- 非常简洁

## 使用

```yaml
steps:
- uses: zweitag/github-actions/global-variables@master
```

## 总结

- 非常有用的action
- 有时要设置的环境变量太多，索性写在文件里
