package model

type User struct {
	UserID      int64  `gorm:"AUTO_INCREMENT:primary_key;unique_index;column:user_id" form:"user_id" json:"userId"`
	Username    string `gorm:"type:varchar(50);column:username" form:"user_name" json:"username"`
	Password    string `gorm:"type:varchar(100);column:password" form:"password" json:"password"`
	Authorities string `gorm:"type:varchar(20);column:authorities" form:"authorities" json:"authorities"`
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}

type UserFinished struct {
	UserID  int64 `gorm:"column:user_id" form:"user_id"`
	TaskID  int64 `gorm:"column:task_id" form:"task_id"`
	ImageID int64 `gorm:"column:image_id" form:"image_id"`
}

func (UserFinished) TableName() string {
	return "userfinished"
}

type UserInfo struct {
	Username string
	UserID   int64
	Labeled  int64
}
