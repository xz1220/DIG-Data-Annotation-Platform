package com.uestc.labelproject.config;

import com.uestc.labelproject.entity.SelfUserDetails;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;
import org.springframework.util.AntPathMatcher;

import javax.servlet.http.HttpServletRequest;
import java.util.Collection;
import java.util.HashSet;
import java.util.Iterator;
import java.util.Set;
import java.util.concurrent.CancellationException;

/**
 * 在此解释一下component注解，在spring中泛指各种组件，这个不属于@services @controller等的时候，我们就可以使用这个
 * @Auther: kiritoghy
 * @Desc:角色控制
 * @Date: 19-7-23 下午8:11
 */
@Slf4j
@Component("rbacauthorityservice")
public class RbacAuthorityService {
    public boolean hasPermission(HttpServletRequest request, Authentication authentication) {
        log.info("try get permission");
        Object userInfo = authentication.getPrincipal();

        boolean hasPermission  = false;

        if (userInfo instanceof UserDetails) {

            Collection<? extends GrantedAuthority> set = ((UserDetails) userInfo).getAuthorities();
            Iterator<? extends GrantedAuthority> iterator = set.iterator();
            while(iterator.hasNext()){
                String role = iterator.next().getAuthority();
                // 这些 url 都是要登录后才能访问，且其他的 url 都不能访问！
                Set<String> urls = new HashSet();

                //获取资源
                switch (role){
                    case "ROLE_ADMIN":
                        urls.add("/api/admin/**");
                        break;
                    case "ROLE_USER":
                        urls.add("/api/user/**");
                        break;
                    case "ROLE_REVIEWER":
                        urls.add("/api/reviewer/**");
                        break;
                }

                AntPathMatcher antPathMatcher = new AntPathMatcher();
                for (String url : urls) {
                    if (antPathMatcher.match(url, request.getRequestURI())) {
                        hasPermission = true;
                        break;
                    }
                }
            }
            log.info("url:{},取得permission:{}",request.getRequestURI(),hasPermission);
            return hasPermission;
        }
        else {
            log.info("url:{},取得permission:{}",request.getRequestURI(),false);
            return false;
        }
    }
}
