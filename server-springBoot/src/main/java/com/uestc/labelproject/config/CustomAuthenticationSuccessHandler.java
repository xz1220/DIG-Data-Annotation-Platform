package com.uestc.labelproject.config;

import com.alibaba.fastjson.JSON;
import com.uestc.labelproject.entity.SelfUserDetails;
import com.uestc.labelproject.entity.User;
import com.uestc.labelproject.service.AdminUserService;
import com.uestc.labelproject.utils.AccessAddressUtil;
import com.uestc.labelproject.utils.JwtTokenUtil;
import com.uestc.labelproject.utils.RedisUtil;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;
import org.springframework.stereotype.Component;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * @Auther: kiritoghy
 * @Desc:登录成功处理器
 * @Date: 19-7-23 下午6:56
 */
@Slf4j
@Component
public class CustomAuthenticationSuccessHandler implements AuthenticationSuccessHandler {

    @Value("${token.expirationSeconds}")
    private int expirationSeconds;

    @Value("${token.validTime}")
    private int validTime;

    @Autowired
    RedisUtil redisUtil;
    @Autowired
    AdminUserService adminUserService;

    @Override
    public void onAuthenticationSuccess(HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse, Authentication authentication) throws IOException, ServletException {
        String ip = AccessAddressUtil.getIpAddress(httpServletRequest);
        Map<String,Object> map = new HashMap<>();
        map.put("ip",ip);

        SelfUserDetails selfUserDetails = (SelfUserDetails) authentication.getPrincipal();

        String token = JwtTokenUtil.generateToken(selfUserDetails.getUsername(), expirationSeconds, map);

        Integer expire = validTime*24*60*60*1000;
        User user = adminUserService.getUserById(selfUserDetails.getId());
        user.setPassword(null);
        String currentIp = AccessAddressUtil.getIpAddress(httpServletRequest);
        redisUtil.setTokenRefresh(token,selfUserDetails.getUsername(),currentIp);
        log.info("用户{}登录成功，信息已保存至redis",selfUserDetails.getUsername());
        Map<String,Object> res = new HashMap<>();
        res.put("token", token);
        res.put("user", user);
        httpServletResponse.setContentType("application/json;charset=utf-8");
        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genSuccessResult(res)));
    }
}
