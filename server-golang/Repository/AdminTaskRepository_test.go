package repository

import (
	"fmt"
	"labelproject-back/common"
	"labelproject-back/model"
	"testing"
)

func TestGetTaskList(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.GetTaskList()
	if err != nil {
		panic("Error")
	}
	fmt.Println("task_id: ", tasks[0].TaskID)
}

func TestGetTaskByID(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	task, err := adminTaskRepositoryInstance.GetTaskByID(1)
	if err != nil {
		panic("Error")
	}
	fmt.Println("task_id: ", task.TaskID)
}

func TestAddTaskUserIDs(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	err := adminTaskRepositoryInstance.AddTaskUserIds([]int64{100, 200}, 1)
	if err != nil {
		panic("Error")
	}
}

func TestTest(t *testing.T) {
	task := &model.Task{TaskID: 2}
	if task.TaskName == "" {
		fmt.Println("Success")
	}
	fmt.Println("name: ", task.TaskName)
}

func TestTask(t *testing.T) {
	common.InitConfig("/root/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	fmt.Println("successfully connect mysql")

	task := model.Task{TaskID: 10001, TaskName: "Test"}
	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	err := adminTaskRepositoryInstance.AddTask(task)
	if err != nil {
		panic("Add Task Error")
	}

	newTask := task
	newTask.TaskName = "NewTest"
	lastRecord, err := adminTaskRepositoryInstance.LastRecord()
	if err != nil {
		panic("Get last Record Error")
	}
	newTask.TaskID = lastRecord.TaskID + 1
	err = adminTaskRepositoryInstance.AddTask(newTask)
	if err != nil {
		panic("Add New Task Error")
	}

	task.TaskDesc = "This is a Test"
	err = adminTaskRepositoryInstance.UpdateTask(task)
	if err != nil {
		panic("Update Task Error")
	}

	err = adminTaskRepositoryInstance.DeleteTask(task.TaskID)
	if err != nil {
		panic("Delete Task Error")
	}

	err = adminTaskRepositoryInstance.DeleteTask(newTask.TaskID)
	if err != nil {
		panic("Delete Task Error")
	}
}

func TestGetTaskNameByImageID(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	taskName, err := adminTaskRepositoryInstance.GetTaskNameByImageID(2)
	if err != nil {
		panic("Error")
	}
	fmt.Println("TaskName : ", taskName)

}

func TestGetTaskListByID(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.GetTaskListByID([]int64{10002, 10001, 1})
	if err != nil {
		panic("Error")
	}
	for index := range tasks {
		fmt.Println("tasks.Name: ", tasks[index].TaskName)
	}

}

func TestTaskList(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	tasks, err := adminTaskRepositoryInstance.TaskList()
	if err != nil {
		panic("Error")
	}
	for index := range tasks {
		fmt.Println("tasks.Name: ", tasks[index].TaskName)
	}

}

func TestHasData(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	fmt.Println("successfully connect mysql")

	adminTaskRepositoryInstance := AdminTaskRepositoryInstance(db)
	count, err := adminTaskRepositoryInstance.HasData(1)
	if err != nil {
		panic("Error")
	}
	fmt.Print("Count : ", count)
}


