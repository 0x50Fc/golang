package lookin

import (
	"github.com/hailongz/golang/db"
)

type Lookin struct {
	db.Object
	Tid	int64	`json:"tid" name:"tid" title:"目标ID" index:"ASC"`
	Iid	int64	`json:"iid" name:"iid" title:"项ID" index:"ASC"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Fuid	int64	`json:"fuid" name:"fuid" title:"好友ID" index:"ASC"`
	Fcode	string	`json:"fcode" name:"fcode" title:"好友推荐码"`
	Flevel	int32	`json:"flevel" name:"flevel" title:"关系级别" index:"ASC"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1" jsonType:"true"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Lookin) GetName() string {
	return "lookin"
}

func (O *Lookin) GetTitle() string {
	return "在看"
}

