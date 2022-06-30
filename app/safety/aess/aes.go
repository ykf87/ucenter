package aess

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"log"
)

var AESKEY []byte

func init() {
	AESKEY = []byte{106, 33, 61, 105, 49, 64, 126, 119, 61, 51, 70, 67, 63, 97, 57, 89}
}

func EcbDecrypt(str string, key []byte) string {
	if key == nil {
		key = AESKEY
	}
	block, _ := aes.NewCipher(key)
	data, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		log.Println(err, str)
		return ""
	}
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	rrs := PKCS5UnPadding(decrypted)
	if rrs == nil {
		return ""
	}
	return string(rrs)
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

	return base64.URLEncoding.EncodeToString(decrypted)
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
	mxx := length - unpadding
	if mxx < 1 {
		return nil
	}
	return origData[:(length - unpadding)]
}
