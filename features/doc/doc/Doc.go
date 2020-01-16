package doc

import (
	"github.com/hailongz/golang/db"
)

type Doc struct {
	db.Object
	Pid	int64	`json:"pid" name:"pid" title:"父级ID" index:"ASC"`
	Title	string	`json:"title" name:"title" title:"标题" length:"2048"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Type	int32	`json:"type" name:"type" title:"类型" index:"ASC"`
	Path	string	`json:"path" name:"path" title:"路径" length:"2048"`
	Keyword	string	`json:"keyword" name:"keyword" title:"搜索关键字" length:"2048"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ext	string	`json:"ext" name:"ext" title:"扩展名" length:"32" index:"DESC"`
	Mtime	int64	`json:"mtime" name:"mtime" title:"最近修改时间" index:"DESC"`
	Atime	int64	`json:"atime" name:"atime" title:"最近访问时间" index:"DESC"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Doc) GetName() string {
	return "doc"
}

func (O *Doc) GetTitle() string {
	return "文档"
}

