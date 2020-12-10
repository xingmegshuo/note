/***************************
@File        : database.go
@Time        : 2020/12/10 15:09:43
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : sqllite数据库
****************************/
package Mydb

import (
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var orm *xorm.Engine

func SetEngin(db string) *xorm.Engine {
	var err error
	orm, err = xorm.NewEngine("sqlite3", db)
	if err != nil {
		panic(err)
	}
	orm.TZLocation = time.Local
	return orm
}
