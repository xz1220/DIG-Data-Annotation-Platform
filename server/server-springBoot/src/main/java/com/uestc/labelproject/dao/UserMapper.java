package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.SelfUserDetails;
import org.springframework.stereotype.Repository;


/**
 * @Auther: kiritoghy
 * @Desc:用于权限认证和登录认证
 * @Date: 19-7-23 下午7:25
 */
@Repository
public interface UserMapper {
    SelfUserDetails getUser(String username);
    String getAuthoritiesById(Long id);
}
