package repository

import (
	"fmt"
	"labelproject-back/model"
	"log"

	"github.com/jinzhu/gorm"
)

// AdminUserReposity defines functions for model.User
type AdminUserReposity interface {
	//获取用户列表
	GetUserList() ([]*model.User, error)

	// 编辑用户
	EditUser(user model.User) error

	//添加用户
	AddUser(user model.User) error

	//删除用户
	DeleteUser(userID int64) error

	//通过用户名找用户
	FindByUserName(username string) (model.User, error)

	//通过用户ＩＤ找用户
	GetUserByID(UserID int64) (model.User, error)

	//获取该图片的标记用户
	GetPendingUserList(ImageID int64) ([]*model.User, error)

	//获取视频标记用户
	GetVideoPendingUserList(VideoID int64) ([]*model.User, error)

	//获取标记用户
	GetListUser() ([]*model.User, error)

	//获取审核用户
	GetListReviewer() ([]*model.User, error)

	//获取任务数量
	GetTaskCount() (int, error)

	//获取审核任务数量
	GetReviewerCount() (int, error)

	//获取标记用户数量
	GetUserCount() (int, error)
}

type adminUserReposity struct {
	/** 数据库连接对象 **/
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminUserInstance = &adminUserReposity{}

//AdminUserReposityInstance returns the instance of adminUserReposity
func AdminUserReposityInstance(db *gorm.DB) AdminUserReposity {
	adminUserInstance.db = db
	return adminUserInstance
}

func (r *adminUserReposity) AddUser(user model.User) error {
	_, err := r.FindByUserName(user.Username)
	//  判断User 是否存在
	if err != nil {
		err := r.db.Create(&user).Error
		return err
	}

	err = fmt.Errorf("用户已存在")
	log.Println(err)
	return err
}

func (r *adminUserReposity) EditUser(user model.User) error {

	userTemp, err := r.FindByUserName(user.Username)
	//  判断User 是否存在
	if err != nil {
		log.Println(err)
		return err
	}

	// 判断是否传参错误
	if err == nil && userTemp.UserID != user.UserID {
		err = fmt.Errorf("userID not matching")
		log.Println(err)
		return err
	}

	err = r.db.Model(&user).Where("user_id = ?", user.UserID).Updates(model.User{Username: user.Username, Password: user.Password, Authorities: user.Authorities}).Error
	return err
}

func (r *adminUserReposity) DeleteUser(userID int64) error {
	err := r.db.Where("user_id = ? ", userID).Delete(&model.User{}).Error
	return err
}

func (r *adminUserReposity) GetUserList() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Find(&users).Error
	return users, err
}

//注意用户名唯一
func (r *adminUserReposity) FindByUserName(username string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *adminUserReposity) GetUserByID(UserID int64) (model.User, error) {
	var user model.User
	err := r.db.Where("user_id = ?", UserID).First(&user).Error
	return user, err
}

func (r *adminUserReposity) GetPendingUserList(ImageID int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Raw("select * from user u where u.user_id in (select t.user_id from taskuserinfo t where t.task_id = (select i.task_id from image i where i.image_id = ?))", ImageID).Scan(&users).Error
	return users, err
}

func (r *adminUserReposity) GetVideoPendingUserList(VideoID int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Raw("select * from user u where u.user_id in (select t.user_id from taskuserinfo t where t.task_id = (select i.task_id from video i where i.video_id = ?))", VideoID).Scan(&users).Error
	return users, err
}

func (r *adminUserReposity) GetListUser() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Where("authorities= 'ROLE_USER'").Find(&users).Error
	return users, err
}

func (r *adminUserReposity) GetListReviewer() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Where("authorities = 'ROLE_REVIEWER'").Find(&users).Error
	return users, err
}

func (r *adminUserReposity) GetTaskCount() (int, error) {
	var count int
	err := r.db.Table("task").Count(&count).Error
	return count, err
}

func (r *adminUserReposity) GetUserCount() (int, error) {
	var count int
	err := r.db.Model(&model.User{}).Where("authorities = 'ROLE_USER'").Count(&count).Error
	return count, err
}

func (r *adminUserReposity) GetReviewerCount() (int, error) {
	var count int
	err := r.db.Model(&model.User{}).Where("authorities = 'ROLE_REVIEWER'").Count(&count).Error
	return count, err
}
