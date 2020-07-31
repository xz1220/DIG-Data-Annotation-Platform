package repository

import (
	"fmt"
	"labelproject-back/model"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

// AdminImageRepository is ...
type AdminImageRepository interface {
	//插入图片
	AddImage(image model.Image) error

	//批量插入图片
	AddImages(images []*model.Image) error

	//批量更新图片
	UpdateImages(images []*model.Image) error

	//更新图片的宽和高
	UpdateImageWH(image model.Image) error

	//获取图片列表
	GetImageList(taskID int64) ([]*model.Image, error)

	//获取DATAID
	GetDataIDs(userID, imageID int64) ([]int64, error)

	//删除数据
	DeleteDatas(userID, imageID int64) error

	//删除数据点
	DeletePoint(userID, imageID int64) error

	//添加Data
	AddData(data model.ImageData) error

	//添加point
	AddPoints(data []*model.ImageDataPoints) error

	//获取data列表
	GetDatas(userID, imageID int64) ([]*model.DataForResponse, error)

	//获取图片
	GetImage(imageID int64) (model.Image, error)

	//用于png图片压缩后转为jpg更新名字
	EditImageByImageID(imageID int64, imageName string) error

	//确认最终版本
	SetFinalVersion(imageID int64, userConfirmID int64) error

	//用户完成标记
	Finish(userID, imageID int64) error

	//获取图片ID
	GetImageIDs(taskID int64) ([]int64, error)

	//获取对应图片的data
	GetDataIDByImageID(imageIDs []int64) ([]int64, error)

	//删除任务时删除图片
	DeleteImagesByTaskID(taskID int64) error

	//删除对应图片的data
	DeleteDatasByImageID(imageIds []int64) error

	DeleteFromImageByImageID(imageID int64) error
	DeleteFromImageDataByImageID(imageID int64) error
	DeleteFromImagePointsByImageID(imageID int64) error

	//批量删除Data
	DeletePoints(dataIDs []int64) error

	//删除该任务下的用户完成
	DeleteFinish(taskID int64) error

	//删除用户完成
	DeleteFinishByID(userID, imageID int64) error

	//获取已标记图片ID
	GetLabeledImageIDs(taskID int64, userID int64) ([]int64, error)

	//拆分任务时更新图片所属ID
	UpdateImagesTaskID(images []*model.Image, taskID int64) error

	//添加RLE数据
	AddRle(rleDatas []*model.Imagedatarle, userID, imageID int64) error

	//删除RLE数据
	DeleteRle(userID, imageID int64) error

	//获取Rle数据
	GetTempRleData(dataID int64) (model.Imagedatarle, error)

	//删除最终版本
	DeleteFinalVersion(imageID int64) error

	//
	SaveLabel(dataList model.LabelData, dataIDs []int64) error
}

type adminImageRepository struct {
	db *gorm.DB
}

// This statement verifies interface compliance.
var adminImageRepositoryInstance = &adminImageRepository{}

// AdminImageRepositoryInstance returns the instance of AdminImageRepository
func AdminImageRepositoryInstance(db *gorm.DB) AdminImageRepository {
	adminImageRepositoryInstance.db = db
	return adminImageRepositoryInstance
}

//添加图片
func (r *adminImageRepository) AddImage(image model.Image) error {
	err := r.db.Create(&image).Error
	return err
}

//批量插入图片
func (r *adminImageRepository) AddImages(images []*model.Image) error {

	for _, image := range images {
		err := r.db.Create(image).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//批量更新图片
func (r *adminImageRepository) UpdateImages(images []*model.Image) error {

	for _, image := range images {
		err := r.db.Model(&image).Where("image_id = ?", image.ImageID).Updates(model.Image{ImageID: image.ImageID, ImageName: image.ImageName, ImageThumb: image.ImageThumb, UserComfirmID: image.UserComfirmID, TaskID: image.TaskID, Width: image.Width, Height: image.Height}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//更新图片的宽和高
func (r *adminImageRepository) UpdateImageWH(image model.Image) error {

	err := r.db.Model(&image).Where("image_id = ?", image.ImageID).Updates(model.Image{ImageID: image.ImageID, Width: image.Width, Height: image.Height}).Error
	if err != nil {
		return err
	}

	return nil
}

//获取图片列表
func (r *adminImageRepository) GetImageList(taskID int64) ([]*model.Image, error) {
	var images []*model.Image
	err := r.db.Where("task_id = ?", taskID).Find(&images).Error
	return images, err
}

//获取DATAID
func (r *adminImageRepository) GetDataIDs(userID, imageID int64) ([]int64, error) {
	var imageDatas []*model.ImageData
	err := r.db.Where("image_id = ? AND user_id = ?", imageID, userID).Find(&imageDatas).Error
	if err != nil {
		return nil, err
	}

	var datas []int64
	for _, imageData := range imageDatas {
		datas = append(datas, imageData.DataID)
	}
	return datas, nil
}

//删除数据
func (r *adminImageRepository) DeleteDatas(userID, imageID int64) error {
	err := r.db.Where("image_id = ? AND user_id = ?", imageID, userID).Delete(&model.ImageData{}).Error
	return err
}

//删除数据点
func (r *adminImageRepository) DeletePoint(userID, imageID int64) error {
	err := r.db.Where("image_id = ? AND user_id = ?", imageID, userID).Delete(&model.ImageDataPoints{}).Error
	return err
}

//添加Data
func (r *adminImageRepository) AddData(data model.ImageData) error {
	err := r.db.Create(&data).Error
	return err
}

//添加point
func (r *adminImageRepository) AddPoints(datas []*model.ImageDataPoints) error {
	for _, data := range datas {
		err := r.db.Create(data).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//获取data列表
func (r *adminImageRepository) GetDatas(userID, imageID int64) ([]*model.DataForResponse, error) {
	Datas := make([]*model.Data, 0)
	err := r.db.Raw("select d.*,dp.order,dp.x,dp.y from imagedata d left join imagedatapoints dp on dp.data_id = d.data_id where d.user_id = ? and d.image_id = ?", userID, imageID).Scan(&Datas).Error
	if err != nil {
		log.Println("GetDatas: ", err)
		return nil, err
	}

	var dataForResponses []*model.DataForResponse
	for _, data := range Datas {

		var dataForResponse model.DataForResponse
		dataForResponse.DataID = data.DataID
		if len(dataForResponses) > 0 && dataForResponse.DataID == dataForResponses[len(dataForResponses)-1].DataID {
			var point model.Points
			point.X = data.X
			point.Y = data.Y
			point.Order = data.Order

			dataForResponses[len(dataForResponses)-1].Point = append(dataForResponses[len(dataForResponses)-1].Point, point)
			continue
		}

		dataForResponse.ImageID = data.ImageID
		dataForResponse.LabelID = data.LabelID
		dataForResponse.UserID = data.UserID
		dataForResponse.DataDesc = data.DataDesc
		dataForResponse.LabelType = data.LabelType
		dataForResponse.Iscrowd = data.Iscrowd

		var point model.Points
		point.X = data.X
		point.Y = data.Y
		point.Order = data.Order

		dataForResponse.Point = append(dataForResponse.Point, point)

		dataForResponses = append(dataForResponses, &dataForResponse)
	}

	return dataForResponses, err
}

//获取图片
func (r *adminImageRepository) GetImage(imageID int64) (model.Image, error) {
	var image model.Image
	err := r.db.Where("image_id = ?", imageID).First(&image).Error
	return image, err
}

//用于png图片压缩后转为jpg更新名字
func (r *adminImageRepository) EditImageByImageID(imageID int64, imageName string) error {
	err := r.db.Model(&model.Image{}).Where("image_id = ?", imageID).Updates(model.Image{ImageName: imageName}).Error
	if err != nil {
		return err
	}
	return nil
}

//确认最终版本
func (r *adminImageRepository) SetFinalVersion(imageID int64, userConfirmID int64) error {
	err := r.db.Model(&model.Image{}).Where("image_id = ?", imageID).Updates(model.Image{UserComfirmID: userConfirmID}).Error
	if err != nil {
		return err
	}
	return nil
}

//用户完成标记
func (r *adminImageRepository) Finish(userID, imageID int64) error {
	var image model.Image
	err := r.db.Table("image").Select("task_id").Where("image_id = ?", imageID).First(&image).Error
	if err != nil {
		return err
	}
	err = r.db.Create(&model.UserFinished{UserID: userID, TaskID: image.ImageID, ImageID: imageID}).Error
	if err != nil {
		return err
	}
	return nil
}

//获取图片ID
func (r *adminImageRepository) GetImageIDs(taskID int64) ([]int64, error) {
	var imageID []int64
	var images []*model.Image
	err := r.db.Where("task_id = ?", taskID).Find(&images).Error
	if err != nil {
		return nil, err
	}
	for _, image := range images {
		imageID = append(imageID, image.ImageID)
	}
	return imageID, err
}

//获取对应图片的data
func (r *adminImageRepository) GetDataIDByImageID(imageIDs []int64) ([]int64, error) {
	var ImageIDs []int64
	var image []*model.ImageData
	for _, imageID := range imageIDs {
		err := r.db.Where("image_id = ?", imageID).Find(&image).Error
		if err != nil {
			return nil, err
		}
		for _, instance := range image {
			ImageIDs = append(ImageIDs, instance.DataID)
		}
		image = nil
	}
	return ImageIDs, nil
}

//删除任务时删除图片
func (r *adminImageRepository) DeleteImagesByTaskID(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.Image{}).Error
	return err
}

//删除对应图片的data
func (r *adminImageRepository) DeleteDatasByImageID(imageIDs []int64) error {
	for _, imageID := range imageIDs {
		err := r.db.Where("image_id = ?", imageID).Delete(&model.ImageData{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *adminImageRepository) DeleteFromImageByImageID(imageID int64) error {
	err := r.db.Where("image_id = ?", imageID).Delete(&model.Image{}).Error
	return err
}
func (r *adminImageRepository) DeleteFromImageDataByImageID(imageID int64) error {
	err := r.db.Where("image_id = ?", imageID).Delete(&model.ImageData{}).Error
	return err
}
func (r *adminImageRepository) DeleteFromImagePointsByImageID(imageID int64) error {
	err := r.db.Where("image_id = ?", imageID).Delete(&model.ImageDataPoints{}).Error
	return err
}

//批量删除Data
func (r *adminImageRepository) DeletePoints(dataIDs []int64) error {
	for _, dataID := range dataIDs {
		err := r.db.Where("data_id = ?", dataID).Delete(&model.ImageDataPoints{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//删除该任务下的用户完成
func (r *adminImageRepository) DeleteFinish(taskID int64) error {
	err := r.db.Where("task_id = ?", taskID).Delete(&model.UserFinished{}).Error
	return err
}

//删除用户完成
func (r *adminImageRepository) DeleteFinishByID(userID, imageID int64) error {
	err := r.db.Where("user_id = ? AND image_id = ?", userID, imageID).Delete(&model.UserFinished{}).Error
	return err
}

//获取已标记图片ID
func (r *adminImageRepository) GetLabeledImageIDs(taskID int64, userID int64) ([]int64, error) {
	var imageIDs []int64
	var userfinished []*model.UserFinished
	if userID != 0 {
		err := r.db.Where("task_id = ? AND user_id = ?", taskID, userID).Find(&userfinished).Error
		if err != nil {
			return nil, err
		}
		for _, user := range userfinished {
			imageIDs = append(imageIDs, user.ImageID)
		}
	} else {
		err := r.db.Where("task_id = ?", taskID).Find(&userfinished).Error
		if err != nil {
			return nil, err
		}
		for _, user := range userfinished {
			imageIDs = append(imageIDs, user.ImageID)
		}
	}
	return imageIDs, nil
}

//拆分任务时更新图片所属ID
func (r *adminImageRepository) UpdateImagesTaskID(images []*model.Image, taskID int64) error {
	for _, image := range images {
		image.TaskID = taskID
		err := r.db.Model(&model.Image{}).Updates(&image).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//添加RLE数据
func (r *adminImageRepository) AddRle(rleDatas []*model.Imagedatarle, userID, imageID int64) error {
	for _, imagedatarle := range rleDatas {
		err := r.db.Create(imagedatarle).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//删除RLE数据
func (r *adminImageRepository) DeleteRle(userID, imageID int64) error {
	err := r.db.Where("image_id = ? AND user_id = ?", imageID, userID).Delete(&model.Imagedatarle{}).Error
	return err
}

//获取Rle数据
func (r *adminImageRepository) GetTempRleData(dataID int64) (model.Imagedatarle, error) {
	var imagedatarle model.Imagedatarle
	err := r.db.Where("data_id = ?", dataID).First(&imagedatarle).Error
	return imagedatarle, err
}

//删除最终版本
func (r *adminImageRepository) DeleteFinalVersion(imageID int64) error {
	err := r.db.Model(&model.Image{}).Where("image_id = ?", imageID).Update(model.Image{UserComfirmID: 0}).Error
	return err
}

func (r *adminImageRepository) SaveLabel(dataList model.LabelData, dataIDs []int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		imageID, err := strconv.ParseInt(dataList.ImageIDString, 10, 64)
		if err != nil {
			return fmt.Errorf("字符串转Int64 出错")
		}
		if len(dataIDs) > 0 {
			if err = r.DeleteDatas(dataList.UserID, imageID); err != nil {
				return err
			}
			if err = r.DeletePoint(dataList.UserID, imageID); err != nil {
				return err
			}
			if err = r.DeleteFinishByID(dataList.UserID, imageID); err != nil {
				return err
			}
		}

		datas := dataList.Data
		if datas != nil && len(datas) > 0 {
			for _, data := range datas {
				imageData := model.ImageData{
					ImageID:   imageID,
					LabelID:   data.LabelID,
					LabelType: data.LabelType,
					UserID:    dataList.UserID,
					DataDesc:  data.DataDesc,
					Iscrowd:   data.IScrowd,
				}

				// 添加数据
				if err = r.AddData(imageData); err != nil {
					return err
				}

				var lastRecord model.ImageData
				if err = tx.Where("image_id = ? AND label_id = ? AND user_id = ? ", imageData.ImageID, imageData.LabelID, imageData.UserID).Last(&lastRecord).Error; err != nil {
					return err
				}

				//获取其中的Points数据
				var imageDataPoints []*model.ImageDataPoints
				points := data.Point
				for _, point := range points {
					var imageDataPoint = &model.ImageDataPoints{
						DataID:  lastRecord.DataID,
						Order:   point.Order,
						X:       point.X,
						Y:       point.Y,
						ImageID: imageID,
						UserID:  dataList.UserID,
					}
					imageDataPoints = append(imageDataPoints, imageDataPoint)
				}

				//添加
				if imageDataPoints != nil && len(imageDataPoints) > 0 {
					if err = r.AddPoints(imageDataPoints); err != nil {
						log.Println("Add Points Error")
						return err
					}
				}

			}

			if err = r.Finish(dataList.UserID, imageID); err != nil {
				return err
			}

			if err = r.DeleteFinalVersion(imageID); err != nil {
				return err
			}

		}
		return nil

	})
}
