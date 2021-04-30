/***************************
@File        : active.go
@Time        : 2021/04/21 15:45:07
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 活动内容接口
****************************/

package handler

import (
	"errors"
	"huakai/models"
	"huakai/pagination"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
	* @api {Get} /active 获取活动信息
	* @apiVersion 0.1.0
	* @apiName 获取活动
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

// 获取全部活动

func GetActive(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		actives []models.Active
		count   int64
	)
	db.Model(&models.Active{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	db.Model(&models.Active{}).Scopes(pagination.PaginationScope(query)).Find(&actives)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   actives,
		"count":  count,
		"page":   query.PageNum,
	})
}

/**
	* @api {Post} /active 新建活动信息
	* @apiVersion 0.1.0
	* @apiName 新建活动
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

func CreateActive(c *gin.Context) {
	db := DB
	filename := Upload(c)
	active := models.Active{
		Name:    c.PostForm("name"),
		Content: c.PostForm("content"),
		Image:   filename,
		Address: c.PostForm("address"),
		Contact: c.PostForm("contact"),
		Time:    c.PostForm("time"),
	}
	res := models.Active{}
	result := db.Where(&active).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && len(filename) > 0 {
		db.Create(&active)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "新建活动成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "新建活动失败",
		})
	}
}

/**
	* @api {Put} /active 修改活动信息
	* @apiVersion 0.1.0
	* @apiName 修改活动
	* @apiGroup All
 	* @apiParam {String} name 活动名称
	* @apiParam {String} content 活动内容
	* @apiParam {String} address 活动地址
	* @apiParam {string} contact 联系方式
	* @apiParam {string} Time 活动时间
	* @apiParam {file} image 图片
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
	*	"message":"修改活动成功",
	*}
	* @apiErrorExample {json} Response-Example
	*{
	*	"status":"error",
	* 	"message":"修改活动失败",
	*}
 **/

func UpdateActive(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	res := models.Active{}
	result := db.First(&res, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有此活动",
		})
	} else {
		filename := Upload(c)
		newActive := models.Active{}
		if filename != "" {
			newActive = models.Active{
				Name:    c.PostForm("name"),
				Content: c.PostForm("content"),
				Image:   filename,
				Address: c.PostForm("address"),
				Contact: c.PostForm("contact"),
			}
		} else {
			newActive = models.Active{
				Name:    c.PostForm("name"),
				Content: c.PostForm("content"),
				Address: c.PostForm("address"),
				Contact: c.PostForm("contact"),
			}
		}

		db.Model(&res).Updates(&newActive)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "更新成功",
			"data":    res,
		})
	}
}

// 图片上传

func Upload(c *gin.Context) string {
	file, _, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		//c.String(http.StatusBadRequest, "Bad request")
		return ""
	}
	//文件的名称
	filename := time.Now().Format("20060102150405") + ".png"
	out, err := os.Create("static/img/" + filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filename = "static/img/" + filename
	return filename
}
