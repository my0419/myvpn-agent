package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"errors"
)

func EncryptAES(plaintext []byte, aeskey string) (string, error) {
	block, err := aes.NewCipher([]byte(aeskey))
	if err != nil {
		return "", err
	}
	plaintext, err = pkcs7Pad(plaintext, block.BlockSize())
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize)
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext, plaintext)
	return string(ciphertext), nil

}

func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("Invalid block size")
	}
	if b == nil || len(b) == 0 {
		return nil, errors.New("Invalid pkcs data")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}
