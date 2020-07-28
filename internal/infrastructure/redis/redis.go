package redis

import (
	"bm-novel/internal/config"
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	config.LoadConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.IPAddress,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(errors.WithMessage(err, "failed to connect to redis"))
	}
}

// Cacher 缓存
type Cacher struct {
}

// GetChcher 获取缓存
func GetChcher() *Cacher {
	return &Cacher{}
}

// Get 获取单值
func (r Cacher) Get(key string) ([]byte, error) {
	return rdb.Get(ctx, key).Bytes()
}

// Put 写入单值
func (r Cacher) Put(key string, data []byte, expiration time.Duration) error {
	return rdb.Set(ctx, key, data, expiration).Err()
}

// HGet 获取redis
func (r Cacher) HGet(key string, field string) ([]byte, error) {
	return rdb.HGet(ctx, key, field).Bytes()
}

// HPut 写入redis
func (r Cacher) HPut(key string, field string, data []byte, expiration time.Duration) error {
	err := rdb.HSetNX(ctx, key, field, data).Err()

	if err == nil {
		err = rdb.Expire(ctx, key, expiration).Err()
	}

	return err
}

// HMPut 批量写入redis
func (r Cacher) HMPut(key string, expiration time.Duration, values ...interface{}) error {
	err := rdb.HMSet(ctx, key, values...).Err()

	if err == nil {
		err = rdb.Expire(ctx, key, expiration).Err()
	}

	return err
}

// Delete 删除
func (r Cacher) Delete(key string) error {
	return rdb.Del(ctx, key).Err()
}

// Exists 是否存在
func (r Cacher) Exists(keys ...string) (bool, error) {
	result, err := rdb.Exists(ctx, keys...).Result()

	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, nil
	}

	return true, nil
}

// HExists 是否存在
func (r Cacher) HExists(key, field string) (bool, error) {
	return rdb.HExists(ctx, key, field).Result()
}
