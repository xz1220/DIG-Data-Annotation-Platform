#API
##1. 通用url
管理员发送的请求统一为 /api/admin/\*\*  
用户发送的请求统一为   /api/user/\*\*  
审核员发送的请求统一为  /api/reviewer/\*\*  
登录直接为/login  
登出直接为/logout  
登录成功后会返回token，以后的请求需携带token来验证权限  
需要在request的请求头的Authorization中添加Bearer Token  
例如Bearer eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJhZG1pbiIsImlwIjoiMDowOjA6MDowOjA6MDoxIiwiZXhwIjoxNTYzOTcxNDAyfQ.cBHAqfgW_BkCm4Z9UbsWq4180XL-mlCgF0Y78H2ARVc4EXvr4adnuzqEfjd2CGvvFMyrx12K5o_-ZfVutfT1nMPERKAIi-q32fPqHMAlb029cKa_xbD-SKe_L_nRr0iVDeGHTGfBoerfmSxmS0j2iHeNKTsptebB75zRsY4tAYw
token有效期为1小时（可根据需要修改），持续时间为7天（可修改）  
失效后该token会加入黑名单，如果该token在持续时间内，则刷新token，在response的header中的Authorization获取新的token，若token失效且token不在持续时间内，则需重新登录获取token   
当token失效时会在response的Authorization添加Bearer Token来返回，需要从response中获取该token然后添加在request中

图片原文件路径 "/home/kiritoghy/labelprojectdata/image/" + taskName + imageName  
图片缩略图路径 "/home/kiritoghy/labelprojectdata/images/" + taskName + tmageThumb

##2. 后台返回格式
返回为json格式的数据


|返回参数|描述|  
|------|----|  
|code|返回状态代码|  
|message|返回信息|  
|data|返回数据|  

###1. 登录返回
登录成功后,会在data里携带token，以json的格式返回
```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJhZG1pbiIsImlwIjoiMDowOjA6MDowOjA6MDoxIiwiZXhwIjoxNTYzOTcxNDAyfQ.cBHAqfgW_BkCm4Z9UbsWq4180XL-mlCgF0Y78H2ARVc4EXvr4adnuzqEfjd2CGvvFMyrx12K5o_-ZfVutfT1nMPERKAIi-q32fPqHMAlb029cKa_xbD-SKe_L_nRr0iVDeGHTGfBoerfmSxmS0j2iHeNKTsptebB75zRsY4tAYw"
    },
    "message": "SUCCESS"
}
```
需要保存该token
###2. 传递方法
前后台数据
前台登录页面以post方法通过表单传递用户名密码  
其余无数据时采用get方法  
有数据时以json的格式通过post方法传递数据

