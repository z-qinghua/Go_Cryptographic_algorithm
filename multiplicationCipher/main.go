// @program:     Go_Cryptographic_algorithm
// @file:        main.go.go
// @create:      2023-03-21 20:21
// @description:

package main

import (
	"fmt"
	"strings"
)

//加密

func MultiplicationEn(plainText string, k, q int) string {
	//将明文转换为大写
	plainText = strings.ToUpper(plainText)

	//定义空字符串存储密文
	cipherText := ""

	//    遍历明文每个字符
	for _, v := range plainText {
		//对字母进行加密
		if v >= 'A' && v <= 'Z' {
			l := int(v - 'A')
			f := (l * k) % q
			//将密文字母追加到密文字符串
			cipherText += string(f + 'A')
		} else {
			//    如果不是字母，原样追加到密文序列
			cipherText += string(v)
		}
	}
	return cipherText
}

//解密

func MultiplicationDe(cipherText string, invKey, q int) string {
	//将密文转换为大写
	cipherText = strings.ToUpper(cipherText)

	//定义空字符串存储明文
	plainText := ""

	//    遍历密文每个字符
	for _, v := range cipherText {
		//对字母进行解密
		if v >= 'A' && v <= 'Z' {
			l := int(v - 'A')
			f := (l * invKey) % q
			//将密文字母追加到明文字符串
			plainText += string(f + 'A')
		} else {
			//    如果不是字母，原样追加到明文序列
			plainText += string(v)
		}
	}
	return plainText
}

// 扩展欧几里得算法求逆元
func extGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extGcd(b, a%b)
	x0 := y1
	y0 := x1 - (a/b)*y1
	return gcd, x0, y0
}

// 求a在模p下的逆元，如果不存在返回-1
func modInverse(a, q int) int {
	gcd, x, _ := extGcd(a, q)
	if gcd != 1 {
		return -1
	} else {
		return (x%q + q) % q
	}
}

//测试

func main() {
	fmt.Print("输入明文：")

	plainText := ""
	fmt.Scanln(&plainText)

	fmt.Print("输入密钥：")
	k := 0

	fmt.Scanln(&k)

	fmt.Print("输入模数：")
	q := 0
	fmt.Scanln(&q)

	fmt.Println("明文："+plainText, "密钥：", +k, "模数：", +q)

	cipherText := MultiplicationEn(plainText, k, q)
	fmt.Println("输出密文：" + cipherText)

	invKey := modInverse(k, q)

	plainTextDE := MultiplicationDe(cipherText, invKey, q)
	fmt.Println("输出明文：" + plainTextDE)

}
