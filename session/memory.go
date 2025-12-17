package session

import (
	"context"
	"sync"
	"time"
)

// MemoryStorageEntry 内存存储条目（包含过期时间）
type MemoryStorageEntry struct {
	Value     string
	ExpiresAt int64
}

// MemoryStorage 内存存储
type MemoryStorage struct {
	mu     sync.RWMutex
	data   map[string]*MemoryStorageEntry
	ctx    context.Context
	cancel context.CancelFunc
}

// NewMemoryStorage 创建内存存储
func NewMemoryStorage() *MemoryStorage {
	ctx, cancel := context.WithCancel(context.Background())
	storage := &MemoryStorage{
		data:   make(map[string]*MemoryStorageEntry),
		ctx:    ctx,
		cancel: cancel,
	}

	// 启动清理协程，定期清理过期条目
	go storage.cleanup()

	return storage
}

// Get 获取值
func (m *MemoryStorage) Get(key string) (string, error) {
	// 性能优化：使用读锁提升并发性能
	m.mu.RLock()
	entry, ok := m.data[key]
	if !ok {
		m.mu.RUnlock()
		return "", nil
	}

	// 检查是否过期（使用毫秒级精度）
	now := time.Now().UnixNano() / 1e6
	if entry.ExpiresAt > 0 && now > entry.ExpiresAt {
		m.mu.RUnlock()
		// 过期条目需要写锁删除，但这里先返回空值避免阻塞读操作
		// 实际删除由cleanup协程处理
		return "", nil
	}

	value := entry.Value
	m.mu.RUnlock()
	return value, nil
}

// Set 设置值
func (m *MemoryStorage) Set(key, value string, ttl time.Duration) error {
	if value == "" {
		return nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	expiresAt := int64(0)
	if ttl > 0 {
		// 使用毫秒级精度
		expiresAt = time.Now().Add(ttl).UnixNano() / 1e6
	}

	m.data[key] = &MemoryStorageEntry{
		Value:     value,
		ExpiresAt: expiresAt,
	}

	return nil
}

// Delete 删除值
func (m *MemoryStorage) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
	return nil
}

// cleanup 定期清理过期条目
func (m *MemoryStorage) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.cleanupExpired()
		case <-m.ctx.Done():
			return
		}
	}
}

// cleanupExpired 清理过期条目
func (m *MemoryStorage) cleanupExpired() {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 使用毫秒级精度
	for key, entry := range m.data {
		if entry.ExpiresAt > 0 && now > entry.ExpiresAt {
			delete(m.data, key)
		}
	}
}

// ForceCleanup 强制清理（用于测试或性能紧急情况）
func (m *MemoryStorage) ForceCleanup() {
	m.cleanupExpired()
}

// Close 关闭存储，停止清理 goroutine
func (m *MemoryStorage) Close() error {
	m.cancel()
	return nil
}
