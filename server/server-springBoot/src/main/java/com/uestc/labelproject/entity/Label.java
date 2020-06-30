package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午9:30
 */
public class Label {

    private Long labelId;
    private String labelName;
    private int labelType;
    private String labelColor;

    public Long getLabelId() {
        return labelId;
    }

    public void setLabelId(Long labelId) {
        this.labelId = labelId;
    }

    public String getLabelName() {
        return labelName;
    }

    public void setLabelName(String labelName) {
        this.labelName = labelName;
    }

    public int getLabelType() {
        return labelType;
    }

    public void setLabelType(int labelType) {
        this.labelType = labelType;
    }

    public String getLabelColor() {
        return labelColor;
    }

    public void setLabelColor(String labelColor) {
        this.labelColor = labelColor;
    }

    @Override
    public String toString() {
        return "Label{" +
                "labelId=" + labelId +
                ", labelName='" + labelName + '\'' +
                ", labelType=" + labelType +
                ", labelColor='" + labelColor + '\'' +
                '}';
    }
}
