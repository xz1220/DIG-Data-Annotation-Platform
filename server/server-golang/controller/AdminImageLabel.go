package controller

import (
	"encoding/json"
	"fmt"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"

	"github.com/gin-gonic/gin"
)

func GetLabelList(ctx *gin.Context) {
	db := common.GetDB()
	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	log.Println("try to get Label List")
	labels, err := adminImageLabelInstance.GetLabelList()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "GetLabelList error!!!")
		return
	}

	log.Println("GetLabelList Successfully")
	var dataMap map[string]interface{}
	dataMap = make(map[string]interface{})
	dataMap["labelList"] = labels
	util.Success(ctx, dataMap, "SUCCESS")
}

func EditLabel(ctx *gin.Context) {
	var tempData model.Imagelabel
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.LabelID == 0 {
		util.ManagerInstance.FailWithoutData(ctx, "Bind Label Request Data error!!!")
		return
	}

	db := common.GetDB()
	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err := adminImageLabelInstance.EditLabel(tempData)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " Edit Label error!!!")
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")

}

func AddLabel(ctx *gin.Context) {
	type data struct {
		LabelName  string `json:"labelName"`
		LabelType  int64  `json:"labelType"`
		LabelColor string `json:"labelColor"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.LabelName == "" {
		util.ManagerInstance.FailWithoutData(ctx, "Bind Label Request Data error!!!")
		return
	}

	db := common.GetDB()
	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err := adminImageLabelInstance.AddLabel(model.Imagelabel{LabelName: tempData.LabelName, LabelType: int(tempData.LabelType), LabelColor: tempData.LabelColor})
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "Add Label Data error!!!")
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}

func DeleteLabel(ctx *gin.Context) {
	type data struct {
		LabelID int64 `json:"labelId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.LabelID == 0 {
		util.ManagerInstance.FailWithoutData(ctx, " Bind Label Request Data error!!!")
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	taskIDs, err := adminTaskRepositoryInstance.GetTaskIDsByLabelID(tempData.LabelID, 1)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " GetTaskIDsByLabelID error!!!")
		return
	}

	if len(taskIDs) > 0 {
		util.ManagerInstance.FailWithoutData(ctx, fmt.Sprint(" : label has been used by task -", taskIDs))
		return
	}

	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err = adminImageLabelInstance.DeleteLabel(tempData.LabelID)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, "DeleteLabel error!!!")
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}

// func SearchLabel(ctx *gin.Context) {

// }
