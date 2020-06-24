package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.Task;
import com.uestc.labelproject.entity.UserInfo;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Set;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午10:07
 */
@Repository
public interface AdminTaskMapper {

    /**
     * 获取任务名
     * @return
     */
    Set<String> getImageTaskNameList();
    Set<String> getVideoTaskNameList();

    /**
     * 添加任务
     * @param task
     * @return
     */
    int addTask(Task task);

    /**
     * 获取任务列表
     * @return
     */
    List<Task> getTaskList();

    /**
     * 添加标记用户
     * @param userIds
     * @param taskId
     * @return
     */
    int addTaskUserIds(@Param("userIds") List<Long> userIds, @Param("taskId")  Long taskId);

    /**
     * 获取任务名
     * @param taskId
     * @return
     */
    String getTaskNameById(Long taskId);

    /**
     * 添加任务标签
     * @param labelIds
     * @param taskId
     * @return
     */
    int addTaskLabelIds(@Param("labelIds") List<Long> labelIds,@Param("taskId") Long taskId);

    /**
     * 添加审核用户
     * @param labelIds
     * @param taskId
     * @return
     */
    int addTaskReviewerIds(@Param("reviewerIds") List<Long> labelIds,@Param("taskId") Long taskId);

    /**
     * 删除标记用户
     * @param taskId
     * @return
     */
    int deleteTaskUserIds(Long taskId);

    /**
     * 删除任务标签
     * @param taskId
     * @return
     */
    int deleteTaskLabelIds(Long taskId);

    /**
     * 删除审核用户
     * @param taskId
     * @return
     */
    int deleteTaskReviewerIds(Long taskId);

    /**
     * 更新任务
     * @param task
     * @return
     */
    int updateTask(Task task);

    /**
     * 删除任务
     * @param taskId
     * @return
     */
    int deleteTask(Long taskId);

    /**
     * 获取任务
     * @param taskId
     * @return
     */
    Task getTaskById(@Param("taskId") Long taskId);

    /**
     * 获取任务名
     * @param imageId
     * @return
     */
    String getTaskNameByImageId(Long imageId);

    /**
     * 获取任务id
     * @param userId
     * @return
     */
    List<Long> getTaskIds(Long userId);

    /**
     * 通过任务id列表获取任务列表
     * @param taskIds
     * @return
     */
    List<Task> getTaskListById(@Param("taskIds") List<Long> taskIds);

    /**
     * 获取审核任务列表
     * @param reviewerId
     * @return
     */
    List<Long> getTaskIdsByReviewerId(Long reviewerId);

    /**
     * 获取任务列表
     * @return
     */
    List<Task> taskList();
    List<UserInfo> getUserInfo(@Param("taskId") Long taskId, @Param("userId") Long userId);
    List<String> getReviewerInfo(@Param("taskId") Long taskId, @Param("reviewerId") Long reviewerId);

    /**
     * 获取任务列表
     * @param taskIds
     * @return
     */
    List<Task> taskListById(@Param("taskIds") List<Long> taskIds);

    List<Long> getTaskIdsByLabelId(Long labelId, int taskType);

    List<Task> searchTask(String keyword);
    List<Task> getNewTaskList();
    int updateTaskType(Long taskId,int taskType);
    int hasData(Long taskId);
}
