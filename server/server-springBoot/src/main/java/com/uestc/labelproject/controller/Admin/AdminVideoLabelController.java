package com.uestc.labelproject.controller.Admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.uestc.labelproject.entity.*;
import com.uestc.labelproject.service.AdminTaskService;
import com.uestc.labelproject.service.AdminVideoLabelService;
import com.uestc.labelproject.utils.Result;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-6 下午5:48
 */
@RestController
@Slf4j
@RequestMapping("/admin")
public class AdminVideoLabelController {

    @Autowired
    AdminVideoLabelService adminVideoLabelService;
    @Autowired
    AdminTaskService adminTaskService;

    @GetMapping("/getVideoLabelList")
    public Result getVideoLabel(){
        List<TempVideoLabel> tempVideoLabels = adminVideoLabelService.getVideoLabelList();
        if (tempVideoLabels == null) {
            return ResultGenerator.genFailResult("获取失败，请重试！");
        }
        List<VideoLabel> videoLabels = new ArrayList<>();
        for(TempVideoLabel tempVideoLabel : tempVideoLabels){
            VideoLabel videoLabel = new VideoLabel();
            videoLabel.setLabelId(tempVideoLabel.getLabelId());
            videoLabel.setQuestion(tempVideoLabel.getQuestion());
            videoLabel.setType(tempVideoLabel.getType());
            videoLabel.setSelector(tempVideoLabel.getSelector() == null ? null : JSONObject.parseObject(tempVideoLabel.getSelector(), List.class));
            videoLabels.add(videoLabel);
        }
        log.info("查询VideoLabel成功");
        return ResultGenerator.genSuccessResult(videoLabels);
    }

    @PostMapping("/addVideoLabel")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result addVideoLabel(@RequestBody VideoLabel videoLabel){
        log.info("收到的参数{}",videoLabel);
        if(videoLabel == null)return ResultGenerator.genFailResult("参数错误，添加失败");
        TempVideoLabel tempVideoLabel = new TempVideoLabel();
        tempVideoLabel.setQuestion(videoLabel.getQuestion());
        tempVideoLabel.setType(videoLabel.getType());
        tempVideoLabel.setSelector(videoLabel.getSelector() == null ? null : JSON.toJSONString(videoLabel.getSelector()));
        if(adminVideoLabelService.addVideoLabel(tempVideoLabel) > 0){
            log.info("添加模板成功");
            return ResultGenerator.genSuccessResult();
        }
        else return ResultGenerator.genFailResult("添加失败，请重试！");
    }

    @PostMapping("/editVideoLabel")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result editVideoLabel(@RequestBody VideoLabel videoLabel){
        log.info("收到的参数{}",videoLabel);
        if(videoLabel == null)return ResultGenerator.genFailResult("参数错误，修改失败");
        TempVideoLabel tempVideoLabel = new TempVideoLabel();
        tempVideoLabel.setLabelId(videoLabel.getLabelId());
        tempVideoLabel.setQuestion(videoLabel.getQuestion());
        tempVideoLabel.setType(videoLabel.getType());
        tempVideoLabel.setSelector(videoLabel.getSelector() == null ? null : JSON.toJSONString(videoLabel.getSelector()));
        if (adminVideoLabelService.editVideoLabel(tempVideoLabel) > 0) {
            log.info("修改模板成功");
            return ResultGenerator.genSuccessResult();
        }
        else return ResultGenerator.genFailResult("修改失败，请重试！");
    }

    @PostMapping("deleteVideoLabel")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result deleteVideoLabel(@RequestBody VideoLabel videoLabel){
        log.info("收到的参数{}",videoLabel);
        if(videoLabel == null)return ResultGenerator.genFailResult("参数错误，修改失败");
        List<Long> taskIds = adminTaskService.getTaskIdsByLabelId(videoLabel.getLabelId(), 2);
        if(taskIds.size() > 0) return ResultGenerator.genFailResult("该标签已被使用，删除失败");
        if(adminVideoLabelService.deleteVideoLabel(videoLabel) > 0){
            log.info("删除模板成功");
            return ResultGenerator.genSuccessResult();
        }
        else return ResultGenerator.genFailResult("删除失败，请重试！");
    }
}
