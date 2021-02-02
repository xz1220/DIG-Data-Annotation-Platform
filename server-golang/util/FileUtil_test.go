package util

import (
	"fmt"
	"os"
	"testing"
)

func TestOsOperation(t *testing.T) {
	// var path string
	// if os.IsPathSeparator('\\') {
	// 	path = "\\"
	// } else {
	// 	path = "/"
	// }
	// pwd, _ := os.Getwd()
	// // err := os.Mkdir(pwd+path+"tmp3", os.ModePerm)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // err = os.Mkdir(pwd+path+"tmp4", os.ModePerm)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // os.Rename(pwd+path+"tmp2000", "/home/xingzheng/labelproject-back/TestFolder/tmp")
	// os.Remove(pwd + path + "tmp3")

	file, err := os.Open("/home/xingzheng/data/labelproject/home2/kiritoghy/labelprojectdata/image")
	if err != nil {
		panic("open file Error")
	}

	fileInfos, err := file.Readdir(-1)
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.IsDir())
	}

}

func TestImage(t *testing.T) {
	src := "/mnt/c/Users/30249/labelproject-back/test/03.png"
	imageName := "03.png"
	dest := "/mnt/c/Users/30249/labelproject-back/test/"

	fileUtilInstance := FileUtilInstance()
	thumbName, _, _, err := fileUtilInstance.Thumb(src, dest, imageName)
	if err != nil {
		panic("Error")
	}
	fmt.Println(thumbName)

}
