package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wechatpy/wechatgo/session"
)

func TestNewBaseClient(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	assert.NotNil(t, client)
	assert.Equal(t, "test_appid", client.AppID)
	assert.Equal(t, "https://api.example.com", client.apiBaseURL)
	assert.True(t, client.autoRetry)
	assert.NotNil(t, client.httpClient)
}

func TestNewBaseClient_NilStorage(t *testing.T) {
	client := NewBaseClient("test_appid", nil, "https://api.example.com")

	assert.NotNil(t, client)
	assert.NotNil(t, client.session)
}

func TestBaseClient_WithLogger(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	// WithLogger应该返回新的client实例
	result := client.WithLogger(nil)
	assert.NotNil(t, result)
	assert.Equal(t, client, result)
}

func TestAccessTokenKey(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	key := client.accessTokenKey()
	assert.Contains(t, key, "test_appid")
	assert.Contains(t, key, "access_token")
}

func TestExpiresAtKey(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	key := client.expiresAtKey()
	assert.Contains(t, key, "test_appid")
	assert.Contains(t, key, "expires_at")
}

func TestGetAccessToken_NotFound(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	token, err := client.GetAccessToken()
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestSetAccessToken(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	err := client.SetAccessToken("test_token", 7200)
	assert.NoError(t, err)

	// 验证token已设置
	storedToken, err := client.GetAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, "test_token", storedToken)
}

func TestSetAccessToken_EmptyValue(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	// 空值应该被忽略
	err := client.SetAccessToken("", 7200)
	assert.NoError(t, err)

	// 验证没有设置token
	storedToken, err := client.GetAccessToken()
	assert.Error(t, err)
	assert.Empty(t, storedToken)
}

func TestMarshalJSON_Cache(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := TestStruct{Name: "test", Age: 25}

	// 第一次序列化
	result1, err := client.marshalJSON(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, result1)

	// 第二次序列化应该从缓存获取
	result2, err := client.marshalJSON(data)
	assert.NoError(t, err)
	assert.Equal(t, result1, result2)
}

func TestMarshalJSON_SimpleString(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	// 简单字符串不使用缓存
	result, err := client.marshalJSON("simple_string")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestMarshalJSON_Error(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	// 测试无效JSON
	type InvalidStruct struct {
		Channel chan int `json:"channel"`
	}

	data := InvalidStruct{Channel: make(chan int)}

	_, err := client.marshalJSON(data)
	assert.Error(t, err)
}

func TestMarshalJSON_CacheSizeLimit(t *testing.T) {
	storage := session.NewMemoryStorage()
	client := NewBaseClient("test_appid", storage, "https://api.example.com")

	// 添加超过缓存限制的数据
	for i := 0; i < 105; i++ {
		data := map[string]int{"key": i}
		_, err := client.marshalJSON(data)
		assert.NoError(t, err)
	}

	// 验证缓存大小不超过限制
	assert.LessOrEqual(t, len(client.jsonMarshalCache), 100)
}
