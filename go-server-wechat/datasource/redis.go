package datasource

import (
	"context"
	"fmt"
	"serverwechat/config"
	"serverwechat/logger"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	logger.Logger.Info("初始化redis start")
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.BASE_CONFIG.Redis.Host, config.BASE_CONFIG.Redis.Port),
		Password: config.BASE_CONFIG.Redis.Password, // no password set
		DB:       0,                                 // use default DB
		PoolSize: config.BASE_CONFIG.Redis.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, ok := RDB.Ping(ctx).Result()
	if ok != nil {
		logger.Logger.Error(fmt.Sprintf("初始化redis 失败 %v:", ok))
		return
	}
	logger.Logger.Info("初始化redis success")
}

func SetRedisByString(key string, value string, timeOut time.Duration) bool {
	err := RDB.Set(ctx, key, value, timeOut).Err()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("redis key %v 设置失败， %v", key, err))
		return false
	}
	return true
}

func GetRedisByString(key string) (bool, string) {
	count, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		logger.Logger.Info(fmt.Sprintf("redis key %v不存在", key))
		return false, "empty"
	} else if err != nil {
		logger.Logger.Error(fmt.Sprintf("redis key %v 获取失败 %v", key, err))
		return false, "error"
	} else {
		return true, count
	}
}

func DelRedisByString(key string) (bool, int64) {
	count, err := RDB.Del(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		logger.Logger.Info(fmt.Sprintf("redis key %v不存在", key))
		return false, 0
	} else if err != nil {
		logger.Logger.Error(fmt.Sprintf("redis key %v 获取失败 %v", key, err))
		return false, 0
	} else {
		return true, count
	}
}
