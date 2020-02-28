package comment

import (
	"github.com/hailongz/golang/db"
)

type Comment struct {
	db.Object
	Pid     int64       `json:"pid" name:"pid" title:"父级ID" index:"ASC"`
	Eid     int64       `json:"eid" name:"eid" title:"评论目标ID" index:"ASC"`
	Uid     int64       `json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Body    string      `json:"body" name:"body" title:"内容" length:"-1"`
	Options interface{} `json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加" length:"-1"`
	Ctime   int64       `json:"ctime" name:"ctime" title:"创建时间"`
	State   int32       `json:"state" name:"state" title:"回收状态" index:"ASC"`
	Path    string      `json:"path" name:"path" title:"一个评论下的所有回复的路径" index:"ASC"`
}

func (O *Comment) GetName() string {
	return "comment"
}

func (O *Comment) GetTitle() string {
	return "评论"
}
