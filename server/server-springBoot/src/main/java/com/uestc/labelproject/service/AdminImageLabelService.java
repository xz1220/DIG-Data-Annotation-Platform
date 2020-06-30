package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.Label;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午9:36
 */
public interface AdminImageLabelService {

    List<Label> getLabelList();

    int editLabel(Label label) throws Exception;

    int addLabel(Label label) throws Exception;

    int deleteLabel(Long labelId);

    List<Label> getLabelListByImageId(Long imageId);

    List<Label> searchLabel(String keyword);
}
