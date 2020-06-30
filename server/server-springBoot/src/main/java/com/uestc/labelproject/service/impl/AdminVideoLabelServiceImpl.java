package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminVideoLabelMapper;
import com.uestc.labelproject.entity.TempVideoLabel;
import com.uestc.labelproject.entity.VideoLabel;
import com.uestc.labelproject.service.AdminVideoLabelService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午6:01
 */
@Service
public class AdminVideoLabelServiceImpl implements AdminVideoLabelService {

    @Autowired
    AdminVideoLabelMapper adminVideoLabelMapper;

    @Override
    public List<TempVideoLabel> getVideoLabelList() {
        return adminVideoLabelMapper.getVideoLabelList();
    }

    @Override
    public int addVideoLabel(TempVideoLabel tempVideoLabel) {
        return adminVideoLabelMapper.addVideoLabel(tempVideoLabel);
    }

    @Override
    public int editVideoLabel(TempVideoLabel tempVideoLabel) {
        return adminVideoLabelMapper.editVideoLabel(tempVideoLabel);
    }

    @Override
    public int deleteVideoLabel(VideoLabel videoLabel) {
        return adminVideoLabelMapper.deleteVideoLabel(videoLabel);
    }

    @Override
    public List<VideoLabel> getVideoLabelsByVideoId(Long videoId) {
        return adminVideoLabelMapper.getVideoLabelsByVideoId(videoId);
    }
}
