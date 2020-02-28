package media

import (
	"github.com/hailongz/golang/db"
)

type Media struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Type	string	`json:"type" name:"type" title:"类型" length:"32" index:"ASC"`
	Title	string	`json:"title" name:"title" title:"标题" length:"2048"`
	Keyword	string	`json:"keyword" name:"keyword" title:"关键字" length:"4096"`
	Path	string	`json:"path" name:"path" title:"存储路径" length:"2048"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Media) GetName() string {
	return "media"
}

func (O *Media) GetTitle() string {
	return "媒体"
}

