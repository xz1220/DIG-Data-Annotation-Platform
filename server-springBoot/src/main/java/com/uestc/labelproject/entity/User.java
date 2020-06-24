package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 上午11:07
 */
public class User {

    private Long userId;
    private String username;
    private String password;
    private String authorities;

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getAuthorities() {
        return authorities;
    }

    public void setAuthorities(String authorities) {
        this.authorities = authorities;
    }

    @Override
    public String toString() {
        return "User{" +
                "userId=" + userId +
                ", username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", authorities='" + authorities + '\'' +
                '}';
    }
}
