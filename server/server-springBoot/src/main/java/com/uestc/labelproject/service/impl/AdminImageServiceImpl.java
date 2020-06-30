package com.uestc.labelproject.service.impl;

import com.uestc.labelproject.dao.AdminImageMapper;
import com.uestc.labelproject.entity.Data;
import com.uestc.labelproject.entity.Image;
import com.uestc.labelproject.entity.RleData;
import com.uestc.labelproject.entity.TempRleData;
import com.uestc.labelproject.service.AdminImageService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;

//这个类操作了image，

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 上午9:48
 */
@Service
@Slf4j
public class AdminImageServiceImpl implements AdminImageService {

    @Autowired
    AdminImageMapper adminImageMapper;

    @Override
    public Long addImage(Image image) {
        return adminImageMapper.addImage(image);
    }

    @Override
    public int addImages(List<Image> images) {
        return adminImageMapper.addImages(images);
    }

    @Override
    public List<Image> getImageList(Long taskId) {
        return adminImageMapper.getImageList(taskId);
    }

    @Override
    public int updateImages(List<Image> images) {
        return adminImageMapper.updateImages(images);
    }


    @Override
    @Transactional(isolation = Isolation.READ_COMMITTED)
    public boolean saveLabel(List<Data> dataList, Long userId, Long imageId, List<Long> dataIds) throws Exception{

        //删除原有的数据，添加新的数据
        if(dataIds.size() > 0){
            adminImageMapper.deleteDatas(userId, imageId); //imagedata
            adminImageMapper.deletePoint(userId,imageId); //imagedatapoint
            adminImageMapper.deleteFinishById(userId,imageId); //userfinished
        }
        if(dataList != null && dataList.size() > 0){
            for(Data data: dataList){
                adminImageMapper.addData(data, userId, imageId); //imagedata
                if(data.getPoint()!= null && data.getPoint().size() > 0)
                if(adminImageMapper.addPoint(data,userId, imageId) <= 0) {
                    log.info("addPoint Wrong");
                    throw new Exception("addPoint Wrong");
                }
            }
            adminImageMapper.finish(userId, imageId);  //userfinished
            adminImageMapper.deleteFinalVersion(imageId);   //更新数据库
            return true;
        }
        return true;
    }

    @Override
    public Image getImage(Long imageId) {
        return adminImageMapper.getImage(imageId);
    }

    @Override
    public List<Data> getDatas(Long imageId, Long userId) {
        return adminImageMapper.getDatas(userId, imageId);
    }

    //修改文件名
    @Override
    public int editImageByImageId(Long imageId, String imageName) {
        return adminImageMapper.editImageByImageId(imageId, imageName);
    }

    @Override
    public int setFinalVersion(Long imageId, Long userConfirmId) {
        return adminImageMapper.setFinalVersion(imageId,userConfirmId);
    }

    @Override
    public int finish(Long UserId, Long imageId) {
        return adminImageMapper.finish(UserId, imageId);
    }

    @Override
    public List<Long> getImageIds(Long taskId) {
        return adminImageMapper.getImageIds(taskId);
    }

    @Override
    public List<Long> getDataIds(List<Long> imageIds) {
        return adminImageMapper.getDataIdsByImageId(imageIds);
    }

    @Override
    public int deleteImagesByTaskId(Long taskId) {
        return adminImageMapper.deleteImagesByTaskId(taskId);
    }

    @Override
    public int deleteDatasByImageId(List<Long> imageIds) {
        return adminImageMapper.deleteDatasByImageId(imageIds);
    }

    @Override
    public int deleteFromImageByImageId(Long imageId) {
        return adminImageMapper.deleteFromImageByImageId(imageId);
    }

    @Override
    public int deleteFromImageDataByImageId(Long imageId) {
        return adminImageMapper.deleteFromImageDataByImageId(imageId);
    }

    @Override
    public int deleteFromImageDataPointsByImageId(Long imageId) {
        return adminImageMapper.deleteFromImageDataPointsByImageId(imageId);
    }



    @Override
    public int deletePoints(List<Long> dataIds) {
        return adminImageMapper.deletePoints(dataIds);
    }

    @Override
    public int deleteFinish(Long taskId) {
        return adminImageMapper.deleteFinish(taskId);
    }

    @Override
    public List<Long> getLabeledImageIds(Long taskId, Long userId) {
        return adminImageMapper.getLabeledImageIds(taskId, userId);
    }

    @Override
    public int updateImagesTaskId(List<Image> images, Long taskId) {
        return adminImageMapper.updateImagesTaskId(images, taskId);
    }

    @Override
    public List<Long> getDataIds(Long userId, Long imageId) {
        return adminImageMapper.getDataIds(userId, imageId);
    }

    @Override
    public int updateImageWH(Image image) {
        return adminImageMapper.updateImageWH(image);
    }

    @Override
    public int addRle(List<TempRleData> tempRleDatas, Long userId, Long imageId) {
        return adminImageMapper.addRle(tempRleDatas, userId, imageId);
    }

    @Override
    public int deleteRle(Long userId, Long imageId) {
        return adminImageMapper.deleteRle(userId, imageId);
    }

    @Override
    public TempRleData getTempRleData(Long dataId) {
        return adminImageMapper.getTempRleData(dataId);
    }

    @Override
    public int deleteFinalVersion(Long imageId) {
        return adminImageMapper.deleteFinalVersion(imageId);
    }
}
