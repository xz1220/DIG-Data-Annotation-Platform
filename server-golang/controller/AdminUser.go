package controller

import (
	"encoding/json"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/middleware"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type userDto struct {
	Username    string `json:"username"`
	UserID      int64  `json:"userid"`
	Authorities string `json:"authorities"`
}

// Login check the username & password. If everything is right, writing successful message to context.
func Login(ctx *gin.Context) {

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
	redisUtilInstance := util.RedisUtilInstance(common.GetCache())

	var requestMap = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	user, err := adminUserReposityInstance.FindByUserName(requestMap.Username)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "用户不存在")
		return
	}

	if strings.Compare(user.Password, requestMap.Password) == -1 {
		util.ManagerInstance.FailWithoutData(ctx, "密码错误")
		return
	}

	token, err := middleware.ReleaseToken(ctx, user)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "系统异常")
		return
	}

	// cookie := http.Cookie{Name: "request_token", Value: "6MIhycayVQizGoweGhRvUFVARhAARiTyJ1NS6YNfiuQJ1ZHU", Expires: time.Now().AddDate(0, 0, 1)}
	// http.SetCookie(ctx.Writer, &cookie)

	err = redisUtilInstance.AddTokenTORedis(token, requestMap.Username, ctx.Request.RemoteAddr)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "add token to redis error")
		return
	}

	util.Success(ctx, gin.H{"user": userDto{
		Username:    user.Username,
		UserID:      user.UserID,
		Authorities: user.Authorities,
	}, "token": token}, "SUCCESS")

	log.Println("登录成功！")
}

// Logout retract authorization from request and put it into blacklist in Redis.
func Logout(ctx *gin.Context) {

	Authorization := ctx.GetHeader("Authorization")
	redisUtilInstance := util.RedisUtilInstance(common.GetCache())

	if Authorization != "" && strings.HasPrefix(Authorization, "Bearer ") {
		authToken := strings.TrimLeft(Authorization, "Bearer ")
		redisUtilInstance.HSet("blacklist", authToken, time.Now().String())
		log.Println("用户登出成功！Token 加入黑名单!")
	}
	util.Success(ctx, gin.H{}, "Logout Success")
}

// GetCount returns amount of tasks, users, and reviewers.
// And it ignores the possible errors from the repository layer.
func GetCount(ctx *gin.Context) {
	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
	taskCout, _ := adminUserReposityInstance.GetTaskCount()
	reviewerCount, _ := adminUserReposityInstance.GetReviewerCount()
	userCount, _ := adminUserReposityInstance.GetUserCount()

	util.Success(ctx, gin.H{"taskCount": taskCout, "userCount": userCount, "reviewerCount": reviewerCount}, "SUCCESS")
}

// GetUserList returns content of users with paramater userlist in the response.
func GetUserList(ctx *gin.Context) {
	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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
	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
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

	adminUserReposityInstance := repository.AdminUserReposityInstance(common.GetDB())
	users, err := adminUserReposityInstance.GetListReviewer()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Get List User Error!!!")
		return
	}

	log.Println("Get List User Success!!!")
	util.Success(ctx, users, "SUCCESS")
}
