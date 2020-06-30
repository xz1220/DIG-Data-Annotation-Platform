package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午3:51
 */
public class CocoAnnotation {
    private Long id;
    private Long image_id;
    private Long category_id;
    private double area;
    private short iscrowd;
    private Object segmentation;
    private List<Double> bbox;
    private String desc;

    public void setDesc(String desc){ this.desc=desc;}

    public String getDesc() {
        return desc;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getImage_id() {
        return image_id;
    }

    public void setImage_id(Long image_id) {
        this.image_id = image_id;
    }

    public Long getCategory_id() {
        return category_id;
    }

    public void setCategory_id(Long category_id) {
        this.category_id = category_id;
    }

    public double getArea() {
        return area;
    }

    public void setArea(double area) {
        this.area = area;
    }

    public short getIscrowd() {
        return iscrowd;
    }

    public void setIscrowd(short iscrowd) {
        this.iscrowd = iscrowd;
    }

    public Object getSegmentation() {
        return segmentation;
    }

    public void setSegmentation(Object segmentation) {
        this.segmentation = segmentation;
    }

    public List<Double> getBbox() {
        return bbox;
    }

    public void setBbox(List<Double> bbox) {
        this.bbox = bbox;
    }
}
