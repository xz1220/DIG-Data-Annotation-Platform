# DIG Data Annotation Platform

![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

基于前后端分离的数据标注平台与容器监控系统，支持docker-compose 快速一键部署
> [English](./doc/README_En.md) | 中文

##  :crystal_ball: **Visuals**

**Annotation Platform**

<img src="./doc/Annotation-Platform.png" alt="Annotation-Platform" style="zoom:50%;" />

**Architecture**

<img src="./doc/server-golang.png" alt="server-golang" style="zoom:50%;" />


**Monitor**

<img src="./doc/monitor.png" alt="monitor" style="zoom:50%;" />

##  🍕 **Requirements**

### Monitor

- docker-ce
- docker-compose

### Annotation Platform

#### Go + Vue.js

- Golang version >= 1.13
- Gin v1
- Gorm v1
- Mysql Version == 5.7 or 8.0


##  🚍 **Installation**

### 🚀 Quick Start (local)

####  Annotation Platform

**Preparation**

- 确保安装docker 以及 docker-compose

- 克隆代码库

```shell
git clone https://github.com/xz1220/DIG-Data-Annotation-Platform.git
# 修改前端配置并运行
cd DIG-Data-Annotation-Platform/front-end/src/model/ && vim Service.js // 修改HOST 对应后端IP地址 
cnpm install && cnpm run build 
# 修改后端配置并运行
cd DIG-Data-Annotation-Platform/server-golang/ && vim main.go
# 修改第107行 r := CollectRoute(gin.New(), "http://127.0.0.1:9999")， 将IP替换为前端IP
docker-compose -f ./doc/labelproject-golang.yml # 启动mysql & redis 镜像
go run main.go # 启动后端程序
```

**Installation By docker-compose**
在front-end和server-golang的目录下，都存放着Dockerfile文件，方便容器化前后端。可自定义修改labelproject-golang.yml文件，实现一键部署。
```shell
docker build -t <your imageName:tag> .
```
##### Features



#### Monitor 

**Preparation** 

- 确保安装docker 以及 docker-compose

**Installation**

```shell
git clone https://github.com/xz1220/LabelDoc.git 
cd LabelDoc/monitor
docker-compose -f monitor.yml up
```



##  🚩 **Usage**

#### 🖼 Annotation Platform

- 初始化用户名：admin 密码：admin

### 🖥 Monitor

- 入口 ： http://localhost:8888
- 初始化数据库
  - URL：http://172.23.0.2:8086
  - 用户名免密为空
- 选取默认面板进入系统



## Reference

[Docker Document](https://docs.docker.com/)

[Golang Document](https://golang.org/doc/)

