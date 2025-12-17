package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
)

// Cipher AES 加密器接口
type Cipher interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

// CBCCipher AES CBC 模式加密器
type CBCCipher struct {
	block cipher.Block
	iv    []byte
}

// NewCBCCipher 创建 CBC 模式加密器
func NewCBCCipher(key, iv []byte) (*CBCCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if iv == nil {
		iv = key[:16]
	}
	return &CBCCipher{block: block, iv: iv}, nil
}

// Encrypt 加密
func (c *CBCCipher) Encrypt(plaintext []byte) ([]byte, error) {
	mode := cipher.NewCBCEncrypter(c.block, c.iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// Decrypt 解密
func (c *CBCCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	mode := cipher.NewCBCDecrypter(c.block, c.iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext, nil
}

// ECBCipher AES ECB 模式加密器
type ECBCipher struct {
	block cipher.Block
}

// NewECBCipher 创建 ECB 模式加密器
func NewECBCipher(key []byte) (*ECBCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &ECBCipher{block: block}, nil
}

// Encrypt 加密
func (c *ECBCipher) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext := make([]byte, len(plaintext))
	blockSize := c.block.BlockSize()
	for i := 0; i < len(plaintext); i += blockSize {
		c.block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}
	return ciphertext, nil
}

// Decrypt 解密
func (c *ECBCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext := make([]byte, len(ciphertext))
	blockSize := c.block.BlockSize()
	for i := 0; i < len(ciphertext); i += blockSize {
		c.block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}
	return plaintext, nil
}

// PrpCrypto 微信消息加密器
type PrpCrypto struct {
	cipher Cipher
}

// NewPrpCrypto 创建消息加密器
func NewPrpCrypto(key []byte) (*PrpCrypto, error) {
	c, err := NewCBCCipher(key, nil)
	if err != nil {
		return nil, err
	}
	return &PrpCrypto{cipher: c}, nil
}

// Encrypt 加密消息
func (p *PrpCrypto) Encrypt(text, id string) (string, error) {
	randomStr := RandomString(16)
	textBytes := []byte(text)
	idBytes := []byte(id)

	// 构建消息: random(16) + length(4) + text + id
	buf := make([]byte, 0, 16+4+len(textBytes)+len(idBytes))
	buf = append(buf, []byte(randomStr)...)

	// 添加长度（网络字节序）
	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(textBytes)))
	buf = append(buf, lengthBuf...)

	buf = append(buf, textBytes...)
	buf = append(buf, idBytes...)

	// PKCS7 填充
	buf = PKCS7Encode(buf)

	// 加密
	ciphertext, err := p.cipher.Encrypt(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密消息
func (p *PrpCrypto) Decrypt(text, id string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	plaintext, err := p.cipher.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}

	// 去除填充
	plaintext = PKCS7Decode(plaintext)

	// 解析: random(16) + length(4) + text + id
	if len(plaintext) < 20 {
		return "", errors.New("invalid plaintext length")
	}

	content := plaintext[16:]
	xmlLength := binary.BigEndian.Uint32(content[:4])
	if len(content) < int(4+xmlLength) {
		return "", errors.New("invalid content length")
	}

	xmlContent := string(content[4 : 4+xmlLength])
	fromID := string(content[4+xmlLength:])

	if fromID != id {
		return "", errors.New("invalid app id")
	}

	return xmlContent, nil
}

// RefundCrypto 退款加密器
type RefundCrypto struct {
	cipher Cipher
}

// NewRefundCrypto 创建退款加密器
func NewRefundCrypto(key []byte) (*RefundCrypto, error) {
	c, err := NewECBCipher(key)
	if err != nil {
		return nil, err
	}
	return &RefundCrypto{cipher: c}, nil
}

// Encrypt 加密
func (r *RefundCrypto) Encrypt(text string) (string, error) {
	textBytes := []byte(text)
	textBytes = PKCS7Encode(textBytes)

	ciphertext, err := r.cipher.Encrypt(textBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密
func (r *RefundCrypto) Decrypt(text string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	plaintext, err := r.cipher.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}

	plaintext = PKCS7Decode(plaintext)
	return string(plaintext), nil
}
