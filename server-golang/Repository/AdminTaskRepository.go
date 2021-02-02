package repository

import (
	"labelproject-back/model"

	"github.com/jinzhu/gorm"
)

// AdminTaskRepository defines functions for model.Task adn reliative tables.
type AdminTaskRepository interface {

	//
	GetImageTaskNameList() ([]string, error)
	GetVideoTaskNameList() ([]string, error)

	//
	AddTask(task model.Task) error

	//
	GetTaskList() ([]*model.TaskList, error)

	//
	AddTaskUserIds(userIDs []int64, taskID int64) error

	//
	GetTaskNameByID(taskID int64) (string, error)

	//
	AddTaskLabelIDs(labelIDs []int64, taskID int64) error

	//
	AddTaskReviewerIDs(labelIDs []int64, taskID int64) error

	//
	DeleteTaskUserIDs(taskID int64) error

	//
	DeleteTaskLabelIDs(taskID int64) error

	//
	DeleteTaskReviewerIDs(taskID int64) error

	//
	UpdateTask(task model.Task) error

	//
	DeleteTask(taskID int64) error

	//
	GetTaskByID(taskID int64) (model.Task, error)

	//
	GetTaskNameByImageID(imageID int64) (string, error)

	//
	GetTaskIDs(userID int64) ([]int64, error)

	//
	GetTaskListByID(taskIDs []int64) ([]*model.Task, error)

	//
	GetTaskIDByReviewerID(reviewerID int64) ([]int64, error)

	//
	TaskList() ([]*model.TaskList, error)

	//
	GetUserInfo(taskID int64) ([]*model.TaskUserInfo, error)

	//
	GetUserIDsFromUserInfo(taskID int64) ([]int64, error)

	//
	GetReviewerInfo(taskID int64) ([]*model.TaskReviewerInfo, error)

	//
	GetReviewerIDsFromReviewerInfo(taskID int64) ([]int64, error)

	//
	GetLabelInfo(taskID int64) ([]*model.Tasklabelinfo, error)

	//
	GetLabelIDsFromLabelInfo(taskID int64) ([]int64, error)
	//
	TaskListByID(taskIDs []int64) ([]*model.Task, error)

	//
	GetTaskIDsByLabelID(labelID int64, taskType int64) ([]int64, error)

	//
	SearchTask(keywords string) ([]*model.TaskList, error)

	//
	GetNewTaskList() ([]*model.Task, error)

	//
	UpdateTaskType(taskID int64, taskType int64) error

	//
	HasData(taskID int64) (int, error)

	// get last insert record
	LastRecord() (model.Task, error)
}

