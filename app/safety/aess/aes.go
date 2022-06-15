package aess

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

var AESKEY []byte

func init() {
	AESKEY = []byte{117, 71, 55, 104, 70, 36, 103, 75, 33, 90, 89, 109, 112, 52, 111, 107}
}

func EcbDecrypt(str string, key []byte) string {
	if key == nil {
		key = AESKEY
	}
	block, _ := aes.NewCipher(key)
	data, _ := base64.StdEncoding.DecodeString(str)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return string(PKCS5UnPadding(decrypted))
}

func EcbEncrypt(str string, key []byte) string {
	if key == nil {
		key = AESKEY
	}
	block, _ := aes.NewCipher(key)
	data := []byte(str)
	data = PKCS5Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return base64.StdEncoding.EncodeToString(decrypted)
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
