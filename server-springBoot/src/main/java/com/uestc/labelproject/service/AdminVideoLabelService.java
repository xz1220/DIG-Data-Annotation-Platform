package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.TempVideoLabel;
import com.uestc.labelproject.entity.VideoLabel;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午5:59
 */
public interface AdminVideoLabelService {

    List<TempVideoLabel> getVideoLabelList();

    int addVideoLabel(TempVideoLabel tempVideoLabel);

    int editVideoLabel(TempVideoLabel tempVideoLabel);

    int deleteVideoLabel(VideoLabel videoLabel);

    List<VideoLabel> getVideoLabelsByVideoId(Long videoId);
}
