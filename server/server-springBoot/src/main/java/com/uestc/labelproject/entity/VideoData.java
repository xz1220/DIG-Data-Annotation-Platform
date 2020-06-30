package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-7 下午8:44
 */
public class VideoData {
    private Long dataId;
    private Long videoId;
    private Long labelId;
    private Long userId;
    private Double startTime;
    private Double endTime;
    private int type;
    private String sentence;

    public Long getDataId() {
        return dataId;
    }

    public void setDataId(Long dataId) {
        this.dataId = dataId;
    }

    public Long getVideoId() {
        return videoId;
    }

    public void setVideoId(Long videoId) {
        this.videoId = videoId;
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

    public Double getStartTime() {
        return startTime;
    }

    public void setStartTime(Double startTime) {
        this.startTime = startTime;
    }

    public Double getEndTime() {
        return endTime;
    }

    public void setEndTime(Double endTime) {
        this.endTime = endTime;
    }

    public int getType() {
        return type;
    }

    public void setType(int type) {
        this.type = type;
    }

    public String getSentence() {
        return sentence;
    }

    public void setSentence(String sentence) {
        this.sentence = sentence;
    }

    @Override
    public String toString() {
        return "VideoData{" +
                "dataId=" + dataId +
                ", videoId=" + videoId +
                ", labelId=" + labelId +
                ", userId=" + userId +
                ", startTime=" + startTime +
                ", endTime=" + endTime +
                ", type=" + type +
                ", sentence='" + sentence + '\'' +
                '}';
    }
}