type adminTaskRepository struct {
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminTaskRepositoryInstance = &adminTaskRepository{}

// AdminTaskRepositoryInstance returns the instance of the adminTaskRepository.
func AdminTaskRepositoryInstance(db *gorm.DB) AdminTaskRepository {
	adminTaskRepositoryInstance.db = db
	return adminTaskRepositoryInstance
}

//
func (r *adminTaskRepository) GetImageTaskNameList() ([]string, error) {
	var taskNames []string
	var tasks []*model.Task

	err := r.db.Where("task_type >=1 AND task_type<=4").Find(&tasks).Error
	if err != nil {
		return nil, nil
	}

	for _, task := range tasks {
		taskNames = append(taskNames, task.TaskName)
	}
	return taskNames, nil
}
func (r *adminTaskRepository) GetVideoTaskNameList() ([]string, error) {
	var taskNames []string
	var tasks []*model.Task

	err := r.db.Where("task_type = 5").Find(&tasks).Error
	if err != nil {
		return nil, nil
	}

	for _, task := range tasks {
		taskNames = append(taskNames, task.TaskName)
	}
	return taskNames, nil
}

//
func (r *adminTaskRepository) AddTask(task model.Task) error {
	err := r.db.Create(task).Error
	return err
}

//
func (r *adminTaskRepository) GetTaskList() ([]*model.TaskList, error) {
	var tasks []*model.TaskList
	err := r.db.Raw("select t.task_id,t.task_name,t.image_number,t.task_type,(select count(*) from image i where i.task_id = t.task_id and i.user_confirm_id is not null) finish,t.task_desc from task t where t.is_created = 1").Scan(&tasks).Error
	return tasks, err
}

//
func (r *adminTaskRepository) AddTaskUserIds(userIDs []int64, taskID int64) error {
	for _, userID := range userIDs {
		err := r.db.Create(&model.TaskUserInfo{TaskID: taskID, UserID: userID}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//
func (r *adminTaskRepository) GetTaskNameByID(taskID int64) (string, error) {
	var task model.Task
	err := r.db.Where("task_id = ?", taskID).First(&task).Error
	if err != nil {
		return "", nil
	}
	return task.TaskName, nil
}

//
func (r *adminTaskRepository) AddTaskLabelIDs(labelIDs []int64, taskID int64) error {
	for _, labelID := range labelIDs {
		err := r.db.Create(&model.Tasklabelinfo{TaskID: taskID, LabelID: labelID}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//
func (r *adminTaskRepository) AddTaskReviewerIDs(reviewerIDs []int64, taskID int64) error {
	for _, reviewerID := range reviewerIDs {
		err := r.db.Create(&model.TaskReviewerInfo{TaskID: taskID, ReviewerID: reviewerID}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//
func (r *adminTaskRepository) DeleteTaskUserIDs(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.TaskUserInfo{}).Error
	return err
}

//
func (r *adminTaskRepository) DeleteTaskLabelIDs(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.Tasklabelinfo{}).Error
	return err
}

//
func (r *adminTaskRepository) DeleteTaskReviewerIDs(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.TaskReviewerInfo{}).Error
	return err
}

//
func (r *adminTaskRepository) UpdateTask(task model.Task) error {
	err := r.db.Model(&task).Where("task_id = ?", task.TaskID).Updates(task).Error
	return err
}

//
func (r *adminTaskRepository) DeleteTask(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.Task{}).Error
	return err
}

//
func (r *adminTaskRepository) GetTaskByID(taskID int64) (model.Task, error) {
	var task model.Task
	err := r.db.Raw("select t.task_id,t.task_name,t.image_number,t.task_desc,t.task_type,(select count(*) from image i where i.task_id = t.task_id and i.user_confirm_id is not null) finish from task t where t.task_id = ?", taskID).Scan(&task).Error
	return task, err
}

//
func (r *adminTaskRepository) GetTaskNameByImageID(imageID int64) (string, error) {
	var task model.Task
	err := r.db.Where("task_id = ?", r.db.Table("image").Select("task_id").Where("image_id = ?", imageID).SubQuery()).First(&task).Error
	return task.TaskName, err
}

//
func (r *adminTaskRepository) GetTaskIDs(userID int64) ([]int64, error) {
	var taskIDs []int64
	var taskUserInfos []*model.TaskUserInfo
	err := r.db.Where("user_id = ?", userID).Find(&taskUserInfos).Error
	if err != nil {
		return nil, err
	}
	for _, taskUserInfo := range taskUserInfos {
		taskIDs = append(taskIDs, taskUserInfo.TaskID)
	}
	return taskIDs, nil
}

//
func (r *adminTaskRepository) GetTaskListByID(taskIDs []int64) ([]*model.Task, error) {
	var tasks []*model.Task

	for _, taskID := range taskIDs {
		var task model.Task
		err := r.db.Raw("select t.task_id,t.task_name,t.image_number,t.task_type,tl.label_id,t.task_desc from task t left join tasklabelinfo tl on tl.task_id = t.task_id where t.task_id = ?", taskID).Scan(&task).Error
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

//
func (r *adminTaskRepository) GetTaskIDByReviewerID(reviewerID int64) ([]int64, error) {
	var taskReviewerInfos []*model.TaskReviewerInfo
	var taskIDs []int64
	err := r.db.Where("reviewer_id = ?", reviewerID).Find(&taskReviewerInfos).Error
	if err != nil {
		return nil, err
	}

	for _, taskReviewerInfo := range taskReviewerInfos {
		taskIDs = append(taskIDs, taskReviewerInfo.TaskID)
	}
	return taskIDs, nil
}

//
func (r *adminTaskRepository) TaskList() ([]*model.TaskList, error) {
	var tasks []*model.TaskList
	err := r.db.Raw("select t.task_id task_id,t.task_name,t.image_number,t.task_type,if(task_type = 5,(select count(*) from video v where v.task_id = t.task_id and v.user_confirm_id !=0),(select count(*) from image i where i.task_id = t.task_id and i.user_confirm_id !=0))finish,t.task_desc from task t where t.is_created = 1").Scan(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//
func (r *adminTaskRepository) GetUserInfo(taskID int64) ([]*model.TaskUserInfo, error) {
	var userInfos []*model.TaskUserInfo
	err := r.db.Where("task_id = ?", taskID).Find(&userInfos).Error
	return userInfos, err
}

//
func (r *adminTaskRepository) GetUserIDsFromUserInfo(taskID int64) ([]int64, error) {
	var userInfos []*model.TaskUserInfo
	err := r.db.Where("task_id = ?", taskID).Find(&userInfos).Error
	if err != nil {
		return nil, err
	}

	var userIDs []int64
	for _, userInfo := range userInfos {
		userIDs = append(userIDs, userInfo.UserID)
	}
	return userIDs, nil
}

//Not used
func (r *adminTaskRepository) GetReviewerInfo(taskID int64) ([]*model.TaskReviewerInfo, error) {
	var reviewerInfos []*model.TaskReviewerInfo
	err := r.db.Where("task_id = ?", taskID).Find(&reviewerInfos).Error
	return reviewerInfos, err
}

func (r *adminTaskRepository) GetReviewerIDsFromReviewerInfo(taskID int64) ([]int64, error) {
	var reviewerInfos []*model.TaskReviewerInfo
	err := r.db.Where("task_id = ?", taskID).Find(&reviewerInfos).Error

	if err != nil {
		return nil, err
	}

	var reviewerIDs []int64
	for _, reviewerInfo := range reviewerInfos {
		reviewerIDs = append(reviewerIDs, reviewerInfo.ReviewerID)
	}
	return reviewerIDs, nil
}

//
func (r *adminTaskRepository) GetLabelInfo(taskID int64) ([]*model.Tasklabelinfo, error) {
	var labelInfos []*model.Tasklabelinfo
	err := r.db.Where("task_id = ?", taskID).Find(&labelInfos).Error
	return labelInfos, err
}

func (r *adminTaskRepository) GetLabelIDsFromLabelInfo(taskID int64) ([]int64, error) {
	var labelInfos []*model.Tasklabelinfo
	err := r.db.Where("task_id = ?", taskID).Find(&labelInfos).Error

	if err != nil {
		return nil, err
	}

	var labelIDs []int64
	for _, labelInfo := range labelInfos {
		labelIDs = append(labelIDs, labelInfo.LabelID)
	}
	return labelIDs, nil
}

//
func (r *adminTaskRepository) TaskListByID(taskIDs []int64) ([]*model.Task, error) {
	var tasks []*model.Task

	for _, taskID := range taskIDs {
		var task model.Task
		err := r.db.Raw("select t.task_id task_id,t.task_name,t.image_number,t.task_type,(select count(*) from image i where i.task_id = t.task_id and i.user_confirm_id is not null) finish,t.task_desc from task t where t.task_id = ?", taskID).Scan(&task).Error
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

//
func (r *adminTaskRepository) GetTaskIDsByLabelID(labelID int64, taskType int64) ([]int64, error) {
	var tasks []*model.Task
	err := r.db.Raw("select t.task_id from task t where t.task_type = ? and t.task_id in (select l.task_id from tasklabelinfo l where label_id = ?)", taskType, labelID).Scan(&tasks).Error
	if err != nil {
		return nil, err
	}

	var taskIDs []int64
	for _, task := range tasks {
		taskIDs = append(taskIDs, task.TaskID)
	}
	return taskIDs, nil
}

//
func (r *adminTaskRepository) SearchTask(keywords string) ([]*model.TaskList, error) {
	var tasks []*model.TaskList
	err := r.db.Raw("select t.task_id,t.task_name,t.image_number,t.task_type,(select count(*) from image i where i.task_id = t.task_id and i.user_confirm_id is not null) finish from task t where t.is_created = 1 and t.task_name like concat('%',?,'%')", keywords).Scan(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//
func (r *adminTaskRepository) GetNewTaskList() ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.Raw("select task_id,task_type,task_name from task where is_created = 0").Scan(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//
func (r *adminTaskRepository) UpdateTaskType(taskID int64, taskType int64) error {
	err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).Updates(model.Task{TaskType: taskType, IsCreated: 1}).Error
	return err
}

//
func (r *adminTaskRepository) HasData(taskID int64) (int, error) {
	var imageDatas []*model.ImageData
	err := r.db.Raw("select * from imagedata where image_id in (select i.image_id from image i where i.task_id = ?) limit 1", taskID).Scan(&imageDatas).Error
	if err != nil {
		return -1, err
	}
	return len(imageDatas), nil
}

func (r *adminTaskRepository) LastRecord() (model.Task, error) {
	var task model.Task
	err := r.db.Last(&task).Error
	return task, err
}
