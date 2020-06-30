package com.uestc.labelproject.service;

import com.uestc.labelproject.entity.Data;
import com.uestc.labelproject.entity.Image;
import com.uestc.labelproject.entity.RleData;
import com.uestc.labelproject.entity.TempRleData;
import org.apache.ibatis.annotations.Param;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:48
 */
public interface AdminImageService {

    Long addImage(Image image);

    int addImages(List<Image> images);

    List<Image> getImageList(Long taskId);

    int updateImages(List<Image> images);

    int updateImageWH(Image image);

    @Transactional
    boolean saveLabel(List<Data> dataList, Long userId, Long imageId, List<Long> dataIds) throws Exception;

    Image getImage(Long imageId);

    List<Data> getDatas(Long imageId,Long userId);

    int editImageByImageId(Long imageId, String imageName);

    int setFinalVersion(Long imageId, Long userConfirmId);

    int finish(Long UserId, Long imageId);

    List<Long> getImageIds(Long taskId);

    List<Long> getDataIds(List<Long> imageIds);

    int deleteImagesByTaskId(Long taskId);

    int deleteDatasByImageId(List<Long> imageIds);

    /**
     * 为接下来的在界面上的删除操作做准备，记住删除的时候不仅仅要删除数据库，还要删除
     * @param imageId
     * @return
     */
    int deleteFromImageByImageId(Long imageId);

    int deleteFromImageDataByImageId(Long imageId);

    int deleteFromImageDataPointsByImageId(Long imageId);

    int deletePoints(List<Long> dataIds);

    int deleteFinish(Long taskId);

    List<Long> getLabeledImageIds(Long taskId, Long userId);

    int updateImagesTaskId(List<Image> images, Long taskId);

    List<Long> getDataIds(Long userId, Long imageId);

    int addRle(List<TempRleData> tempRleDatas, Long userId, Long imageId);

    int deleteRle(Long userId, Long imageId);

    TempRleData getTempRleData(Long dataId);

    int deleteFinalVersion(Long imageId);
}


