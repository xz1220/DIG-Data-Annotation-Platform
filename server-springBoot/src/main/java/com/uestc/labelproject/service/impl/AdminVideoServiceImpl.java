package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminImageMapper;
import com.uestc.labelproject.dao.AdminVideoMapper;
import com.uestc.labelproject.entity.Video;
import com.uestc.labelproject.entity.VideoData;
import com.uestc.labelproject.entity.VideoLabel;
import com.uestc.labelproject.service.AdminVideoService;
import org.springframework.beans.CachedIntrospectionResults;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 上午10:30
 */
@Service
public class AdminVideoServiceImpl implements AdminVideoService {
    @Autowired
    AdminVideoMapper adminVideoMapper;
    @Autowired
    AdminImageMapper adminImageMapper;

    @Override
    public int addVideos(List<Video> videos) {
        return adminVideoMapper.addVideos(videos);
    }

    @Override
    public List<Video> getVideoList(Long taskId) {
        return adminVideoMapper.getVideoList(taskId);
    }

    @Override
    public int updateVideos(List<Video> videos) {
        return adminVideoMapper.updateVideos(videos);
    }

    @Override
    public List<Long> getLabeledVideoIds(Long taskIds, Long userId) {
        return adminImageMapper.getLabeledImageIds(taskIds, userId);
    }

    @Override
    public List<Long> getDataIds(Long userId, Long videoId) {
        return adminVideoMapper.getDataIds(userId, videoId);
    }

    @Override
    public int deleteVideoData(Long userId, Long videoId) {
        return adminVideoMapper.deleteVideoData(userId, videoId);
    }

    @Override
    public int deleteFinishById(Long userId, Long videoId) {
        return adminVideoMapper.deleteFinishById(userId,videoId);
    }

    @Override
    public int finishVideo(Long userId, Long videoId) {
        return adminVideoMapper.finishVideo(userId,videoId);
    }

    @Override
    public int addData(List<VideoData> videoDatas, Long userId, Long videoId) {
        return adminVideoMapper.addData(videoDatas,userId, videoId);
    }

    @Override
    public Video getVideo(Long videoId) {
        return adminVideoMapper.getVideo(videoId);
    }


    @Override
    public List<VideoData> getVideoData(Long videoId, Long userId) {
        return adminVideoMapper.getVideoData(videoId, userId);
    }

    @Override
    public int setVideoFinalVersion(Long videoId, Long userConfirmId) {
        return adminVideoMapper.setVideoFinalVersion(videoId, userConfirmId);
    }

    @Override
    public List<Long> getVideoIds(Long taskId) {
        return adminVideoMapper.getVideoIds(taskId);
    }

    @Override
    public List<Long> getDataIds(List<Long> videoIds) {
        return adminVideoMapper.getDataIdsByVideoId(videoIds);
    }

    @Override
    public int deleteVideosByTaskId(Long taskId) {
        return adminVideoMapper.deleteVideosByTaskId(taskId);
    }

    @Override
    public int deleteDatasByVideoId(List<Long> videoIds) {
        return adminVideoMapper.deleteDatasByVideoId(videoIds);
    }

    @Override
    public int updateVideoTaskId(List<Video> videos, Long TaskId) {
        return adminVideoMapper.updateVideoTaskId(videos,TaskId);
    }
}
