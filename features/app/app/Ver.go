package app

import (
	"github.com/hailongz/golang/db"
)

type Ver struct {
	db.Object
	Appid	int64	`json:"appid" name:"appid" title:"用户ID" index:"ASC"`
	Ver	int32	`json:"ver" name:"ver" title:"版本号" index:"DESC"`
	Info	interface{}	`json:"info,omitempty" name:"info" title:"应用信息 JSON" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Ver) GetName() string {
	return "ver"
}

func (O *Ver) GetTitle() string {
	return "Ver"
}

