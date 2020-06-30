package repository

import (
	"fmt"
	"labelproject-back/common"
	"testing"
)

func TestSearchLabel(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	adminImageLabelInstance := AdminImageLabelRepositoryInstance(db)

	imagelabel, err := adminImageLabelInstance.SearchLabel("Blue")
	if err != nil {
		panic("Error")
	}
	fmt.Println("ImageLabel: ", imagelabel[0].LabelName)
}

func TestGetLabelByImageID(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	adminImageLabelInstance := AdminImageLabelRepositoryInstance(db)

	imagelabel, err := adminImageLabelInstance.GetLabelByImageID(2)
	if err != nil {
		panic("Error")
	}
	fmt.Println("ImageLabel: ", imagelabel[0].LabelName)
}

func TestFindByLabelName(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	db := common.GetDB()
	fmt.Println("successfully connect mysql")
	adminImageLabelInstance := AdminImageLabelRepositoryInstance(db)

	imagelabel, err := adminImageLabelInstance.FindByLabelName("TESTTTTTT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(imagelabel)
}
