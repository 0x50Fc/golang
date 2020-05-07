package article

import (
	"github.com/hailongz/golang/db"
)

type Outbox struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"发布者ID"`
	Mid	int64	`json:"mid" name:"mid" title:"动态ID"`
	Body	string	`json:"body" name:"body" title:"内容" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Outbox) GetName() string {
	return "outbox"
}

func (O *Outbox) GetTitle() string {
	return "发布的动态"
}

