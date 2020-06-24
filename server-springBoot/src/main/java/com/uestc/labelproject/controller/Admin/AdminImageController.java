package com.uestc.labelproject.controller.Admin;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.uestc.labelproject.entity.*;
import com.uestc.labelproject.service.AdminImageService;
import com.uestc.labelproject.service.AdminImageLabelService;
import com.uestc.labelproject.service.AdminTaskService;
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
import javax.servlet.http.HttpServletRequest;
import java.io.File;
import java.io.IOException;
import java.util.*;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:49
 */
@RestController
@Slf4j
@RequestMapping("/admin")
public class AdminImageController {

    @Autowired
    AdminImageService adminImageService;
    @Autowired
    AdminTaskService adminTaskService;
    @Autowired
    AdminImageLabelService adminLabelService;

    /**
     * 获取图片列表，分页，并在此处生成缩略图
     * @param param
     * @return
     * @throws IOException
     */
    @PostMapping("getImgList")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getImgList(@RequestBody Map<String, Object> param) throws IOException {
        log.info("收到的参数{}",param);
        int page = (int)param.get("page");
        int limit = (int)param.get("limit");
        Long taskId = Long.parseLong(String.valueOf(param.get("taskId"))); //得到用户ID
        String taskName = adminTaskService.getTaskNameById(taskId);  //通过任务ID　得到任务的名字
        if (taskName == null) {
            return ResultGenerator.genFailResult("任务不存在");
        }
        log.info("page:{}, limit:{}, taskName:{}",page, limit, taskName);
        PageHelper.startPage(page,limit);
        List<Image> list = adminImageService.getImageList(taskId); //通过任务id来得到任务id
        PageInfo pageInfo = new PageInfo(list);
        if (list.size() > 0) { //有图片
            for(Image image: list){ //通过固定路径读取文件
                String src = FileUtil.IMAGE_DIC+taskName+"/"+image.getImageName();
                String dest = FileUtil.IMAGE_S_DIC+taskName;
                File file = new File(src);
                if(image.getImageThumb() == null){ //如果数据库中不存在缩略图
                    String thumb = FileUtil.thumb(file, dest); //通过工具类生成缩略图
                    image.setImageThumb(thumb); //通过上一个工具类返回的文件名更新数据库的信息
                }
            }
            adminImageService.updateImages(list); //更新数据库信息
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

    /**
     * 保存标注
     * @param jsonStr
     * @return
     */
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
            if(adminImageService.saveLabel(dataList, userId, imageId,dataIds))
                return ResultGenerator.genSuccessResult();
            else return ResultGenerator.genFailResult("保存失败");
        } catch (Exception e) {
            //TransactionAspectSupport.currentTransactionStatus().setRollbackOnly();
            return ResultGenerator.genFailResult(e.getLocalizedMessage());
        }

    }

    /**
     * 获取图片，并在此处压缩过大图片
     * @param param
     * @return
     * @throws IOException
     */
    @PostMapping("/getImg")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result getImg(@RequestBody Map<String,Object> param) throws IOException {
        log.info("收到的参数{}",param);
        Long imageId = Long.parseLong(String.valueOf(param.get("imageId")));
        Long userId = Long.parseLong(String.valueOf(param.get("userId")));
        Image image = adminImageService.getImage(imageId);
        if(image == null) return ResultGenerator.genFailResult("图片不存在！");
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
                /**
                 * 以下代码存在问题，如果图片稍微大一点，那么读取就会出错。
                 */
                try {

                    /**通过后缀来生成读取器*/
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

    @PostMapping("/deleteImageById")
    public Result deleteUser(@RequestBody Image image){
        log.info("收到的参数{}",image);
        if(image == null) return ResultGenerator.genFailResult("参数接受错误，删除失败");
        if(adminImageService.deleteFromImageByImageId(image.getImageId()) > 0
                && adminImageService.deleteFromImageDataByImageId(image.getImageId())>0
                && adminImageService.deleteFromImageDataPointsByImageId(image.getImageId())>0
                                                                                                ){
            //log.info("管理员{} 删除 {}成功", LogUtil.getUsername(request), image);
            return ResultGenerator.genSuccessResult("删除成功");
        }
        return ResultGenerator.genFailResult("删除失败");
    }


    /**
     * 确认最终版本，并在此生成RLE数据，避免在下载时花费大量时间生成，将生成时间分散
     * @param param
     * @return
     */
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

        // 这一块逻辑并没有用
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
}
