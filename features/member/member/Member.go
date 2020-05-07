package member

import (
	"github.com/hailongz/golang/db"
)

type Member struct {
	db.Object
	Bid	int64	`json:"bid" name:"bid" title:"\b\b商户ID" index:"ASC"`
	Uid	int64	`json:"uid" name:"uid" title:"成员ID" index:"ASC"`
	Title	string	`json:"title" name:"title" title:"备注名" length:"255"`
	Keyword	string	`json:"keyword" name:"keyword" title:"搜索关键字" length:"2048"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1" jsonType:"true"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Member) GetName() string {
	return "member"
}

func (O *Member) GetTitle() string {
	return "成员"
}

