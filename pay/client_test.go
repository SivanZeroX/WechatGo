package pay

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockHTTPClient mock HTTP客户端实现api.HTTPClient接口
type mockHTTPClient struct{}

func (m *mockHTTPClient) Post(url string, data []byte, headers map[string]string) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       http.NoBody,
	}
	return resp, nil
}

func (m *mockHTTPClient) Get(url string) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       http.NoBody,
	}
	return resp, nil
}

func TestNewClient(t *testing.T) {
	httpClient := &mockHTTPClient{}

	client := NewClient("appid", "api_key", "mch_id", "cert_path", "key_path", httpClient)

	assert.NotNil(t, client)
	assert.Equal(t, "appid", client.AppID)
	assert.Equal(t, "api_key", client.APIKey)
	assert.Equal(t, "mch_id", client.MchID)
	assert.Equal(t, "cert_path", client.CertPath)
	assert.Equal(t, "key_path", client.KeyPath)
	assert.NotNil(t, client.httpClient)
}

func TestNewClient_NilHTTPClient(t *testing.T) {
	client := NewClient("appid", "api_key", "mch_id", "", "", nil)

	assert.NotNil(t, client)
	assert.Nil(t, client.httpClient)
}

func TestClient_GetAPIKey(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	apiKey := client.GetAPIKey()
	assert.Equal(t, "api_key", apiKey)
}

func TestClient_GetMchID(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	mchID := client.GetMchID()
	assert.Equal(t, "mch_id", mchID)
}

func TestClient_GetAppID(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	appID := client.GetAppID()
	assert.Equal(t, "appid", appID)
}

func TestClient_GetHTTPClient(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	retrieved := client.GetHTTPClient()
	assert.Equal(t, httpClient, retrieved)
}

func TestClient_Get(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	resp, err := client.Get("http://example.com")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_Post(t *testing.T) {
	httpClient := &mockHTTPClient{}
	client := NewClient("appid", "api_key", "mch_id", "", "", httpClient)

	data := []byte("test data")
	headers := map[string]string{"Content-Type": "application/json"}

	resp, err := client.Post("http://example.com", data, headers)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
