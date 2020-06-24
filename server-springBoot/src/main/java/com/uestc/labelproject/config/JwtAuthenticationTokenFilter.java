package com.uestc.labelproject.config;

import com.alibaba.fastjson.JSON;
import com.uestc.labelproject.service.impl.SelfUserServiceImpl;
import com.uestc.labelproject.utils.*;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * @Auther: kiritoghy
 * @Desc:验证token，判断是否登录
 * @Date: 19-7-23 下午8:40
 */
@Component
@Slf4j
public class JwtAuthenticationTokenFilter extends OncePerRequestFilter {
    @Value("${token.expirationSeconds}")
    private int expirationSeconds;

    @Value("${token.validTime}")
    private int validTime;

    @Autowired
    SelfUserServiceImpl selfUserService;

    @Autowired
    RedisUtil redisUtil;
    @Override
    protected void doFilterInternal(HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse, FilterChain filterChain) throws ServletException, IOException {
        String authHeader = httpServletRequest.getHeader("Authorization");
        //获取请求的ip地址
        String currentIp = AccessAddressUtil.getIpAddress(httpServletRequest);

        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            String authToken = authHeader.substring("Bearer ".length());
            String username = JwtTokenUtil.parseToken(authToken);
            String ip = (String)JwtTokenUtil.getClaims(authToken).get("ip");

            //进入黑名单验证
            if (redisUtil.isBlackList(authToken)) {
                log.info("用户：{}的token：{}在黑名单之中，拒绝访问",username,authToken);
                httpServletResponse.setContentType("application/json;charset=utf-8");
                httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genFailResult("Token失效，请重新登录")));
                return;
            }

            //判断token是否过期
            /*
             * 过期的话，从redis中读取有效时间（比如七天登录有效），再refreshToken（根据以后业务加入，现在直接refresh）
             * 同时，已过期的token加入黑名单
             */
            if (redisUtil.hasKey(authToken)) {//判断redis是否有保存
                String expirationTime = redisUtil.hget(authToken,"expirationTime").toString();
                if (JwtTokenUtil.isExpiration(expirationTime)) {
                    //获得redis中用户的token刷新时效
                    String tokenValidTime = (String) redisUtil.getTokenValidTimeByToken(authToken);
                    String currentTime = DateUtil.getTime();
                    //这个token已作废，加入黑名单
                    log.info("{}已作废，加入黑名单",authToken);
                    redisUtil.hset("blacklist", authToken, DateUtil.getTime());

                    if (DateUtil.compareDate(currentTime, tokenValidTime)) {
                        //超过有效期，不予刷新
                        log.info("{}已超过有效期，不予刷新",authToken);
                        httpServletResponse.setContentType("application/json;charset=utf-8");
                        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genFailResult("Token失效，请重新登录")));
                        return;
                    } else {//仍在刷新时间内，则刷新token，放入请求头中
                        String usernameByToken = (String) redisUtil.getUsernameByToken(authToken);
                        username = usernameByToken;//更新username

                        ip = (String) redisUtil.getIPByToken(authToken);//更新ip

                        //获取请求的ip地址
                        Map<String, Object> map = new HashMap<>();
                        map.put("ip", ip);
                        String jwtToken = JwtTokenUtil.generateToken(usernameByToken, expirationSeconds, map);


                        //更新redis
                        Integer expire = validTime * 24 * 60 * 60 * 1000;//刷新时间
                        redisUtil.setTokenRefresh(jwtToken,usernameByToken,ip);
                        //删除旧的token保存的redis
                        redisUtil.deleteKey(authToken);
                        //新的token保存到redis中
                        redisUtil.setTokenRefresh(jwtToken,username,ip);

                        log.info("token仍在有效期，redis已删除旧token：{},\n新token：{}已更新redis",authToken,jwtToken);
                        authToken = jwtToken;//更新token，为了后面
                        httpServletResponse.setHeader("Authorization", "Bearer " + jwtToken);
                    }
                }

            }

            if (username != null && SecurityContextHolder.getContext().getAuthentication() == null) {

                /*
                 * 加入对ip的验证
                 * 如果ip不正确，进入黑名单验证
                 */
                if (!StringUtil.equals(ip, currentIp)) {//地址不正确
                    log.info("用户：{}的ip地址变动，进入黑名单校验",username);
                    //进入黑名单验证
                    if (redisUtil.isBlackList(authToken)) {
                        log.info("用户：{}的token：{}在黑名单之中，拒绝访问",username,authToken);
                        httpServletResponse.setContentType("application/json;charset=utf-8");
                        httpServletResponse.getWriter().write(JSON.toJSONString(ResultGenerator.genFailResult("Token失效，请重新登录")));
                        return;
                    }

                }//黑名单没有则继续，如果黑名单存在就退出后面


                UserDetails userDetails = selfUserService.loadUserByUsername(username);
                if (userDetails != null) {
                    UsernamePasswordAuthenticationToken authentication =
                            new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
                    authentication.setDetails(new WebAuthenticationDetailsSource().buildDetails(httpServletRequest));

                    SecurityContextHolder.getContext().setAuthentication(authentication);
                }
            }
        }
        filterChain.doFilter(httpServletRequest, httpServletResponse);
    }
}

