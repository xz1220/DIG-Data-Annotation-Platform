# DIG Data Annotation Platform

![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

A efficient Data Annotation Platform for Computer Vision Tasks with a container monitoring system.

##  :crystal_ball: **Visuals**

**Annotation Platform**

![Annotation-Platform](C:/Users/30249/LabelDoc/meida/Annotation-Platform.png)



**Monitor**

![monitor](C:/Users/30249/LabelDoc/meida/monitor.png)

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

####  Annotation Platform ： SpringBoot + Vue.js

**Preparation**

- make sure you have installed docker-ce and docker-compose

- Clone library from Github and build a new image

```shell
git clone https://github.com/xz1220/labelproject-foreground-spring.git
cd src/model/ && vim Service.js // modify HOST to the IP address of back-end 
cnpm install && cnpm run build 
docker build -t <image_name> .
vim compose/labelproject-java.yml // modify web-fore.image to the new fore-end image name
```

**Installation By docker-compose**

```shell
docker-compose -f compose/labelproject-java.yml up // back-end: bind port to 8887 fore-end: bind port to 8889 
```

##### Features

- Database (labelproject) and related tables will be created automatically after starting MYSQL container.
- labelproject-back(spring-boot) will automatically create a directory to hold the files and connect to the MYSQL and Redis after cteated.

#### Monitor 

**Preparation** 

- make sure you have installed docker-ce and docker-compose

**Installation**

```shell
git clone https://github.com/xz1220/LabelDoc.git 
cd LabelDoc/monitor
docker-compose -f monitor.yml up
```



##  🚩 **Usage**

#### 🖼 Annotation Platform

- Fore-end URL  : http://localhost:8889
- Initialized user name ：admin  password ：admin

### 🖥 Monitor

-  Fore-end URL： http://localhost:8888
-  Database initialization parameters
  - URL：http://172.23.0.2:8086
  - Username and password are empty
- Select the default dashboard

## Reference

[Docker Document](https://docs.docker.com/)

[Golang Document](https://golang.org/doc/)

