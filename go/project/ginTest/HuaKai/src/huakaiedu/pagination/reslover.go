package pagination

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

//
// 使用部分
//

// PaginationScope解析PaginationQuery到gorm.DB Scopes
// 示例：
//          query := &PaginationQuery{}
// 查询例子:
//          var result []yourModel
//          db.Scopes(PaginationScope(query)).Find(&result)
// 查询总数示例:
//          var c int
//          db.Scopes(PaginationScope(query, Count)).Count(&c)

func PaginationScope(query *PaginationQuery, options ...Option) func(db *gorm.DB) *gorm.DB {
	if Count.in(options) {
		// 只返回查询总条数的组合条件
		return func(db *gorm.DB) *gorm.DB {
			d := &paginationQueryDB{db, query}
			d.parseQuery()
			return d.db
		}
	}
	// 否则返回查询的所有条件
	return func(db *gorm.DB) *gorm.DB {
		d := &paginationQueryDB{db, query}
		d.parseQuery().parseLimit().parsePagination().parseOrder()
		return d.db
	}
}

//
// 查询条件解析部分
//

// paginationQueryDB 包含gorm.DB的结构体
type paginationQueryDB struct {
	db    *gorm.DB
	query *PaginationQuery
}

//解析查询条件
func (d *paginationQueryDB) parseQuery() *paginationQueryDB {
	if len(d.query.Search) > 0 {
		for _, search := range d.query.Search {
			var expression = ""
			if !search.IsChildQuery {
				expression = _subSearchExpress(search.SearchTerm)
			} else {
				var expressChildren = make([]string, 0)
				for _, childSearch := range search.ChildSearch {
					expressChildren = append(expressChildren, _subSearchExpress(childSearch))
				}
				if search.ChildRelation == ChildSearchOr {
					expression = strings.Join(expressChildren, " or ")
				} else if search.ChildRelation == ChildSearchAnd {
					expression = strings.Join(expressChildren, " and ")
				}
			}
			d.db = d.db.Where(expression)
		}
	}
	return d
}

//解析Limit

func (d *paginationQueryDB) parseLimit() *paginationQueryDB {
	if d.query.Limit > 0 {
		d.db = d.db.Limit(d.query.Limit)
	}
	return d
}

//解析分页条件

func (d *paginationQueryDB) parsePagination() *paginationQueryDB {
	if d.query.PageSize > 0 {
		d.db = d.db.Limit(d.query.PageSize)
		if d.query.PageNum > 0 {
			d.db = d.db.Offset((d.query.PageNum - 1) * d.query.PageSize)
		}
	}
	return d
}

//解析排序
func (d *paginationQueryDB) parseOrder() *paginationQueryDB {
	if len(d.query.OrderBy) > 0 {
		if len(d.query.Scope) == 0 {
			if strings.ToLower(d.query.Order) == "asc" {
				d.db = d.db.Order(fmt.Sprintf("%s asc", d.query.OrderBy))
			} else {
				d.db = d.db.Order(fmt.Sprintf("%s desc", d.query.OrderBy))
			}
		} else {
			if strings.ToLower(d.query.Order) == "asc" {
				d.db = d.db.Order(fmt.Sprintf("%s.%s asc", d.query.Scope, d.query.OrderBy))
			} else {
				d.db = d.db.Order(fmt.Sprintf("%s.%s desc", d.query.Scope, d.query.OrderBy))
			}
		}
	}
	return d
}

func _subSearchExpress(term SearchTerm) (sql string) {
	if term.Comparision == "" {
		term.Comparision = "eq"
	}
	comparisionSymbol, ok := ComparisionMap[term.Comparision]
	if !ok {
		comparisionSymbol = "="
	}
	// 逗号分隔，or关系连接
	values := strings.Split(term.Value, ",")
	vNum := len(values)
	if len(term.Scope) > 0 {
		if term.Exactly {
			// todo: 遗留数据库字段隐式转换类型，无法使用索引问题
			if comparisionSymbol == "in" || comparisionSymbol == "not in" {
				sql += "("
				sql += fmt.Sprintf("`%s`.`%s` %s (%s)", term.Scope, term.Key, comparisionSymbol, array2String(values))
				sql += ")"
			} else {
				for index, value := range values {
					if index == 0 {
						sql += "("
					}
					sql += fmt.Sprintf("(`%s`.`%s` %s \"%s\")", term.Scope, term.Key, comparisionSymbol, value)
					if index < vNum-1 {
						sql += " or "
					}
					if index == vNum-1 {
						sql += ")"
					}
				}
			}
		} else {
			for index, value := range values {
				if index == 0 {
					sql += "("
				}
				sql += fmt.Sprintf("(`%s`.`%s` like \"%s\")", term.Scope, term.Key, "%"+value+"%")
				if index < vNum-1 {
					sql += " or "
				}
				if index == vNum-1 {
					sql += ")"
				}
			}
		}
	} else {
		if term.Exactly {
			if comparisionSymbol == "in" || comparisionSymbol == "not in" {
				sql += "("
				sql += fmt.Sprintf("`%s` %s (%s)", term.Key, comparisionSymbol, array2String(values))
				sql += ")"
			} else {
				for index, value := range values {
					if index == 0 {
						sql += "("
					}
					sql += fmt.Sprintf("(`%s` %s \"%s\")", term.Key, comparisionSymbol, value)
					if index < vNum-1 {
						sql += " or "
					}
					if index == vNum-1 {
						sql += ")"
					}
				}
			}
		} else {
			for index, value := range values {
				if index == 0 {
					sql += "("
				}
				sql += fmt.Sprintf("(`%s` like \"%s\")", term.Key, "%"+value+"%")
				if index < vNum-1 {
					sql += " or "
				}
				if index == vNum-1 {
					sql += ")"
				}
			}
		}
	}
	return
}

func array2String(values []string) (s string) {
	for i, v := range values {
		s += fmt.Sprintf("\"%s\"", v)
		if i < len(values)-1 {
			s += ","
		}
	}
	return
}
