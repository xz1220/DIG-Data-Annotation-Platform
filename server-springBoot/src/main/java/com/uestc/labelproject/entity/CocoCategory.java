package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午3:53
 */
public class CocoCategory {
    private String supercategory;
    private Long id;
    private String name;


    public String getSupercategory() {
        return supercategory;
    }

    public void setSupercategory(String supercategory) {
        this.supercategory = supercategory;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }


}
