package toolkit

import (
	"fmt"
	"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	//生成密钥对，保存到文件
	GenerateRSAKey(2048)
}

func TestConsumerTest(t *testing.T) {
	//加密
	data := []byte("hello world")
	encrypt, err := RSAEncrypt(data, "../keypair/public.pem")
	fmt.Println(string(encrypt), err)

	// 解密
	decrypt, err := RSADecrypt(encrypt)
	fmt.Println(string(decrypt), err)
}

// 普通的测试
func TestGenShortID(t *testing.T) {
	//测试加密数据
	shortID := MD5([]byte("123456"))
	if shortID == "" {
		t.Error("GenShortID failed")
	}
}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MD5([]byte("123456"))
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortId := MD5([]byte("123456"))
	if shortId == "" {
		b.Error("GenShortID failed")
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		MD5([]byte("123456"))
	}
}
