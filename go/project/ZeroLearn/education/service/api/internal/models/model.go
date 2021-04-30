/***************************
@File        : model.go
@Time        : 2021/04/16 15:35:59
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 数据库表
****************************/

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string
	Gender   string
	Avatar   string
	City     string
	Phone    string
	Mail     string
	Iden     string
	OpenId   string
}

