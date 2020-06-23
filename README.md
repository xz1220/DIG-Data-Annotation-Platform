# DIG Data Annotation Platform

![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)

文档主要用来详细说明前后端的设计实现思路，便于后来者快速上手进行二次开发和维护。

## 部署

### java 版本

#### 前期准备

1. 确保安装docker 以及 docker-compose
2. 克隆前端库，创建镜像
```
git clone https://github.com/xz1220/labelproject-foreground-spring.git
// cd src/model/ && vim Service.js && cnpm install && cnpm run build // 修改HOST 对应后端IP地址 
docker build -t <image_name> .
// vim compose/labelproject-java.yml // 修改compose配置文件，修改 web-fore.image 为新创建镜像
```

#### 一键部署
```
docker-compose -f compose/labelproject-java.yml up 
```
*特点*
- mysql 容器启动后 数据表自动创建，绑定主机 3306 端口
- labelproject-back (sping 后端程序) 容器启动后 图片数据存放目录自动创建, 自动连接mysql数据库与redis数据库，绑定主机8887端口
- 前端绑定8889端口


## 设计思路
这一部分我们详细说明前后的设计思路。主要说明前后端内各个模块的作用。以及前后端API接口。
### 前端

### 后端

## 实现

### 前端
基于Vue.js 实现
### 后端
后端部分有两个实现版本，基于SpringBoot的JAVA版本和基于GIN框架的Golang版本，二者设计思路完全一致，是上述设计的不同实现。
#### 基于SpringBoot

### 基于GIN

## 基础知识
记录一些背景知识以及学习资料，非必看。