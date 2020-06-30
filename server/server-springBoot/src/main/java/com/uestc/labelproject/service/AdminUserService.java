package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.User;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-24 下午6:43
 */
public interface AdminUserService {

    List<User> getUserList();

    int editUser(User user);

    int addUser(User user);

    int deleteUser(Long userId);

    List<User> getPendingUserList(Long imageId);

    List<User> getVideoPendingUserList(Long videoId);

    List<User> getListUser();

    List<User> getListReviewer();

    User getUserById(Long userId);

    int getTaskCount();

    int getReviewerCount();

    int getUserCount();
}
