package article

import (
	"github.com/hailongz/golang/db"
)

type Article struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"发布者ID"`
	Body	string	`json:"body" name:"body" title:"内容" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
	State	int32	`json:"state" name:"state" title:"状态" index:"ASC"`
}

func (O *Article) GetName() string {
	return "article"
}

func (O *Article) GetTitle() string {
	return "文章"
}

