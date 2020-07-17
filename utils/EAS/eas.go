package EAS

import (
	"crypto/aes"
	"fmt"
	"crypto/cipher"
	"encoding/hex"
	"encoding/base64"
	"bytes"
	"os"
)

// 需要定义key和偏移量
var (
	key,_ = hex.DecodeString(os.Getenv("AES_KEY")) // "MZ7^5f$bjRP#mL$Q"
	iv, _  = hex.DecodeString(os.Getenv("AES_IV"))  // "*JyxSM^Svfzp7wd&"
)

func Encrypt(encodeByte []byte) (cipherText string, err error){

	ckey, err := aes.NewCipher(key)
	if nil != err {
		fmt.Println("钥匙创建错误:", err)
		return "", err
	}

	blockSize := ckey.BlockSize()

	fmt.Println("加密的字符串", string(encodeByte), "\n加密钥匙", key, "\n向量IV", string(iv))

	fmt.Println("加密前的字节：", encodeByte, "\n")

	encrypter := cipher.NewCBCEncrypter(ckey, iv)

	// PKCS7补码
	encodeByte = PKCS7Padding(encodeByte, blockSize)
	out := make([]byte, len(encodeByte))

	encrypter.CryptBlocks(out, encodeByte)
	fmt.Println("加密后字节：", out)


	// hex 兼容nodejs cropty-js包
	cipherText = hex.EncodeToString(out)
	return cipherText, nil
}

func Decrypt(encodeStr string) (origByte []byte, err error) {
	ckey, err := aes.NewCipher(key)
	if nil != err {
		fmt.Println("钥匙创建错误:", err)
		return origByte, err
	}

	base64Str,err := hex.DecodeString(encodeStr)
	if err != nil {
		return origByte, err
	}
	base64Out := base64.URLEncoding.EncodeToString(base64Str)


	//fmt.Println("\n开始解码")
	decrypter := cipher.NewCBCDecrypter(ckey, iv)

	base64In, err := base64.URLEncoding.DecodeString(base64Out)

	if err != nil {
		return origByte, err
	}

	in := make([]byte, len(base64In))

	decrypter.CryptBlocks(in, base64In)

	//fmt.Println("解密后的字节：", in)

	// 去除PKCS7补码
	in = UnPKCS7Padding(in)

	//fmt.Println("去PKCS7补码：", in)
	//fmt.Println("解密：", string(in))
	return in,nil
}

/**
 *	PKCS7补码
 */
func PKCS7Padding(data []byte, blockSize int) []byte {
	//blockSize := 16
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)

}

/**
 *	去除PKCS7的补码
 */
func UnPKCS7Padding(data []byte) []byte {
	length := len(data)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}