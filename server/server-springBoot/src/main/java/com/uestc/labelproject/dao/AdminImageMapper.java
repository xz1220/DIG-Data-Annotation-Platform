package com.uestc.labelproject.dao;

import com.uestc.labelproject.entity.Data;
import com.uestc.labelproject.entity.Image;
import com.uestc.labelproject.entity.RleData;
import com.uestc.labelproject.entity.TempRleData;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 下午7:01
 */
@Repository
public interface AdminImageMapper {
    /**
     * 插入图片
     * @param image
     * @return
     */
    Long addImage(Image image);

    /**
     * 批量插入图片
     * @param images
     * @return
     */
    int addImages(@Param("images") List<Image> images);

    /**
     * 批量更新图片
     * @param images
     * @return
     */
    int updateImages(@Param("images")List<Image> images);

    /**
     * 更新图片width和height
     * @param image
     * @return
     */
    int updateImageWH(Image image);

    /**
     * 获取图片列表
     * @param taskId
     * @return
     */
    List<Image> getImageList(Long taskId);

    /**
     * 获取dataId
     * @param userId
     * @param imageId
     * @return
     */
    List<Long> getDataIds(Long userId, Long imageId);

    /**
     * 删除data
     * @param userId
     * @param imageId
     * @return
     */
    int deleteDatas(Long userId, Long imageId);

    /**
     * 删除point
     * @param userId
     * @param imageId
     * @return
     */
    int deletePoint(Long userId,Long imageId);

    /**
     * 添加data
     * @param data
     * @param userId
     * @param imageId
     * @return
     */
    int addData(Data data, Long userId, Long imageId);

    /**
     * 添加point
     * @param data
     * @param userId
     * @param imageId
     * @return
     */
    int addPoint(@Param("data") Data data, Long userId, Long imageId);

    /**
     * 获取data列表
     * @param userId
     * @param imageId
     * @return
     */
    List<Data> getDatas(Long userId, Long imageId);

    /**
     * 获取图片
     * @param imageId
     * @return
     */
    Image getImage(Long imageId);

    /**
     * 用于png图片压缩后转为jpg更新名字
     * @param imageId
     * @param imageName
     * @return
     */
    int editImageByImageId(Long imageId, String imageName);

    /**
     * 确认最终版本
     * @param imageId
     * @param userConfirmId
     * @return
     */
    int setFinalVersion(Long imageId, Long userConfirmId);

    /**
     * 用户完成标记
     * @param userId
     * @param imageId
     * @return
     */
    int finish(Long userId, Long imageId);

    /**
     * 获取图片id
     * @param taskId
     * @return
     */
    List<Long> getImageIds(Long taskId);

    /**
     * 获取对应图片的data
     * @param imageIds
     * @return
     */
    List<Long> getDataIdsByImageId(@Param("imageIds") List<Long> imageIds);

    /**
     * 删除任务时删除图片
     * @param taskId
     * @return
     */
    int deleteImagesByTaskId(Long taskId);

    /**
     * 删除对应图片的data
     * @param imageIds
     * @return
     */
    int deleteDatasByImageId(@Param("imageIds")List<Long> imageIds);


    int deleteFromImageByImageId(@Param("imageId") Long imageId);

    int deleteFromImageDataByImageId(@Param("imageId") Long imageId);

    int deleteFromImageDataPointsByImageId(@Param("imageId") Long imageId);

    /**
     * 批量删除data
     * @param dataIds
     * @return
     */
    int deletePoints(@Param("dataIds")List<Long> dataIds);

    /**
     * 删除该任务下的用户完成
     * @param taskId
     * @return
     */
    int deleteFinish(Long taskId);

    /**
     * 删除用户完成
     * @param userId
     * @param imageId
     * @return
     */
    int deleteFinishById(Long userId, Long imageId);

    /**
     * 获取已标记图片id
     * @param taskId
     * @param userId
     * @return
     */
    List<Long> getLabeledImageIds(Long taskId, Long userId);

    /**
     * 拆分任务时更新图片所属任务id
     * @param images
     * @param taskId
     * @return
     */
    int updateImagesTaskId(@Param("images") List<Image> images, @Param("taskId") Long taskId);

    /**
     * 添加RLE数据
     * @param tempRleDatas
     * @param userId
     * @param imageId
     * @return
     */
    int addRle(@Param("tempRleDatas")List<TempRleData> tempRleDatas, Long userId, Long imageId);

    /**
     * 删除RLE数据
     * @param userId
     * @param imageId
     * @return
     */
    int deleteRle(Long userId, Long imageId);

    /**
     * 获取RLE数据
     * @param dataId
     * @return
     */
    TempRleData getTempRleData(Long dataId);

    int deleteFinalVersion(Long imageId);
}
