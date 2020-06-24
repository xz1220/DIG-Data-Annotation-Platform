package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:17
 */
public class Image {
    private Long imageId;
    private String imageName;
    private String imageThumb;
    private Long taskId;
    private Long userConfirmId;
    private Integer width;
    private Integer height;

    public Long getImageId() {
        return imageId;
    }

    public void setImageId(Long imageId) {
        this.imageId = imageId;
    }

    public String getImageName() {
        return imageName;
    }

    public void setImageName(String imageName) {
        this.imageName = imageName;
    }

    public String getImageThumb() {
        return imageThumb;
    }

    public void setImageThumb(String imageThumb) {
        this.imageThumb = imageThumb;
    }

    public Long getTaskId() {
        return taskId;
    }

    public void setTaskId(Long taskId) {
        this.taskId = taskId;
    }


    public Long getUserConfirmId() {
        return userConfirmId;
    }

    public void setUserConfirmId(Long userConfirmId) {
        this.userConfirmId = userConfirmId;
    }

    public Integer getWidth() {
        return width;
    }

    public void setWidth(Integer width) {
        this.width = width;
    }

    public Integer getHeight() {
        return height;
    }

    public void setHeight(Integer height) {
        this.height = height;
    }

    @Override
    public String toString() {
        return "Image{" +
                "imageId=" + imageId +
                ", imageName='" + imageName + '\'' +
                ", imageThumb='" + imageThumb + '\'' +
                ", taskId=" + taskId +
                ", userConfirmId=" + userConfirmId +
                ", width=" + width +
                ", height=" + height +
                '}';
    }
}
