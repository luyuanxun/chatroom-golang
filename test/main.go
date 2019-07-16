package main

import (
	"fmt"
	_ "helper"
	"regexp"
	"strings"
)

func main() {
	//req := &helper.Request{}
	//url := "https://login.wx.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1476606163580"
	//ret, err := req.Do("GET", url, nil)
	//if err != nil{
	//	fmt.Println("请求失败!", err)
	//	return
	//}
	//fmt.Println("请求成功：", string(ret),string(ret1))

	//reg := regexp.MustCompile(`/window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+?)"/`)
	//ret1 := reg.Find(ret)
	//reg := regexp.MustCompile(`[\d]{3}`)
	reg := regexp.MustCompile(`[\d]{3}`)
	text := "window.QRLogin.code = 200; window.QRLogin.uuid = \"wb6SUWrEfQ==\""
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	//ss := strings.Split(text, "\"")
	//fmt.Println("uuid:", ss[0],ss[1], ss[2])
	ss := strings.Index(text, "=")
	fmt.Println("ss:",ss)

}
