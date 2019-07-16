---
    hugo go实现的静态站点生产工具
---

## 介绍/使用

- [hello hugo](/hugo/hello-hugo.md)

# 源码分析

分析脚程：
- [第1个pr,添加项目介绍文件,第一个最小实现](/hugo/source/1.md),这是第一次主干更新,实现了一个hugo的最小集
- [第2个pr,将配置和front matter的格式由json改为yaml](/hugo/source/2.md),将配置和front matter的默认格式改为yaml
- [第3个pr,格式支持toml/yaml/json](/hugo/source/3.md),配置文件支持多种格式,启动参数支持长短格式,部分启动参数的功能也做了添加
- [各个地方的优化和bug修复](/hugo/source/4.md),content目录在代码和概念上统一,添加输出目录,用于将static和输出分开
- [10 pr, 解决windows上section标签导致的异常](/hugo/source/5.md), 解决不同平台url分隔符的异常,解决无content时的异常
- [前一页后一页,索引支持](/hugo/source/6.md), 发布0.8.0,修复bug,添加新功能:支持索引,前一页后一页




使用到的库:
- [github.com/pkg/errors](/pion-webrtc/lib/pkg-errors.md)


思想、读后感：


技术知识：
