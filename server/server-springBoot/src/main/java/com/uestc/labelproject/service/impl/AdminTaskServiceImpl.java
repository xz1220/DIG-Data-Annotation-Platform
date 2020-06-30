package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminTaskMapper;
import com.uestc.labelproject.entity.Task;
import com.uestc.labelproject.service.AdminTaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Set;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:47
 */
@Service
public class AdminTaskServiceImpl implements AdminTaskService {

    @Autowired
    AdminTaskMapper adminTaskMapper;

    @Override
    public Set<String> getImageTaskNameList() {
        return adminTaskMapper.getImageTaskNameList();
    }

    @Override
    public Set<String> getVideoTaskNameList() {
        return adminTaskMapper.getVideoTaskNameList();
    }

    @Override
    public int addTask(Task task) {
        return adminTaskMapper.addTask(task);
    }

    @Override
    public List<Task> getTaskList() {
        return adminTaskMapper.getTaskList();
    }


    @Override
    public String getTaskNameById(Long taskId) {
        return adminTaskMapper.getTaskNameById(taskId);
    }


    @Override
    public Task getTaskById(Long taskId) {
        return adminTaskMapper.getTaskById(taskId);
    }

    @Override
    public int addTaskUserIds(List<Long> userIds, Long taskId) {
        return adminTaskMapper.addTaskUserIds(userIds, taskId);
    }

    @Override
    public int addTaskLabelIds(List<Long> labelIds, Long taskId) {
        return adminTaskMapper.addTaskLabelIds(labelIds, taskId);
    }

    @Override
    public int deleteTaskUserIds(Long taskId) {
        return adminTaskMapper.deleteTaskUserIds(taskId);
    }

    @Override
    public int deleteTaskLabelIds(Long taskId) {
        return adminTaskMapper.deleteTaskLabelIds(taskId);
    }

    @Override
    public int addTaskReviewerIds(List<Long> reviewerIds, Long taskdId) {
        return adminTaskMapper.addTaskReviewerIds(reviewerIds, taskdId);
    }

    @Override
    public int deleteTaskReviewerIds(Long taskId) {
        return adminTaskMapper.deleteTaskReviewerIds(taskId);
    }

    @Override
    public int updateTask(Task task) {
        return adminTaskMapper.updateTask(task);
    }

    @Override
    public int deleteTask(Long taskId) {
        return adminTaskMapper.deleteTask(taskId);
    }

    @Override
    public String getTaskNameByImageId(Long imageId) {
        return adminTaskMapper.getTaskNameByImageId(imageId);
    }

    @Override
    public List<Long> getTaskIds(Long userId) {
        return adminTaskMapper.getTaskIds(userId);
    }

    @Override
    public List<Task> getTaskListById(List<Long> taskIds) {
        return adminTaskMapper.getTaskListById(taskIds);
    }

    @Override
    public List<Long> getTaskIdsByReviewerId(Long reviewerId) {
        return adminTaskMapper.getTaskIdsByReviewerId(reviewerId);
    }

    @Override
    public List<Task> taskList() {
        return adminTaskMapper.taskList();
    }

    @Override
    public List<Task> taskListById(List<Long> taskIds) {
        return adminTaskMapper.taskListById(taskIds);
    }

    @Override
    public List<Long> getTaskIdsByLabelId(Long labelId, int taskType) {
        return adminTaskMapper.getTaskIdsByLabelId(labelId, taskType);
    }

    @Override
    public List<Task> searchTask(String keyword) {
        return adminTaskMapper.searchTask(keyword);
    }

    @Override
    public List<Task> getNewTaskList() {
        return adminTaskMapper.getNewTaskList();
    }

    @Override
    public int updateTaskType(Long taskId, int taskType) {
        return adminTaskMapper.updateTaskType(taskId, taskType);
    }

    @Override
    public int hasData(Long taskId) {
        return adminTaskMapper.hasData(taskId);
    }
}
