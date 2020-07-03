package common

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 关系型数据库
var DB *gorm.DB

// Cache 缓存
var Cache *redis.Client

// InitDB 初始化数据库
func InitDB() {
	//redis
	redisURL := viper.GetString("redis.host") + ":" + viper.GetString("redis.port")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

RedisConnection:
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	if err != nil {
		time.Sleep(10000000000)
		goto RedisConnection
	} else {
		log.Println("Redis Connect success!")
	}

	//mysql
	dirverName := viper.GetString("datasource.dirverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	fmt.Println(args)

MysqlConnection:
	db, err := gorm.Open(dirverName, args)
	if err != nil {
		// panic("failed to connect database,err:" + err.Error())
		time.Sleep(10000000000)
		goto MysqlConnection
	} else {
		log.Println("Mysql Connected success!")
	}

	// db.AutoMigrate(&model.User{})
	DB = db
	Cache = client
}

// GetDB 获取关系型数据库
func GetDB() *gorm.DB {
	return DB
}

// GetCache 获取缓存
func GetCache() *redis.Client {
	return Cache
}
