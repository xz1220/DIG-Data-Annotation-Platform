package com.uestc.labelproject.controller.Admin;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.serializer.AfterFilter;
import com.alibaba.fastjson.serializer.NameFilter;
import com.alibaba.fastjson.serializer.PropertyFilter;
import com.alibaba.fastjson.serializer.SerializeFilter;
import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.uestc.labelproject.entity.*;
import com.uestc.labelproject.service.*;
import com.uestc.labelproject.utils.*;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.io.FileUtils;
import org.bytedeco.javacv.FrameGrabber;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.FileSystemResource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.*;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午8:38
 */
@RestController
@Slf4j
@RequestMapping("/admin")
public class AdminTaskController {

    @Autowired
    AdminTaskService adminTaskService;
    @Autowired
    AdminImageService adminImageService;
    @Autowired
    AdminImageLabelService adminLabelService;
    @Autowired
    AdminVideoService adminVideoService;
    @Autowired
    AdminVideoLabelService adminVideoLabelService;

    /**
     * 获取任务列表，并刷新新任务
     * @return
     */
    @PostMapping("/getTaskList")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getTaskList(@RequestBody Map<String, Object> param) {
        log.info("收到的参数{}",param);
        int page = (int)param.get("page");
        int limit = (int)param.get("limit");
        Map<String, Object> map = new HashMap<>();
        PageHelper.startPage(page,limit);
        List<Task> tasks = adminTaskService.getTaskList();
        PageInfo<Task> pageInfo = new PageInfo<>(tasks);
        map.put("page", page);
        map.put("limit", limit);
        map.put("totalpages",pageInfo.getPages());
        map.put("taskList", tasks);
        return ResultGenerator.genSuccessResult(map);
    }

    /**
     * 更新任务
     * @param jsonStr
     * @return
     */
    @PostMapping("/updateTask")
    @Transactional(isolation = Isolation.SERIALIZABLE)
    public Result updateTask(@RequestBody String jsonStr){
        log.info("收到的参数{}",jsonStr);
        JSONObject jsonObject = JSONObject.parseObject(jsonStr);
        List<Long> userIds = JSONArray.parseArray(jsonObject.getString("userIds"), Long.class);
        List<Long> labelIds = JSONArray.parseArray(jsonObject.getString("labelIds"), Long.class);
        List<Long> reviewerIds = JSONArray.parseArray(jsonObject.getString("reviewerIds"), Long.class);
        if(userIds == null || labelIds == null || reviewerIds == null) return ResultGenerator.genFailResult("获取参数失败");
        String desc = JSONObject.parseObject(jsonObject.getString("taskDesc"),String.class);
        String taskName = jsonObject.getString("taskName");
        Long taskId = JSONObject.parseObject(jsonObject.getString("taskId"),Long.class);
        Task task = adminTaskService.getTaskById(taskId);
        if(task == null)return ResultGenerator.genFailResult("未选中任务");
        List<Long> originUserIds = task.getUserIds();
        List<Long> originLabelIds = task.getLabelIds();
        List<Long> originReviewerIds = task.getReviewerIds();

        if(adminTaskService.hasData(taskId) > 0){
            if(labelIds.size() != originLabelIds.size())return ResultGenerator.genFailResult("该任务已有标注，请确认删除后再修改");
            for(int i = 0; i < labelIds.size(); ++i){
                if(!labelIds.get(i).equals(originLabelIds.get(i)))
                    return ResultGenerator.genFailResult("该任务已有标注，请确认删除后再修改");
            }
        }
        if(originUserIds.size() > 0)
            if(adminTaskService.deleteTaskUserIds(taskId) <= 0)return ResultGenerator.genFailResult("修改失败");
        if(originLabelIds.size() > 0)
            if(adminTaskService.deleteTaskLabelIds(taskId) <= 0) return ResultGenerator.genFailResult("修改失败");
        if(originReviewerIds.size() > 0)
            if(adminTaskService.deleteTaskReviewerIds(taskId) <= 0)return ResultGenerator.genFailResult("修改失败");
        if(userIds.size() > 0)
            if(adminTaskService.addTaskUserIds(userIds, taskId) <= 0)return ResultGenerator.genFailResult("修改失败");
        if(labelIds.size() > 0)
            if(adminTaskService.addTaskLabelIds(labelIds,taskId) <= 0)return ResultGenerator.genFailResult("修改失败");
        if(reviewerIds.size() > 0)
            if(adminTaskService.addTaskReviewerIds(reviewerIds,taskId) <= 0)return ResultGenerator.genFailResult("修改失败");
        if(!taskName.equals(task.getTaskName())){
            if(FileUtil.rename(FileUtil.IMAGE_DIC + task.getTaskName(), FileUtil.IMAGE_DIC+taskName))
                task.setTaskName(taskName);
            else return ResultGenerator.genFailResult("修改任务名失败");
        }
        task.setTaskDesc(desc);
        adminTaskService.updateTask(task);
        return ResultGenerator.genSuccessResult();
    }

