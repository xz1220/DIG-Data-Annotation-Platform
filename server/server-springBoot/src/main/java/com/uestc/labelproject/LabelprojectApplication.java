package com.uestc.labelproject;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.transaction.annotation.EnableTransactionManagement;

@SpringBootApplication
@EnableTransactionManagement
@MapperScan("com.uestc.labelproject.dao")
public class LabelprojectApplication {

    public static void main(String[] args) {
        SpringApplication.run(LabelprojectApplication.class, args);
    }

}
