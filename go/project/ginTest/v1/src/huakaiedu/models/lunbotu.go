/***************************
@File        : lunbotu.go
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 轮播图管理
****************************/

package models

import "gorm.io/gorm"

// 轮播图管理

type ImgList struct {
	gorm.Model
	Ord    string
	ImgUrl string `bson:"imgUrl"`
	IsShow string `gorm:"default:true"`
}
