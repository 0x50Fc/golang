package feed

import (
	"github.com/hailongz/golang/db"
)

type Feed struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"发布者ID"`
	Body	string	`json:"body" name:"body" title:"内容" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Status	int32	`json:"status" name:"status" title:"状态" index:"ASC"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
	State	int32	`json:"state" name:"state" title:"回收状态" index:"ASC"`
}

func (O *Feed) GetName() string {
	return "feed"
}

func (O *Feed) GetTitle() string {
	return "动态"
}

