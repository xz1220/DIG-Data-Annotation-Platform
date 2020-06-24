package com.uestc.labelproject.config;

import com.alibaba.fastjson.JSON;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.stereotype.Component;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * 用来解决匿名用户访问无权限资源时的异常
 * @Auther: kiritoghy
 * @Desc：入口处理，用于未登录等
 * @Date: 19-7-23 下午6:50
 */
@Component
@Slf4j
public class CustomAuthenticationEntryPoint implements AuthenticationEntryPoint {
    @Override
    public void commence(HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse, AuthenticationException e) throws IOException, ServletException {
        log.info("This is CustomAuthenticationEntryPoint");
        /**
         * 设置httpresponse的类型
         */
        httpServletResponse.setContentType("application/json;charset=utf-8");
        /**
         * 从里到外，先调用工具类ResultGenerator的genFailResult方法生成一个java对象，转成json写入httpresponse里面
         */
        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genFailResult("用户未认证")));
    }
}
