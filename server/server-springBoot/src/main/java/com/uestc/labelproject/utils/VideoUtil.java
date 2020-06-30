package com.uestc.labelproject.utils;

import org.bytedeco.javacv.FFmpegFrameGrabber;
import org.bytedeco.javacv.Frame;
import org.bytedeco.javacv.FrameGrabber;
import org.bytedeco.javacv.Java2DFrameConverter;

import javax.imageio.ImageIO;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

/**
 * @Auther: kiritoghy
 * @Date: 19-10-7 下午5:35
 */
public class VideoUtil {
    public static Double getDuration(File videoFile) throws FrameGrabber.Exception {
        Double duration = 0.0;
        FFmpegFrameGrabber ff = new FFmpegFrameGrabber(videoFile);
        ff.start();
        duration = ff.getLengthInTime() / (1000.0 * 1000.0);
        ff.stop();
        return duration;
    }
    public static String getThumb(File videoFile, String picPath) throws IOException {
        FFmpegFrameGrabber ff = new FFmpegFrameGrabber(videoFile);
        ff.start();

        // 截取中间帧图片(具体依实际情况而定)
        int i = 0;
        int length = ff.getLengthInFrames();
        Frame frame = null;
        while (i < length) {
            frame = ff.grabFrame();
            if ((i > 10) && (frame.image != null)) {
                break;
            }
            i++;
        }
        Java2DFrameConverter converter = new Java2DFrameConverter();
        BufferedImage srcImage = converter.getBufferedImage(frame);
        int srcImageWidth = srcImage.getWidth();
        int srcImageHeight = srcImage.getHeight();

        // 对截图进行等比例缩放(缩略图)
        int width = 480;
        int height = (int) (((double) width / srcImageWidth) * srcImageHeight);
        BufferedImage thumbnailImage = new BufferedImage(width, height, BufferedImage.TYPE_3BYTE_BGR);
        thumbnailImage.getGraphics().drawImage(srcImage.getScaledInstance(width, height, Image.SCALE_SMOOTH), 0, 0, null);

        File picFile = new File(picPath);
        if(!picFile.exists() || !picFile.isDirectory())picFile.mkdir();
        String name = videoFile.getName().substring(0,videoFile.getName().lastIndexOf("."));
        File destFile = new File(picPath + "/" + "thumb_" + name + ".jpg");
        ImageIO.write(thumbnailImage, "jpg", destFile);
        ff.stop();
        return "thumb_" + name + ".jpg";
    }
}
