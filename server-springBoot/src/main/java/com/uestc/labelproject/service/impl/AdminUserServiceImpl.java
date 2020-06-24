package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminUserMapper;
import com.uestc.labelproject.entity.User;
import com.uestc.labelproject.service.AdminUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-24 下午6:43
 */
@Service
public class AdminUserServiceImpl implements AdminUserService {

    @Autowired
    AdminUserMapper adminUserMapper;

    @Override
    public List<User> getUserList() {
        return adminUserMapper.getUserList();
    }

    @Override
    public int editUser(User user) {
        User temp = adminUserMapper.findByUsername(user.getUsername());
        if(temp != null && temp.getUserId() != user.getUserId())return 0;
        if(adminUserMapper.editUser(user) > 0) return 1;
        else return 0;
    }

    @Override
    public int addUser(User user) {
        if(adminUserMapper.findByUsername(user.getUsername()) != null)return 0;
        if(adminUserMapper.addUser(user) > 0) return 1;
        else return 0;
    }

    @Override
    public int deleteUser(Long userId) {
        return adminUserMapper.deleteUser(userId);
    }
    @Override
    public List<User> getPendingUserList(Long imageId) {
        return adminUserMapper.getPendingUserList(imageId);
    }

    @Override
    public List<User> getVideoPendingUserList(Long videoId) {
        return adminUserMapper.getVideoPendingUserList(videoId);
    }

    @Override
    public List<User> getListUser() {
        return adminUserMapper.getListUser();
    }

    @Override
    public List<User> getListReviewer() {
        return adminUserMapper.getListReviewer();
    }

    @Override
    public User getUserById(Long userId) {
        return adminUserMapper.getUserById(userId);
    }

    @Override
    public int getTaskCount() {
        return adminUserMapper.getTaskCount();
    }

    @Override
    public int getReviewerCount() {
        return adminUserMapper.getReviewerCount();
    }

    @Override
    public int getUserCount() {
        return adminUserMapper.getUserCount();
    }
}
