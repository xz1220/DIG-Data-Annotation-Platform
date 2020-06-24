package com.uestc.labelproject.config;

import com.uestc.labelproject.entity.SelfUserDetails;
import com.uestc.labelproject.service.impl.SelfUserServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

/**
 * @Auther: kiritoghy
 * @Desc:权限配置
 * @Date: 19-7-23 下午1:32
 */

/**
 *该注解和 @Configuration 注解一起使用, 注解 WebSecurityConfigurer 类型的类，或者利用@EnableWebSecurity 注解继承 WebSecurityConfigurerAdapter的类，这样就构成了 Spring Security 的配置。
 *WebSecurityConfigurerAdapter 提供了一种便利的方式去创建 WebSecurityConfigurer的实例，只需要重写 WebSecurityConfigurerAdapter 的方法，即可配置拦截什么URL、设置什么权限等安全控制。
 */
@EnableWebSecurity
@Configuration
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

    @Autowired
    CustomAuthenticationEntryPoint customAuthenticationEntryPoint;
    @Autowired
    CustomAuthenticationSuccessHandler customAuthenticationSuccessHandler;
    @Autowired
    CustomAuthenticationFailureHandler customAuthenticationFailureHandler;
    @Autowired
    JwtAuthenticationTokenFilter jwtAuthenticationTokenFilter;
    @Autowired
    SelfUserServiceImpl selfUserService;
    @Autowired
    CustomAuthenticationProvider customAuthenticationProvider;
    @Autowired
    CustomAccessDeniedHandler customAccessDeniedHandler;
    @Autowired
    CustomLogoutSuccessHandler customLogoutSuccessHandler;

    @Override
    public void configure(WebSecurity web) throws Exception {
        //ignore
        web.ignoring().antMatchers("/test/**");
        web.ignoring().antMatchers("/image/**","/thumb/**", "/video/**", "/videos/**");
    }

    /**
     * AuthenticationManagerBuilder 用于创建一个 AuthenticationManager，让我能够轻松的实现内存验证、LADP验证、基于JDBC的验证、添加UserDetailsService、添加AuthenticationProvider。
     * 在这里实现了基于数据库的用户名密码的认证，通过重写了customAuthenticationProvider中的authenticate方法
     * 然后再selfUserService中调用DAO 层接口，通过之前验证过的用户名来获取数据库中对应用户的权限，然后进行授权，返回User
     * @param auth
     * @throws Exception
     */
    @Override
    public void configure(AuthenticationManagerBuilder auth) throws Exception {
        auth.authenticationProvider(customAuthenticationProvider);
        auth.userDetailsService(selfUserService).passwordEncoder(new BCryptPasswordEncoder());
    }
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        //http.authorizeRequests().antMatchers("/test/**").permitAll();
        http.csrf().disable()//不禁用有时候会弹出一些其他的奇怪的东西
                .sessionManagement() //配置会话管理
                /**
                 * SessionCreationPolicy几种属性
                 * always 保存httpsession 状态，每次session都保存，可能会导致内存溢出
                 * never 不会创建httpsession
                 * if_required 仅在需要的时候船舰httpsession
                 * staleless 不会保存session状态 （默认配置）
                 */
                .sessionCreationPolicy(SessionCreationPolicy.STATELESS) // 使用 JWT，关闭token
                .and()
                /**
                 * Basic认证是一种较为简单的HTTP认证方式，客户端通过明文（Base64编码格式）传输用户名和密码到服务端进行认证，通常需要配合HTTPS来保证信息传输的安全。
                 */
                .httpBasic() //配置 Http Basic 验证
                .authenticationEntryPoint(customAuthenticationEntryPoint) /** 匿名用户访问资源处理异常*/

                .and()
                .authorizeRequests() //定义哪些URL需要被保护、哪些不需要被保护
                .anyRequest().access("@rbacauthorityservice.hasPermission(request, authentication)") /** 每个请求都要进行验证权限*/
                .and()
                .formLogin().loginProcessingUrl("/login")
                .successHandler(customAuthenticationSuccessHandler) /**验证权限成功之后的处理*/
                .failureHandler(customAuthenticationFailureHandler) /** 失败以后的处理*/
                .permitAll() //登录相关的请求不需要验证

                .and()
                .logout()//默认注销行为为logout
                .logoutUrl("/logout")
                .logoutSuccessHandler(customLogoutSuccessHandler)
                .permitAll(); //退出的操作也不做任何验证
        http.exceptionHandling().accessDeniedHandler(customAccessDeniedHandler); // 无权访问 JSON 格式的数据 /*意外处理
        http.addFilterBefore(jwtAuthenticationTokenFilter, UsernamePasswordAuthenticationFilter.class);
        http.cors();
    }
}
