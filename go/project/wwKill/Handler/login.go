/***************************
@File        : login.go
@Time        : 2020/12/21 14:33:25
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 解析登录
****************************/

package Handler

import (
	"log"
)

// 用户登录函数处理
func Login(mes []byte) {
	// var user Mydb.User
	// err := json.Unmarshal(mes, &user)
	// if err != nil {
	// 	log.Println("err:", err)
	// }
	log.Println(mes)
}
