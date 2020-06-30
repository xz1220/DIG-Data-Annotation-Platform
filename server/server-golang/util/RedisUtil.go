package util

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisUtil interface {
	AddTokenTORedis(token string, userName string, ip string) error
	HSetWithExpirationTime(key string, filed string, value interface{}, expirationTime time.Duration) error
	HSet(key string, filed string, value interface{}) error
	HGet(key string, field string) (result string, err error)
	HDeleteKey(key string) error
	HHasKey(key string, field string) (bool, error)
	HasKey(key string) (bool, error)
	IsBlackList(token string) (bool, error)
	AddBlackList(token string) error
	GetTokenValidTimeByToken(token string) (result string, err error)
	GetUsernameByToken(token string) (result string, err error)
	GetIPByToken(token string) (result string, err error)
}

type redisUtil struct {
	cache *redis.Client
}

var redisUtilInstance = &redisUtil{}

func RedisUtilInstance(cache *redis.Client) RedisUtil {
	redisUtilInstance.cache = cache
	return redisUtilInstance
}

func (Redis *redisUtil) AddTokenTORedis(token string, userName string, ip string) error {

	err := Redis.HSetWithExpirationTime(token, "tokenValidTime", time.Now().AddDate(0, 0, 7).String(), time.Duration(time.Hour*24*7))
	if err != nil {
		return err
	}

	err = Redis.HSetWithExpirationTime(token, "expirationTime", time.Now().Add(time.Hour).String(), time.Duration(time.Hour*24*7))
	if err != nil {
		return err
	}

	err = Redis.HSetWithExpirationTime(token, "username", userName, time.Duration(time.Hour*24*7))
	if err != nil {
		return err
	}

	err = Redis.HSetWithExpirationTime(token, "ip", ip, time.Duration(time.Hour*24*7))
	if err != nil {
		return err
	}

	return nil
}

func (Redis *redisUtil) HSetWithExpirationTime(key string, filed string, value interface{}, expirationTime time.Duration) error {
	err := Redis.cache.HSet(key, filed, value).Err()
	if err != nil {
		return err
	}
	err = Redis.cache.Expire(key, expirationTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func (Redis *redisUtil) HSet(key string, filed string, value interface{}) error {
	err := Redis.cache.HSet(key, filed, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (Redis *redisUtil) HGet(key string, field string) (result string, err error) {
	result, err = Redis.cache.HGet(key, field).Result()
	return result, err
}

func (Redis *redisUtil) HDeleteKey(key string) error {
	err := Redis.cache.Del(key).Err()
	return err
}

func (Redis *redisUtil) HHasKey(key string, field string) (bool, error) {
	isExit, err := Redis.cache.HExists(key, field).Result()
	if err != nil {
		return false, err
	}
	return isExit, err
}

func (Redis *redisUtil) HasKey(key string) (bool, error) {
	isExit, err := Redis.cache.Exists(key).Result()
	if isExit == 0 {
		return false, err
	}
	return true, err
}

func (Redis *redisUtil) IsBlackList(token string) (bool, error) {
	return Redis.HHasKey("blacklist", token)
}

func (Redis *redisUtil) AddBlackList(token string) error {
	return Redis.HSet("blacklist", token, true)
}

func (Redis *redisUtil) GetTokenValidTimeByToken(token string) (result string, err error) {
	return Redis.HGet(token, "tokenValidTime")
}

func (Redis *redisUtil) GetUsernameByToken(token string) (result string, err error) {
	return Redis.HGet(token, "username")
}

func (Redis *redisUtil) GetIPByToken(token string) (result string, err error) {
	return Redis.HGet(token, "ip")
}
