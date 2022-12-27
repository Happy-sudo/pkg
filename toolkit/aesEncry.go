package toolkit

import (
	"encoding/base64"
	"fmt"
)

var (
	key = "M5MU^OM7BUWI&BQF"
	IV  = "S4^AX&PFRFVJL73Z"
)

func AesCBCEncrypt(intput string) (string, error) {

	var aesCrypt = AesCrypt{
		Key: []byte(key),
		Iv:  []byte(IV),
	}

	result, err := aesCrypt.Encrypt([]byte(intput))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}

func AesCBCDecrypt(intput string) (string, error) {
	var aesCrypt = AesCrypt{
		Key: []byte(key),
		Iv:  []byte(IV),
	}

	a, err := base64.StdEncoding.DecodeString(intput)
	if err != nil {
		return "", err
	}
	plainText, err := aesCrypt.Decrypt(a)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(plainText), nil
}
