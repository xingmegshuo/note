/********************
@File        : buddy
@Time        : 2020-12-15 17:21:05
@Author      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 狼人杀好友关系
**********************/

package Mydb

type Buddy struct{
	Id int64
	User int `xorm:"foregin key(user) references user(userid)"`
	// Own_user int `xorm:"foregin key(user) references user(userid)"`
	Agree int  `xorm:"integer"`
	buddys string `xorm:"text"`
}