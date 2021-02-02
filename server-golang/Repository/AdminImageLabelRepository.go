package repository

import (
	"fmt"
	"labelproject-back/model"
	"log"

	"github.com/jinzhu/gorm"
)

// AdminImageLabelRepository defines functions for model.ImageLabel
type AdminImageLabelRepository interface {

	//获取标签列表
	GetLabelList() ([]*model.Imagelabel, error)

	//修改标签
	EditLabel(model.Imagelabel) error

	//添加标签
	AddLabel(model.Imagelabel) error

	//删除标签
	DeleteLabel(int64) error

	//通过标签名找标签
	FindByLabelName(string) (model.Imagelabel, error)

	//获取该图片的标签
	GetLabelByImageID(int64) ([]*model.Imagelabel, error)

	//
	SearchLabel(string) ([]*model.Imagelabel, error)
}

type adminImageLabelRepository struct {
	/** 数据库连接对象 **/
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminImageLabelInstance = &adminImageLabelRepository{}

// AdminImageLabelRepositoryInstance returen a instance of AdminImageLabelRepository
func AdminImageLabelRepositoryInstance(db *gorm.DB) AdminImageLabelRepository {
	adminImageLabelInstance.db = db
	return adminImageLabelInstance
}

//获取标签列表
func (r *adminImageLabelRepository) GetLabelList() ([]*model.Imagelabel, error) {
	var imageLabels []*model.Imagelabel
	err := r.db.Model(&model.Imagelabel{}).Find(&imageLabels).Error
	if err != nil {
		return nil, err
	}
	return imageLabels, err
}

//修改标签
func (r *adminImageLabelRepository) EditLabel(imageLabel model.Imagelabel) error {
	temp, err := r.FindByLabelName(imageLabel.LabelName)
	// if err != nil {
	// 	return fmt.Errorf("获取标签名称失败")
	// }

	if temp.LabelID != imageLabel.LabelID {
		return fmt.Errorf("标签已存在，修改失败")
	}
	err = r.db.Model(&model.Imagelabel{}).Where("label_id = ?", imageLabel.LabelID).Updates(model.Imagelabel{LabelName: imageLabel.LabelName, LabelType: imageLabel.LabelType, LabelColor: imageLabel.LabelColor}).Error
	if err != nil {
		log.Println("修改失败请重试")
		return err
	}
	return nil
}

//添加标签
func (r *adminImageLabelRepository) AddLabel(imageLabel model.Imagelabel) error {
	_, err := r.FindByLabelName(imageLabel.LabelName)
	if err == nil {
		log.Println("标签已存在 修改失败")
		return fmt.Errorf("标签已存在 修改失败")
	}
	err = r.db.Create(&imageLabel).Error
	if err != nil {
		log.Println("添加失败请重试")
		return err
	}
	return nil
}

//删除标签
func (r *adminImageLabelRepository) DeleteLabel(labelID int64) error {
	err := r.db.Where("label_id = ?", labelID).Delete(&model.Imagelabel{}).Error
	return err
}

//通过标签名找标签
func (r *adminImageLabelRepository) FindByLabelName(labelName string) (model.Imagelabel, error) {
	var imageLabel model.Imagelabel
	err := r.db.Where("label_name = ?", labelName).First(&imageLabel).Error
	return imageLabel, err
}

//获取该图片的标签
func (r *adminImageLabelRepository) GetLabelByImageID(ImageID int64) ([]*model.Imagelabel, error) {
	var imageLabels []*model.Imagelabel
	err := r.db.Raw("select * from imagelabel where label_id in (select label_id from tasklabelinfo where task_id = (select task_id from image where image_id = ?))", ImageID).Scan(&imageLabels).Error
	return imageLabels, err
}

//
func (r *adminImageLabelRepository) SearchLabel(keywords string) ([]*model.Imagelabel, error) {
	var imageLabels []*model.Imagelabel
	err := r.db.Raw("select * from imagelabel where label_name like concat('%',?,'%')", keywords).Scan(&imageLabels).Error
	return imageLabels, err
}
