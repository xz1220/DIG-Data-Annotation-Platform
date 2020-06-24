package com.uestc.labelproject.controller;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.uestc.labelproject.entity.*;
import com.uestc.labelproject.service.*;
import com.uestc.labelproject.utils.*;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.transaction.interceptor.TransactionAspectSupport;
import org.springframework.web.bind.annotation.*;

import javax.imageio.ImageIO;
import javax.imageio.ImageReader;
import javax.imageio.stream.ImageInputStream;
import java.io.File;
import java.io.IOException;
import java.util.*;

/**
 * @Auther: kiritoghy
 * @Desc:接口实现基本与admin相同，只显示与其对应权限的数据
 * @Date: 19-7-30 下午5:26
 */
@RestController
@Slf4j
@RequestMapping("/reviewer")
public class ReviewerController {
    @Autowired
    AdminTaskService adminTaskService;
    @Autowired
    AdminImageService adminImageService;
    @Autowired
    AdminImageLabelService adminLabelService;
    @Autowired
    AdminUserService adminUserService;
    @Autowired
    AdminVideoService adminVideoService;
    @Autowired
    AdminVideoLabelService adminVideoLabelService;

    @PostMapping("/getTaskList")
    public Result getTaskList(@RequestBody Map<String,Object> param){
        log.info("收到的参数{}",param);
        int page = (int)param.get("page");
        int limit = (int)param.get("limit");
        Long reviewerId = Long.parseLong(String.valueOf(param.get("reviewerId")));
        List<Long> taskIds = adminTaskService.getTaskIdsByReviewerId(reviewerId);
        Map<String, Object> map = new HashMap<>();
        PageHelper.startPage(page,limit);
        List<Task> tasks = adminTaskService.getTaskListById(taskIds);
        PageInfo pageInfo = new PageInfo(tasks);
        map.put("page", page);
        map.put("limit", limit);
        map.put("totalpages",pageInfo.getPages());
        map.put("taskList", tasks);
        return ResultGenerator.genSuccessResult(map);
    }

    @PostMapping("/getImgList")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getImgList(@RequestBody Map<String, Object> param) throws IOException {
        log.info("收到的参数{}",param);
        int page = (int)param.get("page");
        int limit = (int)param.get("limit");
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        String taskName = adminTaskService.getTaskNameById(taskId);
        PageHelper.startPage(page,limit);
        List<Image> list = adminImageService.getImageList(taskId);
        PageInfo pageInfo = new PageInfo(list);
        if (list.size() > 0) {
            for(Image image: list){
                if(image.getImageThumb() == null){
                    String src = FileUtil.IMAGE_DIC+taskName+"/"+image.getImageName();
                    String dest = FileUtil.IMAGE_S_DIC+taskName;
                    File file = new File(src);
                    String thumb = FileUtil.thumb(file, dest);
                    image.setImageThumb(thumb);
                }
            }
            adminImageService.updateImages(list);
        }
        List<Long> labelImageIds = adminImageService.getLabeledImageIds(taskId,null);
        int totalPages = pageInfo.getPages();
        Map<String,Object> data = new HashMap<>();
        data.put("page", page);
        data.put("limit", limit);
        data.put("totalpages",totalPages);
        data.put("images", list);
        data.put("labelImageIds", labelImageIds);
        return ResultGenerator.genSuccessResult(data);
    }

