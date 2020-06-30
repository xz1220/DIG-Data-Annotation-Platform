# labelproject
## 开发环境  
idea2019.3  
jdk1.8  
SpringBoot + SpringSecurity + redis + mysql8 + mybatis + hikari线程池 + maven  

## 使用
### 1. 配置环境
 - 安装mysql5.7或8.0 并在properties中设置好连接参数
 - 安装redis 并在properties中设置好连接参数
 - 如果想要修改任务目录，请修改util/FileUtil中的目录地址
 - properties中的参数根据环境可自行配置
### 2. 使用
clone代码  
使用idea导入maven项目  
等待maven下载好依赖的包  
开始运行  

## 项目结构
main  
├── java  
│   └── com  
│       └── uestc  
│           └── labelproject  
│               ├── config  
│               │   ├── CustomAccessDeniedHandler.java  
│               │   ├── CustomAuthenticationEntryPoint.java  
│               │   ├── CustomAuthenticationFailureHandler.java  
│               │   ├── CustomAuthenticationProvider.java  
│               │   ├── CustomAuthenticationSuccessHandler.java  
│               │   ├── CustomLogoutSuccessHandler.java  
│               │   ├── GlobalCorsConfig.java  
│               │   ├── JwtAuthenticationTokenFilter.java  
│               │   ├── RbacAuthorityService.java  
│               │   ├── WebMvcConfig.java  
│               │   └── WebSecurityConfig.java  
│               ├── controller  
│               │   ├── Admin  
│               │   │   ├── AdminImageController.java  
│               │   │   ├── AdminLabelController.java  
│               │   │   ├── AdminTaskController.java  
│               │   │   └── AdminUserController.java  
│               │   ├── ReviewerController.java  
│               │   ├── TestLogController.java  
│               │   └── UserController.java  
│               ├── dao  
│               │   ├── AdminImageMapper.java  
│               │   ├── AdminLabelMapper.java  
│               │   ├── AdminTaskMapper.java  
│               │   ├── AdminUserMapper.java  
│               │   ├── TestMapper.java  
│               │   └── UserMapper.java  
│               ├── entity  
│               │   ├── CocoAnnotation.java  
│               │   ├── CocoCategory.java  
│               │   ├── CocoDataSet.java  
│               │   ├── CocoImage.java  
│               │   ├── CocoInfo.java  
│               │   ├── Data.java  
│               │   ├── Image.java  
│               │   ├── Label.java  
│               │   ├── Point.java  
│               │   ├── RleData.java  
│               │   ├── SelfUserDetails.java  
│               │   ├── Task.java  
│               │   ├── TempRleData.java  
│               │   ├── Test.java  
│               │   ├── UserInfo.java  
│               │   └── User.java  
│               ├── LabelprojectApplication.java  
│               ├── service  
│               │   ├── AdminImageService.java  
│               │   ├── AdminLabelService.java  
│               │   ├── AdminTaskService.java  
│               │   ├── AdminUserService.java  
│               │   ├── impl  
│               │   │   ├── AdminImageServiceImpl.java  
│               │   │   ├── AdminLabelServiceImpl.java  
│               │   │   ├── AdminTaskServiceImpl.java  
│               │   │   ├── AdminUserServiceImpl.java  
│               │   │   ├── SelfUserServiceImpl.java  
│               │   │   └── TestServiceImpl.java  
│               │   └── TestService.java  
│               └── utils  
│                   ├── AccessAddressUtil.java  
│                   ├── DataGeneratorUtil.java  
│                   ├── DateUtil.java  
│                   ├── FileUtil.java  
│                   ├── JwtTokenUtil.java  
│                   ├── LogUtil.java  
│                   ├── RedisUtil.java  
│                   ├── ResultGenerator.java  
│                   ├── Result.java  
│                   ├── StringToList.java  
│                   └── StringUtil.java  
└── resources  
    ├── application-dev.properties  
    ├── application-proc.properties  
    ├── application.properties  
    ├── image  
    ├── jwt.jks  
    ├── labelproject.sql  
    ├── logback-spring.xml  
    ├── mapper   
    │   ├── AdminImageMapper.xml  
    │   ├── AdminLabelMapper.xml  
    │   ├── AdminTaskMapper.xml  
    │   ├── AdminUserMapper.xml  
    │   ├── TestMapper.xml  
    │   └── UserMapper.xml  
    ├── META-INF  
    │   └── MANIFEST.MF  
    └── thumb  
  
17 directories, 76 files  
