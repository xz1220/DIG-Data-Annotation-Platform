package model

// User represents the Table user in database labelproject.
// UserID starts at 1.
// The length of Username and Password should be more than 1.
// Authorities has three values: ROLE_ADMIN ROLE_USER ROLE_REVIEWER
type User struct {
	UserID      int64  `gorm:"AUTO_INCREMENT:primary_key;unique_index;column:user_id" form:"user_id" json:"userId"`
	Username    string `gorm:"type:varchar(50);column:username" form:"user_name" json:"username"`
	Password    string `gorm:"type:varchar(100);column:password" form:"password" json:"password"`
	Authorities string `gorm:"type:varchar(20);column:authorities" form:"authorities" json:"authorities"`
}

// TableName reset the Table field
func (User) TableName() string {
	return "user"
}

// UserFinished the Table userfinished in database labelproject.
// UserID starts at 1.
// TaskID starts at 1.
// ImageID starts at 1.
type UserFinished struct {
	UserID  int64 `gorm:"column:user_id" form:"user_id"`
	TaskID  int64 `gorm:"column:task_id" form:"task_id"`
	ImageID int64 `gorm:"column:image_id" form:"image_id"`
}

// TableName reset the Table field
func (UserFinished) TableName() string {
	return "userfinished"
}
