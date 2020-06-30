package com.uestc.labelproject.controller;

import com.uestc.labelproject.entity.SelfUserDetails;
import com.uestc.labelproject.service.TestService;
import com.uestc.labelproject.utils.*;
import com.uestc.labelproject.entity.Test;
import lombok.extern.slf4j.Slf4j;
import net.coobird.thumbnailator.Thumbnails;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.transaction.interceptor.TransactionAspectSupport;
import org.springframework.web.bind.annotation.*;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * @Auther:kiritoghy
 * @Date:19-7-23 下午12:26
 */
@Slf4j
@RequestMapping("/test")
@RestController
public class TestLogController {

    @Autowired
    TestService testService;

    @GetMapping("/log")
    public void log(){
        BCryptPasswordEncoder bCryptPasswordEncoder = new BCryptPasswordEncoder();
        System.out.println(bCryptPasswordEncoder.encode("admin"));
        log.info("info test");
        log.error("error test");
    }
    @GetMapping("/sql")
    public Result sql(){
        Result result = new Result();
        List<Test> t = testService.get();
        if(t != null){
            result.setCode(200);
            result.setMessage("success");
            result.setData(t);
            return result;
        }
        result.setCode(500);
        result.setMessage("fali");
        return result;
    }

    @GetMapping("/getParam")
    public void getParam(SelfUserDetails selfUserDetails/*int id, String username, String password, String authorities*/){
        System.out.println(selfUserDetails.getUsername());
//        System.out.println(id);
//        System.out.println(username);
//        System.out.println(password);
//        System.out.println(authorities);
    }

    @GetMapping("/imageTest")
    public void imageTest() throws IOException {
        File imageDic = new File(FileUtil.IMAGE_DIC);
        if(imageDic == null || imageDic.isFile()) return ;
        if(imageDic.isDirectory()){
            File[] imagefiles = imageDic.listFiles();
            for(File imagefile: imagefiles){
                if(imagefile.isDirectory()){
                    System.out.println(imagefile.getName());
                    File[] images = imagefile.listFiles();
                    for(File image: images){
                        long time1=System.currentTimeMillis();
                        System.out.println(image.length() / 1024f);
                        Thumbnails.of(image).scale(0.25f).toFile(FileUtil.IMAGE_S_DIC+image.getName());
                        long time2=System.currentTimeMillis();
                        System.out.println("当前程序耗时："+(time2-time1)+"ms");
                        //System.out.println(image.getName());
                        /*System.out.println(image.getAbsolutePath());
                        System.out.println(image.length());*/
                    }
                }
            }
        }
    }

    @GetMapping("listTest")
    public Result listTest(){
        List<Long> list= new ArrayList<>();
        list.add(Long.parseLong("1"));
        list.add(Long.parseLong("2"));
        list.add(Long.parseLong("3"));
        list.add(Long.parseLong("4"));
        Test test = new Test();
        String l = list.toString();
        test.setTestLong(l);
        test.setTest("Test Long");
        System.out.println(StringToList.longList(l));
        System.out.println(list.toString());
        if(testService.addTest(test) > 0) return ResultGenerator.genSuccessResult(list);
        return ResultGenerator.genFailResult();
    }

    @GetMapping("/testList")
    public Result testList(){
        List<Long> list1 = new ArrayList<>();
        List<Long> list2 = new ArrayList<>();

        list1.add(Long.parseLong("1"));
        list1.add(Long.parseLong("2"));
        list1.add(Long.parseLong("3"));
        list1.add(Long.parseLong("4"));
        list2.add(Long.parseLong("3"));
        list2.add(Long.parseLong("4"));
        list2.add(Long.parseLong("5"));
        System.out.println(list1);
        System.out.println(list2);
        list1.removeAll(list2);
        System.out.println(list1);
        return ResultGenerator.genSuccessResult();
    }

    @GetMapping("/testEditDicName")
    public Result testEditDicName(){
        File file = new File(FileUtil.IMAGE_DIC);
        File[] dics = file.listFiles();
        for(File x : dics){
            if(x.isDirectory()) x.renameTo(new File(FileUtil.IMAGE_DIC + x.getName()+"_test"));
        }

        return ResultGenerator.genSuccessResult();
    }

    @GetMapping("tranTest")
    @Transactional(isolation = Isolation.SERIALIZABLE)
    public Result tranTest(){
        List<Long> list= new ArrayList<>();
        list.add(Long.parseLong("1"));
        list.add(Long.parseLong("2"));
        list.add(Long.parseLong("3"));
        list.add(Long.parseLong("4"));
        Test test = new Test();
        /*String l = list.toString();
        test.setTestLong(l);
        test.setTest("Test Long");
        testService.addTest(test);
        testService.addTest(test);
        testService.addTest(test);
        testService.addTest(test);*/
        if(testService.addTest(test) > 0){
            try {
                throw new RuntimeException("test");
            } catch (Exception e) {
                TransactionAspectSupport.currentTransactionStatus().setRollbackOnly();
                return ResultGenerator.genFailResult(e.getMessage());
            }
        }
        return ResultGenerator.genSuccessResult();
    }

    @GetMapping("/testVideo")
    public Result testVideo() throws IOException {
        File file = new File("/home/kiritoghy/labelprojectdata/video/videotest");
        if(!file.exists())return ResultGenerator.genFailResult();
        File[] videos = file.listFiles();
        for(File video: videos){
            System.out.println(VideoUtil.getDuration(video));
            //System.out.println(VideoUtil.getThumb(video, FileUtil.VIDEO_S_DIC + "videotest"));
        }
        return ResultGenerator.genSuccessResult();
    }
}
