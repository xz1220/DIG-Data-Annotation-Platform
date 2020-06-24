package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-12 下午1:54
 */
public class VideoDataSet {
    private String VideoName;
    private Double duration;
    private List<List<Double>> timeStamps;
    private List<String> sentences;

    public String getVideoName() {
        return VideoName;
    }

    public void setVideoName(String videoName) {
        VideoName = videoName;
    }

    public Double getDuration() {
        return duration;
    }

    public void setDuration(Double duration) {
        this.duration = duration;
    }

    public List<List<Double>> getTimeStamps() {
        return timeStamps;
    }

    public void setTimeStamps(List<List<Double>> timeStamps) {
        this.timeStamps = timeStamps;
    }

    public List<String> getSentences() {
        return sentences;
    }

    public void setSentences(List<String> sentences) {
        this.sentences = sentences;
    }

    @Override
    public String toString() {
        return "VideoDataSet{" +
                "VideoName='" + VideoName + '\'' +
                ", duration=" + duration +
                ", timeStamps=" + timeStamps +
                ", sentences=" + sentences +
                '}';
    }
}
