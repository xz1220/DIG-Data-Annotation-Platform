package com.uestc.labelproject.entity;

import java.io.Serializable;
import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-29 下午5:02
 */
public class Data implements Serializable {
    private Long dataId;
    private Long imageId;
    private Long labelId;
    private int labelType;
    private Long userId;
    private String dataDesc;
    private List<Point> point;
    private short iscrowd;

    public Long getDataId() {
        return dataId;
    }

    public void setDataId(Long dataId) {
        this.dataId = dataId;
    }

    public Long getImageId() {
        return imageId;
    }

    public void setImageId(Long imageId) {
        this.imageId = imageId;
    }

    public Long getLabelId() {
        return labelId;
    }

    public void setLabelId(Long labelId) {
        this.labelId = labelId;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public String getDataDesc() {
        return dataDesc;
    }

    public void setDataDesc(String dataDesc) {
        this.dataDesc = dataDesc;
    }

    public List<Point> getPoint() {
        return point;
    }

    public void setPoint(List<Point> point) {
        this.point = point;
    }

    public int getLabelType() {
        return labelType;
    }

    public void setLabelType(int labelType) {
        this.labelType = labelType;
    }

    public short getIscrowd() {
        return iscrowd;
    }

    public void setIscrowd(short iscrowd) {
        this.iscrowd = iscrowd;
    }

    @Override
    public String toString() {
        return "Data{" +
                "dataId=" + dataId +
                ", imageId=" + imageId +
                ", labelId=" + labelId +
                ", labelType=" + labelType +
                ", userId=" + userId +
                ", dataDesc='" + dataDesc + '\'' +
                ", point=" + point +
                '}';
    }
}
