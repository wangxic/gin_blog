package redis

import (
	"blog/conf"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var MyRedis *redis.Client
var ctx = context.Background()

func init() {
	var err error
	MyRedis = redis.NewClient(&redis.Options{
		Addr:     conf.GetConfig("redis.host"),     // 指定
		Password: conf.GetConfig("redis.password"), //密码
		DB:       0,                                // redis一共16个库，指定其中一个库即可
	})
	_, err = MyRedis.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Fail to open redis: %v", err)
	}

}

func Set(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second

	if err := MyRedis.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true

}

func Get(key string) string {
	result, err := MyRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func Del(key string) bool {
	_, err := MyRedis.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//更新过期时间
func Expire(key string, t int64) bool {
	// 延长过期时间
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
