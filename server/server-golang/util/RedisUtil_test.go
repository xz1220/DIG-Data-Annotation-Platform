package util

import (
	"fmt"
	"labelproject-back/common"
	"testing"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func Test(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	redis := common.GetCache()
	redisUtilInstance := RedisUtilInstance(redis)
	err := redisUtilInstance.AddTokenTORedis("Test", "Test_User", "127.0.0.1")
	if err != nil {
		fmt.Println("Add Data False!")
	} else {
		fmt.Println("Add Data Success!")
	}

	result, err := redisUtilInstance.HGet("Test", "expirationTime")
	if err != nil {
		fmt.Println("Get Record Error!")
	} else {
		fmt.Println("Record is:", result)
	}

	expirationTime, _ := time.ParseInLocation("2006-01-02 15:04:05", result[:19], time.Local)
	now := time.Now()

	fmt.Println(time.Now().Format(result))
	fmt.Println(now.String())
	fmt.Println(result[:19])
	fmt.Println(expirationTime)
	fmt.Println(now.After(expirationTime))
	// timeParse, err := parseWithLocation("Asia/Shanghai", "2020-06-19 08:44:52.190766 +0000 UTC")
	// if err != nil {
	// 	panic("Error")
	// }
	// fmt.Println(timeParse)
}

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

func TestHasKey(t *testing.T) {
	common.InitConfig("/home/xingzheng/labelproject-back")
	common.InitDB()
	redis := common.GetCache()
	redisUtilInstance := RedisUtilInstance(redis)

	isExit, err := redisUtilInstance.HasKey("Test")
	if err != nil {
		fmt.Println("err!=nil isExit:", isExit)
	} else {
		fmt.Println("err==nil isExit:", isExit)
	}

}

func TestForTime(t *testing.T) {
	fmt.Println(time.Now().AddDate(0, 0, 7).String())
}
