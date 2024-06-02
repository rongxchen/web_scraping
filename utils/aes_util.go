package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

var YoudaoAES = map[string]string{
	"secretKey": "fsdsogkndfokasodnaso",
	"aesKey":    "ydsecret://query/key/B*RGygVywfNBwpmBaZg*WT7SIOUP2T0C9WHMZN39j^DAdaZhAnxvGcCY6VYFwnHl",
	"aesIv":     "ydsecret://query/iv/C@lZe2YzHtZ2CYgaXKSVfsb7Y4QWHjITPPZ0nQp87fBeJ!Iv6v^6fvi2WN@bYpJ4",
}

func GenerateRandomKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}

func AESEncrypt(plaintext []byte, secretKey, aesKey, aesIv string) string {
	key := []byte(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	_, err = base64.StdEncoding.DecodeString(aesKey)
	if err != nil {
		panic(err)
	}

	decodedIv, err := base64.StdEncoding.DecodeString(aesIv)
	if err != nil {
		panic(err)
	}

	encrypter := cipher.NewCBCEncrypter(block, decodedIv)
	ciphertext := make([]byte, len(plaintext))
	encrypter.CryptBlocks(ciphertext, plaintext)

	encrypted := append(decodedIv, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded
}

func AESDecrypt(encrypted string, secretKey, aesKey, aesIv string) string {
	key := []byte(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	decoded, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		panic(err)
	}

	_, err = base64.StdEncoding.DecodeString(aesKey)
	if err != nil {
		panic(err)
	}

	_, err = base64.StdEncoding.DecodeString(aesIv)
	if err != nil {
		panic(err)
	}

	iv := decoded[:aes.BlockSize]
	ciphertext := decoded[aes.BlockSize:]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)

	return string(decrypted)
}
