# cds-snc/github-actions/seekret

- docker action
- seekret是一个go实现的命令行工具，也提供了库
- 通过seekret可以扫描github仓库和目录中的敏感信息
- 这个action就是通过seekret工具来查找可能泄漏安全信息

## 仓库分析

- readme是常规操作
- 没有license文件 
- 是一个标准的docker action
- 元数据并不是通过action.yml展示的，而是写在Dockerfile中

```Dockerfile
FROM cdssnc/seekret:5b62fc5

LABEL "name"="Seekret"
LABEL "maintainer"="Max Neuvians <max.neuvians@cds-snc.ca>"
LABEL "version"="1.0.0"

LABEL "com.github.actions.name"="Seekret"
LABEL "com.github.actions.description"="Uses seekret to scan for secrets in your code"
LABEL "com.github.actions.icon"="lock"
LABEL "com.github.actions.color"="orange"

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["bash", "/entrypoint.sh"]
```

## action 分析

- 入口是一个shell命令

```shell
#!/bin/bash
data=()

echo ""
echo "Searching for secrets ..."

while IFS= read -r line; do
  data+=( "$line" )
done < <( sh -c "seekret $* dir $GITHUB_WORKSPACE" )

if [ ${#data[@]} -eq 1 ]; then
    echo ""
    echo "No secrets found!"
    exit 0
else
    echo ""
    echo "Found the following secrets:"
    echo "---------------------------"
    printf '%s\n' "${data[@]}"
    echo ""
    exit 1
fi
```

## 使用

```yaml
# 最简单的使用 

steps:
- uses: cds-snc/github-actions/seekret@master
```

## 总结

- action没有测试 workflow
- 这个action还是非常有用的，可以在每个仓库中都添加上这个
