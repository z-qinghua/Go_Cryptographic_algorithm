// @program:     Go_Cryptographic_algorithm
// @file:        main.go.go
// @create:      2023-03-16 20:31
// @description:

package main

import (
    "fmt"
)

/*
   1. 设置明文和移步长度（密文）
   2. 将清晰的文本转换为小写，准备清晰的文本字节切片和密文切片
   3. 将每个明文字符根据位移长做好位移，存储在密文切片中
   4. 返回密文
*/
//凯撒密码加密

func caesarEn(strRaw string, step int32) string {
    //对位移长度进行判断
    //位移也就是密钥有效为25位
    stepMove := step % 26
    if stepMove <= 0 {
        return strRaw
    }
    //将字符串转化为明文字符切片
    srcSlice := []rune(strRaw) //[]rune底层是[]int32

    //定义一个密文切片
    dstSlice := srcSlice

    //对字符循环进行位移转化
    for i := 0; i < len(srcSlice); i++ {
        if srcSlice[i] >= 'a' && srcSlice[i] <= 'z' {
            dstSlice[i] = (srcSlice[i]+stepMove-97)%26 + 97
        }
        if srcSlice[i] >= 'A' && srcSlice[i] <= 'Z' {
            dstSlice[i] = (srcSlice[i]+stepMove-65)%26 + 65
        }
    }
    return string(dstSlice)
}

/*
   1. 设置密文和位移步骤
   2. 准备密文字符切片和明文字符切片
   3. 每个密文字符根据位移长做位移，并存储在明文切片中
   4. 返回明文
*/
//凯撒密码解密

func caesarDe(strCipher string, step int32) string {
    stepMove := step % 26
    if stepMove <= 0 {
        return strCipher
    }
    //将字符串转化为密文字符切片
    srcSlice := []rune(strCipher) //[]rune底层是[]int32

    //定义一个明文切片
    dstSlice := srcSlice

    //对字符循环进行位移转化
    for i := 0; i < len(srcSlice); i++ {
        if srcSlice[i] >= 'a' && srcSlice[i] <= 'z' {
            if srcSlice[i] >= 97+stepMove {
                dstSlice[i] = srcSlice[i] - stepMove
            } else {
                dstSlice[i] = srcSlice[i] + 26 - stepMove
            }
        }
        if srcSlice[i] >= 'A' && srcSlice[i] <= 'Z' {
            if srcSlice[i] >= 65+stepMove {
                dstSlice[i] = srcSlice[i] - stepMove
            } else {
                dstSlice[i] = srcSlice[i] + 26 - stepMove
            }
        }
        //	if srcSlice[i] >= 'a' && srcSlice[i] <= 'z' {
        //		flag := srcSlice[i] - 97 - stepMove
        //		if flag < 0 {
        //			dstSlice[i] = srcSlice[i] - stepMove + 26
        //		} else {
        //			dstSlice[i] = srcSlice[i] - stepMove
        //		}
        //	}
        //	if srcSlice[i] >= 'A' && srcSlice[i] <= 'Z' {
        //		flag := srcSlice[i] - 65 - stepMove
        //		if flag < 0 {
        //			dstSlice[i] = srcSlice[i] - stepMove + 26
        //		} else {
        //			dstSlice[i] = srcSlice[i] - stepMove
        //		}
        //	}
    }
    return string(dstSlice)
}

func main() {
    plainText := ""
    fmt.Scanln(&plainText)
    //fmt.Println(plainText)
    cipherText := caesarEn(plainText, 3)
    fmt.Println("cipherText:", cipherText)

    plainText = caesarDe(cipherText, 3)
    fmt.Println("caesar plainTex:", plainText)
}
