/***************************
@File        : db_test.go
@Time        : 2020/12/10 15:24:54
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试数据库
****************************/

package main

import (
	"fmt"
	"wwKill/Mydb"
)

func main() {
	Mydb.Db("./wwKill")
	orm := Mydb.SetEngin("./wwKill")
	fmt.Println(orm)
}
