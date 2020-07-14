package main

import (
	"fmt"
	"encoding/hex"
	"bytes"
)

func main() {
	// 注意"Hello"与"encodedStr"不相等，encodedStr是用字符串来表示16进制
	src := []byte("Hello")
	//把[]byte表示成16进制（用String的形式）
	//[]byte -> String
	encodedStr := hex.EncodeToString(src)
	// [72 101 108 108 111]
	fmt.Println(src)
	// 48656c6c6f -> 48(4*16+8=72) 65(6*16+5=101) 6c 6c 6f
	fmt.Println(encodedStr)

	//String -> []byte
	test, _ := hex.DecodeString(encodedStr)
	fmt.Println(bytes.Compare(test, src)) // 0
}