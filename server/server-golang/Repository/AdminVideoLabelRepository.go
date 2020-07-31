package repository

import (
	"labelproject-back/model"

	"github.com/jinzhu/gorm"
)

// AdminVideoLabelRepository defines functions for model.VidelLabel
type AdminVideoLabelRepository interface {
	//
	GetVideoLabelList() ([]*model.VideoLabel, error)

	//
	AddVideoLabel(videoLabel model.VideoLabel) error

	//
	EditVideoLabel(videoLabel model.VideoLabel) error

	//
	DeleteVideoLabel(videoLabel model.VideoLabel) error

	//
	GetVideoLabelsByVideoID(videoID int64) ([]*model.VideoLabelMulti, error)
}

type adminVideoLabelRepository struct {
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminVideoLabelRepositoryInstance = &adminVideoLabelRepository{}

// AdminVideoLabelRepositoryInstance returns the instance of adminVideoLabelRepository
func AdminVideoLabelRepositoryInstance(db *gorm.DB) AdminVideoLabelRepository {
	adminVideoLabelRepositoryInstance.db = db
	return adminVideoLabelRepositoryInstance
}

//
func (r *adminVideoLabelRepository) GetVideoLabelList() ([]*model.VideoLabel, error) {
	var videoLabel []*model.VideoLabel
	err := r.db.Find(&videoLabel).Error
	if err != nil {
		return nil, err
	}
	return videoLabel, nil
}

//
func (r *adminVideoLabelRepository) AddVideoLabel(videoLabel model.VideoLabel) error {
	err := r.db.Create(&videoLabel).Error
	return err
}

//
func (r *adminVideoLabelRepository) EditVideoLabel(videoLabel model.VideoLabel) error {
	err := r.db.Model(&model.VideoLabel{}).Where("label_id = ?", videoLabel.LabelID).Updates(videoLabel).Error
	return err
}

//
func (r *adminVideoLabelRepository) DeleteVideoLabel(videoLabel model.VideoLabel) error {
	err := r.db.Where("label_id = ?", videoLabel.LabelID).Delete(&model.VideoLabel{}).Error
	return err
}

//
func (r *adminVideoLabelRepository) GetVideoLabelsByVideoID(videoID int64) ([]*model.VideoLabelMulti, error) {
	var videoLabel []*model.VideoLabelMulti
	err := r.db.Raw("select * from videolabel where label_id in (select label_id from tasklabelinfo where task_id = (select task_id from video where video_id = ?))", videoID).Scan(&videoLabel).Error
	if err != nil {
		return nil, err
	}
	return videoLabel, nil
}
