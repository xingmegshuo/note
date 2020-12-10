/***************************
@File        : table.go
@Time        : 2020/12/10 15:53:52
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :  创建数据库
****************************/

package Mydb

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)


// 生成创建数据表的sql
func table(n int) (s string){
	switch {
		case n== 0:
			s := `create table users (userId integer, uname text);`
			return s		
		default:
			s := `show tables;`
			return s
	}
}

// 初始化数据库
func Db(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName+".db")
	if err != nil {
		fmt.Printf("错误:%s", err)
	}
	for i:=0;i<2;i++{
		s := table(i)
		db.Exec(s)
	}
	return db
}

