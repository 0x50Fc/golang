package inbox

import (
	"github.com/hailongz/golang/db"
)

type Inbox struct {
	db.Object
	Type	int64	`json:"type" name:"type" title:"收件类型" index:"desc"`
	Uid	int64	`json:"uid" name:"uid" title:"接受者ID" index:"desc"`
	Fuid	int64	`json:"fuid" name:"fuid" title:"发布者ID" index:"desc"`
	Mid	int64	`json:"mid" name:"mid" title:"内容ID" index:"desc"`
	Iid	int64	`json:"iid" name:"iid" title:"内容项ID" index:"desc"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间" index:"desc"`
}

func (O *Inbox) GetName() string {
	return "inbox"
}

func (O *Inbox) GetTitle() string {
	return "收件箱"
}

