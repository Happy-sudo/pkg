package toolkit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type AesCrypt struct {
	Key []byte
	Iv  []byte
}

func (a *AesCrypt) Encrypt(data []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher(a.Key)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	content := pKCS5Padding(data, aesBlockEncrypt.BlockSize())
	cipherBytes := make([]byte, len(content))
	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, a.Iv)
	aesEncrypt.CryptBlocks(cipherBytes, content)
	return cipherBytes, nil
}

func (a *AesCrypt) Decrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err = aes.NewCipher(a.Key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, a.Iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return pKCS5Trimming(decrypted), nil
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
