/*
** Task Controller
** GetTaskList
** UpdateTaskType
** UpdateTask
** DeleteTask
** SplitTask
** GetNewTaskList
** SearchTask
** downloadTask : not completes
 */

package controller

import (
	"encoding/json"
	"fmt"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/model"
	"log"
	"os"
	"strings"

	"labelproject-back/util"

	"github.com/gin-gonic/gin"
)

// type AdminTaskController interface {
// 	GetTaskList(ctx *gin.Context)
// 	UpdateTaskType(ctx *gin.Context)
// 	UpdateTask(ctx *gin.Context)

// 	SplitTask(ctx *gin.Context)
// 	DeleteTask(ctx *gin.Context)
// 	GetNewTaskList(ctx *gin.Context)
// 	SearchTask(ctx *gin.Context)
// }

// type adminTaskController struct {
// }

// var adminTaskControllerInstance = adminTaskController{}

// func AdminTaskControllerInstance() AdminTaskController {
// 	return adminTaskControllerInstance
// }

// GetTaskList Return Task List
func GetTaskList(ctx *gin.Context) {
	// type data struct {
	// 	Page  int64 `json:"page"`
	// 	Limit int64 `json:"limit"`
	// }

	// type response struct {
	// 	TaskID      int64         `json:"taskId"`
	// 	TaskName    string        `json:"taskName"`
	// 	TaskDesc    string        `json:"taskDesc"`
	// 	ImageNumber int64         `json:"imageNumber"`
	// 	TaskType    int64         `json:"taskType"`
	// 	Finish      int64         `json:"finish"`
	// 	UserIDs     []int64       `json:"userIds"`
	// 	ReviewerIDs []int64       `json:"reviewerIds"`
	// 	LabelIDs    []int64       `json:"labelIds"`
	// 	Users       []*model.User `json:"users"`
	// 	Reviewer    []*model.User `json:"reviewers"`
	// }

	var temp = model.PageData{}
	json.NewDecoder(ctx.Request.Body).Decode(&temp)
	if temp.Page == 0 {
		log.Println("Bind Error!!!")
		util.Fail(ctx, gin.H{}, "Bind Error!!!")
		return
	}
	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	tasks, err := adminTaskRepositoryInstance.GetTaskList()
	if err != nil {
		util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get Task List Error!!!")
		util.Fail(ctx, gin.H{}, "Get Task List Error!!!")
		return
	}

	var responseTemps []model.TaskResponse
	for _, task := range tasks {
		userInfos, err := adminTaskRepositoryInstance.GetUserInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get UserInfo List Error!!!")
			util.Fail(ctx, gin.H{}, "Get UserInfo List Error!!!")
			return
		}

		var users []*model.User
		var userIds []int64
		userIds = []int64{}
		for _, userInfo := range userInfos {
			userIds = append(userIds, userInfo.UserID)
			user, err := adminUserReposityInstance.GetUserByID(userInfo.UserID)
			if err != nil {
				util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get User Error!!!")
				util.Fail(ctx, gin.H{}, "Get User Error!!!")
				return
			}
			users = append(users, &user)
		}

		reviewerInfos, err := adminTaskRepositoryInstance.GetReviewerInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get ReviewerInfo List Error!!!")
			util.Fail(ctx, gin.H{}, "Get ReviewerInfo List Error!!!")
			return
		}

		var reviewers []*model.User
		var reviewerIds []int64
		reviewerIds = []int64{}
		for _, userInfo := range reviewerInfos {
			reviewerIds = append(reviewerIds, userInfo.ReviewerID)
			user, err := adminUserReposityInstance.GetUserByID(userInfo.ReviewerID)
			if err != nil {
				util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get Reviewers Error!!!")
				util.Fail(ctx, gin.H{}, "Get Reviewers Error!!!")
				return
			}
			reviewers = append(reviewers, &user)
		}

		labelinfos, err := adminTaskRepositoryInstance.GetLabelInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.SendError(ctx.Request.URL.String(), "Get LabelInfo List Error!!!")
			util.Fail(ctx, gin.H{}, "Get LabelInfo List Error!!!")
			return
		}
		var labelIDs []int64
		labelIDs = []int64{}
		for _, labelinfo := range labelinfos {
			labelIDs = append(labelIDs, labelinfo.LabelID)
		}
		var responseTemp = model.TaskResponse{
			TaskID:      task.TaskID,
			TaskName:    task.TaskName,
			TaskDesc:    task.TaskDesc,
			TaskType:    task.TaskType,
			Finish:      task.Finish,
			ImageNumber: task.ImageNumber,
			UserIDs:     userIds,
			ReviewerIDs: reviewerIds,
			LabelIDs:    labelIDs,
			/**
			*  不得不说接口设计的糟糕，文档中未注明需要一下两个字段，但是在实际前后段交互的时候还是出现了，并且是nil值。
			*  如果返回了一下两个字段的真实值，那么无法显示TaskList，且未报错
			**/
			// Users:       users,
			// Reviewer:    reviewers,
		}

		responseTemps = append(responseTemps, responseTemp)
	}

	// responseJson, _ := json.Marshal(responseTemps)
	// log.Println("Response Json :", string(responseJson))
	totalpages := (int64(len(responseTemps)) + temp.Limit) / temp.Limit
	if temp.Page == totalpages {
		responseTemps = responseTemps[(temp.Page-1)*30:]
	} else {
		responseTemps = responseTemps[(temp.Page-1)*30 : temp.Page*30]
	}
	util.Success(ctx, gin.H{"limit": temp.Limit, "page": temp.Page, "totalpages": totalpages, "taskList": responseTemps}, "SUCCESS")
}

