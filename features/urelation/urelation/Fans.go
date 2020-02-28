package urelation

import (
	"github.com/hailongz/golang/db"
)

type Fans struct {
	db.Object
	Uid	int64	`json:"uid,omitempty" title:"用户ID" index:"ASC"`
	Fuid	int64	`json:"fuid,omitempty" title:"好友ID" index:"ASC"`
	Type	int32	`json:"type,omitempty" title:"类型"`
	Ctime	int64	`json:"ctime,omitempty" title:"创建时间"`
}

func (O *Fans) GetName() string {
	return "fans"
}

func (O *Fans) GetTitle() string {
	return "关系表基类"
}