    /**
     * 删除任务
     * @param param
     * @return
     * @throws IOException
     */
    @PostMapping("/deleteTask")
    @Transactional(isolation = Isolation.SERIALIZABLE)
    public Result deleteTask(@RequestBody Map<String,Object>param) throws IOException {
        log.info("收到的参数{}",param);
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        if (taskId == null) {
            return ResultGenerator.genFailResult("未选中任务");
        }
        Task task = adminTaskService.getTaskById(taskId);
        if(task == null) return ResultGenerator.genFailResult("未选中任务");
        switch (task.getTaskType()){
            case 1:
            case 2:
            case 3:
            case 4:{
                List<Long> imageIds = adminImageService.getImageIds(taskId);
                List<Long> dataIds = adminImageService.getDataIds(imageIds);
                if(imageIds.size() > 0)
                    if(adminImageService.deleteImagesByTaskId(taskId) <= 0)return ResultGenerator.genFailResult("删除image失败");
                if(imageIds.size() > 0 && dataIds.size() > 0)
                    if(adminImageService.deleteDatasByImageId(imageIds) <= 0)return ResultGenerator.genFailResult("删除data失败");
                if(dataIds.size() > 0)
                    if(adminImageService.deletePoints(dataIds) < 0)return ResultGenerator.genFailResult("删除point失败");
                adminImageService.deleteFinish(taskId);
                if(task.getUserIds().size() > 0)
                    if(adminTaskService.deleteTaskUserIds(taskId) <= 0)return ResultGenerator.genFailResult("删除user失败");
                if(task.getLabelIds().size() > 0)
                    if(adminTaskService.deleteTaskLabelIds(taskId) <= 0) return ResultGenerator.genFailResult("删label除失败");
                if(task.getReviewerIds().size() > 0)
                    if(adminTaskService.deleteTaskReviewerIds(taskId) <= 0)return ResultGenerator.genFailResult("删除reviewer失败");
                if(adminTaskService.deleteTask(taskId) <= 0)return ResultGenerator.genFailResult("删除失败");

                File src = new File(FileUtil.IMAGE_DIC + task.getTaskName());
                File thumb = new File(FileUtil.IMAGE_S_DIC + task.getTaskName());
                File dest = new File(FileUtil.IMAGE_DELETE_DIC);
                FileUtils.moveDirectoryToDirectory(src,dest,true);
                FileUtils.deleteDirectory(thumb);
                break;
            }
            case 5:{
                List<Long> videoIds = adminVideoService.getVideoIds(taskId);
                List<Long> dataIds = adminImageService.getDataIds(videoIds);
                if (videoIds.size() > 0) {
                    if (adminVideoService.deleteVideosByTaskId(taskId) <= 0) {
                        return ResultGenerator.genFailResult("删除video失败");
                    }
                }
                if (videoIds.size() > 0 && dataIds.size() > 0){
                    if (adminVideoService.deleteDatasByVideoId(videoIds) <= 0) {
                        return ResultGenerator.genFailResult("删除data失败");
                    }
                }
                adminImageService.deleteFinish(taskId);
                if(task.getUserIds().size() > 0)
                    if(adminTaskService.deleteTaskUserIds(taskId) <= 0)return ResultGenerator.genFailResult("删除user失败");
                if(task.getLabelIds().size() > 0)
                    if(adminTaskService.deleteTaskLabelIds(taskId) <= 0) return ResultGenerator.genFailResult("删label除失败");
                if(task.getReviewerIds().size() > 0)
                    if(adminTaskService.deleteTaskReviewerIds(taskId) <= 0)return ResultGenerator.genFailResult("删除reviewer失败");
                if(adminTaskService.deleteTask(taskId) <= 0)return ResultGenerator.genFailResult("删除任务失败");
                File src = new File(FileUtil.VIDEO_DIC + task.getTaskName());
                File dest = new File(FileUtil.VIDEO_D_DIC);
                File thumb = new File(FileUtil.VIDEO_S_DIC + task.getTaskName());
                FileUtils.moveDirectoryToDirectory(src,dest,true);
                FileUtils.deleteDirectory(thumb);
                break;
            }
        }
        return ResultGenerator.genSuccessResult();
    }

