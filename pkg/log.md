# log

日志库，可以和系统的日志服务对接，
syslog库是log的一部分，可以指定日志设备和日志等级

## logging levels

日志等级，严重等级
- debug
- info
- notice
- warning
- err
- crit
- alert
- emerg

## logging facilities

日志类别，系统中有很多类别，如果发送一个日志信息，但其他中类别
在系统中找不到，意味和这条信息会被忽略，被丢弃

所有的unix系统都有一个单独的服务，用于接收日志数据，并持久化到文件中，
unix系统中有很多类似的服务，有两个是非常常见的：
- syslogd(8)   配置文件一般是/etc/syslog.conf
- rsyslogd(8)   配置文件一般是/etc/rsyslog.conf

ubuntu18.04中，使用的是rsyslogd,日志类别配置在rsyslog.conf中，
也配置了自配置目录，

unix并不是配置了所有的日志类别，我们可以配置一些自定义的，
有很多日志信息，如果类别在指定unix系统中没指定，就会被丢弃。

## source analyze

- every log message is output on a separate line
- Fatal() will call os.Exit(1) after writing the log message
- Painc() wile call panic() after writing the log message

output format:
- prefix
- year/month/day hour:mimute:second.123456
- [file path | file name] : line : 

## example analyze

- unit testing
- black-box testing
- benchmark testing
- example (demo for usage)

## summary

- package log is very sample
- default logger output to stderr, format is : year/month/day hour:mimute:second `message`

