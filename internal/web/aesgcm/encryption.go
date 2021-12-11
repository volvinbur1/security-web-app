package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const keySize = 32

func EncryptUserData(data, keyStr string) (string, error) {
	key, err := hex.DecodeString(keyStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	enc := aesGcm.Seal(nonce, nonce, []byte(data), nil)
	return hex.EncodeToString(enc), nil
}

func DecryptData(encStr, keyStr string) (string, error) {
	key, err := hex.DecodeString(keyStr)
	if err != nil {
		return "", err
	}
	enc, err := hex.DecodeString(encStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, cipherText := enc[:aesGcm.NonceSize()], enc[aesGcm.NonceSize():]
	plaintext, err := aesGcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func GenerateUniqueKey() (string, error) {
	key := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}
