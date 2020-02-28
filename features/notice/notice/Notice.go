package notice

import (
	"github.com/hailongz/golang/db"
)

type Notice struct {
	db.Object
	Type	int32	`json:"type" name:"type" title:"通知类型" index:"ASC"`
	Fid	int64	`json:"fid" name:"fid" title:"消息来源ID" index:"ASC"`
	Iid	int64	`json:"iid" name:"iid" title:"消息来源项ID" index:"ASC"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Body	string	`json:"body" name:"body" title:"通知内容" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Notice) GetName() string {
	return "notice"
}

func (O *Notice) GetTitle() string {
	return "通知"
}

