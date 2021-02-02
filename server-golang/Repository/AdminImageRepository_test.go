package repository

import (
	"encoding/json"
	"fmt"
	"labelproject-back/common"
	"labelproject-back/model"
	"strconv"
	"testing"
)

func TestGetImageList(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	images, err := imageInstance.GetImageList(1)
	if err != nil {
		panic("error")
	}
	fmt.Println("Success!!!")
	for _, image := range images {
		fmt.Println("Image_ID：", image.ImageID)
	}
}

func TestAddImage(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	image := model.Image{
		ImageID:   10000,
		ImageName: "Test",
		Width:     100,
		Height:    100,
	}
	err := imageInstance.AddImage(image)
	if err != nil {
		panic("Error")
	}
}

func TestAddImages(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	images := make([]*model.Image, 2)
	images[0] = &model.Image{
		ImageID:   10001,
		ImageName: "Test2",
		Width:     100,
		Height:    100,
	}
	images[1] = &model.Image{
		ImageID:   10002,
		ImageName: "Test3",
		Width:     100,
		Height:    100,
	}

	err := imageInstance.AddImages(images)

	if err != nil {
		panic("Error!!!")
	}
}

func TestUpdateImages(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	images := make([]*model.Image, 2)
	images[0] = &model.Image{
		ImageID:   10001,
		ImageName: "Test-Change1",
		Width:     100,
		Height:    100,
	}
	images[1] = &model.Image{
		ImageID:   10002,
		ImageName: "Test-change2",
		Width:     100,
		Height:    100,
	}

	err := imageInstance.UpdateImages(images)
	if err != nil {
		panic("Error")
	}
}

func TestUpdateImageWH(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	image := model.Image{
		ImageID: 10000,
		Width:   200,
		Height:  200,
	}
	err := imageInstance.UpdateImageWH(image)
	if err != nil {
		panic("Error")
	}
}

func TestGetDataIDs(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	datas, err := imageInstance.GetDataIDs(2, 1)
	if err != nil {
		panic("Error!")
	}
	fmt.Println("datas:", datas)
}

func TestAddData(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	imageData := model.ImageData{
		DataID:    10000,
		ImageID:   1,
		LabelID:   1,
		UserID:    2,
		DataDesc:  "Test",
		LabelType: 1,
		Iscrowd:   0,
	}

	err := imageInstance.AddData(imageData)

	if err != nil {
		panic("Error!!!")
	}
}

func TestAddPoints(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	imageDataPoints := make([]*model.ImageDataPoints, 4)

	for index := range imageDataPoints {
		imageDataPoints[index] = &model.ImageDataPoints{
			DataID:  10000,
			Order:   index + 1,
			X:       float64(index + 10),
			Y:       float64(index + 10),
			ImageID: 1,
			UserID:  2,
		}
	}

	err := imageInstance.AddPoints(imageDataPoints)
	if err != nil {
		panic("Error")
	}
}

func TestDeletDataAndPoints(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	err := imageInstance.DeleteDatas(2, 1)
	if err != nil {
		panic("Error of delete Datas")
	}

	err = imageInstance.DeletePoint(2, 1)
	if err != nil {
		panic("Error of delete Points")
	}
}

func TestGetDatas(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	datas, err := imageInstance.GetDatas(2, 2)
	if err != nil {
		panic("Error")
	}
	for index := range datas {
		fmt.Println("Datas:", datas[index])
	}

}

func TestGetImage(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)
	image, err := imageInstance.GetImage(2)
	if err != nil {
		panic("Error")
	}
	fmt.Println("Image:", image)
}

func TestSaveLabel(t *testing.T) {
	//data := []byte(`{"imageId":"10004","data":[{"labelId":1,"labelType":0,"dataDesc":"","iscrowd":0,"point":[{"x":394.35598705501616,"y":291.62459546925567,"order":1},{"x":987.546925566343,"y":389.3851132686084,"order":2},{"x":739.0032362459547,"y":719.1197411003236,"order":3},{"x":231.97411003236246,"y":497.0873786407767,"order":4}]}],"userId":2}`)
	data := []byte(`{"imageId":"11111","data":[{"labelId":1,"labelType":0,"dataDesc":"","iscrowd":0,"point":[{"x":203.86270022883295,"y":234.32494279176203,"order":1},{"x":476.8512585812357,"y":255.4141876430206,"order":2},{"x":337.4279176201373,"y":529.5743707093822,"order":3},{"x":105.44622425629291,"y":422.95652173913044,"order":4}]},{"labelId":1,"labelType":0,"dataDesc":"","iscrowd":0,"point":[{"x":681.8855835240274,"y":328.05491990846684,"order":1},{"x":841.2265446224256,"y":520.2013729977117,"order":2},{"x":673.6842105263158,"y":633.8489702517162,"order":3},{"x":538.9473684210526,"y":486.2242562929062,"order":4}]}],"userId":2}`)
	var mapData model.LabelData
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("someting: ", mapData)

	// if mapData.Data == nil {
	// 	fmt.Println("successfully")
	// }

	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	imageInstance := AdminImageRepositoryInstance(db)

	imageID, err := strconv.ParseInt(mapData.ImageIDString, 10, 64)
	dataIDs, err := imageInstance.GetDataIDs(mapData.UserID, imageID)
	err = imageInstance.SaveLabel(mapData, dataIDs)
	if err != nil {
		fmt.Println(err)
	}

}

func TestForTest(t *testing.T) {
	type data struct {
		ID int64
	}

	var Datas []*data
	var i int64
	for i = 0; i < 10; i++ {
		var Data = &data{
			ID: i,
		}

		Datas = append(Datas, Data)
	}

	fmt.Println(len(Datas))
}
