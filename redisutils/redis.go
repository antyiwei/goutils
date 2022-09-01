package redisutils

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedisClient(addr string, pass string, dbNum int) *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pass,  // no password set
		DB:           dbNum, // use default DB
		ReadTimeout:  time.Millisecond * time.Duration(500),
		WriteTimeout: time.Millisecond * time.Duration(500),
		IdleTimeout:  time.Second * time.Duration(60),
		PoolSize:     64,
		MinIdleConns: 16,
	})

}

func DelHashKey(client *redis.Client, hashKey string) {

	ctx := context.TODO()

	defer client.Del(ctx, hashKey)

	var cursor uint64
	for {

		if client == nil {
			break
		}

		var keys []string
		var err error

		keys, cursor, err = client.HScan(ctx, hashKey, cursor, "", 100).Result()
		if err != nil {
			log.Println("client hscan err:", err.Error())
			break
		}

		var delKeys []string
		for i, key := range keys {
			if i%2 == 0 {
				delKeys = append(delKeys, key)
			}
		}

		client.HDel(ctx, hashKey, delKeys...)
		if cursor == 0 {
			break
		}
	}
}
