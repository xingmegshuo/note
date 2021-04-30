/***************************
@File        : lunbotu.go
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 轮播图管理
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

// 新建图片展示

func CreateImage(c *gin.Context) {
	db := DB
	filename := Upload(c)
	imgList := models.ImgList{
		Ord:    c.PostForm("ord"),
		ImgUrl: filename,
		IsShow: c.PostForm("is_show"),
	}
	res := models.ImgList{}
	result := db.Where(&imgList).First(&res)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && len(filename) > 0 {
		db.Create(&imgList)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "新建轮播图成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "新建轮播图失败",
		})
	}
}

// 修改图片

func ChangeImage(c *gin.Context) {
	db := DB
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	res := models.ImgList{}
	result := db.First(&res, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && id != 0 {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "没有此内容",
		})
	} else {
		filename := Upload(c)
		newActive := models.ImgList{}
		if filename != "" {
			newActive = models.ImgList{
				Ord:    c.PostForm("ord"),
				IsShow: c.PostForm("is_show"),
				ImgUrl: filename,
			}
		} else {
			newActive = models.ImgList{
				Ord:    c.PostForm("ord"),
				IsShow: c.PostForm("is_show"),
			}
		}

		db.Model(&res).Updates(&newActive)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "更新成功",
		})
	}
}

// 获取图片

func GetImg(c *gin.Context) {
	db := DB
	paginationQuery, _ := c.Get("PaginationQuery")
	query := paginationQuery.(*pagination.PaginationQuery)
	var (
		actives []models.ImgList
		count   int64
	)
	db.Model(&models.ImgList{}).Scopes(pagination.PaginationScope(query)).Count(&count)
	db.Model(&models.ImgList{}).Scopes(pagination.PaginationScope(query)).Find(&actives)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   actives,
		"count":  count,
		"page":   query.PageNum,
	})
}
