package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 上午10:10
 */
public class Video {
    private Long videoId;
    private String videoName;
    private String videoThumb;
    private Long taskId;
    private Double duration;
    private Long userConfirmId;


    public Long getVideoId() {
        return videoId;
    }

    public void setVideoId(Long videoId) {
        this.videoId = videoId;
    }

    public String getVideoName() {
        return videoName;
    }

    public void setVideoName(String videoName) {
        this.videoName = videoName;
    }

    public String getVideoThumb() {
        return videoThumb;
    }

    public void setVideoThumb(String videoThumb) {
        this.videoThumb = videoThumb;
    }

    public Long getTaskId() {
        return taskId;
    }

    public void setTaskId(Long taskId) {
        this.taskId = taskId;
    }

    public Double getDuration() {
        return duration;
    }

    public void setDuration(Double duration) {
        this.duration = duration;
    }

    public Long getUserConfirmId() {
        return userConfirmId;
    }

    public void setUserConfirmId(Long userConfirmId) {
        this.userConfirmId = userConfirmId;
    }

    @Override
    public String toString() {
        return "Video{" +
                "videoId=" + videoId +
                ", videoName='" + videoName + '\'' +
                ", videoThumb='" + videoThumb + '\'' +
                ", taskId=" + taskId +
                ", duration='" + duration + '\'' +
                ", userConfirmId=" + userConfirmId +
                '}';
    }
}
