/***************************
@File        : user.go
@Time        : 2020/12/10 15:13:40
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : orm user
****************************/

package Mydb

import (
	"log"
	"time"
)

type User struct {
	Id         int
	OpenID     string
	NickName   string
	AvatarURL  string
	CreateDate time.Time
	Login      time.Time
	LastLogin  time.Time
}

// 根据openid返回用户
func (this *User) GetUserByOpenId(openID string) (*User, bool) {
	user := &User{OpenID: openID}
	has, err := orm.Get(user)
	if err != nil {
		log.Panic(err)
	}
	return user, has
}

// 插入用户

func (this *User) Insert() bool {
	_, err := orm.InsertOne(this)
	if err != nil {
		log.Panic(err)
	}
	return true
}

// 删除

func (this *User) Delete() bool {
	_, err := orm.Delete(this)
	if err != nil {
		log.Panic(err)
	}
	return true
}

// 修改
func (this *User) Update() bool {
	_, err := orm.Id(this.Id).Update(this)
	if err != nil {
		log.Panic(err)
	}
	return true
}
