package main

import (
	"labelproject-back/common"
	"labelproject-back/controller"
	"labelproject-back/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CollectRoute(r *gin.Engine, foreIP string) *gin.Engine {
	// r.Use(middleware.CORSMiddleware())
	// r.POST("/api/auth/register", controller.Regsiter)

	r.StaticFS("/api/image", http.Dir("/home/kiritoghy/labelprojectdata/image"))
	r.StaticFS("/api/thumb", http.Dir("/home/kiritoghy/labelprojectdata/images"))
	r.StaticFS("/api/video", http.Dir("/home/kiritoghy/labelprojectdata/video"))
	r.StaticFS("/api/videos", http.Dir("/home/kiritoghy/labelprojectdata/videos"))

	r.Use(middleware.CORSMiddleware(foreIP)).POST("/api/login", controller.Login)
	r.Use(middleware.CORSMiddleware(foreIP)).GET("/api/logout", controller.Logout)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).GET("/api/admin/getCount", controller.GetCount)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).GET("/api/admin/getUserList", controller.GetUserList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/editUser", controller.EditUser)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/addUser", controller.AddUser)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/deleteUser", controller.DeleteUser)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getPendingUserList", controller.GetPendingUserList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getVideoPendingUserList", controller.GetVideoPendingUserList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getListUser", controller.GetListUser)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getListReviewer", controller.GetListReviewer)

	//TaskController
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getTaskList", controller.GetTaskList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/updateTaskType", controller.UpdateTaskType)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/updateTask", controller.UpdateTask)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/deleteTask", controller.DeleteTask)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/splitTask", controller.SplitTask)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).GET("/api/admin/getNewTaskList", controller.GetNewTaskList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/searchTask", controller.SearchTask)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).GET("/api/admin/taskList", controller.TaskList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/downloadDatas", controller.DownloadData)

	//AdminIMageController
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getImgList", controller.GetImageList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/getImg", controller.GetImg)

	//admminImageLabelController
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).GET("/api/admin/getLabelList", controller.GetLabelList)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/addLabel", controller.AddLabel)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/editLabel", controller.EditLabel)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/deleteLabel", controller.DeleteLabel)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/saveLabel", controller.SaveLabel)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/deleteImageById", controller.DeleteImageByID)
	r.Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth()).POST("/api/admin/setFinalVersion", controller.SetFinalVersion)
	// r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}

func main() {
	time.Sleep(5000000000)
	common.InitConfig("main")
	common.InitDB()
	db := common.GetDB()
	cache := common.GetCache()
	log.Println("successfully")
	defer db.Close()
	defer cache.Close()

	// var user model.User
	// // db.Where("Username = ?", "admin").First(&user)
	// db.Where("username = ?", "admin").First(&user)
	// // db.First(&user, 10)
	// fmt.Println(user)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r = CollectRoute(r, "http://localhost:9999")

	// r = CollectRoute(r, "http://127.0.0.1:9999")

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}