    @PostMapping("/getImg")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getImg(@RequestBody Map<String,Object> param) throws IOException {
        log.info("收到的参数{}",param);
        Long imageId = Long.parseLong(String.valueOf(param.get("imageId")));
        Long userId = Long.parseLong(String.valueOf(param.get("userId")));
        Image image = adminImageService.getImage(imageId);
        String taskName = adminTaskService.getTaskNameByImageId(imageId);
        String src = FileUtil.IMAGE_DIC + taskName;
        File imageFile = new File(src + "/" + image.getImageName());
        if(imageFile.exists()){
            if (imageFile.length() > FileUtil.LIMITED_LENGTH) {
                String destdic = FileUtil.IMAGE_L_DIC + taskName;
                FileUtil.moveFile(imageFile,destdic);
                if(FileUtil.resizeImage(src, imageFile)){
                    String name = image.getImageName().substring(0,image.getImageName().lastIndexOf("."));
                    adminImageService.editImageByImageId(imageId,name+".jpg");
                    image = adminImageService.getImage(imageId);
                    imageFile = new File(src + "/" + image.getImageName());
                }
            }
            if(image.getHeight() == null || image.getWidth() == null){
                try {
                    Iterator<ImageReader> readers = ImageIO.getImageReadersByFormatName(image.getImageName().substring(image.getImageName().lastIndexOf(".") + 1));
                    ImageReader reader = (ImageReader) readers.next();
                    ImageInputStream iis = ImageIO.createImageInputStream(imageFile);
                    reader.setInput(iis, true);
                    image.setWidth(reader.getWidth(0));
                    image.setHeight(reader.getHeight(0));
                    adminImageService.updateImageWH(image);
                    image = adminImageService.getImage(imageId);
                } catch (IOException e) {
                    TransactionAspectSupport.currentTransactionStatus().setRollbackOnly();
                    return ResultGenerator.genFailResult("文件打开有误，请重试");
                }
            }
        }

        List<Label> labels = adminLabelService.getLabelListByImageId(imageId);
        List<Data> dataList = adminImageService.getDatas(imageId, userId);
        Map<String,Object> map = new HashMap<>();
        map.put("labels",labels);
        map.put("image",image);
        map.put("datas",dataList);
        return ResultGenerator.genSuccessResult(map);
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

    @PostMapping("/saveLabel")
    //@Transactional
    public Result saveLabel(@RequestBody String jsonStr){
        log.info("收到的参数{}",jsonStr);
        JSONObject jsonObject = JSONObject.parseObject(jsonStr);
        String datas = jsonObject.getString("data");
        Long userId = JSONObject.parseObject(jsonObject.getString("userId"), Long.class);
        Long imageId = JSONObject.parseObject(jsonObject.getString("imageId"), Long.class);
        List<Data> dataList = JSONArray.parseArray(datas, Data.class);
        List<Long> dataIds = adminImageService.getDataIds(userId,imageId);
        try {
            if(adminImageService.saveLabel(dataList, userId, imageId,dataIds)){
                return ResultGenerator.genSuccessResult();
            }

            else return ResultGenerator.genFailResult("保存失败");
        } catch (Exception e) {
            return ResultGenerator.genFailResult("数据出错，请重新保存");
        }
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

    @PostMapping("setFinalVersion")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result setFinalVersion(@RequestBody Map<String,Object> param){
        log.info("收到的参数{}",param);
        Long imageId = Long.parseLong(String.valueOf(param.get("imageId")));
        Long userConfirmId = Long.parseLong(String.valueOf(param.get("userConfirmId")));
        Image image = adminImageService.getImage(imageId);
        if(image.getUserConfirmId() != null){
            adminImageService.deleteRle(image.getUserConfirmId(), imageId);
        }
        List<Data> datas = adminImageService.getDatas(imageId,userConfirmId);
        List<TempRleData> rleDataString = new ArrayList<>();
        if(datas != null && datas.size() > 0){
            for(Data data: datas){
                if (data.getIscrowd() == 1) {
                    Integer width = image.getWidth();
                    Integer height = image.getHeight();
                    RleData rleData = DataGeneratorUtil.genRleData(width, height, data);
                    rleDataString.add(DataGeneratorUtil.rleDataToString(rleData,data.getDataId()));
                }
            }
            if (rleDataString.size() > 0) {
                adminImageService.addRle(rleDataString,userConfirmId,imageId);
            }
        }
        adminImageService.setFinalVersion(imageId, userConfirmId);
        log.info("set finalVersion {}", userConfirmId);
        return ResultGenerator.genSuccessResult();
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

    @PostMapping("/getPendingUserList")
    public Result getPendingUserList(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long imageId = Long.parseLong(String.valueOf(param.get("imageId")));
        List<User> pendingUser = adminUserService.getPendingUserList(imageId);
        for (User user : pendingUser){
            user.setPassword(null);
        }
        return ResultGenerator.genSuccessResult(pendingUser);
    }

    @PostMapping("/getVideoPendingUserList")
    public Result getVideoPendingUserList(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long videoId = Long.parseLong(String.valueOf(param.get("videoId")));
        List<User> pendingUser = adminUserService.getVideoPendingUserList(videoId);
        for (User user : pendingUser){
            user.setPassword(null);
        }
        return ResultGenerator.genSuccessResult(pendingUser);
    }

    @PostMapping("/taskList")
    public Result taskList(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long reviewerId = Long.parseLong(String.valueOf(param.get("reviewerId")));
        List<Long> taskIds = adminTaskService.getTaskIdsByReviewerId(reviewerId);
        List<Task> tasks = adminTaskService.taskListById(taskIds);
        if(tasks == null) return ResultGenerator.genFailResult();
        for(Task task: tasks){
            task.setReviewers(null);
        }
        return ResultGenerator.genSuccessResult(tasks);
    }


    @PostMapping("getVideoList")
    public Result getVideoList(@RequestBody Map<String, Object> param) throws IOException {
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

    @PostMapping("/searchTask")
    public Result searchTask(@RequestBody Map<String, Object> param){
        String keyword = (String)param.get("keyword");
        List<Task> tasks = adminTaskService.searchTask(keyword);
        return ResultGenerator.genSuccessResult(tasks);
    }


}
