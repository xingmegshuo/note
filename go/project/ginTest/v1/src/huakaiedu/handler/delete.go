/***************************
@File        : delete.go
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 删除
****************************/

package handler

import (
	"github.com/gin-gonic/gin"
	"huakai/models"
	"strconv"
)

/**
	* @api {Delete} /delete 删除数据
	* @apiVersion 0.1.0
	* @apiName 删除数据
	* @apiGroup All
	* @apiParam {string} id 主键
	* @apiParam {string} table 表名
	* @apiParamExample {json} Request-Example
	* {
	* 	"id":"1",
	* 	"table":"user",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	*	"message":"删除成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"删除失败",
	*}
**/

// 删除数据

func DeleteData(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	table := c.Query("table")
	//log.Println(id, table)
	if id == 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "删除失败",
		})
		return
	}
	switch table {
	case "user":
		db.Delete(&models.User{}, id)
	case "active":
		db.Delete(&models.Active{}, id)
	case "sign":
		db.Delete(&models.Sign{}, id)
	case "member":
		db.Delete(&models.Member{}, id)
	case "memberRes":
		db.Delete(&models.MemberRes{}, id)
	case "join":
		db.Delete(&models.JoinUs{}, id)
	case "img":
		db.Delete(&models.ImgList{}, id)
	default:
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "删除成功",
	})
}
