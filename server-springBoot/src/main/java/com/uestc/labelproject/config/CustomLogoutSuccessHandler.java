package com.uestc.labelproject.config;

import com.alibaba.fastjson.JSON;
import com.uestc.labelproject.utils.DateUtil;
import com.uestc.labelproject.utils.RedisUtil;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.logout.LogoutSuccessHandler;
import org.springframework.stereotype.Component;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * @Auther: kiritoghy
 * @Desc:登出
 * @Date: 19-7-23 下午8:03
 */
@Slf4j
@Component
public class CustomLogoutSuccessHandler implements LogoutSuccessHandler {

    @Autowired
    RedisUtil redisUtil;

    @Override
    public void onLogoutSuccess(HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse, Authentication authentication) throws IOException, ServletException {
        String authHeader = httpServletRequest.getHeader("Authorization");
        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            final String authToken = authHeader.substring("Bearer ".length());
            //将token放入黑名单中
            redisUtil.hset("blacklist", authToken, DateUtil.getTime());
            log.info("用户登出成功！token：{}已加入redis黑名单",authToken);
        }
        httpServletResponse.setContentType("application/json;charset=utf-8");
        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genSuccessResult("Logout Success!")));
    }
}
