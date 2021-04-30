/***************************
@File        : sign.go
@Time        : 2021/04/21 16:52:18
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 活动报名
****************************/

package handler

import (
	"errors"
	"gorm.io/gorm"
	"huakai/models"
	"huakai/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
	* @api {Get} /sign 获取活动报名信息
	* @apiVersion 0.1.0
	* @apiName 获取活动报名
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
	*			活动数据
	*		}
	*	],
	*	"page": 1,
	*	"status": "success"
	*}
 **/

//获取活动数据

func GetSign(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		signs []models.Sign
		count int64
	)
	db.Model(&models.Sign{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	db.Model(&models.Sign{}).Scopes(pagination.PaginationScope(query)).Find(&signs)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   signs,
		"count":  count,
		"page":   query.PageNum,
	})
}

/**
	* @api {Post} /sign 活动报名
	* @apiVersion 0.1.0
	* @apiName 活动报名
	* @apiGroup All
	* @apiParam {String} activeId 活动id
	* @apiParam {String} id 用户id
	* @apiParamExample {json} Request-Example
	* {
	*  "activeId": "1",//页数
	*	"id":"2",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	* 	"message":"申请活动成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"申请活动失败",
	*}
	*}
 **/

// 申请活动

func CreateSign(c *gin.Context) {
	db := DB
	activeId, _ := strconv.Atoi(c.Query("activeId"))
	id, _ := strconv.Atoi(c.Query("id"))
	sign := models.Sign{
		ActiveID: activeId,
		UserID:   id,
		Status:   "申请",
	}
	res := models.Sign{}
	result := db.Where(&sign).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&sign)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "申请活动成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "申请活动失败,重复申请",
		})
	}
}

/**
	* @api {Put} /sign 取消活动报名
	* @apiVersion 0.1.0
	* @apiName 取消活动报名
	* @apiGroup All
	* @apiParam {String} id id
	* @apiParam {String} status 取消申请
	* @apiParamExample {json} Request-Example
	* {
	*  "id": "1",//页数
	*	"status":"取消申请",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	* 	"message":"取消活动申请",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"没有这个记录",
	*}
	*}
 **/

// 取消申请

func ChangeSign(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	res := models.Sign{}
	result := db.First(&res, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有这个记录",
		})
	} else {
		newSign := models.Sign{
			Status: c.Query("status"),
		}
		db.Model(&res).Updates(&newSign)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "修改活动申请",
		})
	}
}
