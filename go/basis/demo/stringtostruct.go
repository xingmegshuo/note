/***************************
@File        : stringtostruct.go
@Time        : 2021/05/10 17:45:48
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : string to struct
****************************/

package main

import (
	"encoding/json"
	"log"
)

type wx struct {
	Openid     string
	NickName   string
	Sex        string
	Headimgurl string
	City       string
}

func main() {
	text := ` {"openid":"o5NsU6ib3FIqiKkZs9vnJQgwjylY","nickname":"small ant","sex":1,"language":"zh_CN","city":"","province":"","country":"中国","headimgurl":"https:\/\/thirdwx.qlogo.cn\/mmopen\/vi_32\/DYAIOgq83eoT6p5zNH7xdc6libngH4wvfUlDESiaEBhRDL3FMt3eNRfju8lJ6GPldRK0nvQEeVCr1icXzxkFeRR9A\/132","privilege":[]}`
	log.Println(text)
	var wxdata wx
	json.Unmarshal([]byte(text), &wxdata)
	log.Println(wxdata)
}
