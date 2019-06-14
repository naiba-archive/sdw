# 简单网址导航

[![Go Report Card](https://goreportcard.com/badge/github.com/naiba/sdw)](https://goreportcard.com/report/github.com/naiba/sdw) [![Size](https://images.microbadger.com/badges/image/naiba/sdw.svg)](https://microbadger.com/images/naiba/sdw) [![Pulls](https://img.shields.io/docker/pulls/naiba/sdw.svg)](https://microbadger.com/images/naiba/sdw)

:earth_americas: 一个简单的网址导航程序。

## 部署

1. 创建数据目录 `data`
2. 创建配置文件 `data/config.yml` 内容参考代码里面的 `data/config.yml.example`
3. Docker 部署 `docker run -d --restart always --name sdw -p 8083:8080 -v path-to-data:/sdw/data naiba/sdw`

## 感谢

[@gld](https://github.com/gldvip)