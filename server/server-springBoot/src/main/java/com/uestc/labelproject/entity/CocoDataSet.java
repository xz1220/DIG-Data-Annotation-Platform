package com.uestc.labelproject.entity;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午3:49
 */
public class CocoDataSet {
    private CocoInfo info;
    private List<CocoImage> images;
    private List<CocoAnnotation> annotations;
    private List<CocoCategory> categories;

    public CocoInfo getInfo() {
        return info;
    }

    public void setInfo(CocoInfo info) {
        this.info = info;
    }

    public List<CocoImage> getImages() {
        return images;
    }

    public void setImages(List<CocoImage> images) {
        this.images = images;
    }

    public List<CocoAnnotation> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(List<CocoAnnotation> annotations) {
        this.annotations = annotations;
    }

    public List<CocoCategory> getCategories() {
        return categories;
    }

    public void setCategories(List<CocoCategory> categories) {
        this.categories = categories;
    }
}
