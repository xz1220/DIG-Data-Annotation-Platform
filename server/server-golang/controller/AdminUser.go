package controller

import (
	"encoding/json"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/middleware"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	type userDto struct {
		Username    string `json:"username"`
		UserID      int64  `json:"userId"`
		Authorities string `json:"authorities"`
	}

	db := common.GetDB()
	cache := common.GetCache()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)

	//使用结构体
	var requestMap = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestMap) //其中一种

	user, err := adminUserReposityInstance.FindByUserName(requestMap.Username)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "用户不存在")
		return
	}

	if strings.Compare(user.Password, requestMap.Password) == -1 {
		util.ManagerInstance.FailWithoutData(ctx, "密码错误")
		return
	}
	//TOKEN
	token, err := middleware.ReleaseToken(ctx, user) //发放token
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "系统异常")
		return
	}

	cookie := http.Cookie{Name: "request_token", Value: "6MIhycayVQizGoweGhRvUFVARhAARiTyJ1NS6YNfiuQJ1ZHU", Expires: time.Now().AddDate(0, 0, 1)}
	http.SetCookie(ctx.Writer, &cookie)

	redisUtilInstance := util.RedisUtilInstance(cache)
	log.Println(ctx.Request.RemoteAddr)
	err = redisUtilInstance.AddTokenTORedis(token, requestMap.Username, ctx.Request.RemoteAddr)
	//返回结果

	// util.Success(ctx, gin.H{"user": model.ToUserDto(user), "token": token}, "SUCCESS")
	util.Success(ctx, gin.H{"user": userDto{
		Username:    user.Username,
		UserID:      user.UserID,
		Authorities: user.Authorities,
	}, "token": token}, "SUCCESS")
	// util.Success(ctx, gin.H{"token": token}, "SUCCESS")

	log.Println("登录成功！")

}

func Logout(ctx *gin.Context) {
	Authorization := ctx.GetHeader("Authorization")

	cache := common.GetCache()
	redisUtilInstance := util.RedisUtilInstance(cache)

	if Authorization != "" && strings.HasPrefix(Authorization, "Bearer ") {
		authToken := strings.TrimLeft(Authorization, "Bearer ")
		redisUtilInstance.HSet("blacklist", authToken, time.Now().String())
		log.Println("用户登出成功！Token 加入黑名单!")
	}
	util.Success(ctx, gin.H{}, "Logout Success")
}

func GetCount(ctx *gin.Context) {
	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	taskCout, _ := adminUserReposityInstance.GetTaskCount()
	reviewerCount, _ := adminUserReposityInstance.GetReviewerCount()
	userCount, _ := adminUserReposityInstance.GetUserCount()

	util.Success(ctx, gin.H{"taskCount": taskCout, "userCount": userCount, "reviewerCount": reviewerCount}, "SUCCESS")
}

// GetUserList
func GetUserList(ctx *gin.Context) {
	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, _ := adminUserReposityInstance.GetUserList()
	util.Success(ctx, gin.H{"userList": users}, "SUCCESS")
}

// EditUser
func EditUser(ctx *gin.Context) {
	var user = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&user) //其中一种
	if user.Username == "" {
		util.ManagerInstance.FailWithoutData(ctx, "Parameter Error : Bind User Wrong!!")
		return
	}
	log.Println("User Information: ", user.Username)

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	err := adminUserReposityInstance.EditUser(user)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Edit User Error!!!")
		return
	}

	log.Println("Edit User Success!!!")
	util.Success(ctx, gin.H{}, "SUCCESS")

}

// AddUser
func AddUser(ctx *gin.Context) {
	var user = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&user) //其中一种
	if user.Username == "" {
		util.ManagerInstance.FailWithoutData(ctx, "Parameter Error : Bind User Wrong!!")
		return
	}
	log.Println("User Information: ", user.Username)

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	err := adminUserReposityInstance.AddUser(user)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Add User Error!!!")
		return
	}

	log.Println("Add User Success!!!")
	util.Success(ctx, gin.H{}, "SUCCESS")

}

// DeleteUser
func DeleteUser(ctx *gin.Context) {
	type Temp struct {
		UserID int64 `json:"userId"`
	}
	var tempuser = Temp{}
	json.NewDecoder(ctx.Request.Body).Decode(&tempuser) //其中一种
	if tempuser.UserID == 0 {
		util.ManagerInstance.FailWithoutData(ctx, "Parameter Error : Bind User Wrong!!")
		return
	}
	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	user, err := adminUserReposityInstance.GetUserByID(tempuser.UserID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Parameter Error : Can't Find the User By ID!!")
		return
	}
	log.Println("User Information: ", user.Username)

	err = adminUserReposityInstance.DeleteUser(user.UserID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Delete User Error!!!")
		return
	}

	log.Println("Delete User Success!!!")
	util.Success(ctx, gin.H{}, "SUCCESS")

}

func GetPendingUserList(ctx *gin.Context) {
	type data struct {
		ImageID string `json:"imageId"`
	}

	var pendingData data
	_ = ctx.ShouldBindJSON(&pendingData)

	imageID, err := strconv.ParseInt(pendingData.ImageID, 10, 64)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Convert string to int Error!!!")
		return
	}

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, err := adminUserReposityInstance.GetPendingUserList(imageID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Get PendingUser List Error!!!")
		return
	}

	log.Println("Delete User Success!!!")
	util.Success(ctx, users, "SUCCESS")
}

func GetVideoPendingUserList(ctx *gin.Context) {
	type data struct {
		ImageID string `json:"imageId"`
	}

	var pendingData = data{}
	json.NewDecoder(ctx.Request.Body).Decode(&pendingData) //其中一种

	imageID, err := strconv.ParseInt(pendingData.ImageID, 10, 64)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Convert string to int Error!!!")
		return
	}

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, err := adminUserReposityInstance.GetVideoPendingUserList(imageID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Get VideoPendingUser List Error!!!")
		return
	}

	log.Println("Delete User Success!!!")
	util.Success(ctx, users, "SUCCESS")
}

func GetListUser(ctx *gin.Context) {
	log.Println("Get List User")

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, err := adminUserReposityInstance.GetListUser()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Get List User Error!!!")
		return
	}

	log.Println("Get List User Success!!!")
	util.Success(ctx, users, "SUCCESS")
}

func GetListReviewer(ctx *gin.Context) {
	log.Println("Get List User")

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, err := adminUserReposityInstance.GetListReviewer()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Get List User Error!!!")
		return
	}

	log.Println("Get List User Success!!!")
	util.Success(ctx, users, "SUCCESS")
}
