/***************************
@File        : main.go
@Time        : 2020/12/03 13:09:51
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 学生信息管理
****************************/

package main

import (
	"fmt"
	"os"
)

// 保存学生变量
var (
	allstudent map[int64]*student //声明
)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 查看功能
func look() {
	fmt.Println("查看信息")
	for k, v := range allstudent {
		fmt.Printf("学号:%d 姓名:%s \n", k, v.name)
	}
}

// 新增功能
func add() {
	fmt.Println("新增信息")
	var (
		id   int64
		name string
	)
	fmt.Println("输入学号:")
	fmt.Scanln(&id)
	fmt.Println("输入姓名:")
	fmt.Scanln(&name)

	//调用newStudent
	newStu := newStudent(id, name)

	allstudent[id] = newStu

}

// 删除功能
func del() {
	fmt.Println("删除信息")
	var (
		deleteID int64
	)
	fmt.Println("输入删除学号:")
	fmt.Scanln(&deleteID)
	delete(allstudent, deleteID)
}

// 退出功能
func exit() {
	fmt.Println("退出操作")
	os.Exit(1)
}

// 默认功能
func default_demo() {
	fmt.Println("操作失败")
}
func main() {
	allstudent = make(map[int64]*student, 2) //开辟内存空间
	fmt.Println("欢迎")
	fmt.Println(`
		1.查看
		2.新增
		3.删除
		4.退出
	`)
	for {
		fmt.Println("输入操作：")
		// 等待输入
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d选项 \n", choice)
		switch choice {
		case 1:
			look()
		case 2:
			add()
		case 3:
			del()
		case 4:
			exit()
		default:
			default_demo()
		}

	}

}
