/***************************
@File        : join.go
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 会员申请表
****************************/

package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"huakai/models"
	"huakai/pagination"
	"strconv"
)

/**
	* @api {Get} /join 获取会员申请表
	* @apiVersion 0.1.0
	* @apiName 获取会员申请
	* @apiGroup All
	* @apiParam {String} PageNum 页数 //默认为1
	* @apiParam {String} Search 搜索内容
	* @apiParam {String} OrderBy 排序方式
	* @apiParamExample {json} Request-Example
	* {
	*  "PageNum": 10,//页数
	*	"orderBy":"asc",//排序方式
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"count": 1,
	*	"data": [
	*		{
	*			申请表数据
	*		}
	*	],
	*	"page": 1,
	*	"status": "success"
	*}
 **/

// 会员申请信息

func GetJoinMes(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		tables []models.JoinUs
		count  int64
	)
	db.Model(&models.JoinUs{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	db.Model(&models.JoinUs{}).Scopes(pagination.PaginationScope(query)).Find(&tables)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   tables,
		"count":  count,
		"page":   query.PageNum,
	})
}

/**
	* @api {Post} /join 会员申请
	* @apiVersion 0.1.0
	* @apiName 会员申请
	* @apiGroup All
	* @apiParam {String} userId 用户ID
	* @apiParam {String} data 申请表格中的内容
	* @apiParamExample {json} Request-Example
	* {
	*  "userId": "1",//页数
	*	"data":"xxxx",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	* 	"message":"会员申请表提交成功,等待审核",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"该用户已经申请",
	*}
	*}
 **/

// 新建会员申请

func CreateJoinTable(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.PostForm("userId"))
	table := models.JoinUs{
		UserID:    id,
		Data:      c.PostForm("data"),
		Status:    "未审核",
		PayStatus: "未缴纳",
	}
	res := models.JoinUs{}
	result := db.Where(&table).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		db.Create(&table)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "会员申请表提交成功，等待审核",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "该用户已经申请",
		})
	}
}

/**
	* @api {Put} /join 修改申请状态
	* @apiVersion 0.1.0
	* @apiName 审核申请
	* @apiGroup All
	* @apiParam {String} id id
	* @apiParam {String} status 审核状态
	* @apiParam {string} payments 支付状态
	* @apiParamExample {json} Request-Example
	* {
	*  "id": "1",//页数
	*	"status":"审核通过",
	* 	"payments":"已经支付",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	* 	"message":"修改成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"没有数据可以修改",
	*}
	*}
 **/

// 修改审核状态和支付状态

func ChangeJoinTable(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	res := models.JoinUs{}
	result := db.First(&res, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有数据可以修改",
		})
	} else {
		newTable := models.JoinUs{
			Status:    c.Query("status"),
			PayStatus: c.Query("payments"),
		}
		if newTable.Status == "审核通过" && newTable.PayStatus == "已经支付" {
			user := models.User{}
			db.First(&user, res.UserID)
			db.Model(&user).Update("Iden", "会员")
			member := models.Member{
				UserID: res.UserID,
				IsShow: "false",
			}
			db.Create(&member)
		}
		db.Model(&res).Updates(&newTable)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "修改成功",
		})
	}
}
