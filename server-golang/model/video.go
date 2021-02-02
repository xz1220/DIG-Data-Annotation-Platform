package model

type Video struct {
	VideoID       int64   `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:video_id" form:"video_id"`
	VideoName     string  `gorm:"type:varchar(1024);column:video_name" form:"video_name"`
	VideoThumb    string  `gorm:"type:varchar(1024);column:video_thumb" form:"video_thumb"`
	TaskID        int64   `gorm:"column:task_id" form:"task_id"`
	Duration      float64 `gorm:"column:duration" form:"duration"`
	UserComfirmID int64   `gorm:"column:user_comfirm_id" form:"user_comfirm_id"`
}

func (Video) TableName() string {
	return "video"
}

type VideoData struct {
	DataID    int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:data_id" form:"data_id"`
	UserID    int64  `gorm:"column:user_id" form:"user_id"`
	LabelID   int64  `gorm:"column:label_id" form:"label_id"`
	Type      int64  `gorm:"column:type" form:"type"`
	VideoID   int64  `gorm:"column:video_id" form:"video_id"`
	StartTime string `gorm:"type:varchar(255);column:start_time" form:"start_time"`
	EndTime   string `gorm:"type:varchar(255);column:end_time" form:"end_time"`
	Sentence  string `gorm:"type:mediumtext;column:sentence"`
}

func (VideoData) TableName() string {
	return "videodata"
}

type VideoLabel struct {
	LabelID  int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:label_id" form:"label_id" json:"labelId"`
	Question string `gorm:"type:varchar(1024);column:question" form:"question" json:"question"`
	Type     int    `gorm:"column:type" form:"type" json:"type"`
	Selector string `gorm:"type:varchar(1024);column:selector" form:"selector" json:"selector"`
}

func (VideoLabel) TableName() string {
	return "videolabel"
}

type VideoLabelMulti struct {
	LabelID   int64
	Question  string
	Type      int
	selectors []string
}
