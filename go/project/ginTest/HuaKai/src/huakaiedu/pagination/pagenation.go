/***************************
@File        : pagenation.go
@Time        : 2021/04/20 14:10:55
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 分页功能
****************************/

package pagination

import (
	"fmt"
	"huakai/config"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomeParser struct{}

//Parser 建议自己实现分页解析接口
// 可以自定义解析GET/POST参数，此处以GET为例子
func (*CustomeParser) Parser(c *gin.Context) *PaginationQuery {
	params := c.Request.URL.Query()

	PageNumStr := c.DefaultQuery("page", "1")
	PageNum, err := strconv.Atoi(PageNumStr)
	if err != nil {
		PageNum = 1
	}
	pageSize := config.Con.PageSize
	orderBy := c.DefaultQuery("order_by", "id")
	order := c.DefaultQuery("order", "asc")
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	scope := c.DefaultQuery("scope", "")
	isChildQueryStr := c.DefaultQuery("is_child", "false")
	isChildQuery, err := strconv.ParseBool(isChildQueryStr)
	if err != nil {
		isChildQuery = false
	}
	childRelationStr := c.DefaultQuery("child_relation", "0")
	childRelation, err := strconv.Atoi(childRelationStr)
	if err != nil {
		childRelation = 0
	}

	pattern := regexp.MustCompile(fmt.Sprintf(`(%s)\.(\d+)\.key`, RequestQueryPrefix))

	// 使用方式 search.1.key = 字段, search.1.value=搜索值
	searchTerms := make([]SearchTerm, 0)
	for key, values := range params {
		if strings.HasPrefix(key, RequestQueryPrefix) {
			matched := pattern.FindAllStringSubmatch(key, -1)
			if len(matched) > 0 {
				suffix := matched[0][2]
				queryField := values[0]
				queryValue := c.DefaultQuery(fmt.Sprintf("%s.%s.%s", RequestQueryPrefix, suffix, "value"), "")
				queryScope := c.DefaultQuery(fmt.Sprintf("%s.%s.%s", RequestQueryPrefix, suffix, "scope"), "")
				queryComparision := c.DefaultQuery(fmt.Sprintf("%s.%s.%s", RequestQueryPrefix, suffix, "comparision"), "eq")
				queryExactlyStr := c.DefaultQuery(fmt.Sprintf("%s.%s.%s", RequestQueryPrefix, suffix, "exactly"), "true")
				queryExactly, err := strconv.ParseBool(queryExactlyStr)
				if err != nil {
					queryExactly = true
				}
				if len(queryField) > 0 && len(queryValue) > 0 {
					searchTerms = append(searchTerms, SearchTerm{
						Key:         queryField,
						Value:       queryValue,
						Scope:       queryScope,
						Comparision: queryComparision,
						Exactly:     queryExactly,
					})
				}
			}
		}
	}
	var search []SearchGroup
	if isChildQuery {
		search = append(search, SearchGroup{
			ChildRelation: SearchRelation(childRelation),
			IsChildQuery:  true,
			ChildSearch:   searchTerms,
		})
	} else {
		for _, term := range searchTerms {
			search = append(search, SearchGroup{
				SearchTerm: term,
			})
		}
	}
	return &PaginationQuery{
		PageSize: pageSize,
		PageNum:  PageNum,
		OrderBy:  orderBy,
		Order:    order,
		Limit:    limit,
		Search:   search,
		Scope:    scope,
	}
}
