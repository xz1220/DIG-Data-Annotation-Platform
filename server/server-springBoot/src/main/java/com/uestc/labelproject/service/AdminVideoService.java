package com.uestc.labelproject.service;

import com.alibaba.druid.sql.visitor.functions.Lpad;
import com.uestc.labelproject.entity.Video;
import com.uestc.labelproject.entity.VideoData;
import com.uestc.labelproject.entity.VideoLabel;
import org.apache.ibatis.annotations.Param;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 上午10:29
 */
public interface AdminVideoService {

    int addVideos(List<Video> videos);

    List<Video> getVideoList(Long taskId);

    int updateVideos(List<Video> videos);

    List<Long> getLabeledVideoIds(Long taskIds, Long userId);

    List<Long> getDataIds(Long userId, Long videoId);

    int deleteVideoData(Long userId, Long videoId);

    int deleteFinishById(Long userId, Long videoId);

    int finishVideo(Long userId, Long videoId);

    int addData(List<VideoData> videoDatas, Long userId, Long videoId);

    Video getVideo(Long videoId);

    List<VideoData> getVideoData(Long videoId, Long userId);

    int setVideoFinalVersion(Long videoId, Long userConfirmId);

    List<Long> getVideoIds(Long taskId);

    List<Long> getDataIds(List<Long> videoIds);

    int deleteVideosByTaskId(Long taskId);

    int deleteDatasByVideoId(List<Long> videoIds);

    int updateVideoTaskId(List<Video> videos,Long TaskId);
}
