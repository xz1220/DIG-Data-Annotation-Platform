package com.uestc.labelproject.utils;

import javax.servlet.http.HttpServletRequest;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午6:47
 */
public class LogUtil {

    public static String getUsername(HttpServletRequest request){
        String authHeader = request.getHeader("Authorization");
        if (authHeader != null && authHeader.startsWith("Bearer ")){
            String authToken = authHeader.substring("Bearer ".length());
            String username = JwtTokenUtil.parseToken(authToken);
            return username;
        }
        return null;
    }
}
