package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther:kiritoghy
 * @Date:19-7-23 下午2:25
 */
public class Test {
    private int id;
    private String test;
    private String testLong;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getTest() {
        return test;
    }

    public void setTest(String test) {
        this.test = test;
    }

    public String getTestLong() {
        return testLong;
    }

    public void setTestLong(String testLong) {
        this.testLong = testLong;
    }

    @Override
    public String toString() {
        return "Test{" +
                "id=" + id +
                ", test='" + test + '\'' +
                ", testLong=" + testLong +
                '}';
    }
}
