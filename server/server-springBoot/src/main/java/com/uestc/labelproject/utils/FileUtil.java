package com.uestc.labelproject.utils;

import lombok.extern.slf4j.Slf4j;
import net.coobird.thumbnailator.Thumbnails;
import org.apache.commons.io.FileUtils;

import java.io.*;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-25 下午7:27
 */
@Slf4j
public class FileUtil {
//    public static final String IMAGE_DIC = "/opt/labelproject/image";
//    public static final String IMAGE_S_DIC = "/opt/labelproject/images";
    public static final String IMAGE_DIC = "/home/kiritoghy/labelprojectdata/image/";
    public static final String IMAGE_S_DIC = "/home/kiritoghy/labelprojectdata/images/";
    public static final String IMAGE_L_DIC = "/home/kiritoghy/labelprojectdata/imagel/";
    public static final String IMAGE_DELETE_DIC = "/home/kiritoghy/labelprojectdata/imaged/";
    public static final String VIDEO_DIC = "/home/kiritoghy/labelprojectdata/video/";
    public static final String VIDEO_D_DIC = "/home/kiritoghy/labelprojectdata/videod/";
    public static final String VIDEO_S_DIC = "/home/kiritoghy/labelprojectdata/videos/";
    public static final long LIMITED_LENGTH = 4194304;//限制为4MB

    /**
     * 压缩图片
     * @param src 图片文件夹路径
     * @param image 需压缩的图片
     * @return
     * @throws IOException
     */
    public static boolean resizeImage(String src, File image) throws IOException {
        String suffix = image.getName().substring(image.getName().lastIndexOf(".") + 1);
        String name = image.getName().substring(0,image.getName().lastIndexOf("."));
        if(!suffix.toLowerCase().equals("jpg") && !suffix.toLowerCase().equals("jpeg")){
            Thumbnails
                    .of(image).scale(1f).outputQuality(0.3f).outputFormat("jpg").toFile(src+"/"+name+".jpg");
            FileUtils.forceDelete(image);
            return true;
        }
        else {
            Thumbnails.of(image).scale(1f).outputQuality(0.3f).outputFormat(suffix).toFile(src+"/"+name+"."+suffix);
            return false;
        }
    }

    /**
     * 移动文件
     * @param src 目标文件
     * @param destdic 移动的目标文件夹路径
     * @throws IOException
     */
    public static void moveFile(File src, String destdic) throws IOException {
        File dic = new File(destdic);
        if(!dic.exists() || !dic.isDirectory())dic.mkdir();
        File dest = new File(destdic+"/"+src.getName());
        FileUtils.copyFile(src, dest);
    }

    /**
     * 生成缩略图
     * @param image 目标图片文件
     * @param thumbPath 缩略图文件夹路径
     * @return
     * @throws IOException
     */
    public static String thumb(File image, String thumbPath) throws IOException {
        File thumb = new File(thumbPath); //读取文件
        String name = image.getName().substring(0,image.getName().lastIndexOf(".")); //获取文件名
        if(!thumb.exists() || !thumb.isDirectory()) thumb.mkdir(); //对意外情况创建文件夹
        Thumbnails.of(image).outputFormat("jpg").scale(0.3f).toFile(thumbPath + "/" + "thumb_"+name + ".jpg"); //生成缩略图
        return "thumb_"+name + ".jpg";
    }

    /**
     * 重命名
     * @param srcp 目标文件路径
     * @param destp 重命名文件路径
     * @return
     */
    public static boolean rename(String srcp, String destp){
        File src = new File(srcp);
        File dest = new File(destp);

        if(src.isDirectory()){
            return src.renameTo(dest);
        }
        return false;
    }
}
