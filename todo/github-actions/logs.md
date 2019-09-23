# github action的日志

这里说的日志主要指github页面查看日志时的打印信息

打印日志的格式：
- echo ::log-command parameter1={data},parameter2={data}::{command value}
- 不区分大小写

具体例子：
- 设置或更新环境变量
  - 区分大小写
  - ::set-env name={name}::{value} 
  - echo ::set-env name=action_state::yellow
- 设置输出参数
  - ::set-output name={name}::{value}
  - echo ::set-output name=action_fruit::strawberry
  - 设置一个元文件(action.yml)中未指定的输出参数，会报错
- 添加一个系统目录
  - ::add-path::{path}
  - echo ::add-path::/path/to/dir
  - 给job设置一个系统目录
  - 这个设置在之后的step中生效，但前action不起效
- 打印一个debug消息
  - ::debug::{message}
  - echo ::debug file=app.js,line=1::Entered octocatAddition method
  - 如果要查看debug级别的信息，需要将ACTIONS_STEP_DEBUG设置为true，在github中设置
- 打印一个warning消息
  - ::warning file={name},line={line},col={col}::{message}
  - echo ::warning file=app.js,line=1,col=5::Missing semicolon
- 打印错误消息
  - ::error file={name},line={line},col={col}::{message}
  - echo ::error file=app.js,line=10,col=15::Something went wrong
- 屏蔽日志中的值
  - ::add-mask::{value}
  - echo ::add-mask::Mona The Octocat
  - 说白了就是简单的 "Mona The Octocat" 变成"\*\*\*"
- 开始或暂停日志命令
  - ::stop-commands::{endtoken}
  - echo ::stop-commands::pause-logging  停止解析日志命令
  - ::{endtoken}::
  - echo ::pause-logging:: 开始解析日志命令
