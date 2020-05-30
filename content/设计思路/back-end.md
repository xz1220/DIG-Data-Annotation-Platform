# Back-End

后端遵循MVC设计模式，采用的RestFul风格实际API。主要分为三个模块，Repository、Service、Controller。

## Repository
Repository 层利用ORM框架进行对数据表的增删改查
|Repository|Describes
|---|--------|
|AdminImage|对图片以及图片数据进行增删改查|
|AdminImageLabel|针对图片标签进行增删改查|
|AdminUser|针对用户进行增删改查|
|AdminTask|针对任务数据进行增删改查|

