# jww

https://github.com/spf13/jwalterweatherman

jwalterweatherman,孟加拉语,翻译是火的热度

- base on package log
- log level:
    - LevelTrace 
    - LevelDebug    
    - LevelInfo     
    - LevelWarn     
    - LevelError    
    - LevelCritical 
    - LevelFatal    

## source analyze

- jww is simple and have a lot fun
- log output somewhere, jww output multiple place
- jww have 8 log.Logger, one is for log, other is for log level
- jww has a defaultNotePad
    - jww.DEBUG.Println()
    - jww.FEEDBACK.Println()
