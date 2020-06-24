package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.UserMapper;
import com.uestc.labelproject.entity.SelfUserDetails;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Component;

import java.util.HashSet;
import java.util.Set;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-23 下午7:24
 */
@Component
public class SelfUserServiceImpl implements UserDetailsService {

    @Autowired
    UserMapper userMapper;

    // 调用DAO层的接口，判断用户是否存在，然后给用户授权
    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {

        SelfUserDetails user = userMapper.getUser(username); //得到用户实体类

        if(user == null){
            throw new BadCredentialsException("该用户不存在");
        }

        String authorities = userMapper.getAuthoritiesById(user.getId());
        Set authoritiesSet = new HashSet();
        GrantedAuthority authority = new SimpleGrantedAuthority(authorities);   //这边的具体功能还不是很熟
        authoritiesSet.add(authority);
        user.setAuthorities(authoritiesSet);

        return user;
    }
}
