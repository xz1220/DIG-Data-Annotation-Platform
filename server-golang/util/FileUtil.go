package util

import (
	"log"
	"os"

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
	LIMITED_LENGTH   int64
}

var fileUtilInstance = &fileUtil{}

func FileUtilInstance() fileUtil {
	fileUtilInstance.IMAGE_DIC = "/home/kiritoghy/labelprojectdata/image/"
	fileUtilInstance.IMAGE_S_DIC = "/home/kiritoghy/labelprojectdata/images/"
	fileUtilInstance.IMAGE_L_DIC = "/home/kiritoghy/labelprojectdata/imagel/"
	fileUtilInstance.IMAGE_DELETE_DIC = "/home/kiritoghy/labelprojectdata/imaged/"
	fileUtilInstance.VIDEO_DIC = "/home/kiritoghy/labelprojectdata/video/"
	fileUtilInstance.VIDEO_D_DIC = "/home/kiritoghy/labelprojectdata/videod/"
	fileUtilInstance.VIDEO_S_DIC = "/home/kiritoghy/labelprojectdata/videos/"
	fileUtilInstance.LIMITED_LENGTH = 4194304
	return *fileUtilInstance
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
