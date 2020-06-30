package com.uestc.labelproject.controller.Admin;

import com.uestc.labelproject.entity.Label;
import com.uestc.labelproject.service.AdminImageLabelService;
import com.uestc.labelproject.service.AdminTaskService;
import com.uestc.labelproject.utils.LogUtil;
import com.uestc.labelproject.utils.Result;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.transaction.interceptor.TransactionAspectSupport;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午9:28
 */
@RestController
@RequestMapping("/admin")
@Slf4j
public class AdminImageLabelController {

    @Autowired
    AdminImageLabelService adminLabelService;
    @Autowired
    AdminTaskService adminTaskService;

    /**
     * 获取标签列表
     * @return
     */
    @GetMapping("/getLabelList")
    public Result getLabelList() {
        Map<String, Object> map = new HashMap<>();
        List<Label> labels = adminLabelService.getLabelList();
        map.put("labelList", labels);
        return ResultGenerator.genSuccessResult(map);
    }

    /**
     * 编辑标签
     * @param label
     * @param request
     * @return
     */
    @PostMapping("/editLabel")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result editLabel(@RequestBody Label label, HttpServletRequest request){
        log.info("收到的参数{}",label);
        if(label == null)return ResultGenerator.genFailResult("参数错误，修改失败");
        int resultCode = 0;
        try {
            resultCode = adminLabelService.editLabel(label);
            if (resultCode == 1) {
                log.info("管理员{} 修改 {}成功", LogUtil.getUsername(request), label);
                return ResultGenerator.genSuccessResult();
            }
        } catch (Exception e) {
            TransactionAspectSupport.currentTransactionStatus().setRollbackOnly();
            return ResultGenerator.genFailResult(e.getLocalizedMessage());
        }
        return ResultGenerator.genFailResult("修改失败");
    }

    /**
     * 添加标签
     * @param label
     * @param request
     * @return
     */
    @PostMapping("/addLabel")
    @Transactional(isolation = Isolation.REPEATABLE_READ)
    public Result addLabel(@RequestBody Label label, HttpServletRequest request){
        log.info("收到的参数{}",label);
        if(label == null) return ResultGenerator.genFailResult("参数错误，添加失败");
        int resultCode = 0;
        try {
            resultCode = adminLabelService.addLabel(label);
            if (resultCode == 1) {
                log.info("管理员{} 添加 {}成功", LogUtil.getUsername(request), label);
                return ResultGenerator.genSuccessResult();
            }
        } catch (Exception e) {
            TransactionAspectSupport.currentTransactionStatus().setRollbackOnly();
            return ResultGenerator.genFailResult(e.getLocalizedMessage());
        }
        return ResultGenerator.genFailResult("添加失败");
    }

    /**
     * 生成标签
     * @param label
     * @param request
     * @return
     */
    @PostMapping("/deleteLabel")
    @Transactional
    public Result deleteLabel(@RequestBody Label label, HttpServletRequest request){
        log.info("收到的参数{}",label);
        if(label == null) return ResultGenerator.genFailResult("参数错误，删除失败");
        List<Long> taskIds = adminTaskService.getTaskIdsByLabelId(label.getLabelId(), 1);
        if(taskIds.size() > 0) return ResultGenerator.genFailResult("该标签已被使用，删除失败");
        if(adminLabelService.deleteLabel(label.getLabelId()) > 0){
            log.info("管理员{} 删除 {}成功", LogUtil.getUsername(request), label);
            return ResultGenerator.genSuccessResult();
        }
        return ResultGenerator.genFailResult("删除失败");
    }

    @PostMapping("/searchLabel")
    public Result searchLabel(@RequestBody Map<String, Object> param){
        log.info("收到的参数{}",param);
        String keyword = (String)param.get("keyword");
        return ResultGenerator.genSuccessResult(adminLabelService.searchLabel(keyword));
    }
}