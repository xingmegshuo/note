/***************************
@File        : member.go
@Time        : 2021/04/19 14:37:05
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 会员
****************************/

package models

import "gorm.io/gorm"

// 会员成果

type MemberRes struct {
	gorm.Model
	Title    string
	Content  string
	ImgUrl   string `bson:"imgUrl"`
	Member   Member `gorm:"ForeignKey:MemberID"`
	MemberID int
}

// 会员成员

type Member struct {
	gorm.Model
	User   User `gorm:"ForeignKey:UserID"`
	UserID int
	IsShow string
}
