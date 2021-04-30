/***************************
@File        : user.go
@Time        : 2021/04/19 17:48:22
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 用户登录
****************************/

package handler

import (
	"errors"
	"huakai/models"
	"huakai/pagination"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB = models.Db

/**
	* @api {Post} /wx/user/login 用户登录
	* @apiVersion 0.1.0
	* @apiName 登录
	* @apiGroup 微信
	* @apiParam {String} nickname 微信昵称
	* @apiParam {String} openId 微信标识
	* @apiParam {String} gender 性别
	* @apiParam {String} avatar 头像地址
	* @apiParam {String} city 城市
	* @apiParamExample {json} Request-Example
	* {
	*  "nickname": "Eve",
	*	"openId":"b9999",
	*	"gender":"男",
	*	"avatar":"https://.....png",
	*	"city":"成都"
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	* 第一次登录
	* {
	*	"status":"success",
	*	"message":"注册登录成功"
	* }
	* 普通登录
	* {
	*	"status":"success",
	*	"message":"登录成功"
	* }
 **/

func Login(c *gin.Context) {
	db := DB
	user := models.User{
		NickName: c.PostForm("nickname"),
		OpenId:   c.PostForm("openId"),
		Gender:   c.PostForm("gender"),
		Avatar:   c.PostForm("avatar"),
		City:     c.PostForm("city"),
		Iden:     c.PostForm("iden"),
	}
	res := models.User{}
	result := db.Where(&user).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&user)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "注册登录成功",
			"data":    user,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "登录成功",
			"data":    res,
		})
	}

}

/**
	* @api {Post} /web/user/login 用户登录
	* @apiVersion 0.1.0
	* @apiName 登录
	* @apiGroup web
	* @apiParam {String} nickname 微信昵称
	* @apiParam {String} openId 微信标识
	* @apiParamExample {json} Request-Example
	* {
	*  "account": "111",//账号
	*	"password:"xxx",//密码
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	* {
	*	"status":"success",
	*	"message":"登录成功"
	* }
	* @apiErrorExample {json} Response-Example
	* {
	*	"status":"error",
	* 	"message":"登录失败"
	* }
 **/

func LoginWeb(c *gin.Context) {
	db := DB
	user := models.User{
		NickName: c.PostForm("account"),
		OpenId:   c.PostForm("password"),
		Iden:     "admin",
	}
	res := models.User{}
	result := db.Where(&user).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "登录失败,联系管理员",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "登录成功",
			"data":    res,
		})
	}
}

/**
	* @api {Get} /web/getUsers 获取用户
	* @apiVersion 0.1.0
	* @apiName 获取用户
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
	*			"ID": 1,
	*			"CreatedAt": "2021-04-20T17:45:07.550026+08:00",
	*			"UpdatedAt": "2021-04-20T17:45:07.550026+08:00",
	*			"DeletedAt": null,
	*			"NickName": "superuser",
	*			"Gender": "",
	*			"Avatar": "",
	*			"City": "",
	*			"Phone": "",
	*			"Mail": "",
	*			"Ident": "admin",
	*			"OpenId": "123456789"
	*		}
	*	],
	*	"page": 1,
	*	"status": "success"
	*}
 **/

func GetUsers(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		users []models.User
		count int64
	)
	db.Model(&models.User{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	// SELECT * FROM `users`  WHERE ((`age` = "5") or (`email` like "%user-1%")) ORDER BY id desc LIMIT 3 OFFSET 0
	db.Model(&models.User{}).Scopes(pagination.PaginationScope(query)).Find(&users)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   users,
		"count":  count,
		"page":   paginationQuery,
	})
}

/**
	* @api {Post} /web/addSuper 创建管理员用户
	* @apiVersion 0.1.0
	* @apiName 创建用户
	* @apiGroup web
	* @apiParam {String} account 账号
	* @apiParam {String} password 密码
	* @apiParamExample {json} Request-Example
	* {
	*  "account": "1111111",//账号
	*	"password":"**********",//密码
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":  "success",
	*	"message": "添加管理员成功",
	*}
	* @apiErrorExample {json} Response-Example
	* {
	*	"status": "error",
	*	"message":"该账号已经存在",
	* }
 **/

func AddSuperUser(c *gin.Context) {
	db := DB
	superuser := models.User{
		NickName: c.PostForm("account"),
		OpenId:   c.PostForm("password"),
		Iden:     "admin",
	}
	res := models.User{}
	result := db.Where(&superuser).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&superuser)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "添加管理员成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "该账号已经存在",
		})
	}
}

/**
	* @api {Put} /web/changeUser 修改用户信息
	* @apiVersion 0.1.0
	* @apiName 修改用户
	* @apiGroup web
	* @apiParam {String} id 主键
	* @apiParam {String} or 其它字段
	* @apiParamExample {json} Request-Example
	* {
	*  "id": "1",//主键id
	*	"nickname":"新名字",//新名字
	* }
	*
	* @apiSuccessExample  {json} Response-Example
	*{
	*	"status":  "success",
	*	"message": "修改成功",
	*}
	* @apiErrorExample {json} Response-Example
	* {
	*	"status": "error",
	*	"message":"没有此用户",
	* }
 **/

func ChangeUser(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	res := models.User{}
	log.Println(id)
	result := db.First(&res, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有此用户",
		})
	} else {
		newUser := models.User{
			Gender:   c.Query("gender"),
			Avatar:   c.Query("avatar"),
			City:     c.Query("city"),
			NickName: c.Query("nickname"),
			Phone:    c.Query("phone"),
			Mail:     c.Query("mail"),
			Iden:     c.Query("iden"),
			OpenId:   c.Query("openid"),
		}
		log.Println(newUser)
		db.Model(&res).Updates(&newUser)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "更新成功",
		})
	}
}
