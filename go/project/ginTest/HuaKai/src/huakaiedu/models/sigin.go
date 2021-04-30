/***************************
@File        : sigin.go
@Time        : 2021/04/19 14:12:03
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 报名表
****************************/

package models

import "gorm.io/gorm"

type Sign struct {
	gorm.Model
	Active   Active `gorm:"ForeignKey:ActiveID"`
	ActiveID int
	User     User `gorm:"ForeignKey:UserID"`
	UserID   int
	Status   string
}
