/***************************
@File        : join.go
@Time        : 2021/04/19 14:21:42
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 入会申请表
****************************/

package models

import "gorm.io/gorm"

// 申请表

type JoinUs struct {
	gorm.Model
	User      User `gorm:"ForeignKey:UserID"`
	UserID    int
	Data      string //数据表格
	Status    string `gorm:"Default:false"` //审核状态
	PayStatus string `gorm:"Default:false"` //缴纳会费状态
}
