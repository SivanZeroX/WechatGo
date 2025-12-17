package session

import "time"

// Storage 会话存储接口
type Storage interface {
	Get(key string) (string, error)
	Set(key, value string, ttl time.Duration) error
	Delete(key string) error
}
