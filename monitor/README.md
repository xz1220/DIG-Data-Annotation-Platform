# Monitor

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)

使用docker-compose 一键部署Linux监控平台，改编自[容器可视化监控解决方案](https://segmentfault.com/a/1190000015548444)

**使用容器组件如下：**

- InfluxDB : 数据存储
- Telegraf ：数据采集
- Chronigraf : 可视化web UI
- Kapacitor : 监控、告警

## Usage

```
docker-compose -f monitor.yml up
```



## Contribution

原文对每一个组件使用单个的docker命令来安装，较为不方便，对于配置文件仍然需要手动进行配置，以及组件容器相互之间的调试需要测试，所以对源文件进行了改进，使用docker-compose 进行部署，在保证端口和网段不冲突的情况下，无需额外配置，一键部署即可。

