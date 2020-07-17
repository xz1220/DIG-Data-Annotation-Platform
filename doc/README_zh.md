# DIG Data Annotation Platform

![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

基于前后端分离的数据标注平台与容器监控系统，支持docker-compose 快速一键部署

##  :crystal_ball: **Visuals**

**Annotation Platform**

![Annotation-Platform](../doc/Annotation-Platform.png)



**Monitor**

![monitor](../doc/monitor.png)

##  🍕 **Requirements**

### Monitor

- docker-ce
- docker-compose

### Annotation Platform

#### SpringBoot+Vue.js

- jdk >=1.8
- Mysql Version == 5.7 or 8.0

#### Gin + Vue.js

- Golang version >= 1.13
- Gin v1
- Gorm v1



##  🚍 **Installation**

### 🚀 Quick Start

####  Annotation Platform

**Preparation**

- 确保安装docker 以及 docker-compose

- 克隆前端库，创建镜像

```shell
git clone https://github.com/xz1220/labelproject-foreground-spring.git
cd src/model/ && vim Service.js // 修改HOST 对应后端IP地址 
cnpm install && cnpm run build 
docker build -t <image_name> .
vim compose/labelproject-<java/golang>.yml // 修改compose配置文件，修改 web-fore.image 为新创建镜像，按需修改容器volume
```

**Installation By docker-compose**

```shell
docker-compose -f compose/labelproject-<java/golang>.yml up // 后端端口绑定8887 前端端口绑定8889 
```
##### Features

- mysql 容器启动后 数据表自动创建，绑定主机 3306 端口
- labelproject-back (sping 后端程序) 容器启动后 图片数据存放目录自动创建, 自动连接mysql数据库与redis数据库，绑定主机8887端口

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

