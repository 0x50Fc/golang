package top

import (
	"github.com/hailongz/golang/db"
)

type Top struct {
	db.Object
	Tid	int64	`json:"tid" name:"tid" title:"目标ID" index:"ASC"`
	Keyword	string	`json:"keyword" name:"keyword" title:"搜索关键字" length:"2048"`
	Sid	int64	`json:"sid" name:"sid" title:"序号 降序" index:"DESC"`
	Rank	int32	`json:"rank" name:"rank" title:"排名" index:"ASC"`
	Fixed	int32	`json:"fixed" name:"fixed" title:"固定排名位置"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
}

func (O *Top) GetName() string {
	return "top"
}

func (O *Top) GetTitle() string {
	return "Top"
}

