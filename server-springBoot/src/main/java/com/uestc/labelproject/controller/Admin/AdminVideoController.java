package com.uestc.labelproject.controller.Admin;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.uestc.labelproject.entity.Video;
import com.uestc.labelproject.entity.VideoData;
import com.uestc.labelproject.entity.VideoLabel;
import com.uestc.labelproject.service.AdminTaskService;
import com.uestc.labelproject.service.AdminVideoLabelService;
import com.uestc.labelproject.service.AdminVideoService;
import com.uestc.labelproject.utils.FileUtil;
import com.uestc.labelproject.utils.Result;
import com.uestc.labelproject.utils.ResultGenerator;
import com.uestc.labelproject.utils.VideoUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.io.File;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-7 下午5:32
 */
@Slf4j
@RestController
@RequestMapping("/admin")
public class AdminVideoController {

    @Autowired
    AdminVideoService adminVideoService;
    @Autowired
    AdminTaskService adminTaskService;
    @Autowired
    AdminVideoLabelService adminVideoLabelService;

    @PostMapping("getVideoList")
    public Result getVideoList(@RequestBody Map<String, Object> param) throws IOException {
        log.info("收到的参数{}",param);
        int page = (int)param.get("page");
        int limit = (int)param.get("limit");
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        String taskName = adminTaskService.getTaskNameById(taskId);
        if (taskName == null) {
            return ResultGenerator.genFailResult("任务不存在");
        }
        log.info("page:{}, limit:{}, taskName:{}",page, limit, taskName);
        PageHelper.startPage(page,limit);
        List<Video> list = adminVideoService.getVideoList(taskId);
        PageInfo pageInfo = new PageInfo(list);
        if(list.size() > 0){
            for(Video video: list){
                String picPath = FileUtil.VIDEO_S_DIC + taskName;
                String src = FileUtil.VIDEO_DIC + taskName + "/" + video.getVideoName();
                File file = new File(src);
                if(video.getVideoThumb() == null){
                    String thumb = VideoUtil.getThumb(file,picPath);
                    video.setVideoThumb(thumb);
                }
            }
            adminVideoService.updateVideos(list);
        }

        List<Long> labelVideoIds = adminVideoService.getLabeledVideoIds(taskId, null);
        int totalPages = pageInfo.getPages();
        Map<String,Object> data = new HashMap<>();
        data.put("page", page);
        data.put("limit", limit);
        data.put("totalpages",totalPages);
        data.put("videos", list);
        data.put("labelVideoIds", labelVideoIds);
        return ResultGenerator.genSuccessResult(data);
    }

    @PostMapping("/saveVideoLabel")
    @Transactional(isolation = Isolation.READ_COMMITTED)
    public Result saveVideoLabel(@RequestBody String jsonStr){
        log.info("收到的参数{}",jsonStr);
        JSONObject jsonObject = JSONObject.parseObject(jsonStr);
        String datas = jsonObject.getString("data");
        List<VideoData> videoDataList = JSONArray.parseArray(datas, VideoData.class);
        Long userId = JSONObject.parseObject(jsonObject.getString("userId"), Long.class);
        Long videoId = JSONObject.parseObject(jsonObject.getString("videoId"), Long.class);

        List<Long> dataIds = adminVideoService.getDataIds(userId, videoId);
        if(dataIds.size() > 0){
            if (adminVideoService.deleteVideoData(userId, videoId) <= 0) {
                return ResultGenerator.genFailResult("保存失败");
            }
            if (adminVideoService.deleteFinishById(userId,videoId) <= 0) {
                return ResultGenerator.genFailResult("保存失败");
            }
        }
        if(videoDataList != null && videoDataList.size() > 0){
            if (adminVideoService.addData(videoDataList, userId, videoId) <= 0) {
                return ResultGenerator.genFailResult("保存失败");
            }
            if (adminVideoService.finishVideo(userId, videoId) <= 0) {
                return ResultGenerator.genFailResult("保存失败");
            }
        }
        return ResultGenerator.genSuccessResult("保存成功");
    }

    @PostMapping("getVideo")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getVideo(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long videoId = Long.parseLong(String.valueOf(param.get("videoId")));
        Long userId = Long.parseLong(String.valueOf(param.get("userId")));
        Video video = adminVideoService.getVideo(videoId);
        List<VideoLabel> videoLabels = adminVideoLabelService.getVideoLabelsByVideoId(videoId);
        List<VideoData> videoDataList = adminVideoService.getVideoData(videoId, userId);
        Map<String,Object> map = new HashMap<>();
        map.put("labels",videoLabels);
        map.put("video",video);
        map.put("datas",videoDataList);
        return ResultGenerator.genSuccessResult(map);
    }

    @PostMapping("setVideoFinalVersion")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result setVideoFinalVersion(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long videoId = Long.parseLong(String.valueOf(param.get("videoId")));
        Long userConfirmId = Long.parseLong(String.valueOf(param.get("userConfirmId")));
        if (adminVideoService.setVideoFinalVersion(videoId,userConfirmId) <= 0) {
            return ResultGenerator.genFailResult("视频参数错误，请重新保存");
        }
        log.info("set finalVersion {}", userConfirmId);
        return ResultGenerator.genSuccessResult();
    }
}
