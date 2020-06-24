package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminImageLabelMapper;
import com.uestc.labelproject.entity.Label;
import com.uestc.labelproject.service.AdminImageLabelService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

//这个类只操作了imagelabel 这个表

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午9:36
 */
@Service
public class AdminImageLabelServiceImpl implements AdminImageLabelService {

    @Autowired
    AdminImageLabelMapper adminLabelMapper;

    @Override
    public List<Label> getLabelList() {
        return adminLabelMapper.getLabelList();
    }

    @Override
    public int editLabel(Label label) throws Exception{
        Label temp = adminLabelMapper.findByLabelName(label.getLabelName());
        //不是很懂第二条件
        if(temp != null && temp.getLabelId() != label.getLabelId())throw new RuntimeException("标签已存在，修改失败");
        if(adminLabelMapper.editLabel(label) > 0) return 1;
        else throw new RuntimeException("修改失败，请重试");
    }

    @Override
    public int addLabel(Label label) throws Exception{
        Label temp = adminLabelMapper.findByLabelName(label.getLabelName());
        //这里就没有第二个条件
        if(temp != null)throw new RuntimeException("标签已存在，修改失败");
        if(adminLabelMapper.addLabel(label) > 0) return 1;
        else throw new RuntimeException("添加失败，请重试");
    }

    @Override
    public int deleteLabel(Long labelId) {
        return adminLabelMapper.deleteLabel(labelId);
    }

    @Override
    public List<Label> getLabelListByImageId(Long imageId) {
        return adminLabelMapper.getLabelByImageId(imageId);
    }

    @Override
    public List<Label> searchLabel(String keyword) {
        return adminLabelMapper.searchLabel(keyword);
    }
}
