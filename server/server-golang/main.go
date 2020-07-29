package main

import (
	"labelproject-back/common"
	"labelproject-back/controller"
	"labelproject-back/middleware"
	"labelproject-back/util"
	"labelproject-back/ws"
	"log"
	"net/http"

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

	admin := r.Group("/api/admin").Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth())
	{
		admin.GET("/getCount", controller.GetCount)
		admin.GET("/getUserList", controller.GetUserList)
		admin.POST("/editUser", controller.EditUser)
		admin.POST("/addUser", controller.AddUser)
		admin.POST("/deleteUser", controller.DeleteUser)
		admin.POST("/getPendingUserList", controller.GetPendingUserList)
		admin.POST("/getVideoPendingUserList", controller.GetVideoPendingUserList)
		admin.POST("/getListUser", controller.GetListUser)
		admin.POST("/getListReviewer", controller.GetListReviewer)

		//TaskController
		admin.POST("/getTaskList", controller.GetTaskList)
		admin.POST("/updateTaskType", controller.UpdateTaskType)
		admin.POST("/updateTask", controller.UpdateTask)
		admin.POST("/deleteTask", controller.DeleteTask)
		admin.POST("/splitTask", controller.SplitTask)
		admin.GET("/getNewTaskList", controller.GetNewTaskList)
		admin.POST("/searchTask", controller.SearchTask)
		admin.GET("/taskList", controller.TaskList)
		admin.POST("/downloadDatas", controller.DownloadData)

		//AdminIMageController
		admin.POST("/getImgList", controller.GetImageList)
		admin.POST("/getImg", controller.GetImg)

		//admminImageLabelController
		admin.GET("/getLabelList", controller.GetLabelList)
		admin.POST("/addLabel", controller.AddLabel)
		admin.POST("/editLabel", controller.EditLabel)
		admin.POST("/deleteLabel", controller.DeleteLabel)
		admin.POST("/saveLabel", controller.SaveLabel)
		admin.POST("/deleteImageById", controller.DeleteImageByID)
		admin.POST("/setFinalVersion", controller.SetFinalVersion)
		// r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

		admin.GET("/getVideoLabelList", controller.GetVideoLabelList)
		admin.POST("/addVideoLabel", controller.AddVideoLabel)
		admin.POST("/editVideoLabel", controller.EditVideoLabel)
		admin.POST("/deleteVideoLabel", controller.DeleteVideoLabel)
	}

	//ReviewerController

	reviewer := r.Group("/api/reviewer").Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth())
	{
		reviewer.POST("/taskList", controller.TaskListReviewer)
		reviewer.POST("/getImgList", controller.GetImageListReviewer)
		reviewer.POST("/getImg", controller.GetImgReviewer)
		reviewer.POST("/getPendingUserList", controller.GetPendingUserListReviewer)
		reviewer.POST("/saveLabel", controller.SaveLabelReviewer)
		reviewer.POST("/setFinalVersion", controller.SetFinalVersionReviewer)
	}

	user := r.Group("/api/user").Use(middleware.CORSMiddleware(foreIP)).Use(middleware.JwtAuth())
	{
		user.POST("/taskList", controller.TaskListUser)
		user.POST("/getImgList", controller.GetImgListUser)
		user.POST("/getImg", controller.GetImgUser)
		user.POST("/saveLabel", controller.SaveLabelUser)
	}

	return r
}

func main() {
	// go ws.WebsocketManager.Start()
	// go ws.WebsocketManager.SendService()
	// go ws.WebsocketManager.SendService()
	// go ws.WebsocketManager.SendGroupService()
	// go ws.WebsocketManager.SendGroupService()
	// go ws.WebsocketManager.SendAllService()
	// go ws.WebsocketManager.SendAllService()
	// go ws.TestSendGroup()
	// go ws.TestSendAll()

	// time.Sleep(10000000000)

	go util.ManagerInstance.Start()

	common.InitConfig("main")
	common.InitDB()
	db := common.GetDB()
	cache := common.GetCache()
	log.Println("successfully")
	defer db.Close()
	defer cache.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r = CollectRoute(r, "http://localhost:9999")
	r.GET("/sockjs-node", ws.WebsocketManager.WsClient)

	// r = CollectRoute(r, "http://127.0.0.1:9999")

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}
