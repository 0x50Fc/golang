package like

import (
	"github.com/hailongz/golang/db"
)

type Like struct {
	db.Object
	Tid	int64	`json:"tid" name:"tid" title:"目标ID" index:"ASC"`
	Iid	int64	`json:"iid" name:"iid" title:"项ID" index:"ASC"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Like) GetName() string {
	return "like"
}

func (O *Like) GetTitle() string {
	return "赞"
}

