package com.uestc.labelproject.entity;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Collection;
import java.util.Set;

/**
 * 实现UserDetails接口
 * 因为Spring Security自己的用户信息中就只有Username、Password、Roles，如果希望加入额外的字段，那么就需要实现这个接口
 * 要注意的是getAytorities这个方法在重写的时候，权限一定是要以ROLE_开头的，因为springsecurity会自动识别这个开头，然后读取后面的字段作为权限。
 * @Auther: kiritoghy
 * @Date: 19-7-23 下午7:20
 */
public class SelfUserDetails implements UserDetails {

    private Long id;
    private String username;
    private String password;
    private Set<? extends GrantedAuthority> authorities;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return this.authorities;
    }

    public void setAuthorities(Set<? extends GrantedAuthority> authorities) {
        this.authorities = authorities;
    }

    @Override
    public String getPassword() { // 最重点Ⅰ
        return this.password;
    }

    @Override
    public String getUsername() { // 最重点Ⅱ
        return this.username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public Long getId() { return id; }

    public void setId(Long id) { this.id = id; }

    //账号是否过期
    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    //账号是否锁定
    @Override
    public boolean isAccountNonLocked() {
        return true;
    }

    //账号凭证是否未过期
    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return true;
    }


}
