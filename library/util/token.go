package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func CkToken(r *http.Request) bool {
	//从http头部获得token
	header := r.Header
	var acc []string = header["Token"]
	token := ""
	for _, v := range acc {
		token = v
	}

	nowTime := time.Now()
	nowTimeStr := nowTime.Format("2006-01-02")
	md5str := MD5(nowTimeStr)
	fmt.Println("token", md5str)
	//比较token
	if token != md5str {
		errors.New("Token 校验错误")
		return false
	}
	return true
}

func CkToken2(r *http.Request) bool {
	//1 从header头里获得token
	header := r.Header
	acc := header["Token"]
	token := acc[0]

	//2 规定生成token的规则 对当天日期进行md5加密
	nowTime := time.Now()
	nowTimeStr := nowTime.Format("2006-01-02")
	tokenStr := MD5(nowTimeStr)
	if token != tokenStr {
		return false
	}
	return true
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
