# elgohr/Github-Release-Action

- 通过action发布一个build号，这个号是有顺序的
- 这个build号放在github的引用中，这个引用并不附加任何额外信息

## 仓库分析

- license/readme是常规操作
- 是一个简单的js action

```yaml
# action.yml

name: 'Build Number Generator'
description: 'Generate sequential build numbers for workflow runs'
author: 'Einar Egilsson'
runs:
  using: 'node12'
  main: 'main.js'
inputs:
  token:
    description: 'GitHub Token to create and delete refs (GITHUB_TOKEN)'
    required: false # Not required when getting the stored build number for later jobs, only in the first jobs when it's generated

outputs:
  build_number:
    description: 'Generated build number'

branding:
  icon: 'hash'
  color: 'green'
```

## action 分析

- 这个js action很简单
- 一个入参是token
- 一个出参是生成的build号
- 所以这个action一般作为其他action的前置使用

## 使用

```yaml
# 最简单的使用，就是打印获取到的build号

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Generate build number
      uses: einaregilsson/build-number@v1
      with:
        token: ${{secrets.github_token}}
    - name: Print new build number
      run: echo Build number is $BUILD_NUMBER

# 常规场景是将输出作为其他action的输入来使用
# 下一个action引用上一个action的输出，还是有一定规律的
# 用这种引用方式的好处是：一个job中的其他step都可以用同样方式来引用这个action

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Generate build number
      id: buildnumber
      uses: einaregilsson/build-number@v1
      with:
        token: ${{secrets.github_token}}

    # Now you can pass ${{ steps.buildnumber.outputs.build_number }} to the next steps.
    - name: Another step as an example
      uses: actions/hello-world-docker-action@v1
      with:
        who-to-greet: ${{ steps.buildnumber.outputs.build_number }}

# 如果是跨job呢，job的虚拟环境可能不是一样的
# 实际情况中，需要将生成的build保存到github
# 这是就需要用到actions/update-artifact@v1 (这个就是官方的上传action)
# 下一次调用build-number action时，需要先下载，使用actions/download-artifact@v1
# 之后再设置到环境变量即可

jobs:
  job1:
    runs-on: ubuntu-latest
    steps:
    - name: Generate build number
      id: buildnumber
      uses: einaregilsson/build-number@v1 
      with:
        token: ${{secrets.github_token}}        
    - name: Upload build number
      uses: actions/upload-artifact@v1
      with:
        name: BUILD_NUMBER
        path: BUILD_NUMBER
          
  job2:
    runs-on: ubuntu-latest
    steps:
    - name: Download build number
      uses: actions/download-artifact@v1
      with:
        name: BUILD_NUMBER
    - name: Restore build number
      id: buildnumber
      uses: einaregilsson/build-number@v1 
    
    # Don't need to add Github token here, since you're only getting an artifact.
    # After this runs you'll again have the $BUILD_NUMBER environment variable, and 
    # the ${{ steps.buildnumber.outputs.build_number }} output.
```

- 如果想要指定build号的初始值，可以用下面的命令设置
- 因为build号就是放在tag上，更具体一点就是 build-number-100 的tag

```shell
git tag build-number-500
git push origin build-number-500
```

## 总结

- 这个action的功能是每次action运行一次，build号就增加一次
- 一般的版本号是v1.2.3.4
- 其中1是主版本号 2 是子版本号 3 是修正版本号 4 是编译版本号
- v1.2.3可通过其他方式获取，编译版本号正好可以利用这个build号
- 现在发布一个release，就差最后两步了：
  - v1.2.3的获取，计划是通过一个action来获取
  - v1.2.3的变更，应该通过一个参数传入，传入之后，build号可以清零
