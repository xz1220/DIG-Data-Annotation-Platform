package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.User;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-24 下午6:44
 */
@Repository
public interface AdminUserMapper {

    /**
     * 获取用户列表
     * @return
     */
    List<User> getUserList();

    /**
     * 编辑用户
     * @param user
     * @return
     */
    int editUser(User user);

    /**
     * 添加用户
     * @param user
     * @return
     */
    int addUser(User user);

    /**
     * 删除用户
     * @param userId
     * @return
     */
    int deleteUser(Long userId);

    /**
     * 通过用户名找用户
     * @param username
     * @return
     */
    User findByUsername(String username);

    /**
     * 获取该图片的标记用户
     * @param imageId
     * @return
     */
    List<User> getPendingUserList(Long imageId);

    List<User> getVideoPendingUserList(Long videoId);

    /**
     * 获取标记用户
     * @return
     */
    List<User> getListUser();

    /**
     * 获取审核用户
     * @return
     */
    List<User> getListReviewer();

    /**
     * 获取用户
     * @param userId
     * @return
     */
    User getUserById(Long userId);

    /**
     * 获取任务数量
     * @return
     */
    int getTaskCount();

    /**
     * 获取审核任务数量
     * @return
     */
    int getReviewerCount();

    /**
     * 获取标记用户数量
     * @return
     */
    int getUserCount();
}
