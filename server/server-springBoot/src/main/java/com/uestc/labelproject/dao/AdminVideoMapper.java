package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.Video;
import com.uestc.labelproject.entity.VideoData;
import com.uestc.labelproject.entity.VideoLabel;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午3:11
 */
@Repository
public interface AdminVideoMapper {
    int addVideos(@Param("videos")List<Video> videos);
    List<Video> getVideoList(Long taskId);
    int updateVideos(@Param("videos") List<Video> videos);
    List<Long> getDataIds(Long userId, Long videoId);
    int deleteVideoData(Long userId, Long videoId);
    int deleteFinishById(Long userId, Long videoId);
    int finishVideo(Long userId, Long videoId);
    int addData(@Param("videoDatas") List<VideoData> videoDatas, Long userId, Long videoId);
    Video getVideo(Long videoId);
    List<VideoData> getVideoData(Long videoId, Long userId);
    int setVideoFinalVersion(Long videoId, Long userConfirmId);
    List<Long> getVideoIds(Long taskId);
    List<Long> getDataIdsByVideoId(@Param("videoIds") List<Long> videoIds);
    int deleteVideosByTaskId(Long taskId);
    int deleteDatasByVideoId(@Param("videoIds") List<Long> videoIds);
    int updateVideoTaskId(@Param("videos") List<Video> videos, Long taskId);
}
