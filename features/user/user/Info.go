package user

import (
	"github.com/hailongz/golang/db"
)

type Info struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Key	string	`json:"key" name:"key" title:"key" length:"64" index:"ASC"`
	Value	string	`json:"value" name:"value" title:"内容" length:"-1"`
}

func (O *Info) GetName() string {
	return "info"
}

func (O *Info) GetTitle() string {
	return "用户信息"
}

