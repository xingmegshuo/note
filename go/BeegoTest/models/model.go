package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// 表设计
type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "sqlite3", "./DB/test.db")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
