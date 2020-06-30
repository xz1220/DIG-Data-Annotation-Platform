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

var AdminTaskRepositoryInstance repository.AdminTaskRepository
var AdminImageRepositoryInstance repository.AdminImageRepository
var AdminImageLabelRepositoryInstance repository.AdminImageLabelRepository
var AdminUserReposityInstance repository.AdminUserReposity
var AdminVideoRepositoryInstance repository.AdminVideoRepository
var AdminVideoLabelRepository repository.AdminVideoLabelRepository

func init() {
	db := common.GetDB()
	AdminTaskRepositoryInstance = repository.AdminTaskRepositoryInstance(db)
	AdminImageRepositoryInstance = repository.AdminImageRepositoryInstance(db)
	AdminImageLabelRepositoryInstance = repository.AdminImageLabelRepositoryInstance(db)
	AdminUserReposityInstance = repository.AdminUserReposityInstance(db)
	AdminVideoRepositoryInstance = repository.AdminVideoRepositoryInstance(db)
	AdminVideoLabelRepository = repository.AdminVideoLabelRepositoryInstance(db)
	log.Println("init Repository Instance Successfully")
}

// GetImageList : PostMapping("getImgList")
func GetImageList(ctx *gin.Context) {
	type data struct {
		TaskID  string `json:"taskId"`
		Page    int64  `json:"page"`
		Limit   int64  `json:"limit"`
		AdminID string `json:"adminId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)
	log.Println("TaskID:", tempData.TaskID, "  Page", tempData.Page, "  Limit:", tempData.Limit, "  AdminID:", tempData.AdminID)
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

//SaveLabel : PostMapping("/saveLabel")
func SaveLabel(ctx *gin.Context) {

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

func GetImg(ctx *gin.Context) {
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

func DeleteImageByID(ctx *gin.Context) {
	type data struct {
		ImageID int64 `json:"imageId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.ImageID == 0 {
		ErrorString := ctx.Request.URL.String() + "Bind Parameter Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminImageRepositoryInstance := repository.AdminImageRepositoryInstance(db)
	if adminImageRepositoryInstance.DeleteFromImageByImageID(tempData.ImageID) != nil || adminImageRepositoryInstance.DeleteFromImageDataByImageID(tempData.ImageID) != nil || adminImageRepositoryInstance.DeleteFromImagePointsByImageID(tempData.ImageID) != nil {
		ErrorString := ctx.Request.URL.String() + "Delete Error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}

func SetFinalVersion(ctx *gin.Context) {
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
