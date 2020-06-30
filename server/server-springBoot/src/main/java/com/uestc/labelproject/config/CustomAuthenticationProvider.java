package com.uestc.labelproject.config;

import com.uestc.labelproject.entity.SelfUserDetails;
import com.uestc.labelproject.service.impl.SelfUserServiceImpl;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Component;

/**
 * @Auther: kiritoghy
 * @Desc:用户登录判断，可用于自定义登录方式
 * @Date: 19-7-23 下午9:27
 */
@Component
@Slf4j
public class CustomAuthenticationProvider  implements AuthenticationProvider {

    @Autowired
    SelfUserServiceImpl selfUserService;

    /**
     * 重写这个方法
     * 首先通过传进来的authentication得到用户名和密码，然后通过用户名来查数据库，如果数据库中的密码和用户传进来的密码一致，那么返回token，如果不是那么gg
     * @param authentication
     * @return
     * @throws AuthenticationException
     */
    @Override
    public Authentication authenticate(Authentication authentication) throws AuthenticationException {
        String username = authentication.getName();
        String password = (String)authentication.getCredentials();
        SelfUserDetails selfUserDetails = (SelfUserDetails) selfUserService.loadUserByUsername(username);
        log.info("用户请求的账户密码:{}/{}",username,password);
        log.info("数据库的密码:{}",selfUserDetails.getPassword());
        if(!password.equals(selfUserDetails.getPassword())){
            throw new BadCredentialsException("密码不正确");
        }
        log.info("用户权限:{}",selfUserDetails.getAuthorities());
        return new UsernamePasswordAuthenticationToken(selfUserDetails,password,selfUserDetails.getAuthorities());
    }

    @Override
    public boolean supports(Class<?> aClass) {
        return aClass.equals(UsernamePasswordAuthenticationToken.class);
    }
}
