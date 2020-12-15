/***************************
@File        : db_test.go
@Time        : 2020/12/10 15:24:54
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试数据库
****************************/

package main

import (
	"wwKill/Mydb"
)

func main() {
	ctrl := Mydb.NewUserCtrl()
	user := Mydb.User{
		OpenID:    "abcd",
		NickName:  "abcd",
		AvatarURL: "ddddd",
		Orther:    "ahhh",
		Money:     300,
	}
	// 插入一个
	ctrl.Insert(user)
}