    /**
     * 获取任务列表，不刷新
     * @return
     */
    @GetMapping("/taskList")
    public Result taskList(){
        List<Task> tasks = adminTaskService.taskList();
        if(tasks != null) return ResultGenerator.genSuccessResult(tasks);
        else return ResultGenerator.genFailResult();
    }

    /**
     * 分割任务
     * @param param
     * @return
     * @throws IOException
     */
    @PostMapping("/splitTask")
    @Transactional(isolation = Isolation.SERIALIZABLE)
    public Result splitTask(@RequestBody Map<String, Object> param) throws IOException {
        log.info("收到的参数{}",param);
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        int quantity = (int)param.get("quantity");

        Task task = adminTaskService.getTaskById(taskId);
        if(task == null) return ResultGenerator.genFailResult("任务不存在！");
        if(task.getImageNumber() < quantity) return ResultGenerator.genFailResult("拆分任务失败，任务数大于图片数！");
        int taskImageNumber = task.getImageNumber() / quantity;
        int lastTaskImageNumber = task.getImageNumber() - ((quantity - 1) * taskImageNumber);

        switch (task.getTaskType()){
            case 1:
            case 2:
            case 3:
            case 4:{
                // 获取文件路径 判断是否存在以及是否是文件夹
                File taskDic = new File(FileUtil.IMAGE_DIC + task.getTaskName());
                File thumbTaskDic = new File(FileUtil.IMAGE_S_DIC + task.getTaskName());
                if(!taskDic.exists() || !taskDic.isDirectory()) return ResultGenerator.genFailResult("拆分任务失败，不存在该任务");

                for(int i = 1; i <= quantity; ++i){
                    // 生成新任务的名称
                    String newTaskName = task.getTaskName() + "_part" + i;

                    // 不确定具体用法，大致是生成任务数的图片
                    if(i < quantity) PageHelper.startPage(1,taskImageNumber);
                    else PageHelper.startPage(1,lastTaskImageNumber);

                    //获得图片数据
                    List<Image> images = adminImageService.getImageList(taskId);

                    //便利
                    for(Image image: images){
                        File dest = new File(FileUtil.IMAGE_DIC + newTaskName);
                        File imageFile = new File(FileUtil.IMAGE_DIC + task.getTaskName() + "/" + image.getImageName());
                        FileUtils.moveFileToDirectory(imageFile, dest, true);
                        if(thumbTaskDic.exists() && thumbTaskDic.isDirectory()){
                            if(image.getImageThumb() != null){
                                File thumbDest = new File(FileUtil.IMAGE_S_DIC + newTaskName);
                                File thumbImageFile = new File(FileUtil.IMAGE_S_DIC + task.getTaskName() + "/" + image.getImageThumb());
                                FileUtils.moveFileToDirectory(thumbImageFile, thumbDest, true);
                            }
                        }

                    }

                    //创建新任务
                    Task newTask = new Task();
                    BeanUtils.copyProperties(task, newTask);   // 将老的赋值给新的
                    newTask.setTaskName(newTaskName);
                    newTask.setImageNumber(images.size());
                    if (adminTaskService.addTask(newTask) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getUserIds().size() > 0 && adminTaskService.addTaskUserIds(newTask.getUserIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getReviewerIds().size() > 0 && adminTaskService.addTaskReviewerIds(newTask.getReviewerIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getLabelIds().size() > 0 && adminTaskService.addTaskLabelIds(newTask.getLabelIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (adminImageService.updateImagesTaskId(images, newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                }

                if (adminTaskService.deleteTask(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskUserIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskReviewerIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskLabelIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }

                FileUtils.deleteDirectory(taskDic);
                FileUtils.deleteDirectory(thumbTaskDic);
                break;
            }

            case 5:{
                File taskDic = new File(FileUtil.VIDEO_DIC + task.getTaskName());
                File thumbTaskDic = new File(FileUtil.VIDEO_S_DIC + task.getTaskName());
                if(!taskDic.exists() || !taskDic.isDirectory()) return ResultGenerator.genFailResult("拆分任务失败，不存在该任务");

                for (int i = 1; i <= quantity; ++i){
                    String newTaskName = task.getTaskName() + "_part" + i;
                    if(i < quantity) PageHelper.startPage(1,taskImageNumber);
                    else PageHelper.startPage(1,lastTaskImageNumber);
                    List<Video> videos = adminVideoService.getVideoList(taskId);

                    for (Video video : videos){
                        File dest = new File(FileUtil.VIDEO_DIC + newTaskName);
                        File videoFile = new File(FileUtil.VIDEO_DIC + task.getTaskName() + "/" + video.getVideoName());
                        FileUtils.moveFileToDirectory(videoFile, dest, true);
                        if(thumbTaskDic.exists() && thumbTaskDic.isDirectory()){
                            if(video.getVideoThumb() != null){
                                File thumbDest = new File(FileUtil.VIDEO_S_DIC + newTaskName);
                                File thumbVideoFile = new File(FileUtil.VIDEO_S_DIC + task.getTaskName() + "/" + video.getVideoThumb());
                                FileUtils.moveFileToDirectory(thumbVideoFile, thumbDest, true);
                            }
                        }
                    }
                    Task newTask = new Task();
                    BeanUtils.copyProperties(task, newTask);
                    newTask.setTaskName(newTaskName);
                    newTask.setImageNumber(videos.size());
                    if (adminTaskService.addTask(newTask) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getUserIds().size() > 0 && adminTaskService.addTaskUserIds(newTask.getUserIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getReviewerIds().size() > 0 && adminTaskService.addTaskReviewerIds(newTask.getReviewerIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (newTask.getLabelIds().size() > 0 && adminTaskService.addTaskLabelIds(newTask.getLabelIds(), newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                    if (adminVideoService.updateVideoTaskId(videos, newTask.getTaskId()) < 0) {
                        return ResultGenerator.genFailResult();
                    }
                }

                if (adminTaskService.deleteTask(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskUserIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskReviewerIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }
                if (adminTaskService.deleteTaskLabelIds(taskId) < 0) {
                    return ResultGenerator.genFailResult();
                }

                FileUtils.deleteDirectory(taskDic);
                FileUtils.deleteDirectory(thumbTaskDic);
                break;
            }
        }


        return ResultGenerator.genSuccessResult("拆分成功!");
    }


    /**
     * 下载数据
     * @param param
     * @return
     * @throws IOException
     */
    @PostMapping("downloadDatas")
    @Transactional(isolation = Isolation.SERIALIZABLE)
    public ResponseEntity downloadDatas(@RequestBody Map<String, Object> param) throws IOException {
        log.info("收到的参数{}",param);
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        Task task = adminTaskService.getTaskById(taskId);
        HttpHeaders headers = new HttpHeaders();
        headers.add("Cache-Control", "no-cache, no-store, must-revalidate");
        headers.add("content-type", "application/json;charset=utf-8");
        if(task == null)return ResponseEntity.ok().headers(headers).body(new Result(500,"任务不存在，下载失败"));
        SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd_HH-mm-ss");
        String date = df.format(new Date());

        switch (task.getTaskType()){
            case 1:{
                log.info("下载图片数据");
                List<Image> images = adminImageService.getImageList(taskId);
                if(images == null || images.size() == 0) return ResponseEntity.ok().headers(headers).body(new Result(500,"图片不存在，下载失败"));
                List<Label> labels = adminLabelService.getLabelListByImageId(images.get(0).getImageId());
                if (labels == null || labels.size() == 0) {
                    return ResponseEntity.ok().headers(headers).body(new Result(500,"标签不存在，下载失败"));
                }
                CocoDataSet cocoDataSet = new CocoDataSet();
                //cocinfo
                CocoInfo cocoInfo = new CocoInfo();
                cocoInfo.setDate_created(date);
                cocoInfo.setYear(Integer.parseInt(date.substring(0,4)));
                cocoDataSet.setInfo(cocoInfo);

                List<CocoAnnotation> cocoAnnotations = new ArrayList<>();
                List<CocoCategory> cocoCategories = new ArrayList<>();
                List<CocoImage> cocoImages = new ArrayList<>();
                for(Image image : images){
                    if(image.getUserConfirmId() == null) continue;
                    List<Data> datas = adminImageService.getDatas(image.getImageId(), image.getUserConfirmId());
                    if (datas == null) continue;
                    CocoImage cocoImage = new CocoImage();
                    cocoImage.setFile_name(image.getImageName());
                    cocoImage.setHeight(image.getHeight());
                    cocoImage.setWidth(image.getWidth());
                    cocoImage.setId(image.getImageId());
                    cocoImages.add(cocoImage);
                    for(Data data: datas){
                        CocoAnnotation cocoAnnotation = new CocoAnnotation();
                        cocoAnnotation.setId(data.getDataId());
                        cocoAnnotation.setImage_id(data.getImageId());
                        cocoAnnotation.setCategory_id(data.getLabelId());
                        cocoAnnotations.add(cocoAnnotation);
                    }
                }
                cocoDataSet.setAnnotations(cocoAnnotations);
                cocoDataSet.setImages(cocoImages);
                for(Label label: labels){
                    CocoCategory cocoCategory = new CocoCategory();
                    cocoCategory.setId(label.getLabelId());
                    cocoCategory.setName(label.getLabelName());
                    cocoCategory.setSupercategory(label.getLabelName());
                    cocoCategories.add(cocoCategory);
                }
                cocoDataSet.setCategories(cocoCategories);
                String fileName = task.getTaskName() + "_"+date + ".json";
                /*String filePath = FileUtil.IMAGE_DIC + task.getTaskName() + "/" + task.getTaskName() + "_" + date + ".json";
                File file = new File(filePath);
                BufferedWriter bw = new BufferedWriter(new FileWriter(file));
                bw.write(JSONObject.toJSONString(cocoDataSet));
                bw.flush();
                bw.close();*/
                headers.add("Content-Disposition", "attachment; filename=" + fileName);
                headers.add("Pragma", "no-cache");
                headers.add("Expires", "0");
                headers.remove("content-type");
                return ResponseEntity.ok().contentType(MediaType.parseMediaType("application/octet-stream")).headers(headers).body(JSONObject.toJSONString(cocoDataSet));
            }
            case 2:
            case 3:{
                log.info("下载图片数据");
                List<Image> images = adminImageService.getImageList(taskId);
                if(images == null || images.size() == 0) return ResponseEntity.ok().headers(headers).body(new Result(500,"图片不存在，下载失败"));
                List<Label> labels = adminLabelService.getLabelListByImageId(images.get(0).getImageId());
                if (labels == null || labels.size() == 0) {
                    return ResponseEntity.ok().headers(headers).body(new Result(500,"标签不存在，下载失败"));
                }
                CocoDataSet cocoDataSet = new CocoDataSet();
                //cocinfo
                CocoInfo cocoInfo = new CocoInfo();
                cocoInfo.setDate_created(date);
                cocoInfo.setYear(Integer.parseInt(date.substring(0,4)));
                cocoDataSet.setInfo(cocoInfo);

                List<CocoAnnotation> cocoAnnotations = new ArrayList<>();
                List<CocoCategory> cocoCategories = new ArrayList<>();
                List<CocoImage> cocoImages = new ArrayList<>();
                for(Image image : images){
                    if(image.getUserConfirmId() == null) continue;  /**需要有人确认*/
                    List<Data> datas = adminImageService.getDatas(image.getImageId(), image.getUserConfirmId());
                    if (datas == null) continue; /**需要的确有数据存在*/
                    CocoImage cocoImage = new CocoImage();
                    cocoImage.setFile_name(image.getImageName());
                    cocoImage.setHeight(image.getHeight());
                    cocoImage.setWidth(image.getWidth());
                    cocoImage.setId(image.getImageId());
                    cocoImages.add(cocoImage);
                    for(Data data: datas){
                        CocoAnnotation cocoAnnotation = new CocoAnnotation();
                        
                        if(data.getIscrowd() == 0) {
                            cocoAnnotation.setSegmentation(DataGeneratorUtil.genPolygonData(data));
                        }
                        //永远不会触发：iscrowd always equals 0 because no function calls for SetIsCrowd
                        if(data.getIscrowd() == 1){
                            TempRleData tempRleData = adminImageService.getTempRleData(data.getDataId());
                            RleData rleData = DataGeneratorUtil.stringToRleData(tempRleData.getData());
                            cocoAnnotation.setSegmentation(rleData);
                        }

                        cocoAnnotation.setId(data.getDataId());
                        cocoAnnotation.setImage_id(data.getImageId());
                        cocoAnnotation.setCategory_id(data.getLabelId());
                        cocoAnnotation.setIscrowd(data.getIscrowd());
                        cocoAnnotation.setBbox(DataGeneratorUtil.getBbox(data));
                        cocoAnnotation.setArea(DataGeneratorUtil.CalculateArea(data.getPoint()));

                        cocoAnnotation.setDesc(data.getDataDesc());

                        cocoAnnotations.add(cocoAnnotation);
                    }
                }
                cocoDataSet.setAnnotations(cocoAnnotations);
                cocoDataSet.setImages(cocoImages);
                /**
                 * 对应json中的categories类
                 */
                for(Label label: labels){
                    CocoCategory cocoCategory = new CocoCategory();
                    cocoCategory.setId(label.getLabelId());
                    cocoCategory.setName(label.getLabelName());
                    cocoCategory.setSupercategory(label.getLabelName());
                    cocoCategories.add(cocoCategory);
                }
                cocoDataSet.setCategories(cocoCategories);
                String fileName = task.getTaskName() + "_"+date + ".json";
                /*String filePath = FileUtil.IMAGE_DIC + task.getTaskName() + "/" + task.getTaskName() + "_" + date + ".json";
                File file = new File(filePath);
                BufferedWriter bw = new BufferedWriter(new FileWriter(file));
                bw.write(JSONObject.toJSONString(cocoDataSet));
                bw.flush();
                bw.close();*/
                /**
                 * 加入HTTP头
                 */
                headers.add("Content-Disposition", "attachment; filename=" + fileName);
                headers.add("Pragma", "no-cache");
                headers.add("Expires", "0");
                headers.remove("content-type");
                return ResponseEntity.ok().contentType(MediaType.parseMediaType("application/octet-stream")).headers(headers).body(JSONObject.toJSONString(cocoDataSet));
            }
            case 4:{
                log.info("下载图片数据");
                List<Image> images = adminImageService.getImageList(taskId);
                if(images == null || images.size() == 0) return ResponseEntity.ok().headers(headers).body(new Result(500,"图片不存在，下载失败"));
                List<Label> labels = adminLabelService.getLabelListByImageId(images.get(0).getImageId());
                if (labels == null || labels.size() == 0) {
                    return ResponseEntity.ok().headers(headers).body(new Result(500,"标签不存在，下载失败"));
                }
                CocoDataSet cocoDataSet = new CocoDataSet();
                //cocinfo
                CocoInfo cocoInfo = new CocoInfo();
                cocoInfo.setDate_created(date);
                cocoInfo.setYear(Integer.parseInt(date.substring(0,4)));
                cocoDataSet.setInfo(cocoInfo);

                List<CocoAnnotation> cocoAnnotations = new ArrayList<>();
                List<CocoCategory> cocoCategories = new ArrayList<>();
                List<CocoImage> cocoImages = new ArrayList<>();
                List<String> keypoints = new ArrayList<>();
                for(Image image : images){
                    if(image.getUserConfirmId() == null) continue;
                    List<Data> datas = adminImageService.getDatas(image.getImageId(), image.getUserConfirmId());
                    if (datas == null) continue;
                    CocoImage cocoImage = new CocoImage();
                    cocoImage.setFile_name(image.getImageName());
                    cocoImage.setHeight(image.getHeight());
                    cocoImage.setWidth(image.getWidth());
                    cocoImage.setId(image.getImageId());
                    cocoImages.add(cocoImage);
                    for(Data data: datas){
                        CocoAnnotation cocoAnnotation = new CocoAnnotation();
                        if(data.getIscrowd() == 0){
                            cocoAnnotation.setSegmentation(DataGeneratorUtil.genPolygonData(data));
                        }
                        if(data.getIscrowd() == 1){
                            TempRleData tempRleData = adminImageService.getTempRleData(data.getDataId());
                            RleData rleData = DataGeneratorUtil.stringToRleData(tempRleData.getData());
                            cocoAnnotation.setSegmentation(rleData);
                        }
                        cocoAnnotation.setId(data.getDataId());
                        cocoAnnotation.setImage_id(data.getImageId());
                        cocoAnnotation.setCategory_id(data.getLabelId());
                        cocoAnnotation.setIscrowd(data.getIscrowd());
                        cocoAnnotation.setBbox(DataGeneratorUtil.getBbox(data));
                        cocoAnnotation.setArea(DataGeneratorUtil.CalculateArea(data.getPoint()));
                        cocoAnnotations.add(cocoAnnotation);
                    }
                }
                cocoDataSet.setAnnotations(cocoAnnotations);
                cocoDataSet.setImages(cocoImages);
                for(Label label: labels){
                    CocoCategory cocoCategory = new CocoCategory();
                    cocoCategory.setId(label.getLabelId());
                    cocoCategory.setName(label.getLabelName());
                    cocoCategory.setSupercategory(label.getLabelName());
                    cocoCategories.add(cocoCategory);
                    keypoints.add(label.getLabelName());
                }
                cocoDataSet.setCategories(cocoCategories);
                String fileName = task.getTaskName() + "_"+date + ".json";
                /*String filePath = FileUtil.IMAGE_DIC + task.getTaskName() + "/" + task.getTaskName() + "_" + date + ".json";
                File file = new File(filePath);
                BufferedWriter bw = new BufferedWriter(new FileWriter(file));*/
                SerializeFilter[] filters = new SerializeFilter[]{
                        new NameFilter() {
                            @Override
                            public String process(Object o, String s, Object o1) {
                                if(s.equals("segmentation"))return "keypoints";
                                return s;
                            }
                        },
                        new AfterFilter() {
                            @Override
                            public void writeAfter(Object o) {
                                if(o instanceof CocoDataSet){
                                    writeKeyValue("keypoints",keypoints);
                                }
                            }
                        },
                        new PropertyFilter() {
                            @Override
                            public boolean apply(Object o, String s, Object o1) {
                                if(s.equals("bbox")) return false;
                                return true;
                            }
                        }
                };
                /*bw.write(JSONObject.toJSONString(cocoDataSet,filters));
                bw.flush();
                bw.close();*/
                headers.add("Content-Disposition", "attachment; filename=" + fileName);
                headers.add("Pragma", "no-cache");
                headers.add("Expires", "0");
                headers.remove("content-type");
                return ResponseEntity.ok().contentType(MediaType.parseMediaType("application/octet-stream")).headers(headers).body(JSONObject.toJSONString(cocoDataSet,filters));
            }
            case 5:{
                log.info("下载视频数据");
                List<Video> videos = adminVideoService.getVideoList(taskId);
                if(videos == null || videos.size() == 0) return ResponseEntity.ok().headers(headers).body(new Result(500,"视频不存在，下载失败"));
                List<VideoDataSet> videoDataSets = new ArrayList<>();
                for (Video video : videos){
                    VideoDataSet videoDataSet = new VideoDataSet();
                    List<VideoData> videoDataList = adminVideoService.getVideoData(video.getVideoId(), video.getUserConfirmId());
                    if(videoDataList == null || videoDataList.size() == 0)continue;
                    videoDataSet.setVideoName(video.getVideoName());
                    videoDataSet.setDuration(video.getDuration());
                    List<String> sentences = new ArrayList<>();
                    List<List<Double>> timeStamps = new ArrayList<>();
                    for (VideoData videoData : videoDataList){
                        sentences.add(videoData.getSentence());
                        List<Double> tmp = new ArrayList<>();
                        tmp.add(videoData.getStartTime());
                        tmp.add(videoData.getEndTime());
                        timeStamps.add(tmp);
                    }
                    videoDataSet.setTimeStamps(timeStamps);
                    videoDataSet.setSentences(sentences);
                    videoDataSets.add(videoDataSet);
                }
                String fileName = task.getTaskName() + "_"+date + ".json";
                /*String filePath = FileUtil.VIDEO_DIC + task.getTaskName() + "/" + task.getTaskName() + "_" + date + ".json";
                File file = new File(filePath);
                BufferedWriter bw = new BufferedWriter(new FileWriter(file));
                bw.write(JSONObject.toJSONString(videoDataSets));
                bw.flush();
                bw.close();*/
                headers.add("Content-Disposition", "attachment; filename=" + fileName);
                headers.add("Pragma", "no-cache");
                headers.add("Expires", "0");
                headers.remove("content-type");
                return ResponseEntity.ok().contentType(MediaType.parseMediaType("application/octet-stream")).headers(headers).body(JSONObject.toJSONString(videoDataSets));
            }
            default:{
                return ResponseEntity.ok().headers(headers).body(new Result(500,"下载失败"));
            }
        }
        //return ResultGenerator.genSuccessResult();
    }

    @PostMapping("/searchTask")
    public Result searchTask(@RequestBody Map<String, Object> param){
        String keyword = (String)param.get("keyword");
        List<Task> tasks = adminTaskService.searchTask(keyword);
        return ResultGenerator.genSuccessResult(tasks);
    }

    @GetMapping("/getNewTaskList")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getNewTaskList() throws FrameGrabber.Exception {
        Set<String> imageNames = adminTaskService.getImageTaskNameList();
        Set<String> videoNames = adminTaskService.getVideoTaskNameList();
        File imageDic = new File(FileUtil.IMAGE_DIC);
        Task temp = new Task();
        if(!imageDic.exists() || imageDic.isFile()) return ResultGenerator.genFailResult("图片路径不存在，文件读取失败");
        //扫描图片
        if(imageDic.isDirectory()){
            File[] imageFiles = imageDic.listFiles();
            if(imageFiles != null && imageFiles.length > 0){
                for(File imageFile: imageFiles){//遍历任务，加入数据库

                    if(imageFile.isDirectory()){
                        File[] images = imageFile.listFiles();
                        String dicName = imageFile.getName();
                        if(!imageNames.contains(dicName)){
                            temp.setTaskName(dicName);
                            temp.setImageNumber(images == null ? 0 : images.length);
                            temp.setTaskType(1);
                            adminTaskService.addTask(temp);
                            if (images != null && images.length > 0) {
                                List<Image> list = new ArrayList<>();
                                for(File image: images){ //遍历图片，加入数据库
                                    String suffix = image.getName().substring(image.getName().lastIndexOf(".") + 1).toLowerCase();
                                    if(!suffix.equals("jpg") && !suffix.equals("jpeg") && !suffix.equals("bmp") && !suffix.equals("png"))continue;
                                    Image temimage = new Image();
                                    temimage.setImageName(image.getName());
                                    temimage.setTaskId(temp.getTaskId());
                                    list.add(temimage);
                                }
                                adminImageService.addImages(list);
                            }
                        }
                    }
                }
            }
        }
        //扫描视频
        File videoDic = new File(FileUtil.VIDEO_DIC);
        if(!videoDic.exists() || videoDic.isFile()) return ResultGenerator.genFailResult("视频路径不存在，文件读取失败");
        if(videoDic.isDirectory()){
            File[] videoFiles = videoDic.listFiles();
            if(videoFiles != null && videoFiles.length > 0){
                for(File videoFile : videoFiles){
                    if(videoFile.isDirectory()){
                        File[] videos = videoFile.listFiles();
                        String dicName = videoFile.getName();
                        if(!videoNames.contains(dicName)){
                            temp.setTaskType(5);
                            temp.setImageNumber(videos == null ? 0 : videos.length);
                            temp.setTaskName(dicName);
                            adminTaskService.addTask(temp);
                            if(videos != null && videos.length > 0){
                                List<Video> videoList = new ArrayList<>();
                                for(File video : videos){
                                    Video tmpVideo = new Video();
                                    tmpVideo.setVideoName(video.getName());
                                    tmpVideo.setTaskId(temp.getTaskId());
                                    tmpVideo.setDuration(VideoUtil.getDuration(video));
                                    videoList.add(tmpVideo);
                                }
                                adminVideoService.addVideos(videoList);
                            }
                        }
                    }
                }
            }
        }
        log.info("刷新任务成功！");
        Map<String, Object> map = new HashMap<>();
        map.put("taskList", adminTaskService.getNewTaskList());
        return ResultGenerator.genSuccessResult(map);
    }

    @PostMapping("/updateTaskType")
    public Result updateTaskType(@RequestBody Map<String, Object> param){
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId")));
        int taskType = (int)param.get("taskType");
        if(adminTaskService.hasData(taskId) > 0) return ResultGenerator.genFailResult("该任务已有标注，请确认删除后再修改");
        if(adminTaskService.updateTaskType(taskId,taskType) > 0)return ResultGenerator.genSuccessResult();
        else return ResultGenerator.genFailResult();
    }
}
