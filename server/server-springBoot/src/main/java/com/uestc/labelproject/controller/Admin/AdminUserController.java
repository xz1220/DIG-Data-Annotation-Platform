package com.uestc.labelproject.controller.Admin;

import com.uestc.labelproject.entity.User;
import com.uestc.labelproject.service.AdminUserService;
import com.uestc.labelproject.utils.LogUtil;
import com.uestc.labelproject.utils.Result;
import com.uestc.labelproject.utils.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.apache.coyote.OutputBuffer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-24 下午6:42
 */
@RestController
@Slf4j
@RequestMapping("/admin")
public class AdminUserController {

    @Autowired
    AdminUserService adminUserService;

    /**
     * 获取用户列表
     * @return
     */
    @GetMapping("/getUserList")
    public Result getUserList(){
        Map<String, Object> map = new HashMap<>();
        List<User> users = adminUserService.getUserList();
        map.put("userList", users);
        return ResultGenerator.genSuccessResult(map);
    }

    /**
     * 编辑用户
     * @param user
     * @param request
     * @return
     */
    @PostMapping("/editUser")
    public Result editUser(@RequestBody User user, HttpServletRequest request){
        log.info("收到的参数{}",user);
        if(user == null)return ResultGenerator.genFailResult("参数接受错误，修改失败");
        int resultCode = adminUserService.editUser(user);
        switch (resultCode){
            case 1:
                log.info("管理员{} 修改 {}成功", LogUtil.getUsername(request), user);
                return ResultGenerator.genSuccessResult();
            case 0:
                return ResultGenerator.genFailResult("用户名已存在，修改失败");
        }
        return ResultGenerator.genFailResult("修改失败");
    }

    /**
     * 添加用户
     * @param user
     * @param request
     * @return
     */
    @PostMapping("/addUser")
    public Result addUser(@RequestBody User user, HttpServletRequest request){
        log.info("收到的参数{}",user);
        if(user == null)return ResultGenerator.genFailResult("参数接受错误，添加失败");
        int resultCode = adminUserService.addUser(user);
        switch (resultCode){
            case 1:
                log.info("管理员{} 添加 {}成功", LogUtil.getUsername(request), user);
                return ResultGenerator.genSuccessResult();
            case 0:
                return ResultGenerator.genFailResult("用户名已存在，添加失败");
        }
        return ResultGenerator.genFailResult("添加失败");
    }

    /**
     * 删除用户
     * @param user
     * @param request
     * @return
     */
    @PostMapping("/deleteUser")
    public Result deleteUser(@RequestBody User user, HttpServletRequest request){
        log.info("收到的参数{}",user);
        if(user == null) return ResultGenerator.genFailResult("参数接受错误，删除失败");
        if(adminUserService.deleteUser(user.getUserId()) > 0){
            log.info("管理员{} 删除 {}成功", LogUtil.getUsername(request), user);
            return ResultGenerator.genSuccessResult();
        }
        return ResultGenerator.genFailResult("删除失败");
    }

    /**
     * 获取标记该图片的用户
     * @param param
     * @return
     */
    @PostMapping("/getPendingUserList")
    public Result getPendingUserList(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long imageId = Long.parseLong(String.valueOf(param.get("imageId")));
        List<User> pendingUser = adminUserService.getPendingUserList(imageId);
        return ResultGenerator.genSuccessResult(pendingUser);
    }

    @PostMapping("/getVideoPendingUserList")
    public Result getVideoPendingUserList(@RequestBody Map<String, Object>param){
        log.info("收到的参数{}",param);
        Long videoId = Long.parseLong(String.valueOf(param.get("videoId")));
        List<User> pendingUser = adminUserService.getVideoPendingUserList(videoId);
        return ResultGenerator.genSuccessResult(pendingUser);
    }

    /**
     * 获取标记用户列表
     * @param param
     * @return
     */
    @PostMapping("/getListUser")
    public Result getListUser(@RequestBody Map<String,Object> param){
        log.info("收到的参数{}",param);
        List<User> user = adminUserService.getListUser();
        return ResultGenerator.genSuccessResult(user);
    }

    /**
     * 获取审核用户列表
     * @param param
     * @return
     */
    @PostMapping("/getListReviewer")
    public Result getListReviewer(@RequestBody Map<String,Object> param){
        log.info("收到的参数{}",param);
        List<User> user = adminUserService.getListReviewer();
        return ResultGenerator.genSuccessResult(user);
    }

    /**
     * 获取数量
     * @return
     */
    @GetMapping("/getCount")
    public Result getCount(){
        int taskCount = adminUserService.getTaskCount();
        int reviewerCount = adminUserService.getReviewerCount();
        int userCount = adminUserService.getUserCount();

        Map<String,Object> map = new HashMap<>();
        map.put("taskCount", taskCount);
        map.put("reviewerCount", reviewerCount);
        map.put("userCount", userCount);

        return ResultGenerator.genSuccessResult(map);
    }
}
