package crypto

import "bytes"

const blockSize = 32

// PKCS7Encode PKCS7 填充
func PKCS7Encode(text []byte) []byte {
	length := len(text)
	paddingCount := blockSize - length%blockSize
	if paddingCount == 0 {
		paddingCount = blockSize
	}
	padding := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(text, padding...)
}

// PKCS7Decode PKCS7 去填充
func PKCS7Decode(decrypted []byte) []byte {
	if len(decrypted) == 0 {
		return decrypted
	}
	padding := int(decrypted[len(decrypted)-1])
	if padding < 1 || padding > 32 {
		padding = 0
	}
	if padding > len(decrypted) {
		return decrypted
	}
	return decrypted[:len(decrypted)-padding]
}
