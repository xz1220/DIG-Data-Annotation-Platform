package controller

import (
	"encoding/json"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
* taskList
* getImgList
* getImg
* saveLabel
 */

func TaskListUser(ctx *gin.Context) {
	type data struct {
		UserID string `json:"userId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	UserID, err := util.String2Int64(tempData.UserID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "TaskID string2int error!!!")
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.TaskList()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "TaskList error!!!")
		return
	}

	var taskResponses []model.TaskResponse
	for _, task := range tasks {
		userIDs, err := adminTaskRepositoryInstance.GetUserIDsFromUserInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.FailWithoutData(ctx, "SearchTask error!!!")
			return
		}

		reviewersIDs, err := adminTaskRepositoryInstance.GetReviewerIDsFromReviewerInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.FailWithoutData(ctx, "GetReviewerIDsFromReviewerInfo Error")
			return
		}

		labelIDs, err := adminTaskRepositoryInstance.GetLabelIDsFromLabelInfo(task.TaskID)
		if err != nil {
			util.ManagerInstance.FailWithoutData(ctx, "GetLabelIDsFromLabelInfo Error")
			return
		}

		if IfInSlice(UserID, userIDs) {
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

	}

	util.Success(ctx, taskResponses, "SUCCESS")
}

func GetImgListUser(ctx *gin.Context) {
	type data struct {
		TaskID string `json:"taskId"`
		Page   int64  `json:"page"`
		Limit  int64  `json:"limit"`
		UserID string `json:"userId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	log.Println("TaskID:", tempData.TaskID, "  Page", tempData.Page, "  Limit:", tempData.Limit, "  AdminID:", tempData.UserID)
	// util.Success(ctx, gin.H{}, "SUCCESS")

	taskID, err := util.String2Int64(tempData.TaskID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "TaskID string2int error!!!")
		return
	}

	// adminID, err := util.String2Int64(tempData.AdminID)
	// if err != nil {
	// 	ErrorString := ctx.Request.URL.String() + "AdminID string2int error!!!"
	// 	log.Println(ErrorString)
	// 	util.Fail(ctx, gin.H{}, ErrorString)
	// 	return
	// }

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)

	taskName, err := adminTaskRepositoryInstance.GetTaskNameByID(taskID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "GetTaskNameByID error!!!")
		return
	}

	if strings.Compare("", taskName) == 0 {
		util.ManagerInstance.FailWithoutData(ctx, "任务不存在!!!")
		return
	}

	imageList, err := adminImageRepositoryInstance.GetImageList(taskID)

	log.Println("taskID:", taskID, "  has", len(imageList), " images")
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "GetImageList Error!!!")
		return
	}

	//page
	var newImageList []*model.Image
	totalpages := (int64(len(imageList)) + tempData.Limit) / tempData.Limit
	if totalpages == tempData.Page {
		newImageList = imageList[(tempData.Page-1)*tempData.Limit:]
	} else {
		newImageList = imageList[(tempData.Page-1)*tempData.Limit : (tempData.Page)*tempData.Limit]
	}

	//有图片
	fileUtilInstance := util.FileUtilInstance()
	if len(newImageList) > 0 {
		for _, image := range newImageList {
			src := fileUtilInstance.IMAGE_DIC + taskName + "/" + image.ImageName
			dest := fileUtilInstance.IMAGE_S_DIC + taskName

			if image.ImageThumb == "" {
				thumb, width, height, err := fileUtilInstance.Thumb(src, dest, image.ImageName)
				if err != nil {
					util.ManagerInstance.FailWithoutData(ctx, "GetImageList Error!!!")
					return
				}
				image.ImageThumb = thumb
				image.Width = int64(width)
				image.Height = int64(height)
			}

		}

		err = adminImageRepositoryInstance.UpdateImages(newImageList)
		if err != nil {
			util.ManagerInstance.FailWithoutData(ctx, "GetImageList Error!!!")
			return
		}
	}

	labelImageIDs, err := adminImageRepositoryInstance.GetLabeledImageIDs(taskID, 0)

	var dataMap map[string]interface{}
	dataMap = make(map[string]interface{})
	dataMap["page"] = tempData.Page
	dataMap["limit"] = tempData.Limit
	dataMap["totalpages"] = totalpages
	dataMap["images"] = newImageList

	if labelImageIDs != nil {
		dataMap["labelImageIds"] = labelImageIDs
	} else {
		dataMap["labelImageIds"] = []int64{}
	}

	util.Success(ctx, dataMap, "SUCCESS")
}

func GetImgUser(ctx *gin.Context) {
	type data struct {
		ImageID string `json:"imageId"`
		UserID  int64  `json:"userId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	imageID, err := util.String2Int64(tempData.ImageID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "imageID string2int error!!!")
		return
	}

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	image, err := adminImageRepositoryInstance.GetImage(imageID)
	if err != nil || image.ImageID == 0 {
		log.Println("Error:  ", err)
		util.ManagerInstance.FailWithoutData(ctx, "   Get Image Error error!!!")
		return
	}

	// adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	// taskName, err := adminTaskRepositoryInstance.GetTaskNameByImageID(imageID)

	// fileUtilInstance := util.FileUtilInstance()
	// src := fileUtilInstance.IMAGE_DIC+taskName

	/** remove the limitation of image && delete some unavailing code **/

	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	labels, err := adminImageLabelInstance.GetLabelByImageID(imageID)
	dataList, err := adminImageRepositoryInstance.GetDatas(tempData.UserID, imageID)
	if len(dataList) == 0 {
		dataList = make([]*model.DataForResponse, 0)
	}

	var dataMap map[string]interface{}
	dataMap = make(map[string]interface{})
	dataMap["labels"] = labels
	dataMap["image"] = image
	dataMap["datas"] = dataList

	util.Success(ctx, dataMap, "SUCCESS")
}

func SaveLabelUser(ctx *gin.Context) {
	var tempData model.LabelData
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.UserID == 0 {
		util.ManagerInstance.FailWithoutData(ctx, " --- "+ctx.Request.Method+"Bind Data Error!!!")
		return
	}

	imageID, err := util.String2Int64(tempData.ImageIDString)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " --- "+ctx.Request.Method+"imageID string2int error!!!")
		return
	}

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	dataIDs, err := adminImageRepositoryInstance.GetDataIDs(tempData.UserID, imageID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " --- "+ctx.Request.Method+"GetDataIDs error!!!")
		return
	}

	err = adminImageRepositoryInstance.SaveLabel(tempData, dataIDs)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " --- "+ctx.Request.Method+"SaveLabel error!!!")
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}
