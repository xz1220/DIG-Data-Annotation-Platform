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
var db *gorm.DB

// cache 缓存
var cache *redis.Client

// InitDB 初始化数据库
// 旧有的函数，
func InitDB() {
	// db.AutoMigrate(&model.User{})
	db = initMysql()
	db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	db.DB().SetMaxOpenConns(30)                   //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

	cache = initRedis()
}

// initMysql
func initMysql() *gorm.DB {
	//mysql
	dirverName := viper.GetString("datasource.dirverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	MysqlArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	fmt.Println(MysqlArgs)

	var connection *gorm.DB
	for {
		TempConnection, err := gorm.Open(dirverName, MysqlArgs)
		if err != nil {
			continue
		} else {
			log.Println("Mysql Connected success!")
			connection = TempConnection
			break
		}
	}

	return connection
}

// initRedis
func initRedis() *redis.Client {
	//redis
	redisURL := viper.GetString("redis.host") + ":" + viper.GetString("redis.port")

	// 创建Redis连接
	var client *redis.Client
	for {
		client = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong, err := client.Ping().Result()
		log.Println(pong, err)
		if err != nil {
			continue
		} else {
			log.Println("Redis Connect success!")
			break
		}
	}

	return client
}

// NewMysqlConnection
func NewMysqlConnection() *gorm.DB {
	// db.AutoMigrate(&model.User{})
	db = initMysql()
	db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	db.DB().SetMaxOpenConns(30)                   //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

	return db
}

// GetDB 获取关系型数据库
func GetDB() *gorm.DB {
	if err := db.DB().Ping(); err != nil {
		db.Close()
		db = NewMysqlConnection()
	}
	return db
}

// NewRedisConnection
func NewRedisConnection() *redis.Client {
	cache = initRedis()
	return cache
}

// Getcache 获取缓存
func GetCache() *redis.Client {
	if _, err := cache.Ping().Result(); err != nil {
		cache.Close()
		cache = NewRedisConnection()
	}
	return cache
}
