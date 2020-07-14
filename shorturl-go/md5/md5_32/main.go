package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string{
	return GetMD5Encode(data)[8:24]
}

func main() {
	source:="hello"
	fmt.Println(GetMD5Encode(source),len(GetMD5Encode(source)))
	fmt.Println(Get16MD5Encode(source),len(Get16MD5Encode(source)))
}