##3. API
###1. ADMIN  
####(1)User
#####getUserList
前台  
url:/api/admin/getUserList  
method:get  
后台
成功
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "userList": [
            {
                "userId": 1,
                "username": "admin",
                "password": "admin",
                "authorities": "ROLE_ADMIN"
            },
            {
                "userId": 2,
                "username": "username",
                "password": "password",
                "authorities": "ROLE_USER"
            },
            {
                "userId": 3,
                "username": "reviewer",
                "password": "reviewer",
                "authorities": "ROLE_REVIEWER"
            }
        ]
    }
}
```
失败返回对应的失败信息
例如
```
{
    "code": 500,
    "message": "用户未认证"
}
```
#####editUser
前台
url:/api/admin/editUser  
method:post  
携带数据  
```
{
	"userId":2,
	"username":"username",
	"password":"editpassword",
	"authorities":"ROLE_USER"
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####addUser
前台  
url:/api/admin/addUser  
method:post  
携带数据
```
{
	"username":"test2",
	"password":"password2",
	"authorities":"ROLE_REVIEWER"
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####deleteUser
前台  
url:/api/admin/deleteUser  
method:post  
携带数据
```
{
    "userId": xxx
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####getPendingUSerList  
前台  
url:/api/admin/getPendingUserList  
method:post
```
{
    "imageId":xxx
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": [
        {
            "userId": 2,
            "username": "username",
            "password": "editpassword",
            "authorities": "ROLE_USER"
        },
        {
            "userId": 4,
            "username": "test1",
            "password": "password2",
            "authorities": "ROLE_USER"
        }
    ]
}
```
#####getListUser
前台  
url:/api/admin/getListUser  
method:post
```
{
    "taskId":xxx
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data":[
        {
            "userId": 2,
            "username": "username",
            "password": "editpassword",
            "authorities": "ROLE_USER"
        },
        {
            "userId": 4,
            "username": "test1",
            "password": "password2",
            "authorities": "ROLE_USER"
        }
    ],
```
#####getListReviewer
前台   
url:/api/admin/getListUser  
method:post  
method:post  
```
{
    "taskId":xxx
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": [
        {
            "userId": 3,
            "username": "reviewer",
            "password": "reviewer",
            "authorities": "ROLE_REVIEWER"
        }
    ]
}
```
####(2)Label
#####getLabelList
前台   
url:/api/admin/getLabelList  
method:get  
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "labelList": [
            {
                "labelId": 1,
                "labelName": "Car",
                "labelType": 0,
                "labelColor": "#3498db"
            },
            {
                "labelId": 2,
                "labelName": "People",
                "labelType": 1,
                "labelColor": "#2ecc71"
            }...
        ]
    }
}
```
#####editLabel
前台  
url:/api/admin/editLabel  
method:post  
```
{
	"labelId":3,
	"labelName": "Tree",
	"labelColor": "#1abc9c",
	"labelType": 0
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####addLabel
前台  
url:/api/admin/addLabel  
method:post
```
{
	"labelName":"LabeladdTest1",
	"labelType":1,
	"labelColor":"#2C5A78"
}
```
后台
```
{
    "code": 500,
    "message": "标签已存在，添加失败",
    "data": null
}
```
#####deleteLabel
前台  
url:/api/admin/deleteLabel  
method:post
```
{
	"labelId":8
}
```
后台
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
####(3)Task
#####getTaskList
前台  
url:/api/admin/getTaskList  
method:get  
后台  
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "taskList": [
            {
                "taskId": 112,
                "taskName": "wlop3",
                "taskDesc": null,
                "imageNumber": 118,
                "userIds": [
                    1,
                    2,
                    3,
                    4
                ],
                "reviewerIds": [
                    3,
                    4,
                    5
                ],
                "labelIds": [
                    5,
                    6,
                    7
                ]
            },
            {
                "taskId": 113,
                "taskName": "task1",
                "taskDesc": null,
                "imageNumber": 6,
                "userIds": [],
                "reviewerIds": [],
                "labelIds": []
            },
            {
                "taskId": 117,
                "taskName": "task2",
                "taskDesc": null,
                "imageNumber": 1,
                "userIds": [],
                "reviewerIds": [],
                "labelIds": []
            }
        ]
    }
}
```
#####updateTask  
url:api/admin/updateTask  
method:post  
```
{
	"taskId":112,
	"taskName":"wlop3",
	"userIds":[1,2,3,4],
	"labelIds":[5,6,7],
	"reviewerIds":[3,4,5],
	"taskDesc":"test"
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####deleteTask  
url:/api/admin/deleteTask  
method:post  
```
{
	"taskId":114
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
####(4)Image  
#####getImgList  
url:/api/admin/getImgList  
method:post  
```
{
	"taskId":112,
	"page":2,
    ""
	"limit":10
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "images": [
            {
                "imageId": 10106,
                "imageName": "7.jpg",
                "imageThumb": "thumb_7.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10107,
                "imageName": "Queens_8k.jpg",
                "imageThumb": "thumb_Queens_8k.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10108,
                "imageName": "37.jpg",
                "imageThumb": "thumb_37.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10109,
                "imageName": "ceremonial2.jpg",
                "imageThumb": "thumb_ceremonial2.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10110,
                "imageName": "5 (2).jpg",
                "imageThumb": "thumb_5 (2).jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10111,
                "imageName": "38(1).jpg",
                "imageThumb": "thumb_38(1).jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10112,
                "imageName": "5.jpg",
                "imageThumb": "thumb_5.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10113,
                "imageName": "40.jpg",
                "imageThumb": "thumb_40.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10114,
                "imageName": "20.jpg",
                "imageThumb": "thumb_20.jpg",
                "taskId": 112,
                "userConfirmId": null
            },
            {
                "imageId": 10115,
                "imageName": "17.jpg",
                "imageThumb": "thumb_17.jpg",
                "taskId": 112,
                "userConfirmId": null
            }
        ],
        "limit": 10,
        "page": 2,
        "totalpages": 12
    }
}
```
#####saveLabel  
url:/adi/admin/saveLabel  
method:Post
```
{
	"userId":4,
	"imageId":10096,
    "type":1,
	"data":[{
		"labelId":1,
		"dataDesc":"test",
        "labelType":1,
		"point":[{
			"x":1,
			"y":2,
			"order":1
		},{
			"x":1.2,
			"y":2.3,
			"order":2
		}]
	},
	{
		"labelId":2,
		"dataDesc":"test2",
        "labelType":1,
		"point":[{
			"x":1.1,
			"y":2.2,
			"order":1
		},{
			"x":1.3,
			"y":2.4,
			"order":2
		}]
	}]
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####getImg
url:/api/admin/getImg  
method:post
```
{
    "imageId":xxx
    "userId":xxx
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "image": {
            "imageId": 10096,
            "imageName": "sin2_8k.jpg",
            "imageThumb": "thumb_sin2_8k.jpg",
            "taskId": 112,
            "userConfirmId": null
        },
        "datas": [
            {
                "dataId": 51,
                "imageId": 10096,
                "labelId": 1,
                "userId": 2,
                "dataDesc": "test",
                "point": [
                    {
                        "x": 1,
                        "y": 2,
                        "order": 1
                    },
                    {
                        "x": 1.2,
                        "y": 2.3,
                        "order": 2
                    }
                ]
            },
            {
                "dataId": 52,
                "imageId": 10096,
                "labelId": 2,
                "userId": 2,
                "dataDesc": "test2",
                "point": [
                    {
                        "x": 1.1,
                        "y": 2.2,
                        "order": 1
                    },
                    {
                        "x": 1.3,
                        "y": 2.4,
                        "order": 2
                    }
                ]
            }
        ],
        "labels": [
            {
                "labelId": 1,
                "labelName": "Car",
                "labelType": 0,
                "labelColor": "#3498db"
            },
            {
                "labelId": 2,
                "labelName": "People",
                "labelType": 1,
                "labelColor": "#2ecc71"
            },
            {
                "labelId": 3,
                "labelName": "Tree",
                "labelType": 0,
                "labelColor": "#1abc9c"
            }
        ]
    }
}
```

#####saveFinalVersion
url:/api/admin/setFinalVersion  
method:post
```
{
    "imageId":xxxx
    "userConfirmId":xxxx
}
```
```
{
    "code": 200,
    "message": "SUCCESS",
    "data": null
}
```
#####downloadDatas
url:/api/admin/downloadDatas  
method:post
```
{
"taskId":xxx
}
```
#####splitTask
url:/api/admin/splitTask  
method:post
```
{
    taskId: taskId,
    quantity: quantity
}
```
####统计接口
url:/api/admin/getCount  
method:get  
{
    "code": 200,
    "message": "SUCCESS",
    "data": {
        "taskCount": 3,
        "userCount": 6,
        "reviewerCount": 2
    }
}



###2.  User
getTaskList  
getImgList  
```
{
	"taskId":112,
	"page":2,
	"limit":10,
	"userId":2
}
```
getImg  
saveLabel  
同Admin  
将url改为/api/user/xxx

###3.  Reviewer
getTaskList  
getImgList  
getImg  
saveLabel  
同Admin  
将url改为/api/reviewer/xxx