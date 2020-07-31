package repository

import (
	"labelproject-back/model"

	"github.com/jinzhu/gorm"
)

// AdminVideoRepository defines functions for model.Video
type AdminVideoRepository interface {
	//
	AddVideo(videos []*model.Video) error

	//
	GetVideoList(TaskID int64) ([]*model.Video, error)

	//
	UpdateVideos(videos []*model.Video) error

	//
	GetDataIDs(userID, videoID int64) ([]int64, error)

	//
	DeleteVideoData(userID, videoID int64) error

	//
	DeleteFinishByID(userID, videoID int64) error

	//
	FinishVideo(userID, videoID int64) error

	//
	AddData(videoDatas []*model.VideoData, userID int64, videoID int64) error

	//
	GetVideo(videoID int64) (model.Video, error)

	//
	GetVideoData(userID, videoID int64) ([]*model.VideoData, error)

	//
	SetVideoFinalVersion(videoID int64, userConfirmID int64) error

	//
	GetVideoIDs(taskID int64) ([]int64, error)

	//
	GetDataIDsByVideoID(videoIDs []int64) ([]int64, error)

	//
	DeleteVideoByTaskID(taskID int64) error

	//
	DeleteDatasByTaskID(videoIDs []int64) error

	//
	UpdateVideoTaskID(videos []*model.Video, taskID int64) error
}

type adminVideoRepository struct {
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminVideoRepositoryInstance = &adminVideoRepository{}

// AdminVideoRepositoryInstance returns the instance of adminVideoRepository
func AdminVideoRepositoryInstance(db *gorm.DB) AdminVideoRepository {
	adminVideoRepositoryInstance.db = db
	return adminVideoRepositoryInstance
}

//
func (r *adminVideoRepository) AddVideo(videos []*model.Video) error {
	for _, video := range videos {
		err := r.db.Create(video).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//
func (r *adminVideoRepository) GetVideoList(taskID int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.Where("task_id = ?", taskID).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// Got some problems about 'When ... the ...'
func (r *adminVideoRepository) UpdateVideos(videos []*model.Video) error { return nil }

//
func (r *adminVideoRepository) GetDataIDs(userID, videoID int64) ([]int64, error) {
	var videodatas []*model.VideoData
	err := r.db.Where("user_id = ? AND video_id = ?", userID, videoID).Find(&videodatas).Error
	if err != nil {
		return nil, err
	}

	var dataIDs []int64
	for _, videodata := range videodatas {
		dataIDs = append(dataIDs, videodata.DataID)
	}
	return dataIDs, nil
}

//
func (r *adminVideoRepository) DeleteVideoData(userID, videoID int64) error {
	err := r.db.Where("video_id = ? AND user_id = ?", videoID, userID).Delete(&model.VideoData{}).Error
	return err
}

//
func (r *adminVideoRepository) DeleteFinishByID(userID, videoID int64) error {
	err := r.db.Where("video_id = ? AND user_id = ?", videoID, userID).Delete(&model.UserFinished{}).Error
	return err
}

//
func (r *adminVideoRepository) FinishVideo(userID, videoID int64) error {
	var video model.Video
	err := r.db.Where("video_id = ?", videoID).First(&video).Error
	if err != nil {
		return err
	}

	userFinished := &model.UserFinished{
		UserID:  userID,
		TaskID:  video.TaskID,
		ImageID: videoID,
	}
	err = r.db.Create(userFinished).Error
	return err
}

//
func (r *adminVideoRepository) AddData(videoDatas []*model.VideoData, userID int64, videoID int64) error {

	for _, videoData := range videoDatas {
		videoData := &model.VideoData{
			UserID:    userID,
			LabelID:   videoData.LabelID,
			Type:      videoData.Type,
			VideoID:   videoID,
			StartTime: videoData.StartTime,
			EndTime:   videoData.EndTime,
			Sentence:  videoData.Sentence,
		}

		err := r.db.Create(videoData).Error
		if err != nil {
			return err
		}
	}
	//err := r.db.Create
	return nil
}

//
func (r *adminVideoRepository) GetVideo(videoID int64) (model.Video, error) {
	var video model.Video
	err := r.db.Where("video_id = ?", videoID).First(&video).Error
	return video, err
}

//
func (r *adminVideoRepository) GetVideoData(userID, videoID int64) ([]*model.VideoData, error) {
	var videoDatas []*model.VideoData
	err := r.db.Where("user_id = ? AND video_id = ?", userID, videoID).Find(&videoDatas).Error
	return videoDatas, err
}

//
func (r *adminVideoRepository) SetVideoFinalVersion(videoID int64, userConfirmID int64) error {
	err := r.db.Model(&model.Video{}).Where("video_id = ?", videoID).Update(model.Video{UserComfirmID: userConfirmID}).Error

	return err
}

//
func (r *adminVideoRepository) GetVideoIDs(taskID int64) ([]int64, error) {
	var videos []*model.Video
	err := r.db.Where("task_id = ?", taskID).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	var videoIDs []int64
	for _, video := range videos {
		videoIDs = append(videoIDs, video.VideoID)
	}

	return videoIDs, err
}

//
func (r *adminVideoRepository) GetDataIDsByVideoID(videoIDs []int64) ([]int64, error) {
	if len(videoIDs) == 0 {
		return nil, nil
	}

	var dataIDs []int64
	for _, videoID := range videoIDs {
		videoData := &model.VideoData{}
		err := r.db.Where("video_id = ?", videoID).First(&videoData).Error
		if err != nil {
			return nil, err
		}
		dataIDs = append(dataIDs, videoData.DataID)
	}
	return dataIDs, nil
}

//
func (r *adminVideoRepository) DeleteVideoByTaskID(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.Video{}).Error
	return err
}

//
func (r *adminVideoRepository) DeleteDatasByTaskID(videoIDs []int64) error {
	for _, videoID := range videoIDs {
		err := r.db.Where("video_id = ?", videoID).Delete(&model.VideoData{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//
func (r *adminVideoRepository) UpdateVideoTaskID(videos []*model.Video, taskID int64) error {
	for _, video := range videos {
		err := r.db.Model(&model.Video{}).Where("video_id = ?", video.VideoID).Update(model.Video{TaskID: taskID}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
