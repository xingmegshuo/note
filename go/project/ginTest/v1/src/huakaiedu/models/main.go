/***************************
@File        : main.go
@Time        : 2021/04/19 14:48:27
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 初始化
****************************/

package models

import (
	"errors"
	"huakai/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func init() {
	Db, err = gorm.Open(sqlite.Open(config.Con.Sql), &gorm.Config{})
	//使用mysql, gorm.Open(“mysql”, “user:pwd@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local”)
	if err != nil {
		log.Fatal("db connect error")
	}
	// defer db.Close() //延时调用函数
	Db.AutoMigrate(&User{}, &Active{}, &Sign{}, &JoinUs{}, &Member{}, &MemberRes{}, &ImgList{})
	sqlDB, err := Db.DB()
	// 配置连接池
	if err == nil {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(1000)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100000)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(-1)
	}
	adminUser := &User{
		NickName: config.Con.SuperAccount,
		OpenId:   config.Con.Password,
		Iden:     "admin",
	}
	result := Db.First(adminUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		Db.Create(&adminUser)
	}
	// for i := 0; i < 1000; i++ {
	// 	Db.Create(&User{
	// 		NickName: config.Con.SuperAccount,
	// 		OpenId:   config.Con.Password,
	// 		Iden:     "admin",
	// 	})
	// }

}
