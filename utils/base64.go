package utils

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：base64加密解密
 * 参考文档：https://learnku.com/docs/build-web-application-with-golang/096-encryption-and-decryption-of-data/3214
 */
import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

// Encrypt 加密
func Encrypt(str string) string {
	debyte := base64Encode([]byte(str))
	return string(debyte)
}

// Decrypt 解密
func Decrypt(str string) string {
	enbyte, err := base64Decode([]byte(str))
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(enbyte)
}
