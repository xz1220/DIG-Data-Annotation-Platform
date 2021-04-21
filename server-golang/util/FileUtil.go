package util

import (
	"labelproject-back/common"
	"log"
	"os"
	"reflect"

	"github.com/disintegration/imaging"
)

type fileUtil struct {
	IMAGE_DIC        string
	IMAGE_S_DIC      string
	IMAGE_L_DIC      string
	IMAGE_DELETE_DIC string
	VIDEO_DIC        string
	VIDEO_D_DIC      string
	VIDEO_S_DIC      string
	LIMITED_LENGTH   int
}

var fileUtilInstance = &fileUtil{}

func FileUtilInstance() *fileUtil {
	fileUtilInstance.IMAGE_DIC = common.IMAGE_DIC
	fileUtilInstance.IMAGE_S_DIC = common.IMAGE_S_DIC
	fileUtilInstance.IMAGE_L_DIC = common.IMAGE_L_DIC
	fileUtilInstance.IMAGE_DELETE_DIC = common.IMAGE_DELETE_DIC
	fileUtilInstance.VIDEO_DIC = common.VIDEO_DIC
	fileUtilInstance.VIDEO_D_DIC = common.VIDEO_D_DIC
	fileUtilInstance.VIDEO_S_DIC = common.VIDEO_S_DIC
	fileUtilInstance.LIMITED_LENGTH = common.LIMITED_LENGTH
	log.Println("FileUtilInstance() :", fileUtilInstance.IMAGE_DIC)
	return fileUtilInstance
}

func (f *fileUtil) Thumb(src, dest, imageName string) (string, int, int, error) {
	image, err := imaging.Open(src)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y
	image = imaging.Resize(image, int(0.2*float64(image.Bounds().Max.X)), 0, imaging.Lanczos)
	_, err = os.Open(dest)
	if os.IsNotExist(err) {
		log.Println(dest, "is not exsit")
		err = os.Mkdir(dest, os.ModePerm)
		if err != nil {
			log.Println("Create Dir Error")
		}
	}

	err = imaging.Save(image, dest+"/"+"thumb_"+imageName)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
		return "", 0, 0, err
	}
	return "thumb_" + imageName, width, height, nil

}

func (f *fileUtil) CreateDir(path interface{}) error {
	err := os.MkdirAll(path.(string), 0777)
	if err != nil {
		return err
	}
	return nil
}

func (f *fileUtil) CreateBaseDir() {

	dirs := FileUtilInstance()
	t := reflect.TypeOf(*dirs)
	v := reflect.ValueOf(*dirs)

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name != "LIMITED_LENGTH" {
			if _, err := os.Stat(v.Field(i).String()); os.IsNotExist(err) {
				f.CreateDir(v.Field(i).Interface())
			}
		}
	}
}
