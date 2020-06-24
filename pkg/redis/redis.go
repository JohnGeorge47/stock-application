package redis

import (
	"context"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/configmanager"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	*redis.Client
}

var Rdb RedisClient

func InitConnection() {
	config := configmanager.GetConfig()
	addr := fmt.Sprintf("%s:%s", config.RedisConfig.Host, config.RedisConfig.Port)
	Rdb.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       config.RedisConfig.Db,
		Password: "",
	})
	fmt.Println("here")
}

func (r RedisClient) Exists(ctx context.Context, key string) {
	val := r.Client.Exists(ctx, key)
	fmt.Println(val)
}

func (r RedisClient) PushArray(ctx context.Context, key string, val string) {
	res := r.Client.LPush(ctx, key, val)
	fmt.Println(res)
}
