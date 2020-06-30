package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-9 下午3:08
 */
public class UserInfo {

    private String username;
    private Long userId;
    private Long labeled;

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public Long getLabeled() {
        return labeled;
    }

    public void setLabeled(Long labeled) {
        this.labeled = labeled;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    @Override
    public String toString() {
        return "UserInfo{" +
                "username='" + username + '\'' +
                ", userId=" + userId +
                ", labeled=" + labeled +
                '}';
    }
}
