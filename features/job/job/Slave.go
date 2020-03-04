package job

import (
	"github.com/hailongz/golang/db"
)

type Slave struct {
	db.Object
	Prefix	string	`json:"prefix" name:"prefix" title:"别名前缀" length:"128"`
	Token	string	`json:"token" name:"token" title:"授权token" length:"32" index:"ASC"`
	State	int32	`json:"state" name:"state" title:"主机状态" index:"ASC"`
	Etime	int64	`json:"etime" name:"etime" title:"超时时间" index:"DESC"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Slave) GetName() string {
	return "slave"
}

func (O *Slave) GetTitle() string {
	return "Slave 主机"
}

