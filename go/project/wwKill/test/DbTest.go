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
	ctrl := Mydb.NewBackCtrl()
	// ctrlUser := Mydb.NewUserCtrl()
	// user := Mydb.User{
	// 	OpenID:    "abcd",
	// 	NickName:  "abcd",
	// 	AvatarURL: "ddddd",
	// 	Orther:    "ahhh",
	// 	Money:     300,
	// }
	// // 插入一个
	// ctrlUser.Insert(user)
	// update := Mydb.User{Id: 1}
	// ctrl.GetUser(update)
	// back := Mydb.Backpack{
	// 	Name: "可爱上衣",
	// 	User: 1,
	// }
	// ctrl.Insert(back)
	back := Mydb.Backpack{
		Id:       1,
		StilTime: "2012-12-12 08：00：00",
	}
	ctrl.Update(back)

}
