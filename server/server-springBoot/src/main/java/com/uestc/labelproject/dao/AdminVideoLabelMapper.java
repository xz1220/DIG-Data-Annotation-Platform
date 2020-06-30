package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.TempVideoLabel;
import com.uestc.labelproject.entity.VideoLabel;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午6:05
 */
@Repository
public interface AdminVideoLabelMapper {
    List<TempVideoLabel> getVideoLabelList();
    int addVideoLabel(TempVideoLabel tempVideoLabel);
    int editVideoLabel(TempVideoLabel tempVideoLabel);
    int deleteVideoLabel(VideoLabel videoLabel);
    List<VideoLabel> getVideoLabelsByVideoId(Long videoId);
}
