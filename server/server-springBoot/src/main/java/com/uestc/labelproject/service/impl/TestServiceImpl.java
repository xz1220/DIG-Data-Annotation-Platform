package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.TestMapper;
import com.uestc.labelproject.entity.Test;
import com.uestc.labelproject.service.TestService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @Auther:kiritoghy
 * @Date:19-7-23 下午2:41
 */
@Service
public class TestServiceImpl implements TestService {

    @Autowired
    TestMapper testMapper;

    @Override
    public List<Test> get() {
        return testMapper.get();
    }

    @Override
    public int addTest(Test test) {
        return testMapper.addTest(test);
    }
}
