package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.Task;
import org.apache.ibatis.annotations.Param;
import org.springframework.security.core.parameters.P;

import java.util.List;
import java.util.Set;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:47
 */
public interface AdminTaskService {

    Set<String> getVideoTaskNameList();

    Set<String> getImageTaskNameList();

    int addTask(Task task);

    List<Task> getTaskList();

    String getTaskNameById(Long taskId);

    int updateTask(Task task);

    int addTaskUserIds(List<Long> userIds,Long taskId);

    int addTaskLabelIds(List<Long> labelIds,Long taskId);


    int addTaskReviewerIds(List<Long> reviewerIds, Long taskdId);

    int deleteTaskUserIds(Long taskId);

    int deleteTaskLabelIds(Long taskId);

    int deleteTaskReviewerIds(Long taskId);

    int deleteTask(Long taskId);

    Task getTaskById(Long taskId);

    String getTaskNameByImageId(Long imageId);

    List<Long> getTaskIds(Long userId);

    List<Task> getTaskListById(List<Long> taskIds);

    List<Long> getTaskIdsByReviewerId(Long reviewerId);

    List<Task> taskList();

    List<Task> taskListById(List<Long> taskIds);

    List<Long> getTaskIdsByLabelId(Long labelId, int taskType);

    List<Task> searchTask(String keyword);

    List<Task> getNewTaskList();

    int updateTaskType(Long taskId,int taskType);

    int hasData(Long taskId);
}
