package common

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//init config file
func InitConfig(workDir string) {
	if workDir == "main" {
		Dir, err := os.Getwd()
		if err != nil {
			panic("路径读取失败")
		}
		workDir = Dir
	}
	viper.AddConfigPath(workDir + "/config")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
	log.Print("读取成功")
}
