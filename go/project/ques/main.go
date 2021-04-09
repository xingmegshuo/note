/***************************
@File        : main.go
@Time        : 2021/04/08 14:51:23
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 运行主程序
****************************/

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "ques/server"
	"time"
)

type dataValue struct {
	Data   []string
	Result []bool
	Err    string
}

var value dataValue

// 客户端调用接口
func Run(data []string) ([]bool, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			/* 不校验服务端证书，直接信任 */
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}

	/* 协议：https，非http */
	resp, err := Send(client, data)
	if err != nil {
		return []bool{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []bool{}, err
	}
	if err := json.Unmarshal([]byte(body), &value); err == nil {
		return value.Result, err
	} else {
		return []bool{}, err
	}
}

// 向服务端发送数据
func Send(client *http.Client, data []string) (*http.Response, error) {
	value = dataValue{
		Data: data,
	}
	jsonStr, e := json.Marshal(value)
	if e != nil {
		log.Println(e)
	}
	resp, err := client.Post("https://127.0.0.1:8000/", "json", bytes.NewBuffer(jsonStr))
	return resp, err
}

func main() {
	// 默认调用接口
	time.Sleep(time.Second * 1)
	r, e := Run([]string{"a", "b", "a"})
	log.Println(r, e)
}
