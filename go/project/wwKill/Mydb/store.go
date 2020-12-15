/********************
@File        : store
@Time        : 2020-12-15 17:10:35
@Author      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 狼人杀个人背包数据库
**********************/

package Mydb

import "log"

type Backpack struct {
	Id       int64
	Name     string `xorm:"varchar(255)"`
	Property int    `xorm:"integer"`
	Num      int    `xorm:"integer"`
	StilTime string `xorm:"text"`
	user     int    `xorm:"foreign key(user) references user(userid)"`
}

// 插入单个物品
func (b Backpack) Insert(a ...interface{}) bool {
	_, err := orm.InsertOne(a[0])
	if err != nil {
		log.Panic(err)
	}
	return true
}
