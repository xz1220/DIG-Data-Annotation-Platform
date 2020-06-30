package com.uestc.labelproject.config;

import com.alibaba.fastjson.JSON;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.web.access.AccessDeniedHandler;
import org.springframework.stereotype.Component;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * 用来解决认证过的用户访问无权限资源时的异常
 * @Auther: kiritoghy
 * @Desc:权限不足处理
 * @Date: 19-7-23 下午9:59
 */
@Component
@Slf4j
public class CustomAccessDeniedHandler implements AccessDeniedHandler {
    @Override
    public void handle(HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse, AccessDeniedException e) throws IOException, ServletException {
        log.info("用户权限不够");
        /**
         * 和CustomAuthenticationEntryPoint这个类一个处理方法
         */
        httpServletResponse.setContentType("application/json;charset=utf-8");
        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genFailResult("用户权限不足")));
    }
}
