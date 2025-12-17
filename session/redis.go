package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisStorage Redis 存储
type RedisStorage struct {
	client *redis.Client
	prefix string
	ctx    context.Context
}

// NewRedisStorage 创建 Redis 存储
func NewRedisStorage(client *redis.Client, prefix string) *RedisStorage {
	if client == nil {
		panic("redis client cannot be nil")
	}
	if prefix == "" {
		prefix = "wechatgo"
	}
	return &RedisStorage{
		client: client,
		prefix: prefix,
		ctx:    context.Background(),
	}
}

// keyName 生成完整的 key 名称
func (r *RedisStorage) keyName(key string) string {
	return fmt.Sprintf("%s:%s", r.prefix, key)
}

// Get 获取值
func (r *RedisStorage) Get(key string) (string, error) {
	fullKey := r.keyName(key)
	value, err := r.client.Get(r.ctx, fullKey).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return value, nil
	}
	return result, nil
}

// Set 设置值
func (r *RedisStorage) Set(key, value string, ttl time.Duration) error {
	if value == "" {
		return nil
	}
	fullKey := r.keyName(key)
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, fullKey, data, ttl).Err()
}

// Delete 删除值
func (r *RedisStorage) Delete(key string) error {
	fullKey := r.keyName(key)
	return r.client.Del(r.ctx, fullKey).Err()
}
