/***************************
@File        : parse.go
@Time        : 2020/12/21 15:56:57
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 解析数据根据不同key来进行区分
****************************/

package Handler

import (
	"encoding/json"
	"log"
)

// 数据格式
type Data struct {
	Name   string
	Values []string
}

// 解析key
func ParseData(con string) {
	var data Data
	oldData := []byte(con)
	err := json.Unmarshal(oldData, &data)
	if err != nil {
		log.Println(err)
	}
	log.Println(data.Name)
}
