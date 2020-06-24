package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.Test;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface TestMapper {
    List<Test> get();
    int addTest(Test test);
}
