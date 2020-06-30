package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.Test;

import java.util.List;

public interface TestService {
    List<Test> get();
    int addTest(Test test);
}
