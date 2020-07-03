package controller

import (
	"encoding/json"
	repository "labelproject-back/Repository"
	"labelproject-back/common"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"

	"github.com/gin-gonic/gin"
)

func GetVideoLabelList(ctx *gin.Context) {
	type VideoLabel struct {
		LabelID  int64    `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:label_id" form:"label_id" json:"labelId"`
		Question string   `gorm:"type:varchar(1024);column:question" form:"question" json:"question"`
		Type     int      `gorm:"column:type" form:"type" json:"type"`
		Selector []string `gorm:"type:varchar(1024);column:selector" form:"selector" json:"selector"`
	}

	db := common.GetDB()
	adminVideoLabelRepositoryInstance := repository.AdminVideoLabelRepositoryInstance(db)
	tempVideoLabels, err := adminVideoLabelRepositoryInstance.GetVideoLabelList()
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " GetVideoLabelList  Error!!!")
		return
	}

	videoLabels := make([]*VideoLabel, 0)
	for _, tempVideoLabel := range tempVideoLabels {
		videoLabel := VideoLabel{
			LabelID:  tempVideoLabel.LabelID,
			Question: tempVideoLabel.Question,
			Type:     tempVideoLabel.Type,
			Selector: []string{tempVideoLabel.Selector},
		}
		videoLabels = append(videoLabels, &videoLabel)
	}
	log.Println("查询VideoLabel成功")
	util.Success(ctx, videoLabels, "SUCCESS")

}

func AddVideoLabel(ctx *gin.Context) {
	type VideoLabel struct {
		LabelID  int64    `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:label_id" form:"label_id" json:"labelId"`
		Question string   `gorm:"type:varchar(1024);column:question" form:"question" json:"question"`
		Type     int      `gorm:"column:type" form:"type" json:"type"`
		Selector []string `gorm:"type:varchar(1024);column:selector" form:"selector" json:"selector"`
	}

	var tempData VideoLabel
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	tempVideoLabel := model.VideoLabel{
		Question: tempData.Question,
		Type:     tempData.Type,
		Selector: tempData.Selector[0],
	}

	db := common.GetDB()
	adminVideoLabelRepositoryInstance := repository.AdminVideoLabelRepositoryInstance(db)
	err := adminVideoLabelRepositoryInstance.AddVideoLabel(tempVideoLabel)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " 添加失败，请重试!!!")
		return
	}
	log.Println("添加模板成功")
	util.Success(ctx, gin.H{}, "SUCCESS")
}

func EditVideoLabel(ctx *gin.Context) {
	type VideoLabel struct {
		LabelID  int64    `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:label_id" form:"label_id" json:"labelId"`
		Question string   `gorm:"type:varchar(1024);column:question" form:"question" json:"question"`
		Type     int      `gorm:"column:type" form:"type" json:"type"`
		Selector []string `gorm:"type:varchar(1024);column:selector" form:"selector" json:"selector"`
	}

	var tempData VideoLabel
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	tempVideoLabel := model.VideoLabel{
		LabelID:  tempData.LabelID,
		Question: tempData.Question,
		Type:     tempData.Type,
		Selector: tempData.Selector[0],
	}

	db := common.GetDB()
	adminVideoLabelRepositoryInstance := repository.AdminVideoLabelRepositoryInstance(db)
	err := adminVideoLabelRepositoryInstance.EditVideoLabel(tempVideoLabel)
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " 修改失败，请重试!!!")
		return
	}
	log.Println("修改模板成功")
	util.Success(ctx, gin.H{}, "SUCCESS")

}

func DeleteVideoLabel(ctx *gin.Context) {

	type data struct {
		LabelID int64 `json:"labelId"`
	}

	var tempData data
	json.NewDecoder(ctx.Request.Body).Decode(&tempData)

	if tempData.LabelID == 0 {
		util.ManagerInstance.FailWithoutData(ctx, " 参数错误，修改失败!!!")
		return
	}

	db := common.GetDB()
	adminTaskRepositoryInstance := repository.AdminTaskRepositoryInstance(db)
	taskIDs, err := adminTaskRepositoryInstance.GetTaskIDsByLabelID(tempData.LabelID, 5)
	if len(taskIDs) > 0 {
		util.ManagerInstance.FailWithoutData(ctx, " 该标签已被使用，删除失败!!!")
		return
	}

	adminVideoLabelRepositoryInstance := repository.AdminVideoLabelRepositoryInstance(db)
	err = adminVideoLabelRepositoryInstance.DeleteVideoLabel(model.VideoLabel{LabelID: tempData.LabelID})
	if err != nil {
		util.ManagerInstance.FailWithoutData(ctx, " 删除失败，请重试!!!")
		return
	}
	log.Println("删除模板成功")
	util.Success(ctx, gin.H{}, "SUCCESS")
}
