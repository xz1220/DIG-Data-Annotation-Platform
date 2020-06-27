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
* getPendingUserList
* saveLabel
* setFinalVersion
 */

func GetImageListReviewer(ctx *gin.Context) {
	type data struct {
		TaskID     string `json:"taskId"`
		Page       int64  `json:"page"`
		Limit      int64  `json:"limit"`
		ReviewerID string `json:"reviewerId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	log.Println("TaskID:", tempData.TaskID, "  Page", tempData.Page, "  Limit:", tempData.Limit, "  AdminID:", tempData.ReviewerID)
	// util.Success(ctx, gin.H{}, "SUCCESS")

	taskID, err := util.String2Int64(tempData.TaskID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "TaskID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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
		ErrorString := ctx.Request.URL.String() + "GetTaskNameByID error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	if strings.Compare("", taskName) == 0 {
		ErrorString := ctx.Request.URL.String() + "任务不存在!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	imageList, err := adminImageRepositoryInstance.GetImageList(taskID)

	log.Println("taskID:", taskID, "  has", len(imageList), " images")
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "GetImageList Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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
					ErrorString := ctx.Request.URL.String() + "GetImageList Error!!!"
					log.Println(ErrorString)
					util.Fail(ctx, gin.H{}, ErrorString)
					return
				}
				image.ImageThumb = thumb
				image.Width = int64(width)
				image.Height = int64(height)
			}

		}

		err = adminImageRepositoryInstance.UpdateImages(newImageList)
		if err != nil {
			ErrorString := ctx.Request.URL.String() + "GetImageList Error!!!"
			log.Println(ErrorString)
			util.Fail(ctx, gin.H{}, ErrorString)
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

func TaskListReviewer(ctx *gin.Context) {

	type data struct {
		ReviewerID string `json:"reviewerId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	ReviewerID, err := util.String2Int64(tempData.ReviewerID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "TaskID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	taskIDs, err := adminTaskRepositoryInstance.GetTaskIDByReviewerID(ReviewerID)
	log.Println("len(TaskIDs):", len(taskIDs))
	tasks, err := adminTaskRepositoryInstance.TaskListByID(taskIDs)
	log.Println("len(Tasks):", len(tasks))
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

		// reviewersIDs, err := adminTaskRepositoryInstance.GetReviewerIDsFromReviewerInfo(task.TaskID)
		// if err != nil {
		// 	ErrorString := ctx.Request.URL.String() + "SearchTask error!!!"
		// 	log.Println(ErrorString)
		// 	util.Fail(ctx, gin.H{}, ErrorString)
		// 	return
		// }

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
			// ReviewerIDs: reviewersIDs,
			LabelIDs: labelIDs,
			// Finish:      task.Finish,
		}

		taskResponses = append(taskResponses, temp)
	}

	util.Success(ctx, taskResponses, "SUCCESS")

}

func GetImgReviewer(ctx *gin.Context) {
	type data struct {
		ImageID string `json:"imageId"`
		UserID  int64  `json:"userId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	imageID, err := util.String2Int64(tempData.ImageID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "imageID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	image, err := adminImageRepositoryInstance.GetImage(imageID)
	if err != nil || image.ImageID == 0 {
		log.Println("Error:  ", err)
		ErrorString := ctx.Request.URL.String() + "   Get Image Error error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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

func GetPendingUserListReviewer(ctx *gin.Context) {
	type data struct {
		ImageID string `json:"imageId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	imageID, err := util.String2Int64(tempData.ImageID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "imageID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminUserReposityInstance := repository.AdminUserReposityInstance(db)
	users, err := adminUserReposityInstance.GetPendingUserList(imageID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "imageID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	for _, user := range users {
		user.Password = ""
	}
	util.Success(ctx, users, "SUCCESS")

}

func SaveLabelReviewer(ctx *gin.Context) {
	var tempData model.LabelData
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.UserID == 0 {
		ErrorString := ctx.Request.URL.String() + " --- " + ctx.Request.Method + "Bind Data Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	imageID, err := util.String2Int64(tempData.ImageIDString)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " --- " + ctx.Request.Method + "imageID string2int error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	dataIDs, err := adminImageRepositoryInstance.GetDataIDs(tempData.UserID, imageID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " --- " + ctx.Request.Method + "GetDataIDs error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	err = adminImageRepositoryInstance.SaveLabel(tempData, dataIDs)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " --- " + ctx.Request.Method + "SaveLabel error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}

func SetFinalVersionReviewer(ctx *gin.Context) {
	type data struct {
		ImageID       int64 `json:"imageId"`
		UserConfirmID int64 `json:"userConfirmId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	// image, err := adminImageRepositoryInstance.GetImage(tempData.ImageID)
	// if err != nil {
	// 	ErrorString := ctx.Request.URL.String() + " Get Image  Error!!!"
	// 	log.Println(ErrorString)
	// 	util.Fail(ctx, gin.H{}, ErrorString)
	// 	return
	// }

	/** TODO: Ignore the process to gen Rle Data **/

	err := adminImageRepositoryInstance.SetFinalVersion(tempData.ImageID, tempData.UserConfirmID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " Set Final Version  Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}
