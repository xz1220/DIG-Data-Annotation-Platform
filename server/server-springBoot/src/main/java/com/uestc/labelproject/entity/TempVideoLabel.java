package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午6:21
 */
public class TempVideoLabel {
    private Long labelId;
    private String question;
    private int type;
    private String selector;

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

    public String getSelector() {
        return selector;
    }

    public void setSelector(String selector) {
        this.selector = selector;
    }

    @Override
    public String toString() {
        return "TempVideoLabel{" +
                "labelId=" + labelId +
                ", question='" + question + '\'' +
                ", type=" + type +
                ", selector='" + selector + '\'' +
                '}';
    }
}
