package rsa

import (
	"encoding/base64"
	"fmt"
	"github.com/agclqq/goencryption"
	"io/ioutil"
)

// PrivateKeyDecrypt  Rsa私钥解密
func PrivateKeyDecrypt(cipherData string) (error, []byte) {

	//1、通过私钥文件获取私钥信息
	keyInfo, err := ioutil.ReadFile("./private.pem")

	if err != nil {
		fmt.Println("ReadFile err:", err)
		return err, nil
	}
	b, err := base64.StdEncoding.DecodeString(cipherData)
	text, err := goencryption.PrvKeyDecrypt(keyInfo, b)

	if err != nil {
		fmt.Println("PrvKeyDecrypt err:", err)
		return err, nil
	}

	return nil, text

}
