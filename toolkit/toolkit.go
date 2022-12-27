package toolkit

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"github.com/pkg/errors"
	"os"
)

// MD5 加密 加密一次
func MD5(str []byte) string {
	h := md5.New()
	h.Write(str) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	toString := hex.EncodeToString(cipherStr)
	return toString
}

// GenerateRSAKey 生成RSA私钥和公钥，保存到文件中 bits 证书大小
//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
//Reader是一个全局、共享的密码用强随机数生成器
func GenerateRSAKey(bits int) {

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	privateFile, err := os.Create("../keypair/private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()

	privateBlock := pem.Block{Type: "PRIVATE KEY", Bytes: X509PrivateKey}

	pem.Encode(privateFile, &privateBlock)

	publicKey := privateKey.PublicKey

	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	publicFile, err := os.Create("../keypair/public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()

	publicBlock := pem.Block{Type: "PUBLIC KEY", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

// RSAEncrypt RSA加密
func RSAEncrypt(plainText []byte, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.Wrap(errors.New("pem.Decode Encoding Error"), "pem.Decode Encoding Error")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "pem.Decode Encoding Error")
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//data, err := base64.StdEncoding.DecodeString(plainText)
	//if err != nil {
	//	return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString Encoding Error")
	//}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText, nil
}

//RSADecrypt 解密 cipherText string, path string
func RSADecrypt(cipherText []byte) ([]byte, error) {
	file, err := os.Open("../keypair/private.pem")

	if err != nil {
		return nil, errors.Wrap(err, "RSADecrypt fail to open the file")
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, errors.Wrap(err, "x509 Failed to encode by key")
	}

	//data, err := base64.StdEncoding.DecodeString(cipherText)
	//if err != nil {
	//	return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString Encoding Error")
	//}

	//对密文进行解密 注意这个函数是否返回错误会泄露秘密信息。
	//如果攻击者可以使该函数重复运行并且了解每个实例是否返回错误然后他们可以解密和伪造签名
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)

	if err != nil {
		return nil, errors.Wrap(err, "DecryptPKCS1v15 Failed to encode by key")
	}

	//返回明文
	return plainText, nil
}
