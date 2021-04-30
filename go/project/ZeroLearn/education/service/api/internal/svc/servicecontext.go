package svc

import (
	"api/internal/config"

	"api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(sqlite.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	db.AutoMigrate(&models.User{})
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
