/***************************
@File        : middlewar_page.go
@Time        : 2021/04/20 15:50:09
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : middleware
****************************/

package pagination

import "github.com/gin-gonic/gin"

//RequestQueryPrefix 参数前缀, 用于`PaginationQuery.SearchGroup`和其他参数区分

var RequestQueryPrefix = "search"
var PaginationQueryContextKey = "PaginationQuery"

//Pagination 实现这个接口，把前端传过来的参数解析成需要的`PaginationQuery`
//接口的`Parse`可以自定义足够复杂，相对应前端需要传的参数越多，该方法的解析越复杂

type Pagination interface {
	Parser(c *gin.Context) *PaginationQuery
}

// PaginationParser 兼容gin中间件
// 从`gin.Context`中解析`PaginationQuery`
// url例子: page=1&&size=10&&search.1.key=age&&search.1.value=20&&search.2.key=gender&&search.2.value=0&&orderfield=age&&orderby=desc&&search.1.exactly=true&&search.2.exactly=false

func New(pagination Pagination) gin.HandlerFunc {
	return func(c *gin.Context) {
		paginationQuery := pagination.Parser(c)
		c.Set(PaginationQueryContextKey, paginationQuery)
		c.Next()
	}
}
