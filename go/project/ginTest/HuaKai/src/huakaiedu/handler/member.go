/***************************
@File        : member.go
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 会员展示
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
	* @api {Get} /api/memberUser 获取会员
	* @apiVersion 0.1.0
	* @apiName 获取会员
	* @apiGroup web
	* @apiParam {String} PageNum 页数
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
	*			会员信息
	*		}
	*	],
	*	"page": 1,
	*	"status": "success"
	*}
 **/

// 获取会员信息

func GetMembers(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		members []models.Member
		count   int64
	)
	db.Model(&models.Member{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	// SELECT * FROM `users`  WHERE ((`age` = "5") or (`email` like "%user-1%")) ORDER BY id desc LIMIT 3 OFFSET 0
	db.Model(&models.Member{}).Scopes(pagination.PaginationScope(query)).Find(&members)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   members,
		"count":  count,
		"page":   query.PageNum,
	})
}

/**
	* @api {Get} /api/memberUser 获取会员成果
	* @apiVersion 0.1.0
	* @apiName 获取成果
	* @apiGroup web
	* @apiParam {String} PageNum 页数
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
	*			会员成果
	*		}
	*	],
	*	"page": 1,
	*	"status": "success"
	*}
 **/

// 获取会员成果

func GetMemberRes(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		memberRes []models.MemberRes
		count     int64
	)
	db.Model(&models.MemberRes{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	// SELECT * FROM `users`  WHERE ((`age` = "5") or (`email` like "%user-1%")) ORDER BY id desc LIMIT 3 OFFSET 0
	db.Model(&models.MemberRes{}).Scopes(pagination.PaginationScope(query)).Find(&memberRes)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   memberRes,
		"count":  count,
		"page":   query.PageNum,
	})
}

/**
	* @api {Post} /api/member 新建会员展示
	* @apiVersion 0.1.0
	* @apiName 新建会员展示
	* @apiGroup All
	* @apiParam {String} name 活动名称
	* @apiParam {String} content 活动内容
	* @apiParam {String} address 活动地址
	* @apiParam {string} contact 联系方式
	* @apiParam {file} image 图片
	* @apiParam {string} Time 活动时间
	* @apiParamExample {json} Request-Example
	* {
	*  "name": "今天是个好日子",
	*	"content":"bug消失术",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	*	"message":"新建活动成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"新建活动失败",
	*}
 **/

// 新建会员成果

func CreateMemberRes(c *gin.Context) {
	db := DB
	filename := Upload(c)
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	if filename != "" && id != 0 {
		member := models.MemberRes{
			Title:    c.PostForm("title"),
			Content:  c.PostForm("content"),
			ImgUrl:   filename,
			MemberID: id,
		}
		db.Create(&member)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "新建会员成果成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "新建活动失败，图片和会员不能为空",
		})
	}
}

/**
	* @api {Put} /active 修改会员成果
	* @apiVersion 0.1.0
	* @apiName 修改会员成果
	* @apiGroup All
 	* @apiParam {String} Title 标题
	* @apiParam {String} content 内容
	* @apiParam {file} image 图片
	* @apiParam {string} memberId 会员
	* @apiParam {string} id 主键
	* @apiParamExample {json} Request-Example
	* {
	*  "id": "1",
	*	"content":"bug消失术",
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":"success",
	*	"message":"更新成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"没有此信息",
	*}
 **/

// 修改会员成果信息

func ChangeMemberRes(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	res := models.MemberRes{}
	result := db.First(&res, id)
	memberID, _ := strconv.Atoi(c.DefaultPostForm("memberId", "0"))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有此信息",
		})
	} else {
		filename := Upload(c)
		newMember := models.MemberRes{}
		if filename != "" {
			newMember = models.MemberRes{
				Title:    c.PostForm("title"),
				Content:  c.PostForm("content"),
				ImgUrl:   filename,
				MemberID: memberID,
			}
		} else {
			newMember = models.MemberRes{
				MemberID: memberID,
				Title:    c.PostForm("title"),
				Content:  c.PostForm("content"),
			}
		}

		db.Model(&res).Updates(&newMember)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "更新成功",
			"data":    res,
		})
	}
}
