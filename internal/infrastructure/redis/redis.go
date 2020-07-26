package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}

type Cacher struct {
}

func GetChcher() *Cacher {
	return &Cacher{}
}

func (r Cacher) Get(key string) ([]byte, error) {
	return rdb.Get(ctx, key).Bytes()
}

func (r Cacher) Put(key string, data []byte, expiration time.Duration) error {
	return rdb.Set(ctx, key, data, expiration).Err()
}

func (r Cacher) HGet(key string, field string) ([]byte, error) {
	return rdb.HGet(ctx, key, field).Bytes()
}

func (r Cacher) HPut(key string, field string, data []byte, expiration time.Duration) error {
	err := rdb.HSetNX(ctx, key, field, data).Err()

	if err == nil {
		err = rdb.Expire(ctx, key, expiration).Err()
	}

	return err
}

func (r Cacher) Delete(key string) error {
	return rdb.Del(ctx, key).Err()
}

func (r Cacher) Exists(keys ...string) error {
	return rdb.Exists(ctx, keys...).Err()
}
