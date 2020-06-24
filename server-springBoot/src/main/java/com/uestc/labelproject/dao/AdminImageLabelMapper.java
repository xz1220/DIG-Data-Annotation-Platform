package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.Label;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午9:30
 */
@Repository
public interface AdminImageLabelMapper {

    /**
     * 获取标签列表
     * @return
     */
    List<Label> getLabelList();

    /**
     * 修改标签
     * @param label
     * @return
     */
    int editLabel(Label label);

    /**
     * 添加标签
     * @param label
     * @return
     */
    int addLabel(Label label);

    /**
     * 删除标签
     * @param labelId
     * @return
     */
    int deleteLabel(Long labelId);

    /**
     * 通过标签名找标签
     * @param labelName
     * @return
     */
    Label findByLabelName(String labelName);

    /**
     * 获取该图片的标签
     * @param imageId
     * @return
     */
    List<Label> getLabelByImageId(Long imageId);

    List<Label> searchLabel(String keyword);
}
