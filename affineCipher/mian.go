package main

import (
    "strings"
    "fmt"
    "os"
)

//加密方法
func affineEncypt(plainText string, a, b int) string {
    //转化为大写
    plainText = strings.ToUpper(plainText)
    cipherText := ""
    for _, m := range plainText {
        if m >= 'A' && m <= 'Z' {
            //	将字符转化为数字
            chNum := int(m - 'A')
            //仿射加密公式
            cipherNum := (a*chNum + b) % 26
            //加密后的字符拼接到密文
            cipherText += string(cipherNum + 'A')
        } else {
            cipherText += string(m)
        }
    }
    return cipherText
}

//解密方法
func affineDecrypt(cipherText string, a, b int) string {
    //转化为大写
    cipherText = strings.ToUpper(cipherText)
    plainText := ""

    //法一:	暴力枚举求逆元
    aInv := 0
    //for i := 0; i < 26; i++ {
    //    if (a*i)%26 == 1 {
    //        aInv = i
    //        break
    //    }
    //}

    //法二: 扩展欧几里得算法求逆元(匿名函数实现)
    q := 26
    var exGcd func(int, int) (int, int, int)
    exGcd = func(a, b int) (int, int, int) {
        if b == 0 {
            return a, 1, 0
        }
        gcd, x1, y1 := exGcd(b, a%b)
        x0 := y1
        y0 := x1 - (a/b)*y1
        return gcd, x0, y0
    }

    gcd, x, _ := exGcd(a, q)
    if gcd != 1 {
        fmt.Print("没有逆元")
        os.Exit(0)
    } else {
        aInv = (x%q + q) % q
    }

    for _, c := range cipherText {
        if c >= 'A' && c <= 'Z' {
            //转为数字
            chNum := int(c - 'A')
            //仿射解密公式
            plainNum := (aInv * (chNum - b + 26)) % 26 //应注意出现负数的情况
            //明文字符拼接到字符串
            plainText += string(plainNum + 'A')
        } else {
            plainText += string(c)
        }
    }
    return plainText
}
func main() {
    fmt.Print("输入明文字符串:")
    plainText := ""
    fmt.Scanln(&plainText)
    a := 0
    b := 0
    fmt.Print("输入a:")
    fmt.Scanln(&a)
    fmt.Print("输入B:")
    fmt.Scanln(&b)

    fmt.Println("plainText:" + plainText)
    //加密
    cipherText := affineEncypt(plainText, a, b)

    fmt.Println("cipherText:", cipherText)

    //解密
    plainText = affineDecrypt(cipherText, a, b)
    fmt.Println("plainText:" + plainText)
}