// UpdateTaskType
func UpdateTaskType(ctx *gin.Context) {
	type temp struct {
		TaskID   int64 `json:"taskId"`
		TaskType int64 `json:"taskType"`
	}

	var tempData temp
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	if isHasData, err := adminTaskRepositoryInstance.HasData(tempData.TaskID); err != nil {
		log.Println("Get Task Label Datas Error!!!")
		util.Fail(ctx, gin.H{}, "Get  Task Label Datas Error!!!")
		return
	} else if isHasData != 0 {
		log.Println("Error: Task already has the label datas, Please delete it!!!")
		util.Fail(ctx, gin.H{}, "Error: Task already has the label datas, Please delete it!!!")
		return
	}

	err := adminTaskRepositoryInstance.UpdateTaskType(tempData.TaskID, tempData.TaskType)
	if err != nil {
		log.Println()
		util.ManagerInstance.SendError(ctx.Request.URL.String(), "Update Task Error!!!")
		util.Fail(ctx, gin.H{}, "Update Task Error!!!")
		return
	}

	log.Println("Update Task Successfully !!!")
	util.Success(ctx, gin.H{}, "SUCCESS")
}

func UpdateTask(ctx *gin.Context) {
	var taskRequest model.TaskRequest
	json.NewDecoder(ctx.Request.Body).Decode(&taskRequest)

	if taskRequest.UserIDs == nil || taskRequest.ReviewerIDs == nil || taskRequest.LabelIDs == nil {
		log.Println("Bind Error!!!")
		util.Fail(ctx, gin.H{}, "Bind Error!!!")
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	oldTask, err := adminTaskRepositoryInstance.GetTaskByID(taskRequest.TaskID)
	if err != nil {
		util.ManagerInstance.SendError(ctx.Request.URL.String(), "Find Task By ID Error!!!")
		util.Fail(ctx, gin.H{}, "Find Task By ID Error!!!")
		return
	}

	oldUserInfo, _ := adminTaskRepositoryInstance.GetUserInfo(taskRequest.TaskID)
	oldReviewerInfo, _ := adminTaskRepositoryInstance.GetReviewerInfo(taskRequest.TaskID)
	oldLabel, _ := adminTaskRepositoryInstance.GetLabelInfo(taskRequest.TaskID)

	if isHasData, _ := adminTaskRepositoryInstance.HasData(taskRequest.TaskID); isHasData != 0 {
		if len(taskRequest.LabelIDs) != len(oldLabel) {
			ErrorString := ctx.Request.URL.String() + "Error: Task already has the label datas, Please delete it!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		for index, label := range taskRequest.LabelIDs {
			if label != oldLabel[index].LabelID {
				ErrorString := ctx.Request.URL.String() + "Error: Task already has the label datas, Please delete it!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}
	}

	// delete all && add all
	if len(oldLabel) > 0 {
		err = adminTaskRepositoryInstance.DeleteTaskLabelIDs(oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Delete Label!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	if len(oldReviewerInfo) > 0 {
		err = adminTaskRepositoryInstance.DeleteTaskReviewerIDs(oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Delete Reviewer!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	if len(oldUserInfo) > 0 {
		err = adminTaskRepositoryInstance.DeleteTaskUserIDs(oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Delete Users!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	if len(taskRequest.UserIDs) > 0 {
		err = adminTaskRepositoryInstance.AddTaskUserIds(taskRequest.UserIDs, oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Add Users!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	if len(taskRequest.ReviewerIDs) > 0 {
		err = adminTaskRepositoryInstance.AddTaskReviewerIDs(taskRequest.ReviewerIDs, oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Add Reviewer!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	if len(taskRequest.LabelIDs) > 0 {
		err = adminTaskRepositoryInstance.AddTaskLabelIDs(taskRequest.LabelIDs, oldTask.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Error: Add Labels!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
	}

	// 修改文件名
	fileUtilInstance := util.FileUtilInstance()
	if strings.Compare(taskRequest.TaskName, oldTask.TaskName) != 0 {
		err = os.Rename(fileUtilInstance.IMAGE_DIC+oldTask.TaskName, fileUtilInstance.IMAGE_DIC+taskRequest.TaskName)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Rename Task Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
		oldTask.TaskName = taskRequest.TaskName
	}

	//update Task
	oldTask.TaskDesc = taskRequest.TaskDesc
	err = adminTaskRepositoryInstance.UpdateTask(oldTask)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Error: updates Task!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")

}

func DeleteTask(ctx *gin.Context) {

	type data struct {
		TaskID int64 `json:"taskId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	if tempData.TaskID == 0 {
		ErrorString := ctx.Request.URL.String() + "Bind Parameter Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	task, err := adminTaskRepositoryInstance.GetTaskByID(tempData.TaskID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Get Task Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)

	if task.TaskType == 1 || task.TaskType == 2 || task.TaskType == 3 || task.TaskType == 4 {

		imageIDs, err := adminImageRepositoryInstance.GetImageIDs(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Get Image By TaskID Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		dataIDs, err := adminImageRepositoryInstance.GetDataIDByImageID(imageIDs)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Get DataIDs By ImageID Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if len(imageIDs) > 0 {
			err = adminImageRepositoryInstance.DeleteImagesByTaskID(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Images By TaskID Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if len(imageIDs) > 0 && len(dataIDs) > 0 {
			err = adminImageRepositoryInstance.DeleteDatasByImageID(imageIDs)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Datas By ImageIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if len(dataIDs) > 0 {
			err = adminImageRepositoryInstance.DeletePoints(dataIDs)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Points By DataIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		err = adminImageRepositoryInstance.DeleteFinish(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Delete Finished Data By TaskID Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if users, _ := adminTaskRepositoryInstance.GetUserInfo(tempData.TaskID); len(users) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskUserIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task UserIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if reviewers, _ := adminTaskRepositoryInstance.GetReviewerInfo(tempData.TaskID); len(reviewers) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskReviewerIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task ReviewersIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if labels, _ := adminTaskRepositoryInstance.GetLabelInfo(tempData.TaskID); len(labels) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskLabelIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task LabelIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		err = adminTaskRepositoryInstance.DeleteTask(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Delete Task Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		fileUtilInstance := util.FileUtilInstance()
		src := fileUtilInstance.IMAGE_DIC + task.TaskName
		thmub := fileUtilInstance.IMAGE_S_DIC + task.TaskName
		dest := fileUtilInstance.IMAGE_DELETE_DIC + task.TaskName

		err = os.Rename(src, dest)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		os.RemoveAll(thmub)
		os.RemoveAll(src)
	} else if task.TaskType == 5 {

		adminVideoRepositoryInstance := repository.AdminVideoRepositoryInstance(db)
		videoIDs, err := adminVideoRepositoryInstance.GetVideoIDs(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Get VideoIDs By TaskID Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		dataIDs, err := adminImageRepositoryInstance.GetImageIDs(tempData.TaskID)
		log.Println("May Error Occur, adminImageRepository ---- videoDataIDs")
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Get Video Data IDs By TaskID Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if len(videoIDs) > 0 {
			if err = adminVideoRepositoryInstance.DeleteVideoByTaskID(tempData.TaskID); err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Video By TaskID Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if len(videoIDs) > 0 && len(dataIDs) > 0 {
			if err = adminVideoRepositoryInstance.DeleteDatasByTaskID(videoIDs); err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Video Datas by VideoIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if err = adminImageRepositoryInstance.DeleteFinish(tempData.TaskID); err != nil {
			ErrorString := ctx.Request.URL.String() + "Delete Video  Finish Datas Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return

		}

		if users, _ := adminTaskRepositoryInstance.GetUserInfo(tempData.TaskID); len(users) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskUserIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task UserIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if reviewers, _ := adminTaskRepositoryInstance.GetReviewerInfo(tempData.TaskID); len(reviewers) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskReviewerIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task ReviewersIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		if labels, _ := adminTaskRepositoryInstance.GetLabelInfo(tempData.TaskID); len(labels) > 0 {
			err = adminTaskRepositoryInstance.DeleteTaskLabelIDs(tempData.TaskID)
			if err != nil {
				ErrorString := ctx.Request.URL.String() + "Delete Task LabelIDs Error!!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
		}

		err = adminTaskRepositoryInstance.DeleteTask(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "Delete Task Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		fileUtilInstance := util.FileUtilInstance()
		src := fileUtilInstance.VIDEO_DIC + task.TaskName
		thmub := fileUtilInstance.VIDEO_S_DIC + task.TaskName
		dest := fileUtilInstance.VIDEO_D_DIC + task.TaskName

		os.Rename(src, dest)
		os.RemoveAll(thmub)
		os.RemoveAll(src)

	}

	log.Println("Delete Task Successfully")
	util.Success(ctx, gin.H{}, "SUCCESS")
}

func SplitTask(ctx *gin.Context) {
	type temp struct {
		TaskId   int64 `json:"taskId"`
		Quantity int64 `json:"quantity"`
	}

	var tempData temp
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	if tempData.TaskId == 0 {
		ErrorString := ctx.Request.URL.String() + "Bind Parameter Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	rx := db.Begin()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(rx)
	task, err := adminTaskRepositoryInstance.GetTaskByID(tempData.TaskId)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Get Task By ID Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	if task.ImageNumber < tempData.Quantity {
		ErrorString := ctx.Request.URL.String() + "Split Task Error: Task Number is more than image Number!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	newTaskImageNumber := task.ImageNumber / tempData.Quantity
	log.Println("task.ImageNumber:", task.ImageNumber, "  tempData.Quantity", tempData.Quantity)
	lastTaskImageNumber := task.ImageNumber - ((tempData.Quantity - 1) * newTaskImageNumber)

	// Get All Images
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(rx)
	images, err := adminImageRepositoryInstance.GetImageList(task.TaskID)

	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Get Image List Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}
	rx.Commit()

	log.Println("总共有", len(images), "张图片", " newTaskImageNumber:", newTaskImageNumber, "  lastTaskImageNumber", lastTaskImageNumber)

	fileUtilInstance := util.FileUtilInstance()
	if task.TaskType == 1 || task.TaskType == 2 || task.TaskType == 3 || task.TaskType == 4 {
		// 获取文件路径 判断是否存在以及是否是文件夹
		taskDic := fileUtilInstance.IMAGE_DIC + task.TaskName
		thumbTaskDic := fileUtilInstance.IMAGE_S_DIC + task.TaskName

		if exit, err := PathExists(taskDic); err != nil && !exit {
			ErrorString := ctx.Request.URL.String() + "Split Task Error : Task Not Exit!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		var index int64
		for index = 1; index <= tempData.Quantity; index++ {
			newTaskName := task.TaskName + fmt.Sprintf("_part%d", index)

			// Calculate Image Number of New Task
			// Get New Slice of Images
			var taskImageNumber int64
			if index < tempData.Quantity {
				taskImageNumber = newTaskImageNumber
			} else {
				taskImageNumber = lastTaskImageNumber
			}
			newImageList := images[(index-1)*newTaskImageNumber : ((index-1)*newTaskImageNumber + taskImageNumber)]

			// Create New Task Based On Old Task && Update NewTask
			newTask := task
			newTask.TaskName = newTaskName
			newTask.ImageNumber = int64(len(newImageList))
			rx = db.Begin()
			adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(rx)
			lastRecord, err := adminTaskRepositoryInstance.LastRecord()
			if err != nil {
				// ErrorString := ctx.Request.URL.String() + "Find Last Record Error!!!"
				// log.Println(ErrorString)
				// util.Fail(ctx, gin.H{}, ErrorString)
				// return
				newTask.TaskID = 1
			} else {
				newTask.TaskID = lastRecord.TaskID + 1

			}

			log.Println("NewTask ID:", newTask.TaskID)

			log.Println("length of newImageList:", len(newImageList))
			for _, image := range newImageList {
				// 移动文件 效率可能可以提升
				// log.Println("taskName:", newTaskName)
				// log.Println("  Image.ImageName", image.ImageName)
				dest := fileUtilInstance.IMAGE_DIC + newTaskName + "/" + image.ImageName
				imageFile := fileUtilInstance.IMAGE_DIC + task.TaskName + "/" + image.ImageName
				// log.Println("ImageFile:", imageFile, "  dest:", dest)
				err = os.Rename(imageFile, dest)
				if err != nil {
					log.Println(err)

					// } else if os.IsNotExist(err) {
					log.Println("尝试创建", dest)
					err = os.Mkdir(fileUtilInstance.IMAGE_DIC+newTaskName, os.ModePerm)
					if err != nil {
						log.Println("创建失败")
						ErrorString := ctx.Request.URL.String() + "创建目录失败 !!!"
						log.Println(ErrorString)
						util.Fail(ctx, gin.H{}, ErrorString)
						return
					}
					os.Rename(imageFile, dest)
				}

				if exit, _ := PathExists(thumbTaskDic); exit {
					if image.ImageThumb != "" {
						// Move thum file to new directory
						dest := fileUtilInstance.IMAGE_S_DIC + newTaskName + "/" + image.ImageThumb
						imageFile := fileUtilInstance.IMAGE_S_DIC + task.TaskName + "/" + image.ImageThumb

						err = os.Rename(imageFile, dest)
						if err != nil {
							log.Println(err)
							ErrorString := ctx.Request.URL.String() + "重命名失败 !!!"
							log.Println(ErrorString)
							// util.Fail(ctx, gin.H{}, ErrorString)
							// return
							// } else if os.IsNotExist(err) {
							err = os.Mkdir(fileUtilInstance.IMAGE_S_DIC+newTaskName, os.ModePerm)
							if err != nil {
								log.Println("创建失败")
								ErrorString := ctx.Request.URL.String() + "创建目录失败 !!!"
								log.Println(ErrorString)
								util.Fail(ctx, gin.H{}, ErrorString)
								return
							}
							err = os.Rename(imageFile, dest)
						}
					}
				}

				image.TaskID = newTask.TaskID

			}

			tx := db.Begin()
			adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(tx)
			if err = adminTaskRepositoryInstance.AddTask(newTask); err != nil {
				tx.Rollback()
				ErrorString := ctx.Request.URL.String() + "Add Task Error !!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
			tx.Commit()

			tx = db.Begin()
			adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
			if userIDs, err := adminTaskRepositoryInstance.GetUserIDsFromUserInfo(task.TaskID); err == nil && len(userIDs) > 0 {
				err = adminTaskRepositoryInstance.AddTaskUserIds(userIDs, newTask.TaskID)
				if err != nil {
					tx.Rollback()
					ErrorString := ctx.Request.URL.String() + "Add UserIDs To NewTask Error !!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}
			}
			tx.Commit()

			tx = db.Begin()
			adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
			if userIDs, err := adminTaskRepositoryInstance.GetReviewerIDsFromReviewerInfo(task.TaskID); err == nil && len(userIDs) > 0 {
				err = adminTaskRepositoryInstance.AddTaskReviewerIDs(userIDs, newTask.TaskID)
				if err != nil {
					tx.Rollback()
					ErrorString := ctx.Request.URL.String() + "Add ReviewerIDs To NewTask Error !!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}
			}
			tx.Commit()

			tx = db.Begin()
			adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
			if userIDs, err := adminTaskRepositoryInstance.GetLabelIDsFromLabelInfo(task.TaskID); err == nil && len(userIDs) > 0 {
				err = adminTaskRepositoryInstance.AddTaskLabelIDs(userIDs, newTask.TaskID)
				if err != nil {
					tx.Rollback()
					ErrorString := ctx.Request.URL.String() + "Add LabelIDs To NewTask Error !!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}
			}
			tx.Commit()

			tx = db.Begin()
			adminImageRepositoryInstance = repository.AdminImageRepositoryInstance(tx)
			if err = adminImageRepositoryInstance.UpdateImagesTaskID(newImageList, newTask.TaskID); err != nil {
				tx.Rollback()
				ErrorString := ctx.Request.URL.String() + "Update NewTask Error !!!"
				log.Println(ErrorString)
				util.Fail(ctx, gin.H{}, ErrorString)
				return
			}
			tx.Commit()

		}

		// Delete All The Information Of Old Task
		tx := db.Begin()
		adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
		if err = adminTaskRepositoryInstance.DeleteTask(task.TaskID); err != nil {
			tx.Rollback()
			ErrorString := ctx.Request.URL.String() + "Update NewTask Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
		tx.Commit()

		tx = db.Begin()
		adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
		if err = adminTaskRepositoryInstance.DeleteTaskUserIDs(task.TaskID); err != nil {
			tx.Rollback()
			ErrorString := ctx.Request.URL.String() + "Update NewTask Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
		tx.Commit()

		tx = db.Begin()
		adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
		if err = adminTaskRepositoryInstance.DeleteTaskReviewerIDs(task.TaskID); err != nil {
			tx.Rollback()
			ErrorString := ctx.Request.URL.String() + "Update NewTask Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
		tx.Commit()

		tx = db.Begin()
		adminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(tx)
		if err = adminTaskRepositoryInstance.DeleteTaskLabelIDs(task.TaskID); err != nil {
			tx.Rollback()
			ErrorString := ctx.Request.URL.String() + "Update NewTask Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}
		tx.Commit()

		err = os.RemoveAll(taskDic)
		if err != nil {
			log.Println(err)
		}
		err = os.RemoveAll(thumbTaskDic)
		if err != nil {
			log.Println(err)
		}
	} else if task.TaskType == 5 {

		// TODO : Create New Function To Reduce Code
		ErrorString := ctx.Request.URL.String() + "Still In progress !!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "拆分成功！")

}

func PathExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return true && file.IsDir(), nil //存在且为目录
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Get : Get New Task List
func GetNewTaskList(ctx *gin.Context) {
	db := common.GetDB()

	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)

	// Get The Names of images and videos
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	imageNames, err := adminTaskRepositoryInstance.GetImageTaskNameList()
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Get Image Task Name List Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	// videoNames, err := adminTaskRepositoryInstance.GetVideoTaskNameList()
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Get Image Task Name List Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	/** 扫描图片 **/
	// 判断是否存在图片目录
	var temp model.Task
	fileUtilInstance := util.FileUtilInstance()
	log.Println("判断是否存在图片目录")
	if exit, _ := PathExists(fileUtilInstance.IMAGE_DIC); !exit {
		ErrorString := ctx.Request.URL.String() + "Image Dic don't exit or isn't a directory!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	// 打开文件目录
	log.Println("打开文件目录")
	ImageDic, err := os.Open(fileUtilInstance.IMAGE_DIC)

	defer ImageDic.Close()
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Open Image Dic File error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	//读取目录下文件
	log.Println("读取目录下文件")
	if fileInfos, err := ImageDic.Readdir(-1); err == nil && len(fileInfos) > 0 {
		for _, fileInfo := range fileInfos {
			log.Println(fileInfo.Name())
			if fileInfo.IsDir() {
				log.Println(fileInfo.Name(), "是目录")
				newImageDic := fileUtilInstance.IMAGE_DIC + fileInfo.Name()
				newImageFile, err := os.Open(newImageDic)
				defer newImageFile.Close()
				if err != nil {
					ErrorString := ctx.Request.URL.String() + "Open Image Dic File error!!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}

				newImageList, err := newImageFile.Readdir(-1)
				// newImageListName, _ := newImageFile.Readdirnames(-1)
				log.Println(fileInfo.Name(), "内有", len(newImageList), "张图片")
				if err != nil {
					ErrorString := ctx.Request.URL.String() + "Read Image List error!!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}

				// log.Println("imageNames:", imageNames, "   fileInfo:", fileInfo.Name())
				if !stringInSlice(imageNames, fileInfo.Name()) {
					// log.Println(fileInfo.Name(), "不在", imageNames)
					temp.TaskName = fileInfo.Name()
					temp.ImageNumber = int64(len(newImageList))
					temp.TaskType = 1
					lastRecord, err := adminTaskRepositoryInstance.LastRecord()
					if err != nil {
						// ErrorString := ctx.Request.URL.String() +"get last record error!!!"
						// log.Println(ErrorString)
						// util.Fail(ctx, gin.H{}, ErrorString)
						// return
						temp.TaskID = 1 // 初始化时数据库中不存在task数据
					} else {
						temp.TaskID = lastRecord.TaskID + 1
					}

					if len(newImageList) > 0 {
						// log.Println(fileInfo.Name(), "内有图片")
						var imageList []*model.Image

						// log.Println("newImageListName:", len(newImageListName))
						// log.Println("newImageList", len(newImageList))
						for _, image := range newImageList {
							// log.Println("imageName:", image)
							if !strings.HasSuffix(image.Name(), ".jpg") && !strings.HasSuffix(image.Name(), ".JPG") && !strings.HasSuffix(image.Name(), ".jpeg") && !strings.HasSuffix(image.Name(), ".bmp") && !strings.HasSuffix(image.Name(), ".png") {
								log.Println("跳过")
								temp.ImageNumber--
								continue
							}

							// log.Println("index:", index, "total:", len(newImageList))
							var tempImage = &model.Image{
								ImageName: image.Name(),
								TaskID:    temp.TaskID,
							}

							imageList = append(imageList, tempImage)
						}

						adminImageRepositoryInstance.AddImages(imageList)
					}

					err = adminTaskRepositoryInstance.AddTask(temp)
					if err != nil {
						ErrorString := ctx.Request.URL.String() + "Add Task error!!!"
						log.Println(ErrorString)
						util.Fail(ctx, gin.H{}, ErrorString)
						return
					}
				}

			}
		}
	}

	/**
	**  go 使用 ffmpeg 需要使用cgo， 并且在编译的时候需要链接静态库
	** 另外支持音视频处理的库  ： https://github.com/nareix/joy5
	**/
	// /** 扫描视频 **/
	// if exit, _ := PathExists(fileUtilInstance.VIDEO_DIC); !exit {
	// 	ErrorString := ctx.Request.URL.String() +"Video Dic don't exit or isn't a directory!!!"
	// 	log.Println(ErrorString)
	// 	util.Fail(ctx, gin.H{}, ErrorString)
	// 	return
	// }

	// // 打开文件目录
	// VideoDic, err := os.Open(fileUtilInstance.VIDEO_DIC)
	// defer VideoDic.Close()
	// if err != nil {
	// 	ErrorString := ctx.Request.URL.String() +"Open VIDEO Dic File error!!!"
	// 	log.Println(ErrorString)
	// 	util.Fail(ctx, gin.H{}, ErrorString)
	// 	return
	// }

	// //读取目录下文件
	// if fileInfos, err := VideoDic.Readdir(-1); err != nil && len(fileInfos) > 0 {
	// 	for _, fileInfo := range fileInfos {
	// 		if fileInfo.IsDir() {
	// 			newVideoDic := fileUtilInstance.VIDEO_DIC + fileInfo.Name()
	// 			newVideoFile, err := os.Open(newVideoDic)
	// 			defer newVideoFile.Close()
	// 			if err != nil {
	// 				ErrorString := ctx.Request.URL.String() +"Open Video Dic File error!!!"
	// 				log.Println(ErrorString)
	// 				util.Fail(ctx, gin.H{}, ErrorString)
	// 				return
	// 			}

	// 			newVideoList, err := newVideoFile.Readdir(-1)
	// 			if err != nil {
	// 				ErrorString := ctx.Request.URL.String() +"Read Video List error!!!"
	// 				log.Println(ErrorString)
	// 				util.Fail(ctx, gin.H{}, ErrorString)
	// 				return
	// 			}

	// 			if !stringInSlice(videoNames, fileInfo.Name()) {
	// 				temp.TaskName = fileInfo.Name()
	// 				temp.ImageNumber = int64(len(newVideoList))
	// 				temp.TaskType = 1
	// 				lastRecord, err := adminTaskRepositoryInstance.LastRecord()
	// 				if err != nil {
	// 					ErrorString := ctx.Request.URL.String() +"get last record error!!!"
	// 					log.Println(ErrorString)
	// 					util.Fail(ctx, gin.H{}, ErrorString)
	// 					return
	// 				}
	// 				temp.TaskID = lastRecord.TaskID + 1
	// 				err = adminTaskRepositoryInstance.AddTask(temp)
	// 				if err != nil {
	// 					ErrorString := ctx.Request.URL.String() +"Add Task error!!!"
	// 					log.Println(ErrorString)
	// 					util.Fail(ctx, gin.H{}, ErrorString)
	// 					return
	// 				}

	// 				if len(newVideoList) > 0 {
	// 					var videoList []*model.Video

	// 					for _, video := range newVideoList {
	// 						if !strings.HasSuffix(video.Name(), ".jpg") && !strings.HasSuffix(video.Name(), ".jpeg") && !strings.HasSuffix(video.Name(), ".bmp") && !strings.HasSuffix(video.Name(), ".png") {
	// 							continue
	// 						}

	// 						var tempVideo *model.Video
	// 						tempVideo.VideoName = video.Name()
	// 						tempVideo.TaskID = temp.TaskID

	// 					}

	// 					adminVideoRepositoryInstance := repository.AdminVideoRepositoryInstance(db)
	// 					adminVideoRepositoryInstance.AddVideo(videoList)
	// 				}

	// 			}

	// 		}
	// 	}
	// }

	log.Println("Refresh the Task Successfully")
	tasks, err := adminTaskRepositoryInstance.GetNewTaskList()

	//TaskResponse
	var taskResponse []model.TaskResponse
	for _, task := range tasks {
		temp := model.TaskResponse{
			TaskID:      task.TaskID,
			TaskName:    task.TaskName,
			TaskDesc:    task.TaskDesc,
			ImageNumber: task.ImageNumber,
			TaskType:    task.TaskType,
		}

		taskResponse = append(taskResponse, temp)
	}

	util.Success(ctx, gin.H{"taskList": taskResponse}, "SUCCESS")

}

func stringInSlice(list []string, s string) bool {
	for _, element := range list {
		if strings.Compare(element, s) == 0 {
			return true
		}
	}
	return false
}

func SearchTask(ctx *gin.Context) {
	type data struct {
		Keyword string `json:"keyword"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	log.Println("keyword:", tempData.Keyword)

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.SearchTask(tempData.Keyword)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	var taskResponses []model.TaskResponse
	for _, task := range tasks {
		userIDs, err := adminTaskRepositoryInstance.GetUserIDsFromUserInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		reviewersIDs, err := adminTaskRepositoryInstance.GetReviewerIDsFromReviewerInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		labelIDs, err := adminTaskRepositoryInstance.GetLabelIDsFromLabelInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		temp := model.TaskResponse{
			TaskID:      task.TaskID,
			TaskType:    task.TaskType,
			TaskName:    task.TaskName,
			TaskDesc:    task.TaskDesc,
			ImageNumber: task.ImageNumber,
			UserIDs:     userIDs,
			ReviewerIDs: reviewersIDs,
			LabelIDs:    labelIDs,
			Finish:      task.Finish,
		}

		taskResponses = append(taskResponses, temp)
	}

	util.Success(ctx, taskResponses, "SUCCESS")

}

func TaskList(ctx *gin.Context) {
	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.TaskList()
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "TaskList error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	var taskResponses []model.TaskResponse
	for _, task := range tasks {
		userIDs, err := adminTaskRepositoryInstance.GetUserIDsFromUserInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		reviewersIDs, err := adminTaskRepositoryInstance.GetReviewerIDsFromReviewerInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		labelIDs, err := adminTaskRepositoryInstance.GetLabelIDsFromLabelInfo(task.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		temp := model.TaskResponse{
			TaskID:      task.TaskID,
			TaskType:    task.TaskType,
			TaskName:    task.TaskName,
			TaskDesc:    task.TaskDesc,
			ImageNumber: task.ImageNumber,
			UserIDs:     userIDs,
			ReviewerIDs: reviewersIDs,
			LabelIDs:    labelIDs,
			Finish:      task.Finish,
		}

		taskResponses = append(taskResponses, temp)
	}

	util.Success(ctx, taskResponses, "SUCCESS")

}

func DownloadData(ctx *gin.Context) {

	type data struct {
		TaskID int64 `json:"taskId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	// ctx.Header("content-type", "application/json;charset=utf-8")
	ctx.Header("Content-Type", "application/octet-stream")

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	task, err := adminTaskRepositoryInstance.GetTaskByID(tempData.TaskID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "GetTaskByID error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	adminImageLabelRepository := repository.AdminImageLabelRepositoryInstance(db)

	// log.Println("tempData.TaskID:", tempData.TaskID)

	switch task.TaskType {
	case 1:
		log.Println(" 开始下载图片数据！ case1")
		images, err := adminImageRepositoryInstance.GetImageList(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + " Case1 : GetImageList  Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if len(images) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 图片不存在!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		labels, err := adminImageLabelRepository.GetLabelByImageID(images[0].ImageID)
		if len(labels) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 标签不存在 下载失败!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		cocoDataSet := model.CocoDataSet{}
		cocoInfo := model.CocoInfo{
			Year:        0,
			DataCreated: "",
		}
		cocoDataSet.Info = cocoInfo

		cocoAnnotations := []model.CocoAnnotation{}
		cocoCategories := []model.CocoCategory{}
		cocoImages := []model.CocoImage{}

		for _, image := range images {
			log.Println("case1: 进入images循环")
			if image.UserComfirmID == 0 {
				continue
			}

			datas, _ := adminImageRepositoryInstance.GetDatas(image.UserComfirmID, image.ImageID)
			if len(datas) == 0 {
				continue
			}
			cocoImage := model.CocoImage{
				FileName: image.ImageName,
				Height:   image.Height,
				Width:    image.Width,
				ID:       image.ImageID,
			}
			cocoImages = append(cocoImages, cocoImage)

			for _, data := range datas {
				log.Println("case1: 进入datas循环")
				cocoAnnotation := model.CocoAnnotation{
					ID:         data.DataID,
					ImageID:    int64(data.ImageID),
					CategoryID: int64(data.LabelID),
				}

				cocoAnnotations = append(cocoAnnotations, cocoAnnotation)
			}
		}

		cocoDataSet.Annotations = cocoAnnotations
		cocoDataSet.Images = cocoImages

		for _, label := range labels {
			cocoCategory := model.CocoCategory{
				ID:            label.LabelID,
				Name:          label.LabelName,
				SuperCategory: label.LabelName,
			}
			cocoCategories = append(cocoCategories, cocoCategory)
		}

		cocoDataSet.Categories = cocoCategories

		filename := task.TaskName + "json"
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")

		util.Success(ctx, cocoDataSet, "SUCCESS")
		return
	case 2, 3:
		log.Println(" case2,3:开始下载图片数据！")
		images, err := adminImageRepositoryInstance.GetImageList(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + " Case1 : GetImageList  Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if len(images) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 图片不存在!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		labels, err := adminImageLabelRepository.GetLabelByImageID(images[0].ImageID)
		if len(labels) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 标签不存在 下载失败!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		cocoDataSet := model.CocoDataSet{}
		cocoInfo := model.CocoInfo{
			Year:        0,
			DataCreated: "",
		}
		cocoDataSet.Info = cocoInfo

		cocoAnnotations := []model.CocoAnnotation{}
		cocoCategories := []model.CocoCategory{}
		cocoImages := []model.CocoImage{}

		for _, image := range images {
			log.Println("case1: 进入images循环")
			if image.UserComfirmID == 0 {
				continue
			}

			datas, _ := adminImageRepositoryInstance.GetDatas(image.UserComfirmID, image.ImageID)
			if len(datas) == 0 {
				continue
			}
			cocoImage := model.CocoImage{
				FileName: image.ImageName,
				Height:   image.Height,
				Width:    image.Width,
				ID:       image.ImageID,
			}
			cocoImages = append(cocoImages, cocoImage)

			for _, data := range datas {
				log.Println("case1: 进入datas循环")
				cocoAnnotation := model.CocoAnnotation{
					ID:           data.DataID,
					ImageID:      int64(data.ImageID),
					CategoryID:   int64(data.LabelID),
					Segmentation: util.GenPolygonData(data.Point),
				}

				cocoAnnotations = append(cocoAnnotations, cocoAnnotation)
			}
		}

		cocoDataSet.Annotations = cocoAnnotations
		cocoDataSet.Images = cocoImages

		for _, label := range labels {
			cocoCategory := model.CocoCategory{
				ID:            label.LabelID,
				Name:          label.LabelName,
				SuperCategory: label.LabelName,
			}
			cocoCategories = append(cocoCategories, cocoCategory)
		}

		cocoDataSet.Categories = cocoCategories

		filename := task.TaskName + "json"
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")

		util.Success(ctx, cocoDataSet, "SUCCESS")
		return
	case 4:
		log.Println(" case4:开始下载图片数据！")
		images, err := adminImageRepositoryInstance.GetImageList(tempData.TaskID)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + " Case1 : GetImageList  Error !!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		if len(images) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 图片不存在!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		labels, err := adminImageLabelRepository.GetLabelByImageID(images[0].ImageID)
		if len(labels) == 0 {
			ErrorString := ctx.Request.URL.String() + " Case1: 标签不存在 下载失败!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
			return
		}

		cocoDataSet := model.CocoDataSet{}
		cocoInfo := model.CocoInfo{
			Year:        0,
			DataCreated: "",
		}
		cocoDataSet.Info = cocoInfo

		cocoAnnotations := []model.CocoAnnotation{}
		cocoCategories := []model.CocoCategory{}
		cocoImages := []model.CocoImage{}

		for _, image := range images {
			log.Println("case1: 进入images循环")
			if image.UserComfirmID == 0 {
				continue
			}

			datas, _ := adminImageRepositoryInstance.GetDatas(image.UserComfirmID, image.ImageID)
			if len(datas) == 0 {
				continue
			}
			cocoImage := model.CocoImage{
				FileName: image.ImageName,
				Height:   image.Height,
				Width:    image.Width,
				ID:       image.ImageID,
			}
			cocoImages = append(cocoImages, cocoImage)

			for _, data := range datas {
				log.Println("case1: 进入datas循环")
				cocoAnnotation := model.CocoAnnotation{
					ID:         data.DataID,
					ImageID:    int64(data.ImageID),
					CategoryID: int64(data.LabelID),
				}

				cocoAnnotations = append(cocoAnnotations, cocoAnnotation)
			}
		}

		cocoDataSet.Annotations = cocoAnnotations
		cocoDataSet.Images = cocoImages

		for _, label := range labels {
			cocoCategory := model.CocoCategory{
				ID:            label.LabelID,
				Name:          label.LabelName,
				SuperCategory: label.LabelName,
			}
			cocoCategories = append(cocoCategories, cocoCategory)
			cocoDataSet.Keypoints = append(cocoDataSet.Keypoints, label.LabelName)
		}

		cocoDataSet.Categories = cocoCategories

		filename := task.TaskName + "json"
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")

		util.Success(ctx, cocoDataSet, "SUCCESS")
		return
	}

	util.Fail(ctx, gin.H{}, "Fail")

	// content := "hello world, 我是一个文件，"
	// strings.NewReader(content)

	// os.Create("./hello.txt")
	// err = ioutil.WriteFile("./hello.txt", []byte(content), os.ModePerm)
	// ctx.Writer.WriteHeader(http.StatusOK)
	// ctx.Header("Content-Disposition", "attachment; filename=hello.txt")
	// ctx.Header("Content-Type", "application/text/plain")
	// ctx.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	// ctx.Header("Content-Type", "application/octet-stream")
	// ctx.Header("Expires", "0")

	// type Category struct {
	// 	ID            int64  `json:"id"`
	// 	Name          string `json:"name"`
	// 	Supercategory string `json:"supercategory"`
	// }

	// type Infor struct {
	// 	DateCreated string `json:"date_created"`
	// 	Year        int64  `json:"year"`
	// }

	// type ResponseTemp struct {
	// 	Annotations []int64    `json:"annotations"`
	// 	Categories  []Category `json:"categories"`
	// 	Images      []int64    `json:"images"`
	// 	Info        Infor
	// }

	// categories := Category{
	// 	ID:            1,
	// 	Name:          "Test1",
	// 	Supercategory: "Test1",
	// }

	// info := Infor{
	// 	DateCreated: "2020-06-22_13-12-18",
	// 	Year:        2020,
	// }

	// responseTemp := ResponseTemp{
	// 	Annotations: []int64{},
	// 	Images:      []int64{},
	// 	Info:        info,
	// }

	// responseTemp.Categories = append(responseTemp.Categories, categories)
	// util.Success(ctx, responseTemp, "SUCCESS")

}
