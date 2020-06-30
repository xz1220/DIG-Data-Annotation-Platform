package com.uestc.labelproject.entity;

import java.io.Serializable;
import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午5:07
 */
public class RleData implements Serializable {
    private List<Integer> counts;
    private List<Integer> size;

    public List<Integer> getCounts() {
        return counts;
    }

    public void setCounts(List<Integer> counts) {
        this.counts = counts;
    }

    public List<Integer> getSize() {
        return size;
    }

    public void setSize(List<Integer> size) {
        this.size = size;
    }
}
