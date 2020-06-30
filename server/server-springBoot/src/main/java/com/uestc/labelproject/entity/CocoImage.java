package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午3:51
 */
public class CocoImage {
    private Long id;
    private String file_name;
    private int width;
    private int height;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getFile_name() {
        return file_name;
    }

    public void setFile_name(String file_name) {
        this.file_name = file_name;
    }

    public int getWidth() {
        return width;
    }

    public void setWidth(int width) {
        this.width = width;
    }

    public int getHeight() {
        return height;
    }

    public void setHeight(int height) {
        this.height = height;
    }
}
