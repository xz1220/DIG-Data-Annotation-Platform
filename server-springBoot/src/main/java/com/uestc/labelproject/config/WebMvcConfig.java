package com.uestc.labelproject.config;

import com.uestc.labelproject.utils.FileUtil;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-30 上午11:26
 * @Desc: 图片路径映射
 */
@Configuration
public class WebMvcConfig implements WebMvcConfigurer {

    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry){
        registry.addResourceHandler("/image/**").addResourceLocations("file:"+ FileUtil.IMAGE_DIC);
        registry.addResourceHandler("/thumb/**").addResourceLocations("file:"+ FileUtil.IMAGE_S_DIC);
        registry.addResourceHandler("/video/**").addResourceLocations("file:" + FileUtil.VIDEO_DIC);
        registry.addResourceHandler("/videos/**").addResourceLocations("file:" + FileUtil.VIDEO_S_DIC);
    }
}
