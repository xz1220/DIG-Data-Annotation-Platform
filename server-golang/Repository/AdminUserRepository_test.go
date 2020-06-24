package repository

import (
	"fmt"
	"labelproject-back/common"
	"labelproject-back/model"
	"testing"
)

func TestGetUserList(t *testing.T) {
	// 	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	userList, _ := userInstance.GetUserList()
	for _, user := range userList {
		fmt.Print("Name: ", user.Username)
	}
	fmt.Println(len(userList))
}

func TestAddUser(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	user := model.User{
		Username:    "xz",
		Password:    "xz",
		Authorities: "ROLE_USER",
	}
	err := userInstance.AddUser(user)
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully")

}

func TestDeleteUser(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	err := userInstance.DeleteUser(9)
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully")
}

func TestEditUser(t *testing.T) {

	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	user := model.User{
		UserID:      2020,
		Username:    "xz",
		Password:    "xz",
		Authorities: "ROLE",
	}
	err := userInstance.EditUser(user)
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully")

}

func TestFindByUserName(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	_, err := userInstance.FindByUserName("")
	if err != nil {
		fmt.Println("Wrong")
	}
	// if user == model.User{} {
	// 	fmt.Println("successfully  None")
	// }
	fmt.Println("successfully")
}

func TestGetUserByID(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	user, err := userInstance.GetUserByID(8)
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", user.UserID)

}

func TestGetPendingUserList(t *testing.T) {
	fmt.Println("waiting ...")
}

func TestGetVideoPendingUserList(t *testing.T) {
	fmt.Println("waiting ...")
}

func TestGetListUser(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	users, err := userInstance.GetListUser()
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", len(users))
}

func TestGetListReviewer(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	users, err := userInstance.GetListReviewer()
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", len(users))
}

func TestGetTaskCount(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	count, err := userInstance.GetTaskCount()
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", count)
}

func TestGetReviewerCount(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	count, err := userInstance.GetReviewerCount()
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", count)
}

func TestGetUserCount(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	userInstance := AdminUserReposityInstance(db)
	count, err := userInstance.GetUserCount()
	if err != nil {
		panic("someting happened!")
	}
	fmt.Println("successfully && ", count)
}

func TestRawSql(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	users := make([]*model.User, 0)
	db.Raw("select * from user u where u.user_id in (select t.user_id from taskuserinfo t where t.task_id = (select i.task_id from image i where i.image_id = ?))", 24).Scan(&users)
	fmt.Println(len(users))
}

// select * from user u where u.user_id in (select t.user_id from taskuserinfo t where t.task_id = (select i.task_id from video i where i.video_id = #{videoId}))
