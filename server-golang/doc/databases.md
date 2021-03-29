#数据库设计
##1. user
|字段名|描述
|---|--------|
|user_id|用户id|
|username|用户名|
|password|密码|
|authorities|权限|

##2. task
|字段名|描述
|---|--------|
|task_id|任务id|
|task_name|任务名称|
|task_desc|任务描述|
|image_number|图片数量

##3. images
|字段名|描述
|---|--------|
|image_id|图片id|
|image_name|图片名|
|image_thumb|缩略图名|
|task_id|任务id|
|user_confirm_id|确认id|
|width|
|height|

##4. imagelabel
|字段名|描述
|---|--------|
|label_id|标签id|
|label_name|标签名称
|label_type|标签类型
|label_color|标签颜色

>表明有哪些标签

##5. data
|字段名|描述
|---|--------|
|data_id|数据id
|image_id|图片id
|label_id|标签id
|user_id|用户id
|data_desc|描述
|label_type|
|iscrowd|

##6.userfinished 用户完成情况
|字段名|描述
|---|--------|
|user_id|用户id|
|task_id|任务id
|image_id|图片id

##7.taskuserinfo
|字段名|描述
|---|--------|
|task_id|任务id
|user_id|用户id

##8.taskreviewerinfo
|字段名|描述
|---|--------|
|task_id|任务id
|reviewer_id(user_id)|审核id

##9. tasklabelinfo
|字段名|描述
|---|--------|
|task_id|任务id
|label_id|标签id
##10. datapoints
|字段名|描述
|---|--------|
|data_id|数据id
|order|顺序
|x|x
|y|y
|image_id|
|user_id|
##11. datarle
|字段名|描述
|---|--------|
|data_id|数据id
|image_id|
|data_rle|
|user_id|







