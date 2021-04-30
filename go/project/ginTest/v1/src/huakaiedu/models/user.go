/***************************
@File        : user.go
@Time        : 2021/04/19 13:58:37
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 用户表
****************************/

package models

import "gorm.io/gorm"

// 存放用户数据的表
type User struct {
	gorm.Model
	NickName string //昵称
	Gender   string //性别
	Avatar   string //头像地址
	City     string //城市
	Phone    string //手机
	Mail     string //邮箱
	Iden     string `gorm:"default:普通用户"` //身份
	OpenId   string //微信标识
}
