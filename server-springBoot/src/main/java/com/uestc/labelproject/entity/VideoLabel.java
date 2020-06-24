package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午4:19
 */
public class VideoLabel {
    private Long labelId;
    private String question;
    private int type;
    private List<String> selector;

    public Long getLabelId() {
        return labelId;
    }

    public void setLabelId(Long labelId) {
        this.labelId = labelId;
    }

    public String getQuestion() {
        return question;
    }

    public void setQuestion(String question) {
        this.question = question;
    }

    public int getType() {
        return type;
    }

    public void setType(int type) {
        this.type = type;
    }

    public List<String> getSelector() {
        return selector;
    }

    public void setSelector(List<String> selector) {
        this.selector = selector;
    }

    @Override
    public String toString() {
        return "VideoLabel{" +
                "labelId=" + labelId +
                ", question='" + question + '\'' +
                ", type=" + type +
                ", selector=" + selector +
                '}';
    }
}
