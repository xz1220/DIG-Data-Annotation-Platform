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
		ErrorString := ctx.Request.URL.String() + "GetLabelList error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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
		ErrorString := ctx.Request.URL.String() + " : Bind Label Request Data error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err := adminImageLabelInstance.EditLabel(tempData)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " : Edit Label error!!!" + err.Error()
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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
		ErrorString := ctx.Request.URL.String() + "Bind Label Request Data error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err := adminImageLabelInstance.AddLabel(model.Imagelabel{LabelName: tempData.LabelName, LabelType: int(tempData.LabelType), LabelColor: tempData.LabelColor})
	if err != nil {
		ErrorString := ctx.Request.URL.String() + "Add Label Data error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
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
		ErrorString := ctx.Request.URL.String() + " : Bind Label Request Data error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	taskIDs, err := adminTaskRepositoryInstance.GetTaskIDsByLabelID(tempData.LabelID, 1)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " : GetTaskIDsByLabelID error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	if len(taskIDs) > 0 {
		ErrorString := ctx.Request.URL.String() + fmt.Sprint(" : label has been used by task -", taskIDs)
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	adminImageLabelInstance := repository.AdminImageLabelRepositoryInstance(db)
	err = adminImageLabelInstance.DeleteLabel(tempData.LabelID)
	if err != nil {
		ErrorString := ctx.Request.URL.String() + " : DeleteLabel error!!!"
		log.Println(ErrorString)
		util.Fail(ctx, gin.H{}, ErrorString)
		return
	}

	util.Success(ctx, gin.H{}, "SUCCESS")
}

func SearchLabel(ctx *gin.Context) {

}
