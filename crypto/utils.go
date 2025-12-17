package crypto

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	randMu     sync.Mutex
)

// RandomString 生成随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randMu.Lock()
	defer randMu.Unlock()

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
