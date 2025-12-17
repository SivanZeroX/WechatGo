package wechatgo

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Signer 微信数据签名器
type Signer struct {
	data      []string
	delimiter string
}

// NewSigner 创建新的签名器
func NewSigner(delimiter string) *Signer {
	return &Signer{
		data:      make([]string, 0),
		delimiter: delimiter,
	}
}

// AddData 添加数据到签名器
func (s *Signer) AddData(args ...string) {
	s.data = append(s.data, args...)
}

// Signature 获取数据签名
func (s *Signer) Signature() string {
	sort.Strings(s.data)
	strToSign := strings.Join(s.data, s.delimiter)
	hash := sha1.Sum([]byte(strToSign))
	return hex.EncodeToString(hash[:])
}

// CheckSignature 检查微信回调签名
func CheckSignature(token, signature, timestamp, nonce string) error {
	signer := NewSigner("")
	signer.AddData(token, timestamp, nonce)
	if signer.Signature() != signature {
		return NewInvalidSignatureError()
	}
	return nil
}

// CheckWxaSignature 校验小程序前端传来的 rawData 签名
func CheckWxaSignature(sessionKey, rawData, clientSignature string) error {
	str2sign := rawData + sessionKey
	hash := sha1.Sum([]byte(str2sign))
	signature := hex.EncodeToString(hash[:])
	if signature != clientSignature {
		return NewInvalidSignatureError()
	}
	return nil
}

// RandomString 生成随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
