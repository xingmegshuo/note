/***************************
@File        : model.go
@Time        : 2021/04/20 15:57:05
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : model
****************************/

package pagination

//
// 定义部分
//

// PaginationQuery 分页查询
type PaginationQuery struct {
	PageSize int
	PageNum  int
	// OrderBy 小写的字段名称
	OrderBy string
	// Order 默认是'desc', 可选的: 'desc', 'asc'
	Order string
	// Limit 当指定pagesize时, limit参数是无效的
	Limit int
	// SearchGroup 其他搜索条件, 各个查询组之间是AND关系
	Search []SearchGroup
	// Scope 仅在复杂查询时候使用(包含多个表查询的时候)
	// 一般来说, Scope是表的名字, 用以表明哪个表用作复杂的匹配
	// 例如: join多个表有相同的字段名的时候
	Scope string
}

// SearchGroup 查询组定义
// 有子查询(ChildSearch) IsChildQuery=true的情况下，会忽略该Search里的查询直接查询子查询
type SearchGroup struct {
	SearchTerm
	// isChildQuery 是否包含子查询
	IsChildQuery bool
	// ChildSearch子查询，嵌套搜索, 仅仅当isChildQuery是true的时候有效
	ChildSearch []SearchTerm
	// ChildRelation, 子查询多个条件之间的关系，仅仅当isChildQuery是true的时候有效，默认值是AND
	ChildRelation SearchRelation
}
type SearchTerm struct {
	// key是搜索的键
	// value是搜索的值，可以是value1,value2, 多个值的时候将会被翻译成 "or" 的关系
	// scope和`PaginationQuery.Scope`相似
	Key, Value, Scope string
	// Comparision key和value之间的关系, 默认eq即等于
	// 参考ComparisionMap
	Comparision string
	// Exactly 精确匹配/模糊配置
	Exactly bool
}

var ComparisionMap = map[string]string{
	"gt":     ">",
	"ge":     ">=",
	"lt":     "<",
	"le":     "<=",
	"eq":     "=",
	"ne":     "!=",
	"in":     "in",
	"not_in": "not in",
}

// SearchRelation 搜索查询添加 And或者Or
type SearchRelation int

const (
	ChildSearchAnd SearchRelation = iota
	ChildSearchOr
)

// Options自定义的选项，用户区分PaginationScope不同的组合查询条件

type Option int

const (
	// Count Option=0的时候，标志返回查询总条数的语句
	Count Option = iota
)

func (o Option) in(options []Option) bool {
	for _, opt := range options {
		if opt == o {
			return true
		}
	}
	return false
}
