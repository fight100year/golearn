# vsoch/pull-request-action

- docker action
- 向某些分支(带有某些前缀)push时，会给某些分支(默认是master)自动创建一个pr
- 适用场景：仓库有workflow，如果想让push的数据最后review并合并到master，就可以使用这个action

## 总结

- 感觉使用的场景太少了
- 而且入参需要配置的东西太多(很多都固定死了),使用并不灵活
- 这个就是检测到push，就创建一个pr
- 后期有遇到类似需求可以再看看
