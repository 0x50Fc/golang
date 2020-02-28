package adv

import (
	"github.com/hailongz/golang/db"
)

type Adv struct {
	db.Object
	Channel	string	`json:"channel" name:"channel" title:"频道" index:"ASC"`
	Title	string	`json:"title" name:"title" title:"标题" index:"ASC"`
	Position	int32	`json:"position" name:"position" title:"广告组位置" index:"ASC"`
	Pic	string	`json:"pic" name:"pic" title:"图片" length:"128"`
	Description	string	`json:"description" name:"description" title:"描述" length:"512"`
	Link	string	`json:"link" name:"link" title:"跳转链接" length:"128"`
	Linktype	int32	`json:"linktype" name:"linktype" title:"跳转类型"`
	Sort	int32	`json:"sort" name:"sort" title:"排序"`
	Starttime	int64	`json:"starttime" name:"starttime" title:"开始时间"`
	Endtime	int64	`json:"endtime" name:"endtime" title:"结束时间"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Adv) GetName() string {
	return "adv"
}

func (O *Adv) GetTitle() string {
	return "广告"
}

